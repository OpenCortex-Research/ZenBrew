/**
 * @file repo.go
 * @description
 * @author
 * @copyright
 */

package repo

import (
	"encoding/json"
	"errors"
	log "log/slog"
	"net/http"
	"strings"

	"OpenCortex/ZenBrew/pkg"
	"OpenCortex/ZenBrew/utils"
)

type Repo struct {
	Name       string `json:"name"`
	Format    string `json:"format"`
	Maintainer string `json:"maintainer"`
	URL        string `json:"url"`
	Packages   []pkg.PackageLink `json:"packages"`
}

func DownloadRepoJson(repo_url string) Repo {
	json_url := repo_url + "repo.json"
	//hash_url := repo_url + "repo.sha256"

	json_bytes := utils.DownloadFile(json_url)
	//hash_bytes := utils.DownloadFile(hash_url)

	//if !utils.CheckHash(json_bytes, hash_bytes) {
	//	log.Error("Hashes do not match.")
	//	panic("Hashes do not match.")
	//}

	var repo Repo
	err := json.Unmarshal(json_bytes, &repo)
	if err != nil {
		log.Error("Failed to unmarshal JSON:", err)
		panic("Failed to unmarshal JSON")
	}

	return repo
}

func (repo Repo) CheckPackage(package_name string) (string, error) {
	if repo.Format == "array" {
		if repo.Packages == nil {
			return "", errors.New("no packages in the repo")
		}
		for _, package_link := range repo.Packages {
			if strings.EqualFold(package_link.Name, package_name) {
				return package_link.URL, nil
			}
		}
	}
	if repo.Format == "files" {
		packageURL := repo.URL + "/packages/" + package_name
		resp, err := http.Head(packageURL)
		if err != nil {
			log.Error("Failed to check package existence:", err)
			return "", errors.New("failed to check package existence")
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			return packageURL, nil
		}
	}
	return "", errors.New("package not found")
}

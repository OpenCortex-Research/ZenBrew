/**
 * @file repo.go
 * @description
 * @author
 * @copyright
 */

package repo

import (
	"encoding/json"
	log "log/slog"
	"net/http"

	"OpenCortex/ZenBrew/utils"
)

type Repo struct {
	Name       string `json:"name"`
	Version    string `json:"version"`
	Maintainer string `json:"maintainer"`
	URL        string `json:"url"`
}

func DownloadRepoJson(repo_url string) Repo {
	json_url := repo_url + "repo.json"
	hash_url := repo_url + "repo.sha256"

	json_bytes := utils.DownloadFile(json_url)
	hash_bytes := utils.DownloadFile(hash_url)

	if !utils.CheckHash(json_bytes, hash_bytes) {
		log.Error("Hashes do not match.")
		panic("Hashes do not match.")
	}

	var repo Repo
	err := json.Unmarshal(json_bytes, &repo)
	if err != nil {
		log.Error("Failed to unmarshal JSON:", err)
		panic("Failed to unmarshal JSON")
	}

	return repo
}

func (repo Repo) CheckPackage(package_name string) bool {
	packageURL := repo.URL + "/packages/" + package_name
	resp, err := http.Head(packageURL)
	if err != nil {
		log.Error("Failed to check package existence:", err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return true
	}
	return false
}
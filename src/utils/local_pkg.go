/**
 * @file installed_packages.go
 * @description
 * @author
 * @copyright
 */

package utils

import (
	"encoding/json"
	log "log/slog"
	"os"
)

type InstalledPackage struct {
	Name    string `json:"name"`
	Version	string `json:"version"`
	Status	string `json:"status"`
	Repository	string `json:"repository"`
}

func GetInstalledPackages() []InstalledPackage {
	// Read the JSON file
	file, err := os.ReadFile("installed_packages.json")
	if os.IsNotExist(err) {
		return []InstalledPackage{}
	} else if err != nil {
		log.Error(err.Error())
		panic(err)
	}

	// Parse the JSON data into the settings structure
	var installed_packages []InstalledPackage
	err = json.Unmarshal(file, &installed_packages)
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}

	return installed_packages
}

func CheckIfPackageInstalled(name string) (bool, string) {
	installed_packages := GetInstalledPackages()
	if len(installed_packages) == 0 {
		return false, ""
	}
	for _, installed_package := range installed_packages {
		if installed_package.Name == name {
			return true, installed_package.Version
		}
	}
	return false, ""
}

func AddInstalledPackage(name string, version string, status string, repo_name string) {
	installed_packages := GetInstalledPackages()
	new_package := InstalledPackage{
		Name: name,
		Version: version,
		Status: status,
		Repository: repo_name,
	}
	installed_packages = append(installed_packages, new_package)
	SaveInstalledPackages(installed_packages)
}

func SaveInstalledPackages(installed_packages []InstalledPackage) {
	// Convert the settings structure to JSON
	json_data, err := json.Marshal(installed_packages)
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}

	// Write the JSON data to the file
	err = os.WriteFile("installed_packages.json", json_data, 0644)
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}
}

func SetPackageStatus(name string, repo string, version string, status string) {
	installed_packages := GetInstalledPackages()
	found := false
	for i, installed_package := range installed_packages {
		if installed_package.Name == name && installed_package.Repository == repo {
			found = true
			installed_packages[i].Version = version
			installed_packages[i].Status = status
		}
	}
	if !found {
		new_package := InstalledPackage{
			Name: name,
			Version: version,
			Status: status,
			Repository: repo,
		}
		installed_packages = append(installed_packages, new_package)
	}
	SaveInstalledPackages(installed_packages)
}
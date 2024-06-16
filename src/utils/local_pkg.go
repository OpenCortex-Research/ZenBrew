/**
 * @file installed_packages.go
 * @description
 * @author
 * @copyright
 */

package utils

import (
	zb_types "OpenCortex/ZenBrew/types"
	"encoding/json"
	log "log/slog"
	"os"
)

type Package struct {
	zb_types.Package
}

type PackageLink struct {
	zb_types.PackageLink
}

type InstalledPackage struct {
	zb_types.InstalledPackage
}

func GetInstalledPackages() []InstalledPackage {
	// Read the JSON file
	file, err := os.ReadFile(Preferences.RootDir + "installed_packages.json")
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
			return true, installed_package.Version.Version
		}
	}
	return false, ""
}

func AddInstalledPackage(installed_package Package, status string, repo_name string, version_index int) {
	installed_packages := GetInstalledPackages()
	new_package := InstalledPackage{}
	new_package.Name = installed_package.Name
	new_package.Version = installed_package.Versions[version_index]
	new_package.Format = installed_package.Format
	new_package.Maintainer = installed_package.Maintainer
	new_package.Status = status
	new_package.Repository = repo_name
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
	err = os.WriteFile(Preferences.RootDir + "installed_packages.json", json_data, 0644)
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}
}

func SetPackageStatus(updated_package Package, status string, version_index int, repo_name string) {
	installed_packages := GetInstalledPackages()
	found := false
	for i, installed_package := range installed_packages {
		if installed_package.Name == updated_package.Name && installed_package.Repository == repo_name {
			found = true
			installed_packages[i].Version = updated_package.Versions[version_index]
			installed_packages[i].Status = status
		}
	}
	if !found {
		new_package := InstalledPackage{}
		new_package.Name = updated_package.Name
		new_package.Version = updated_package.Versions[version_index]
		new_package.Format = updated_package.Format
		new_package.Maintainer = updated_package.Maintainer
		new_package.Status = status
		new_package.Repository = repo_name
		installed_packages = append(installed_packages, new_package)
	}
	SaveInstalledPackages(installed_packages)
}

func RemoveInstalledPackage(name string, repo_name string) {
	installed_packages := GetInstalledPackages()
	for i, p := range installed_packages {
		if p.Name == name && p.Repository == repo_name {
			installed_packages = append(installed_packages[:i], installed_packages[i+1:]...)
			break
		}
	}
	SaveInstalledPackages(installed_packages)
}
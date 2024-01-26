/**
 * @file package.go
 * @description
 * @author
 * @copyright
 */

package main

import (
	"encoding/json"
	"io"
	log "log/slog"
	"net/http"
	"os"
	"path"
)

type Package struct {
	name string "json:string"
	version string "json:string"
	maintainer string "json:string"
	url string "json:string"
}

type PackageLink struct {
	name string
	url string
}

func download_package_metadata(package_link PackageLink) Package {
	json_url := package_link.url + "package.json"
	hash_url := package_link.url + "package.sha256"

	json_bytes := download_file(json_url)
	hash_bytes := download_file(hash_url)

	if !check_hash(json_bytes, hash_bytes) {
		log.Error("Hashes do not match.")
		panic("Hashes do not match.")
	}

	var pkg Package
	err := json.Unmarshal(json_bytes, &pkg)
	if err != nil {
		log.Error("Failed to unmarshal JSON:", err)
		panic("Failed to unmarshal JSON")
	}

	return pkg
}

func (pkg Package) download() {
	package_url := pkg.url + "package.tar.gz"
	package_path := path.Join(settings.root_dir, "zenbrew", pkg.name, pkg.version)

	// Download the package
	resp, err := http.Get(package_url)
	if err != nil {
		log.Error("Failed to download package:", err)
		panic("Failed to download package")
	}
	defer resp.Body.Close()

	// Create the directory if it doesn't exist
	err = os.MkdirAll(package_path, os.ModePerm)
	if err != nil {
		log.Error("Failed to create package directory:", err)
		panic("Failed to create package directory")
	}

	// Create the tar.gz file
	filePath := path.Join(package_path, "package.tar.gz")
	file, err := os.Create(filePath)
	if err != nil {
		log.Error("Failed to create package file:", err)
		panic("Failed to create package file")
	}
	defer file.Close()

	// Copy the response body to the file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Error("Failed to save package file:", err)
		panic("Failed to save package file")
	}

	// Extract the tar.gz file
	err = extract_tar(filePath, package_path)
	if err != nil {
		log.Error("Failed to extract package:", err)
		panic("Failed to extract package")
	}
}

func (pkg Package) install() {
	
}
/**
 * @file repo.go
 * @description
 * @author
 * @copyright
 */

package main

import (
	"encoding/json"
	log "log/slog"
)

type Repo struct {
	Name       string `json:"name"`
	Version    string `json:"version"`
	Maintainer string `json:"maintainer"`
	URL        string `json:"url"`
}

func download_repo_json(repo_url string) Repo {
	json_url := repo_url + "repo.json"
	hash_url := repo_url + "repo.sha256"

	json_bytes := download_file(json_url)
	hash_bytes := download_file(hash_url)

	if !check_hash(json_bytes, hash_bytes) {
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
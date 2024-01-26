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

type RepoLink struct{
	name string
	url string
}

type Repo struct {
	name string "json:string"
	version string "json:string"
	maintainer string "json:string"
	url string "json:string"
}

func download_repo_json(repo_link RepoLink) Repo {
	json_url := repo_link.url + "repo.json"
	hash_url := repo_link.url + "repo.sha256"

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
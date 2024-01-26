/**
 * @file hash.go
 * @description
 * @author
 * @copyright
 */

package main

import (
	"crypto/sha256"
	"fmt"
	log "log/slog"
)

func check_hash(file []byte, hash []byte) bool {

	// Calculate the hash of the file
	file_hash := sha256.Sum256(file)
	hash_from_file := fmt.Sprintf("%x", file_hash)
	hash_from_file = hash_from_file[:len(hash_from_file)-1] // Remove newline character

	if string(hash) == hash_from_file {
		log.Info("File and hash file match.")
		return true
	} else {
		log.Info("File and hash file do not match.")
		return false
	}
}

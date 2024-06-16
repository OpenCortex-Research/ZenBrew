/**
 * @file settings.go
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
type Settings struct {
	AutoCleanup bool   `json:"cleanup"`
	RootDir     string `json:"dir"`
	Repos        []string `json:"repos"`
}

var Preferences Settings

func GetSettings(zenbrew_dir string) {
	// Read the JSON file
	file, err := os.ReadFile(zenbrew_dir + "settings.json")
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}

	// Parse the JSON data into the settings structure
	err = json.Unmarshal(file, &Preferences)
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}
}
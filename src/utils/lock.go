/**
 * @file lock.go
 * @description
 * @author
 * @copyright
 */

package utils

import (
	log "log/slog"
	"os"
)

var SafeLock bool = false

func Lock() {
	// Create a file
	// If the file exists, wait
	// If the file does not exist, create it
	if _, err := os.Stat(Preferences.RootDir + "/ZenBrew/ZenBrew.lock"); err == nil {
		log.Error("ZenBrew is already running.")
		panic("ZenBrew is already running.")
	} else if os.IsNotExist(err) {
		SafeLock = true
		os.Create(Preferences.RootDir + "/ZenBrew/ZenBrew.lock")
	}
}

func Unlock() {
	// Delete the file
	// If the file does not exist, do nothing
	if !SafeLock {
		return
	}
	if _, err := os.Stat(Preferences.RootDir + "/ZenBrew/ZenBrew.lock"); err == nil {
		os.Remove(Preferences.RootDir + "/ZenBrew/ZenBrew.lock")
	} else if os.IsNotExist(err) {
		return
	}
}
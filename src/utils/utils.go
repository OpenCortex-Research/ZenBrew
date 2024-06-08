/**
 * @file download.go
 * @description
 * @author
 * @copyright
 */

package utils

import (
	"archive/tar"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	log "log/slog"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strings"
)

func CheckHash(file []byte, hash []byte) bool {

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

func DownloadFile(raw_url string) ([]byte) {
	log.Info(fmt.Sprintf("Downloading file from: %s", raw_url))
	// Download the file
	resp, err := http.Get(raw_url)
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}
	defer resp.Body.Close()

	// Read the file into a byte array
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}
	return bytes
}

// extract_tar extracts the contents of a tar.gz file to a specified destination directory.
// It takes the source file path and the destination directory path as input parameters.
// The function returns an error if any error occurs during the extraction process.
func ExtractTar(src, dest string) error {
	// Open the source file
	file, err := os.Open(src)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a gzip reader for the source file
	gz_reader, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gz_reader.Close()

	// Create a tar reader for the gzip reader
	tar_reader := tar.NewReader(gz_reader)

	var file_names []string
	// Iterate through each file in the tar archive
	for {
		// Read the next file header
		header, err := tar_reader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if (header.Name != "pax_global_header") && (strings.Split(header.Name, "/")[1] != "") {
			file_names = append(file_names, header.Name)
		}

		// Get the file path for the current file
		file_path := path.Join(dest, header.Name)

		// Handle different types of files
		switch header.Typeflag {
		case tar.TypeDir:
			// Create directories if the file is a directory
			err = os.MkdirAll(file_path, os.ModePerm)
			if err != nil {
				return err
			}
		case tar.TypeReg:
			// Create and write to regular files
			file, err := os.OpenFile(file_path, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}
			defer file.Close()

			// Copy the file contents from the tar reader to the file
			_, err = io.Copy(file, tar_reader)
			if err != nil {
				return err
			}
		}
	}

	// Delete the source tar file
	err = os.Remove(src)
	if err != nil {
		return err
	}

	for _, file_name := range file_names {
		cmd_err := exec.Command("mv", path.Join(dest, file_name), path.Join(dest)).Run()
		if cmd_err != nil {
			log.Error(fmt.Sprintf("Failed to run mv: %s", cmd_err))
			panic("Failed to run mv")
		}

		arr := strings.Split(file_name, "/")
		new_name := strings.Join(arr[1:], "/")

		ch_err := exec.Command("chmod", "+x", path.Join(dest, new_name)).Run()
		if ch_err != nil {
			log.Error(fmt.Sprintf("Failed to run chmod: %s", ch_err))
			panic("Failed to run chmod")
		}
	}
	return nil
}
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
	"path"
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

	return nil
}
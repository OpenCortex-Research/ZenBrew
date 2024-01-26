/**
 * @file download.go
 * @description
 * @author
 * @copyright
 */

package main

import (
	"archive/tar"
	"compress/gzip"
	"io"
	log "log/slog"
	"net/http"
	"os"
	"path"
)

func download_file(raw_url string) ([]byte) {

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

func extract_tar(src, dest string) error {
	file, err := os.Open(src)
	if err != nil {
		return err
	}
	defer file.Close()

	gzReader, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gzReader.Close()

	tarReader := tar.NewReader(gzReader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		filePath := path.Join(dest, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			err = os.MkdirAll(filePath, os.ModePerm)
			if err != nil {
				return err
			}
		case tar.TypeReg:
			file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}
			defer file.Close()

			_, err = io.Copy(file, tarReader)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
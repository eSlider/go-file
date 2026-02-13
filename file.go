// Package file provides cross-platform file system utilities for path resolution,
// existence checks, directory creation, and writability testing.
//
// Designed for use in Go applications and test suites that need to locate
// project root directories, verify file paths, and manage directories.
package file

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

var rootPath string

// GetModRootPath walks up from the current working directory looking for
// a project root (a directory containing "etc" or "data" subdirectories).
// Returns the path and an error if no root marker is found.
func GetModRootPath() (string, error) {
	path, _ := os.Getwd()
	var prevPath string
	for {
		if Exists(path+"/etc") || Exists(path+"/data") {
			rootPath = path
			return rootPath, nil
		}
		if prevPath == path || len(path) < 2 {
			break
		}
		prevPath = path
		path = filepath.Dir(path)
	}
	return path, errors.New("project root not found (no etc/ or data/ directory)")
}

// GetRootPath returns the cached project root path, panicking if it cannot be found.
func GetRootPath() string {
	if rootPath != "" {
		return rootPath
	}
	path, err := GetModRootPath()
	if err != nil {
		panic(err)
	}
	rootPath = path
	return path
}

// GetCmdRootPath returns the current working directory, stripping any /etc/ suffix.
func GetCmdRootPath() string {
	dir, _ := os.Getwd()
	if strings.Contains(dir, "/etc/") {
		dir = strings.Split(dir, "/etc/")[0]
	}
	return dir
}

// IsExist returns true if the path exists and is not a directory.
func IsExist(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// Size returns the file size in bytes, or 0 if the file does not exist.
func Size(path string) int64 {
	info, err := os.Stat(path)
	if err != nil {
		return 0
	}
	return info.Size()
}

// Exists returns true if the path exists (file or directory).
func Exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

// IsWritable tests whether the given directory path is writable by creating
// and immediately removing a temporary file.
func IsWritable(path string) bool {
	file, err := os.OpenFile(path+"/.test", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return false
	}
	defer func() {
		file.Close()
		os.Remove(path + "/.test")
	}()
	return true
}

// PreCreateDirectory ensures the directory at path exists, creating it
// (and any parents) if necessary.
func PreCreateDirectory(path string) error {
	if Exists(path) {
		return nil
	}
	return os.MkdirAll(path, 0755)
}

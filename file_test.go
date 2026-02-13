package file

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExists(t *testing.T) {
	// Create a temp file
	tmp, err := os.CreateTemp("", "go-file-test-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmp.Name())
	tmp.Close()

	if !Exists(tmp.Name()) {
		t.Errorf("Exists(%q) = false, want true", tmp.Name())
	}

	if Exists(tmp.Name() + "-nonexistent") {
		t.Error("Exists returned true for nonexistent path")
	}
}

func TestIsExist(t *testing.T) {
	// File should return true
	tmp, err := os.CreateTemp("", "go-file-test-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmp.Name())
	tmp.Close()

	if !IsExist(tmp.Name()) {
		t.Errorf("IsExist(%q) = false, want true", tmp.Name())
	}

	// Directory should return false
	dir := t.TempDir()
	if IsExist(dir) {
		t.Errorf("IsExist(%q) = true for directory, want false", dir)
	}
}

func TestSize(t *testing.T) {
	tmp, err := os.CreateTemp("", "go-file-test-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmp.Name())

	data := []byte("hello world")
	if _, err := tmp.Write(data); err != nil {
		t.Fatal(err)
	}
	tmp.Close()

	got := Size(tmp.Name())
	if got != int64(len(data)) {
		t.Errorf("Size() = %d, want %d", got, len(data))
	}

	if Size("/nonexistent/file") != 0 {
		t.Error("Size() should return 0 for nonexistent file")
	}
}

func TestIsWritable(t *testing.T) {
	dir := t.TempDir()
	if !IsWritable(dir) {
		t.Errorf("IsWritable(%q) = false, want true", dir)
	}

	if IsWritable("/root/nonexistent-dir-for-test") {
		t.Error("IsWritable should return false for non-writable path")
	}
}

func TestPreCreateDirectory(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "a", "b", "c")

	if Exists(dir) {
		t.Fatal("directory should not exist yet")
	}

	if err := PreCreateDirectory(dir); err != nil {
		t.Fatalf("PreCreateDirectory() error = %v", err)
	}

	if !Exists(dir) {
		t.Error("directory should exist after PreCreateDirectory")
	}

	// Calling again should be a no-op
	if err := PreCreateDirectory(dir); err != nil {
		t.Fatalf("PreCreateDirectory() second call error = %v", err)
	}
}

func TestGetCmdRootPath(t *testing.T) {
	path := GetCmdRootPath()
	if path == "" {
		t.Error("GetCmdRootPath() returned empty string")
	}
}

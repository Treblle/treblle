package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Storage struct {
	Path string
}

// New initializes a new Storage instance.
func New() *Storage {
	return &Storage{}
}

// At sets the current path.
func (s *Storage) At(path string) *Storage {
	s.Path = path
	return s
}

// Exists checks if the path exists.
func (s *Storage) Exists() bool {
	_, err := os.Stat(s.Path)
	return !os.IsNotExist(err)
}

// IsDir checks if the path is a directory.
func (s *Storage) IsDir() bool {
	info, err := os.Stat(s.Path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// CreateDir creates a directory at the current path.
func (s *Storage) CreateDir(perm os.FileMode) error {
	return os.MkdirAll(s.Path, perm)
}

// Write writes data to a file at the current path.
func (s *Storage) Write(data string) error {
	return os.WriteFile(s.Path, []byte(data), 0644)
}

// Read reads data from a file at the current path.
func (s *Storage) Read() (string, error) {
	data, err := os.ReadFile(s.Path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Delete removes the file or directory at the current path.
func (s *Storage) Delete() error {
	return os.RemoveAll(s.Path)
}

// AbsPath returns the absolute path.
func (s *Storage) AbsPath() (string, error) {
	return filepath.Abs(s.Path)
}

// WriteJSON writes a Go value as JSON to the current path.
func (s *Storage) WriteJSON(data interface{}) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return os.WriteFile(s.Path, bytes, 0644)
}

// ReadJSON reads JSON from the current path into a Go value.
func (s *Storage) ReadJSON(data interface{}) error {
	bytes, err := os.ReadFile(s.Path)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, data)
}

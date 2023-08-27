package config

import (
	"errors"

	"gitub.com/treblle/treblle/pkg/storage"
)

type Config struct {
	data map[string]interface{}
}

type ConfigEntry struct {
	value interface{}
	err   error
}

func New() *Config {
	return &Config{
		data: make(map[string]interface{}),
	}
}

// Set sets a value for the given key.
func (c *Config) Set(key string, value interface{}) *Config {
	c.data[key] = value

	return c
}

// Get retrieves the value for the given key.
func (c *Config) Get(key string) *ConfigEntry {
	if value, exists := c.data[key]; exists {
		return &ConfigEntry{
			value: value,
			err:   nil,
		}
	}
	return &ConfigEntry{
		value: nil,
		err:   errors.New("key not found"),
	}
}

// Int retrieves the int value.
func (e *ConfigEntry) Int() (int, error) {
	if e.err != nil {
		return 0, e.err
	}
	if v, ok := e.value.(int); ok {
		return v, nil
	}
	return 0, errors.New("value is not an int")
}

// String retrieves the string value.
func (e *ConfigEntry) String() (string, error) {
	if e.err != nil {
		return "", e.err
	}
	if v, ok := e.value.(string); ok {
		return v, nil
	}
	return "", errors.New("value is not a string")
}

// SaveToFile saves the current configuration to a file.
func (c *Config) SaveToFile(filePath string) error {
	s := storage.New().At(filePath)
	return s.WriteJSON(c.data)
}

// LoadFromFile loads the configuration from a file.
func (c *Config) LoadFromFile(filePath string) error {
	s := storage.New().At(filePath)
	return s.ReadJSON(&c.data)
}

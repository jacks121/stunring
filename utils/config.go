package utils

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/BurntSushi/toml"
)

type Config struct {
	mu      sync.RWMutex
	file    string
	entries map[string]interface{}
}

func NewConfig(file string) *Config {
	c := &Config{
		file:    file,
		entries: make(map[string]interface{}),
	}
	c.load()
	return c
}

func (c *Config) load() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, err := toml.DecodeFile(c.file, &c.entries); err != nil {
		log.Fatalf("Failed to load config: %s", err)
	}
}

func (c *Config) Get(key string) string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	parts := strings.Split(key, ".")
	value := c.entries

	for _, part := range parts {
		if v, ok := value[part].(map[string]interface{}); ok {
			value = v
		} else {
			return fmt.Sprintf("%v", value[part])
		}
	}

	return ""
}

func (c *Config) GetJWTSecret() []byte {
	return []byte(c.Get("JWT_SECRET"))
}

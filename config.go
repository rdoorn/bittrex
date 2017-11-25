package bittrex

import (
	"io/ioutil"
	"strings"
	"sync"
	"time"

	"github.com/BurntSushi/toml"
	yaml "gopkg.in/yaml.v2"
)

const (
	// URL is the main site url
	bittrexURL string = "https://bittrex.com"
	// HTTPTimeout is the http timeout
	httpTimeout time.Duration = 30
)

var (
	config     *KeyConfig
	configLock sync.RWMutex

	apiKey    string
	apiSecret string

	testMode = false
)

// Config holds your main config
type Config struct {
	Key    string `yaml:"key"`
	Secret string `yaml:"secret"`
}

// Config holds your main config
type KeyConfig struct {
	//Logging LoggingConfig
	Bittrex Config `yaml:"bittrex"`
}

// LoadConfig a config file
func LoadConfig(file string) error {
	//log.Println("Loading config")
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	temp := new(KeyConfig)
	f := strings.Split(file, ".")
	switch f[len(f)-1] {
	case "toml":
		_, err = toml.Decode(string(data), temp)
		if err != nil {
			return err
		}
	case "yaml":
		err = yaml.Unmarshal([]byte(data), temp)
		if err != nil {
			return err
		}
	}

	configLock.Lock()
	config = temp
	//log.Println("Config loaded succesfully")
	configLock.Unlock()
	return nil
}

// Get returns the pointer to the latest config loaded
func Get() *KeyConfig {
	configLock.RLock()
	defer configLock.RUnlock()
	return config
}

// Lock the config for Writes
func Lock() {
	configLock.Lock()
}

// Unlock the config for Writes
func Unlock() {
	configLock.Unlock()
}

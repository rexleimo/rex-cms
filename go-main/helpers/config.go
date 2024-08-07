package helpers

import (
	"os"
	"sync"

	"github.com/pelletier/go-toml/v2"
)

var (
	configInstance *ConfigInstance
	configOnce     sync.Once
)

type ConfigInstance struct {
	Config *Config
}

type Config struct {
	Github Github
}

type Github struct {
	Owner string
	Repo  string
	Token string
}

func GetConfigInstance() *ConfigInstance {
	configOnce.Do(func() {
		configInstance = &ConfigInstance{}
		configInstance.New()
	})
	return configInstance
}

func (c *ConfigInstance) New() {
	var contacts Config
	data, err := os.ReadFile("config.toml")
	if err != nil {
		panic(err)
	}
	err = toml.Unmarshal(data, &contacts)
	if err != nil {
		panic(err)
	}
	c.Config = &contacts
}

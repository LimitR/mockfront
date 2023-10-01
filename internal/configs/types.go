package configs

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Api struct {
	Path   string
	Method string
}

type Config struct {
	api []Api
}

func NewConfig(cap uint) *Config {
	return &Config{
		api: make([]Api, 0, cap),
	}
}

func (c *Config) add(path, method string) {
	c.api = append(c.api, Api{
		Path:   path,
		Method: method,
	})
}

func (c *Config) GetAll() []Api {
	return c.api
}

func (c *Config) Init(path string) {
	f, _ := loadConfig(path)
	for path, v := range f["paths"].(map[interface{}]interface{}) {
		for method := range v.(map[interface{}]interface{}) {
			c.add(path.(string), method.(string))
		}
	}
}

func loadConfig(fileName string) (map[interface{}]interface{}, error) {
	m := make(map[interface{}]interface{})
	b, err := os.ReadFile(fileName)
	err = yaml.Unmarshal(b, &m)
	if err != nil {
		return m, err
	}
	return m, nil
}

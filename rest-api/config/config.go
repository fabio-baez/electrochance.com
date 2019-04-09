package config

import (
	"log"
	"io/ioutil"
    "path/filepath"

	"gopkg.in/yaml.v2"
)


// YAML Represents database server and credentials
type YAML struct {
	Server   string `yaml:"server"`
	Database string `yaml:"database"`
	Collection string `yaml:"collection"`
	Port string `yaml:"port"`

}


// ReadYaml and parse the YAML configuration file
func (c *YAML) ReadYaml() {
	filename, _ := filepath.Abs("./config.yml")
    yamlFile, err := ioutil.ReadFile(filename)

    if err != nil {
       log.Fatal(err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
    if err != nil {
        log.Fatal(err)
    }
}


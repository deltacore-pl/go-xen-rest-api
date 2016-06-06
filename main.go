package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Config yaml parameters
type Config struct {
	Main struct {
		APIPort string
	}
}

func main() {
	configFile, _ := filepath.Abs("./config.yml")
	yamlFile, err := ioutil.ReadFile(configFile)

	if err != nil {
		panic(err)
	}

	var config Config
	err = yaml.Unmarshal(yamlFile, &config)

	if err != nil {
		panic(err)
	}

	fmt.Println(config.Main.APIPort)
}

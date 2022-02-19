package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type conf struct {
	Token          string `yaml:"token"`
	RandomImageUrl string `yaml:"random_image_url"`
}

func ParseConfig(path string) *conf {
	c := conf{}
	yamlData, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Read config err: %v ", err)
	}
	if err := yaml.Unmarshal(yamlData, &c); err != nil {
		log.Fatalf("Unmarshal yaml: %v", err)
	}

	return &c
}

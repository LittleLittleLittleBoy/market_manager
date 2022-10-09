package settings

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Setting struct {
	Ocr OCR `yaml:"ocr"`
}

type OCR struct {
	Endpoint string `yaml:"endpoint"`
	Key      string `yaml:"key"`
}

var Config = &Setting{}

func Setup() {

	yfile, err := ioutil.ReadFile("config.yaml")

	if err != nil {
		log.Fatal("config file can not read,", err)
	}

	err = yaml.Unmarshal(yfile, Config)
	if err != nil {
		log.Fatal("config file can not map tp yaml", err)
	}
}

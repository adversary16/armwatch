package confParser

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/BurntSushi/toml"
)

type Configuration map[string]interface{}

var allowedExtensions = []string{
	"toml",
	"json",
}

func Parse(path *os.FileInfo, confPointer *Configuration) {
	if *path == nil {
		return
	}
	configType := filepath.Ext((*path).Name())

	switch configType {
	case ".toml":
		_, err := toml.DecodeFile((*path).Name(), &confPointer)
		if err != nil {
			log.Fatal(err)
		}
	case ".json":
		file, err := os.Open((*path).Name())
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		decoder := json.NewDecoder(file)
		decoder.Decode(&confPointer)
	default:
		return
	}
}

func FindConfFiles() *os.FileInfo {
	regexString := strings.Join([]string{
		"^conf\\.(",
		strings.Join(allowedExtensions, "|"),
		")$",
	}, "")
	regexPattern, err := regexp.Compile(regexString)
	var confPath os.FileInfo

	if err != nil {
		log.Fatal(err)
	}

	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if regexPattern.MatchString(file.Name()) {
			confPath = file
			return &confPath
		}
	}
	return &confPath
}

func Init(confPointer *Configuration) {
	Parse(FindConfFiles(), confPointer)
}

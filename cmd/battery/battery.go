package battery

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var currentCapacity int
var batteryPath string

func FindEndpoint() (string, error) {
	var err error
	var endpointPath string
	var basePath = "/sys/class/power_supply/"

	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.Mode()&os.ModeSymlink != 0 {
			filePath := strings.Join([]string{basePath, f.Name()}, "")
			linkPath, err := filepath.EvalSymlinks(filePath)
			if err != nil {
				log.Fatal(err)
			}

			contents, err := ioutil.ReadDir(linkPath)
			if err != nil {
				log.Fatal(err)
			}

			for _, s := range contents {
				if s.Name() == "capacity" {
					endpointPath = strings.Join([]string{linkPath, s.Name()}, "/")
					return endpointPath, err
				}
			}
		}
	}
	return endpointPath, err
}

func UpdateCapacity() (int, error) {
	var capacity int
	var err error

	FindEndpoint()

	file, err := os.Open(batteryPath)

	if err != nil {
		return capacity, nil
	}
	raw, err := ioutil.ReadAll(file)

	if err != nil {
		return capacity, nil
	}

	capacity, err = strconv.Atoi(string(raw[0:2]))

	if err != nil {
		return capacity, err
	}
	if capacity != 0 {
		currentCapacity = capacity
	}
	fmt.Println(capacity)
	return capacity, nil
}

func Get() int {
	return currentCapacity
}

func Init() {
	batEndpoint, err := FindEndpoint()
	if err != nil {
		log.Fatal(err)
	}
	batteryPath = batEndpoint
	UpdateCapacity()
}

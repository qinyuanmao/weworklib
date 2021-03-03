package weworklib

import (
	"io/ioutil"
	"log"
)

func ReadFile(path string) string {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}

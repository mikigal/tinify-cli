package tinify

import (
	"errors"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"log"
	"strings"
)

var TooManyRequests = errors.New("You exceeded your monthly limit")
var Unauthorized = errors.New("Your API KEY is invalid")
var UnsupportedMediaType = errors.New("Invalid file. Supported file types: png, jpg, jpeg")

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func SaveKey(key string) error {
	home, err := homedir.Dir()
	Check(err)

	if !strings.HasSuffix(home, "/") {
		home += "/"
	}

	err = ioutil.WriteFile(home+".tinify", []byte(key), 0644)
	return err
}

func GetKey() (string, error) {
	home, err := homedir.Dir()
	Check(err)

	if !strings.HasSuffix(home, "/") {
		home += "/"
	}

	bytes, err := ioutil.ReadFile(home + ".tinify")
	if err != nil {
		return "", err
	}

	key := string(bytes)
	return key, nil
}

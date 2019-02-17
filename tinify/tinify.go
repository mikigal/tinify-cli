package tinify

import (
	"errors"
	"fmt"
	"github.com/antonholmquist/jason"
	"io/ioutil"
	"net/http"
	"os"
)

func Upload(key string, file *os.File) (Response, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://api.tinify.com/shrink", file)
	req.SetBasicAuth("api", key)

	resp, err := client.Do(req)
	if err != nil {
		return Response{}, err
	}

	if resp.StatusCode == http.StatusTooManyRequests {
		return Response{}, TooManyRequests
	}
	if resp.StatusCode == http.StatusUnauthorized {
		return Response{}, Unauthorized
	}
	if resp.StatusCode != http.StatusCreated {
		return Response{}, errors.New(fmt.Sprintf("Unknown error. HTTP status code: %d", resp.StatusCode))
	}

	bytes, err := ioutil.ReadAll(resp.Body)

	json, err := jason.NewObjectFromBytes(bytes)
	if err != nil {
		return Response{}, err
	}

	input, _ := json.GetObject("input")
	output, _ := json.GetObject("output")

	inputSize, _ := input.GetInt64("size")
	inputType, _ := input.GetString("type")

	size, _ := output.GetInt64("size")
	typee, _ := output.GetString("type")
	width, _ := output.GetInt64("width")
	height, _ := output.GetInt64("height")
	ratio, _ := output.GetFloat64("ratio")
	url, _ := output.GetString("url")

	return Response {
		InputSize: inputSize,
		InputType: inputType,

		Size: size,
		Type: typee,
		Width: width,
		Height: height,
		Ratio: ratio,
		Url: url,
	}, err
}
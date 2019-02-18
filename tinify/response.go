package tinify

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
)

type Response struct {
	InputSize int64
	InputType string

	Size   int64
	Type   string
	Width  int64
	Height int64
	Ratio  float64
	Url    string
}

func (res *Response) Download(name string, key string) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", res.Url, nil)
	req.SetBasicAuth("api", key)

	resp, err := client.Do(req)
	Check(err)

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusTooManyRequests {
		log.Fatal(TooManyRequests)
	}
	if resp.StatusCode == http.StatusUnauthorized {
		log.Fatal(Unauthorized)
	}
	if resp.StatusCode == http.StatusUnsupportedMediaType {
		log.Fatal(UnsupportedMediaType)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatal(fmt.Sprintf("Unknown download error. HTTP status code: %d. Please report it on GitHub or by mail (mikigal.priv@gmail.com). \n Debug info: %v ", resp.StatusCode, resp))
	}

	out, err := os.Create(name)
	Check(err)

	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	Check(err)
}

func (res *Response) CalcPercent() int64 {
	return int64(math.Round((1 - res.Ratio) * 100))
}

func (res *Response) CalcSizeKB() int64 {
	return (res.InputSize - res.Size) / 1000
}

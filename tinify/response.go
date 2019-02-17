package tinify

import (
	"io"
	"math"
	"net/http"
	"os"
)

type Response struct {
	InputSize int64
	InputType string

	Size int64
	Type string
	Width int64
	Height int64
	Ratio float64
	Url string
}

func (res *Response) Download(name string)  {
	resp, err := http.Get(res.Url)
	Check(err)

	defer resp.Body.Close()

	out, err := os.Create(name)
	Check(err)

	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	Check(err)
}

func (res *Response) CalcPercent() int64{
	return int64(math.Round((1 - res.Ratio) * 100))
}

func (res *Response) CalcSizeKB() int64 {
	return (res.InputSize - res.Size) / 1000
}
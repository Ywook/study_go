package lib

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const base_url string = "https://api.bithumb.com"

func init() {
	fmt.Println("Import Go Bithumb Package!")
}

func Get(path string, timeout int) (*http.Response, []byte, error) {
	uri := base_url + path
	resp, err := http.Get(uri)

	if err != nil {
		return nil, nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, nil, err
	}

	return resp, body, nil
}

func Post() {

}

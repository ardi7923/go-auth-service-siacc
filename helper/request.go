package helper

import (
	"fmt"
	"io"
	"net/http"
)

func RequestUrl(method, url, token string, payload io.Reader) (*http.Response, error) {

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)

	fmt.Println("DATA ==>", method, url, payload)
	fmt.Println("REQUEST ==>", res.StatusCode, res.Status, res.Header, err)
	if err != nil {
		return nil, err
	}
	return res, nil
}

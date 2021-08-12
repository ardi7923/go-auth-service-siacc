package test

import (
	"auth/config"
	"auth/router"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type Body struct {
	Key   string
	Value string
	Type  string
}

type Request struct {
	Method   string
	Header   map[string]string
	Endpoind string
	Body     []Body
}

type Items struct {
	Name           string  `json:"name"`
	Request        Request `json:"request"`
	ResponseStatus []int   `json:"response_status"`
}

type Group struct {
	Group string  `json:"group"`
	Item  []Items `json:"item"`
}

type ListEndpoind struct {
	Router []Group `json:"router"`
}

var json_list_services *os.File
var list_services ListEndpoind

func init() {
	json_list_services, err := os.Open(config.List_services_files)
	byte_val, _ := ioutil.ReadAll(json_list_services)
	json.Unmarshal(byte_val, &list_services)
	if err != nil {
		panic(err)
	}
}

func intInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
func TestRouter(t *testing.T) {
	ts := httptest.NewServer(router.SetupServer())
	defer ts.Close()

	new_req := func(method string, url string, payload io.Reader) *http.Request {
		req, err := http.NewRequest(method, url, payload)
		if err != nil {
			t.Fatal("fatal")
		}
		return req
	}

	for _, val := range list_services.Router {
		fmt.Println(fmt.Sprintf("============== testing group %s ==============", val.Group))
		for _, val_item := range val.Item {
			t.Run("testing_route", func(t *testing.T) {
				body_value := make(map[string]string, len(val_item.Request.Body))
				for _, val_body := range val_item.Request.Body {
					body_value[val_body.Key] = val_body.Value
				}
				json_body, _ := json.Marshal(body_value)
				req := new_req(val_item.Request.Method, ts.URL+val_item.Request.Endpoind, bytes.NewBuffer(json_body))
				for key_header, val_header := range val_item.Request.Header {
					req.Header.Set(key_header, val_header)
				}
				res, err := http.DefaultClient.Do(req)
				if err != nil {
					panic(err)
				}
				if intInSlice(res.StatusCode, val_item.ResponseStatus) == false {
					fmt.Println(fmt.Sprintf("TEST FAIL ===> group:%s ===> name : %s", val.Group, val_item.Name))
					res_body, _ := ioutil.ReadAll(res.Body)
					fmt.Println("response Body:", string(res_body))
					panic(res.Body)
				}
			})
		}

	}
}

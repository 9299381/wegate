package main

import (
	"encoding/json"
	"fmt"
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/tools/convert"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestProxy(t *testing.T) {
	host := "127.0.0.1:8341"
	service := "demo.post"
	params := map[string]interface{}{}
	params["name"] = "wang"
	params["number"] = "1"

	path := "http://" + host + "/" + strings.Replace(service, ".", "/", -1)
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	data := url.Values{}
	for k, v := range params {
		data.Set(k, v.(string))
	}
	resp, err := client.PostForm(path, convert.FormEncode(params))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	ret := &contracts.Response{}
	_ = json.Unmarshal(body, ret)
	fmt.Println(ret)

}

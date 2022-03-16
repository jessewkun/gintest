package utils

import (
	"encoding/json"
	"fmt"
	"gintest/config"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type XHttp struct {
	Url     string            `json:"url"`
	Get     string            `json:"get"`
	Post    string            `json:"post"`
	Timeout int               `json:"timeout"`
	Method  string            `json:"method"`
	Headers map[string]string `json:"headers"`
}

func (xh XHttp) String() string {
	jsonBytes, err := json.Marshal(xh)
	if err != nil {
		return fmt.Sprintf("Url: %s, Get: %s, Post: %s, Timeout: %d, Method: %s, Headers: %s", xh.Url, xh.Get, xh.Post, xh.Timeout, xh.Method, xh.Headers)
	}
	return string(jsonBytes)
}

const DEFAULT_TIMEOUT = 3

func NewXhttp() *XHttp {
	return &XHttp{}
}

// http 请求
func (x XHttp) DoRequest() (string, *APIException) {
	var err error

	url := x.Url + "?" + x.Get
	req, err := http.NewRequest(x.Method, url, strings.NewReader(x.Post))
	if err != nil {
		return "", NewAPIException(config.ERROR_CREATE_REQUEST, err)
	}

	for k, v := range x.Headers {
		req.Header.Set(k, v)
	}

	timeout := DEFAULT_TIMEOUT
	if x.Timeout > 0 {
		timeout = x.Timeout
	}
	client := &http.Client{
		Timeout: time.Duration(timeout * int(time.Second)),
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", NewAPIException(config.ERROR_SEND_REQUEST_FAIL, err)
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", NewAPIException(config.ERROR_READ_RESPONSE_FAIL, err)
	}

	return string(respBytes), nil
}

// 生成 http query
func HttpBuildQuery(data map[string]interface{}) string {
	var uri url.URL
	q := uri.Query()
	for k, v := range data {
		q.Add(k, Strval(v))
	}
	return q.Encode()
}

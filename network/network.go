package network

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	METHOD_GET  = "GET"
	METHOD_POST = "POST"
)

func JSON(url string, headers map[string]string, params map[string]any, cookies []*http.Cookie) Response {
	jsonBytes, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err)
		return NewResponse(0, nil, err)
	}

	request, err := http.NewRequest(METHOD_POST, url, bytes.NewReader(jsonBytes))
	if err != nil {
		fmt.Println(err)
		return NewResponse(0, nil, err)
	}

	for k, v := range headers {
		request.Header.Set(k, v)
	}

	for _, cookie := range cookies {
		request.AddCookie(cookie)
	}

	resp, err := (&http.Client{}).Do(request)
	if err != nil {
		fmt.Println(err)
		return NewResponse(resp.StatusCode, nil, err)
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return NewResponse(resp.StatusCode, nil, err)
	}

	r := NewResponse(resp.StatusCode, data, nil)
	r.Cookies = resp.Cookies()
	return r
}

func GET(url string, headers map[string]string, params map[string]any, cookies []*http.Cookie) Response {
	temp := make([]string, len(params))
	index := 0
	for k, v := range params {
		temp[index] = fmt.Sprintf("%s=%s", k, v)
		index ++
	}
	pString := strings.Join(temp, "&")

	url = strings.TrimSuffix(url, "?")
	url = strings.TrimSuffix(url, "&")
	if strings.Contains(pString, "?") {
		url += "&" + pString
	} else {
		url += "?" + pString
	}
	request, err := http.NewRequest(METHOD_GET, url, nil)
	if err != nil {
		fmt.Println(err)
		return NewResponse(0, nil, err)
	}

	for k, v := range headers {
		request.Header.Set(k, v)
	}

	for _, cookie := range cookies {
		request.AddCookie(cookie)
	}

	return sendRequest(request)
}

func POST(url string, headers map[string]string, params map[string]any, cookies []*http.Cookie) Response {
	temp := make([]string, len(params))
	index := 0
	for k, v := range params {
		temp[index] = fmt.Sprintf("%s=%s", k, v)
		index ++
	}
	pString := strings.Join(temp, "&")

	request, err := http.NewRequest(METHOD_POST, url, strings.NewReader(pString))
	if err != nil {
		fmt.Println(err)
		return NewResponse(0, nil, err)
	}

	for k, v := range headers {
		request.Header.Set(k, v)
	}

	for _, cookie := range cookies {
		request.AddCookie(cookie)
	}

	return sendRequest(request)
}

// 发送请求
func sendRequest(request *http.Request) Response {
	resp, err := (&http.Client{}).Do(request)
	if err != nil {
		fmt.Println(err)
		return NewResponse(resp.StatusCode, nil, err)
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return NewResponse(resp.StatusCode, nil, err)
	}

	r := NewResponse(resp.StatusCode, data, nil)
	r.Cookies = resp.Cookies()
	return r
}

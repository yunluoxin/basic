package network

import (
	"io/ioutil"
	"net/http"
)

type Response struct {
	StatusCode int
	Err error
	Content []byte
	Cookies []*http.Cookie
	ContentType string
	UserAgent string
}

// NewResponse 创建一个新的 Response
func NewResponse(statusCode int, content []byte, err error) Response {
	var r Response
	r.StatusCode = statusCode
	r.Err = err
	r.Content = content
	return r
}

func NewResponseWithHttpResponse(resp http.Response) Response {
	data, err := ioutil.ReadAll(resp.Body)
	var r Response
	r.StatusCode = resp.StatusCode
	r.Err = err
	r.Content = data
	r.Cookies = resp.Cookies()
	r.ContentType = resp.Header.Get("Content-Type")
	r.UserAgent = resp.Header.Get("UserAgent")
	return r
}
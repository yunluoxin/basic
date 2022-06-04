package network

import (
	"log"
	"net/http"
)

type Session struct {
	Headers map[string]string
	Cookies []*http.Cookie
	DisableAutoUpdateCookie bool
}

func (s *Session) JSON(url string, headers map[string]string, params map[string]any) (string, error) {
	newHeaders := updateTwoMap(headers, s.Headers)
	newHeaders["Content-Type"] = "application/json; charset=utf-8"

	resp := JSON(url, newHeaders, params, s.Cookies)
	if resp.Err != nil {
		log.Println(resp.Err)
		return "", resp.Err
	} else {
		if !s.DisableAutoUpdateCookie {
			s.Cookies = append(s.Cookies, resp.Cookies...)
			s.Cookies = removeDuplicateCookies(s.Cookies)
		}
	}
	return string(resp.Content), nil
}


func (s *Session) GET(url string, headers map[string]string, params map[string]any) (string, error) {
	newHeaders := updateTwoMap(headers, s.Headers)
	resp := GET(url, newHeaders, params, s.Cookies)
	if resp.Err != nil {
		log.Println(resp.Err)
		return "", resp.Err
	} else {
		if !s.DisableAutoUpdateCookie {
			s.Cookies = append(s.Cookies, resp.Cookies...)
			s.Cookies = removeDuplicateCookies(s.Cookies)
		}
	}
	return string(resp.Content), nil
}

func (s *Session) POST(url string, headers map[string]string, params map[string]any) (string, error) {
	newHeaders := updateTwoMap(headers, s.Headers)
	resp := POST(url, newHeaders, params, s.Cookies)
	if resp.Err != nil {
		log.Println(resp.Err)
		return "", resp.Err
	} else {
		if !s.DisableAutoUpdateCookie {
			s.Cookies = append(s.Cookies, resp.Cookies...)
			s.Cookies = removeDuplicateCookies(s.Cookies)
		}
	}
	return string(resp.Content), nil
}

// removeDuplicateCookies 移除重复的Cookie
func removeDuplicateCookies(cookies []*http.Cookie) []*http.Cookie {
	temp := make(map[string]*http.Cookie)
	// 倒序遍历
	for i := len(cookies) - 1; i >= 0 ; i-- {
		cookie := cookies[i]
		_, e := temp[cookie.Name]
		if !e {
			// 丢弃已经存在的
			// todo: 时间判断，留下更新的？
			temp[cookie.Name] = cookie
		}
	}

	count := len(temp)
	rs := make([]*http.Cookie, count)
	for _, cookie := range temp {
		rs[count - 1] = cookie
		count --
	}
	return rs
}

// 更新 last -> first
// 两者都有的情况下，last 的会覆盖 first 的
func updateTwoMap(first, last map[string]string) map[string]string {
	var newMap = make(map[string]string)
	for k, v := range first {
		v2, exists := last[k]
		if exists {
			newMap[k] = v2
		} else {
			newMap[k] = v
		}
	}
	for k, v := range last {
		_, exists := first[k]
		if !exists {
			newMap[k] = v
		}
	}
	return newMap
}
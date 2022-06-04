package network

import (
	"errors"
	"github.com/yunluoxin/basic"
	"net/http"
	"strings"
)

// ParseCookiesFromString 从字符串中解析出所有的Cookie
func ParseCookiesFromString(str string) ([]*http.Cookie, error) {
	if basic.IsBlank(str) {
		return nil, errors.New("传入的string为空的")
	}
	components := strings.Split(str, "; ")
	if len(components) < 2 {
		return nil, errors.New("传入的string有问题")
	}
	cs := make([]*http.Cookie, len(components))
	count := 0
	for i, component := range components {
		tmp := strings.Split(component, "=")
		if len(tmp) == 2 {
			count ++
			cs[i] = &http.Cookie{
				Name:  tmp[0],
				Value: tmp[1],
			}
		}
	}
	if count < len(cs) {
		cs = cs[:count]
	}
	return cs, nil
}

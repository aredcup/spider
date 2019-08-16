package types

import "regexp"

var RegexpMap map[string]*regexp.Regexp

func init() {
	RegexpMap = make(map[string]*regexp.Regexp, 0)
	RegexpMap["a[href]"] = regexp.MustCompile(`<a[^>]*href="([^"]+)"[^>]*>`)
}

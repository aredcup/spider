package lagou

import (
	"bufio"
	"fmt"
	"github.com/golang/glog"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"redcup/spider/package/parser"
	"strings"
)

type fetch struct {
	client *http.Client
}

func NewLagouFetch() parser.FetchURL {
	f := &fetch{
		client: parser.NewHTTPClient(),
	}
	return f
}



func (f *fetch) FetchURL(url string) (b []byte, err error) {

	req, err := http.NewRequest("GET",
		fmt.Sprintf("https://gate.lagou.com/v1/entry/positionrec/view/%s?pageno=1&pagesize=5", matches[0][1]),
		strings.NewReader(`{"pageno":1,"pagesize":5}`))

	req.Header.Add("Referer", url)
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("X-L-REQ-HEADER", "{deviceType: 1}")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
	resp, err := f.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		glog.Info(resp.StatusCode)
		return
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := parser.DetermineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader,
		e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

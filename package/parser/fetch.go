package parser

import (
	"bufio"
	"github.com/golang/glog"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var UserAgentList = []string{"Mozilla/5.0 (compatible, MSIE 10.0, Windows NT, DigExt)",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, 360SE)",
	"Mozilla/4.0 (compatible, MSIE 8.0, Windows NT 6.0, Trident/4.0)",
	"Mozilla/5.0 (compatible, MSIE 9.0, Windows NT 6.1, Trident/5.0,",
	"Opera/9.80 (Windows NT 6.1, U, en) Presto/2.8.131 Version/11.11",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, TencentTraveler 4.0)",
	"Mozilla/5.0 (Windows, U, Windows NT 6.1, en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	"Mozilla/5.0 (Macintosh, Intel Mac OS X 10_7_0) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11",
	"Mozilla/5.0 (Macintosh, U, Intel Mac OS X 10_6_8, en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	"Mozilla/5.0 (Linux, U, Android 3.0, en-us, Xoom Build/HRI39) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13",
	"Mozilla/5.0 (iPad, U, CPU OS 4_3_3 like Mac OS X, en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, Trident/4.0, SE 2.X MetaSr 1.0, SE 2.X MetaSr 1.0, .NET CLR 2.0.50727, SE 2.X MetaSr 1.0)",
	"Mozilla/5.0 (iPhone, U, CPU iPhone OS 4_3_3 like Mac OS X, en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
	"MQQBrowser/26 Mozilla/5.0 (Linux, U, Android 2.3.7, zh-cn, MB200 Build/GRJ22, CyanogenMod-7) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1"}

func NewHTTPClient() *http.Client {
	return &http.Client{ //要管理HTTP客户端的头域、重定向策略和其他设置，创建一个Client
		Timeout: 15 * time.Second,
	}
}

type FetchURL interface {
	FetchURL(url string) (b []byte, err error)
}

type FetchBase struct {
	Client *http.Client
}

func (f *FetchBase) FetchURL(url string) (b []byte, err error) {
	req, err := http.NewRequest("GET", url, nil)
	//req.Header.Set("user-agent", UserAgentList[rand.Int31n(13)])

	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
	resp, err := f.Client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		glog.Info(resp.StatusCode)
		return
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := DetermineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader,
		e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func DetermineEncoding(
	r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(
		bytes, "")
	return e
}

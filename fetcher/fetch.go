package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"
)

var ratelimit = time.Tick(1000 * time.Millisecond)
var proxyaddr = "http://47.102.119.88:80/"

// 模拟浏览器发送请求,解决浏览器反爬问题
func Fetch(url string) ([]byte, error) {

	<-ratelimit //等待10ms

	client := &http.Client{
		Timeout: 1000 * time.Second,
	}
	//client := NewHttpClient(proxyaddr)
	req, _ := http.NewRequest("POST", url, nil)
	req.Header.Set("User-Agent", "curl/8.0.1")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("fail to connect webserver")
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := DeterminEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewEncoder())

	return io.ReadAll(utf8Reader)

}

// 无论是什么格式都转换为utf-8
func DeterminEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("fetch error:%s", err)
		return unicode.UTF8
	}
	//截取1024个字节来推断文本格式类型
	e, _, _ := charset.DetermineEncoding(bytes, "")
	//fmt.Println(e)   打印编码
	return e
}

// 设置代理服务器
func NewHttpClient(proxyAddr string) *http.Client {
	proxy, err := url.Parse(proxyAddr)
	if err != nil {
		return nil
	}

	netTransport := &http.Transport{
		//Proxy: http.ProxyFromEnvironment,
		Proxy: http.ProxyURL(proxy),
		Dial: func(netw, addr string) (net.Conn, error) {
			c, err := net.DialTimeout(netw, addr, time.Second*time.Duration(10))
			if err != nil {
				return nil, err
			}
			return c, nil
		},
		MaxIdleConnsPerHost:   10,                             //每个host最大空闲连接
		ResponseHeaderTimeout: time.Second * time.Duration(5), //数据收发5秒超时
	}

	return &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}
}

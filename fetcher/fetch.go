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
	"net/http"
	"time"
)

// 模拟浏览器发送请求,解决浏览器反爬问题
func Fetch(url string) ([]byte, error) {
	client := &http.Client{
		Timeout: 1000 * time.Second,
	}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36")
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
	fmt.Println(e)
	return e
}

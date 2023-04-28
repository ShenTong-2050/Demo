package attack

import "testing"

func BenchmarkSendRequest(b *testing.B) {
	for i:=0;i<b.N;i++ {
		SendRequest("https://www.alsz888.com/api/wap/wapSet")
		SendRequest("https://www.alsz888.com/api/wap/getHomePageAnnouncement")
	}
}

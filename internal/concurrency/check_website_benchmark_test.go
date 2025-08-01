package concurrency

import (
	"testing"
	"time"
)

func slowWebsiteChecker(url string) bool {
	time.Sleep(20*time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	
	for i := range(urls) {
		urls[i] = "a url"
	}

	for b.Loop() {
		CheckWebsites(slowWebsiteChecker ,urls)
	}
}


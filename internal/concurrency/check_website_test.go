package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	return url != "waat://whatever.cod"
}

func TestCheckWebsites(t *testing.T) {
	urls := []string{
		"http://google.com",
		"http://stackoverflow.com",
		"waat://whatever.cod",
	}

	want := map[string]bool{
		"http://google.com":        true,
		"http://stackoverflow.com": true,
		"waat://whatever.cod":      false,
	}
	
	got := CheckWebsites(mockWebsiteChecker, urls)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Got %v want %v", got, want)		
	}

}

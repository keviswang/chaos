package limiter

import (
	"testing"
	"time"
)

func TestHttpRequest(t *testing.T) {
	r := NewRequest(5)
	const max = 20

	for i := 0; i <= max; i++ {
		go r.DoRequest("https://httpbin.org/ip", i)
	}

	time.Sleep(5 * time.Minute)

}

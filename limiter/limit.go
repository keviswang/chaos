package limiter

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type Request struct {
	running  int
	limit    int
	blocking int
	cond     *sync.Cond
	mu       *sync.Mutex
	client   *http.Client
}

func NewRequest(num int) *Request {
	l := new(sync.Mutex)
	return &Request{
		limit: num,
		cond:  sync.NewCond(l),
		mu:    l,
		client: &http.Client{
			Timeout: 6 * time.Second,
		},
	}
}

func (r *Request) Get() {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.running < r.limit {
		r.running++
		fmt.Printf("Now Running: %d\n", r.running)
		return
	}
	r.blocking++
	fmt.Printf("Now Blocking: %d\n", r.blocking)
	for !(r.running < r.limit) {
		r.cond.Wait()
	}
	r.running++
	r.blocking--
	fmt.Printf("Now Get Running: %d\n", r.running)
}

func (r *Request) Release() {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.blocking > 0 {
		r.running--
		r.cond.Signal()
		return
	}
	r.running--
}

func (r *Request) Reset(limit int) {
	r.mu.Lock()
	defer r.mu.Unlock()

	original := r.limit
	r.limit = limit
	blocking := r.blocking
	if limit-original > 0 && blocking > 0 {
		for i := 0; i < limit-original && blocking > 0; i++ {
			r.cond.Signal()
			blocking--
		}
	}
}

func (r *Request) DoRequest(url string, requestID int) {
	r.Get()
	r.doRequest(url, requestID)
	r.Release()
}

func (r *Request) doRequest(url string, requestID int) string {
	resp, _ := r.client.Get(url)
	defer resp.Body.Close()
	res, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("request_id:%d, resp: %v\n", requestID, string(res))
	return string(res)
}

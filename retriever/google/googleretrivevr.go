package google

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	UserAgent string
	timeOut   time.Duration
}

func (r *Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(nil)
	}

	result, err := httputil.DumpResponse(resp, true)

	resp.Body.Close()
	if err != nil {
		panic(err)
	}

	return string(result)
}

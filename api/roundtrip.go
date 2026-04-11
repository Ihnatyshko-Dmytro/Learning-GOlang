package coincap

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type loggingRoundTripper struct {
	logger io.Writer
	next   http.RoundTripper
}

func (l *loggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {

	r.Header.Set("accept", "application/json")
	r.Header.Set("Authorization", "Bearer bfef7cadb8e3a885f95dc4a13ad00d22f7885934b44685ef8e7a31f03bee3553")


	fmt.Fprintf(l.logger, "[%s] %s %s\n", time.Now().Format(time.ANSIC), r.Method, r.URL)
	
	return l.next.RoundTrip(r)
}
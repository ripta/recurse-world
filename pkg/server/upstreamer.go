package server

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Upstreamer chains the request to Host based on the original request path and query string.
type Upstreamer struct {
	Host          string
	OriginalPath  string
	OriginalQuery string
}

// Do performs request upstream.
func (u *Upstreamer) Do(ctx context.Context) (string, error) {
	if err := u.Validate(); err != nil {
		return "", err
	}
	loc := "http://" + u.Host + u.Path()
	if u.OriginalQuery != "" {
		loc += "?" + u.OriginalQuery
	}
	rsp, err := http.Get(loc)
	if err != nil {
		return "", err
	}
	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// Path calculates the upstream request path.
func (u *Upstreamer) Path() string {
	path := u.OriginalPath
	if len(path) < 2 {
		return "/"
	}
	if idx := strings.Index(path[1:], "/"); idx >= 0 {
		return path[idx+1:]
	}
	return "/"
}

// Validate determines whether the upstreamer object is valid or not.
func (u *Upstreamer) Validate() error {
	next := u.Path()
	if len(next) >= len(u.OriginalPath) {
		return fmt.Errorf("BUG: original request path (%q) and upstream path (%q) match, when they should not", u.OriginalPath, next)
	}
	return nil
}

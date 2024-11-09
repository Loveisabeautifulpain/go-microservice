package handlers

import "net/http"

type Hello struct {
}

func (h *Hello) ServeHTTP(rw http.Response, r *http.Request) {

}

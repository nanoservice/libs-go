package web

import (
	"fmt"
	"net/http"
)

type Handler func(*Response, *Request)

type Response struct {
	w *http.ResponseWriter
}

type Request struct {
	r *http.Request
}

func Route(pattern string, handler Handler) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		response := Response{&w}
		request := Request{r}
		handler(&response, &request)
	})
}

func Start() {
	fmt.Printf("Listening on 0.0.0.0%s...\n", port())
	http.ListenAndServe(port(), nil)
}

func (r *Response) Println(text string) {
	fmt.Fprintln(*r.w, text)
}

func port() string {
	return ":8080"
}

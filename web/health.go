package web

var (
	DefaultHealthCheck = Handler(func(w *Response, r *Request) {
		w.Println("OK")
	})
)

func init() {
	Route("/health", healthCheckHandler)
}

func healthCheckHandler(w *Response, r *Request) {
	DefaultHealthCheck(w, r)
}

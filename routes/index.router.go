package routes

import "net/http"

func IndexRootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World Mi adoking!"))
}

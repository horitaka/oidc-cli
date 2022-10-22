package server

import (
	"fmt"
	"net/http"
)

func Token(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r)
}

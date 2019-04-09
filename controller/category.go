package controller

import (
	"fmt"
	"net/http"
)

func CategoryTree(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "hello world")
}
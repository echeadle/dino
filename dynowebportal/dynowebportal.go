package dynowebportal

import (
	"fmt"
	"net/http"
)

//RunWebPortal runs the web server and handles routes
func RunWebPortal(addr string) error {
	http.HandleFunc("/", roothandler)
	return http.ListenAndServe(addr, nil)
}

func roothandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Dino web portal %s", r.RemoteAddr)
}

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fileBytes, _ := ioutil.ReadFile("/proc/cpuinfo")
		html := strings.Replace(fmt.Sprintf("%s", fileBytes), "\n", "<br>", -1)
		fmt.Fprintf(w, "<html><head><title>/proc/cpuinfo demo</title></head><body><h2><b>output from /proc/cpuinfo: </b></h2><br>%s</body></html>", html)
	})

	http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("PORT")), nil)
}

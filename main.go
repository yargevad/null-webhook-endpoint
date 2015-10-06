package main

import (
	"fmt"
	iou "io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		defer r.Body.Close()
		bodyBytes, err := iou.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: failed to read body: %s\n", err)
		} else {
			fmt.Fprintf(os.Stderr, "%s\n", string(bodyBytes))
		}
	})

	log.Fatal(http.ListenAndServe(":3000", nil))
}

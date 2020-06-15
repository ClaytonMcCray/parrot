package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("ERROR!")
		}

		fmt.Println(string(bodyBytes))
	})

	http.ListenAndServe(":80", nil)
}

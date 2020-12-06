// Command http-echo is an HTTP server for examining request headers sent by HTTP clients.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// override the default usage function
	flag.Usage = func() {
		fmt.Println("http-echo is an HTTP server for examining request headers sent by HTTP clients.")
		fmt.Println("\n\tUsage: http-echo [arguments]")
		fmt.Printf("\nThe arguments are:\n\n")
		flag.PrintDefaults()
	}

	// http address argument
	addr := flag.String("http", ":8000", "network address to listen on")

	// parse the arguments
	flag.Parse()

	// define a handler function that repsonds to requests with a
	// JSON object containing the request headers
	handler := func(w http.ResponseWriter, req *http.Request) {

		// get the headers from the request and add a few more bits of useful information
		headers := req.Header.Clone()
		headers["::RequestURI"] = []string{req.RequestURI}
		headers["::Host"] = []string{req.Host}
		headers["::Method"] = []string{req.Method}

		// convert the request headers to a JSON object
		headersJSON, err := json.MarshalIndent(headers, "", "  ")
		if err != nil {
			// set the status response to 500 Internal Server Error
			w.WriteHeader(http.StatusInternalServerError)

			// send the error as the response
			io.WriteString(w, err.Error())

			// log the error
			log.Print(err)
		}

		// set the response content type to JSON
		w.Header().Set("Content-Type", "application/json")

		// send the JSON object as the response
		io.WriteString(w, string(headersJSON))

		// log the JSON object to the console
		log.Println(string(headersJSON))
	}

	// register the hander function
	http.HandleFunc("/", handler)

	// start the http server
	log.Println("http-echo")
	log.Printf("Listening on %s...\n", *addr)
	err := http.ListenAndServe(*addr, nil)
	if err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

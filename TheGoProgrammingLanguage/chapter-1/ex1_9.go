package main

/**
Exercise 1.9: Modify fetch to also print the HTTP status code, found in resp.Status.
 */

import (
	"os"
	"net/http"
	"fmt"
	"io/ioutil"
)

// Fetch prints the content found at a URL.

func main()  {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Fetch: %v\n", err)
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(resp.Body)
		statusCode := resp.StatusCode
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("%s", b)
		fmt.Printf("Status Code: %d", statusCode)

	}
}

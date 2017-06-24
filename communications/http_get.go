package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func main()  {
	r, err := http.Get("http://www.google.com/robots.txt")

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	r.Body.Close()

	if err == nil {
		fmt.Printf("%s", string(b))
	}
}

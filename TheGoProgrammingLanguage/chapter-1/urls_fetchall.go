package main

import (
	"time"
	"os"
	"fmt"
	"net/http"
	"io"
	"io/ioutil"
)

/*
// Fetchall fetches URLs in parallel and reports their times and sizes.
 */

func main()  {

	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch) //start a go co routine.
	}

	for range os.Args[1:] {
		fmt.Println( <- ch) // recieve from channel
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan <- string) {
	start := time.Now()

	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)

	resp.Body.Close() //dont leak resources.

	if err != nil {
		ch <- fmt.Sprintf("While reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()

	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)

}

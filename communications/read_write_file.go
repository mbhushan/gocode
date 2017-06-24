package main

import (
	"os"
	"log"
)

func main()  {
	buf := make([] byte, 1024)

	f,e := os.Open("/etc/passwd")

	if e != nil {
		log.Fatal(e)
	}

	defer f.Close()

	for {
		n, e := f.Read(buf)

		if e != nil {
			log.Fatal(e)
		}

		if n == 0 {
			break
		}

		os.Stdout.Write(buf[:n])

	}
}

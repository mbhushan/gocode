package main

import (
	"os"
	"log"
	"bufio"
)

func main()  {
	buf := make([]byte, 1024)
	f, e := os.Open("/etc/passwd")

	if e != nil {
		log.Fatal(e)
	}

	defer f.Close()

	r := bufio.NewReader(f)
	w := bufio.NewWriter(os.Stdout)

	defer w.Flush()

	for {
		n, e := r.Read(buf)
		if e != nil {
			log.Fatal(e)
		}

		if n == 0 {
			break
		}

		w.Write(buf[0:n])
	}

}

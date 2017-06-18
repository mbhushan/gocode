package main

import (
	"flag"
	"bufio"
	"io"
	"fmt"
	"os"
)

/*
Cat
Write a program which mimics the Unix program cat.
Make it support the -n flag, where each line is numbered.
The solution to the above question given in contains a bug. Can you spot and fix it?
*/

var numberFlag = flag.Bool("n", false, "number each line")

func cat(r *bufio.Reader)  {
	i := 1
	for {
		buf, e := r.ReadBytes('\n')
		if e == io.EOF && string(buf) == "" {
			break
		}
		if *numberFlag {
			fmt.Fprintf(os.Stdout, "%5d %s", i, buf)
			i++
		} else {
			fmt.Fprintf(os.Stdout, "%s", buf)
		}
	}
	return
}

func main()  {
	flag.Parse()
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}

	for i := 0; i < flag.NArg(); i++ {
		f, e := os.Open(flag.Arg(i))
		if e != nil {
			fmt.Fprintf(os.Stderr, "%s: error reading from %s: %s\n",
				os.Args[0], flag.Arg(i), e.Error())
			continue
		}
		cat(bufio.NewReader(f))
	}
}


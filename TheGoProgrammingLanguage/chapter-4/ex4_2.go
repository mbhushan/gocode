package main

import (
	"flag"
	"os"
	"fmt"
	"crypto/sha256"
	"crypto/sha512"
)

/*
Exercise 4.2: Write a program that prints the SHA256 hash of its standard input
by default but supports a command-line flag to print the SHA384 or SHA512 hash instead.
 */

const (
	useSHA256 = iota
	useSHA384
	useSHA512
)

func main() {
	var methodStr string
	
	mySet := flag.NewFlagSet("", flag.ExitOnError)
	mySet.StringVar(&methodStr, "m", "256", "hash func")
	mySet.Parse(os.Args[1:])

	method := useSHA256
	
	if len(os.Args) > 0 {
		switch methodStr {
		case "384":
			method = useSHA384
		case "512":
			method = useSHA512
		default:
			method = useSHA256
		
		}
	}
	
	hashEvaluator(method)
}

func hashEvaluator(method int)  {
	var input string
	fmt.Scan(&input)
	
	if input == "" {
		return
	}
	
	value := [] byte(input)
	
	switch method {
	case useSHA256:
		hashVal := sha256.Sum256(value)
		printHashValue(hashVal[:])
	
	case useSHA384:
		hashVal := sha512.Sum384(value)
		printHashValue(hashVal[:])
	
	case useSHA512:
		hashVal := sha512.Sum512(value)
		printHashValue(hashVal[:])
	}
}

func printHashValue(hashVal []byte) {
	for _, val := range hashVal {
		fmt.Printf("%X", val)
	}
	fmt.Println()
}

/*
Output:
=========
//sha-512
manibhushan
BE76B9FFF166EB741355CD85FD21C659B3652D528B3FD449A5D38BE726C99F5F5C53FE317BDEEB112F866DD2A3D677151664A88455F9A3EB3CC3457154

//sha-384
manibhushan
9FC161D28319D6CFCFEF395A39A5F419519D62F82A79210D3E1376067D7E9D83493B2EBF7C93B88AD399D6FF8

//sha-256
manibhushan
D2F113F04BF16A61823719F9EEC503A98871A1F51D36F402EABCA14044CF26
 */


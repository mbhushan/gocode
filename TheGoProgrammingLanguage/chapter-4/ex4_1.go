package main

/*
Exercise 4.1: Write a function that counts the number of bits that are different in two SHA256 hashes.
*/

var pc [8]byte

func init()  {
	for i:= uint(0); i<8; i++ {
		pc[i] = byte(1 << i)
	}
}

package main

import "os"

func main() {
	y := []byte("y\n")
	for true {
		os.Stdout.Write(y)
	}
}

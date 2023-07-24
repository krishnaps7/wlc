package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func count(r io.Reader, words bool) int {
	fmt.Println("The words is ", words)
	scan := bufio.NewScanner(r)
	count := 0
	fmt.Println("**********Enter Data**********")
	if words {
		scan.Split(bufio.ScanWords)
	}
	for scan.Scan() {
		if scan.Text() == "exit" {
			break
		}
		count++
	}
	// fmt.Println("Number of words is: ", count)
	return count
}

var ch = make(chan string)

func statement(words bool, ch chan string) {
	if words {
		ch <- "The words count is"
	} else {
		ch <- "The lines count is"
	}
}
func main() {
	words := flag.Bool("w", false, "count words")
	flag.Parse()
	go statement(*words, ch)
	fmt.Printf("%v:%v", <-ch, count(os.Stdin, *words))
}

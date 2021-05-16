package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

var lock = sync.Mutex{}
var total int32 = 0

func main() {
	start := time.Now()

	var frequency [26]int32
	for i := 1000; i < 1200; i++ {

		countLetter(fmt.Sprintf("https://www.rfc-editor.org/rfc/rfc%d.txt", i), &frequency)
	}

	fmt.Println("Done.")
	for i, f := range frequency {
		total += f
		fmt.Printf("%s => %d\n", string(allLetters[i]), f)
	}
	fmt.Printf("Total letters: %d\n", total)
	elpased := time.Since(start)
	fmt.Printf("Processing took %s\n", elpased)
}

func countLetter(url string, frequency *[26]int32) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	for _, b := range body {
		c := strings.ToLower(string(b))
		index := strings.Index(allLetters, c)
		if index >= 0 {
			frequency[index] += 1
		}
	}

}

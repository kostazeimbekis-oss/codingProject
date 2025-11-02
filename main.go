package main

import (
	"fmt"
)

func main() {
	fmt.Println(countkmers("ACGAGGTACGA", 3))
}




type Kmers struct{
	sequence string
	count int
}





func countkmers(sequence string, k int) map[string]int {
	var kmers = make(map[string]int)

	for i := 0; i < len(sequence) - k + 1; i++ {
			// accumulate k-mers
	}

	return kmers
}


package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}
	if scanner.Err() != nil {
		log.Printf("Error occured while readng from the input: %v", scanner.Err())
	}

}

func checkDomain() {

}

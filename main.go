package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
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

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecord, err := net.LookupMX(domain)
	//In Go, the LookupMX function is part of the net package, and it is used to perform a DNS MX (Mail Exchange)
	//record lookup for a given domain. MX records are DNS records that specify the mail servers responsible for
	//receiving email on behalf of a domain.
	if err != nil {
		log.Printf("error while fetching MX Record: %v\n", err)
	}
	if len(mxRecord) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)
	//The LookupTXT function in the net package is a part of the Go programming language standard library.
	//It allows you to perform a DNS TXT record lookup for a given domain.
	//A DNS TXT record is a type of resource record in the DNS that is used to associate text information with a domain.
	//This information is typically used for various purposes, such as verifying domain ownership,
	//specifying security policies (like SPF and DMARC), or providing other descriptive information.
	if err != nil {
		log.Printf("error while fetching TXT record: %v\n", err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("error while checking DMARC record: %v", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}
	fmt.Printf(" >>%v\n >>%v\n >>%v\n >>%v\n >>%v\n >>%v \n", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}

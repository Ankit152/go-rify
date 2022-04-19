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
	fmt.Printf("domain, hasMX, hasSPX, sprRecord, hasDMARC, dmarcRecord\n")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Could not read from the input!!! Error: %v!!\n", err)
	}

}

func checkDomain(domain string) {
	var (
		hasMX    bool
		hasDMARC bool
		hasSPF   bool
	)
	var (
		spfRecord   string
		dmarcRecord string
	)

	mxRecord, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	if len(mxRecord) > 0 {
		hasMX = true
	}

	txtRecord, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range txtRecord {
		if strings.HasPrefix(record, domain) {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarc := "_dmarc." + domain
	dmarcRecords, err := net.LookupTXT(dmarc)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("%v, %v, %v, %v, %v\n", hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	if domain == "q" {
		log.Print("Thank You for using this app!")
		os.Exit(0)
	}

	mxRecord, err := net.LookupMX(domain)
	hasMX = false
	if err != nil {
		log.Printf("Error: %v", err)
	} else if len(mxRecord) > 0 {
		hasMX = true
	}

	hasSPF = false
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error: %v", err)
	}
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	hasDMARC = false
	txtDmarc, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error: %v", err)
	}
	for _, record := range txtDmarc {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("hasMx: %v \nhasSPF: %v \nhasDMARC: %v \nspfRecord: %v \ndmarcRecors: %v\n", hasMX, hasSPF, hasDMARC, spfRecord, dmarcRecord)
	fmt.Printf("\nEnter another domain('q' to exit): ")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord Checker\nPress 'q' to quit\n\n")
	fmt.Printf("Enter the domain: ")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(fmt.Sprintf("Error: couldn't read from input: %v\n", err))
	}
}

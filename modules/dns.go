package modules

import (
	"fmt"
	"net"
)

func RunDNS(domain string) {
	fmt.Println("\033[1;34m[*] DNS Lookup:\033[0m")
	aRecords, _ := net.LookupHost(domain)
	fmt.Println("\033[32mA Records:\033[0m", aRecords)

	cname, _ := net.LookupCNAME(domain)
	fmt.Println("\033[32mCNAME:\033[0m", cname)

	mxRecords, _ := net.LookupMX(domain)
	for _, mx := range mxRecords {
		fmt.Printf("\033[32mMX:\033[0m %s %d\n", mx.Host, mx.Pref)
	}

	nsRecords, _ := net.LookupNS(domain)
	for _, ns := range nsRecords {
		fmt.Println("\033[32mNS:\033[0m", ns.Host)
	}

	txtRecords, _ := net.LookupTXT(domain)
	for _, txt := range txtRecords {
		fmt.Println("\033[32mTXT:\033[0m", txt)
	}
	fmt.Println()
}

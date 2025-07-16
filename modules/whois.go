package modules

import (
	"fmt"
	"github.com/likexian/whois"
)

func RunWhois(domain string) {
	fmt.Println("\033[1;34m[*] WHOIS Lookup:\033[0m")
	result, err := whois.Whois(domain)
	if err != nil {
		fmt.Println("\033[31mError:\033[0m", err)
		return
	}
	fmt.Println(result[:300], "...")
	fmt.Println()
}

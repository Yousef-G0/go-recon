package modules

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func RunSubdomains(domain string) {
	fmt.Println("\033[1;34m[*] Subdomain Enumeration (crt.sh):\033[0m")
	url := fmt.Sprintf("https://crt.sh/?q=%%25.%s&output=json", domain)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("\033[31mError:\033[0m", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var results []map[string]interface{}
	err = json.Unmarshal(body, &results)
	if err != nil {
		fmt.Println("\033[31mJSON parse error\033[0m")
		return
	}

	subdomains := make(map[string]bool)
	for _, entry := range results {
		sub := entry["name_value"].(string)
		subdomains[sub] = true
	}

	for sub := range subdomains {
		fmt.Println("\033[32m+\033[0m", sub)
	}
	fmt.Println()
}

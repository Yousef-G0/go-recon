package main

import (
	"fmt"
	"os"
	"github.com/Yousef-G0/go-recon/modules"
)

const banner = "\033[1;36m\n" +
"   ____       ____                          ____                 \n" +
"  / ___| ___ |  _ \\ ___  ___ ___  _ __ ___ |  _ \\ ___  ___  ___  \n" +
" | |  _ / _ \\| |_) / _ \\/ __/ _ \\| '_ ` _ \\| |_) / _ \\/ __|/ _ \\\n" +
" | |_| | (_) |  _ <  __/ (_| (_) | | | | | |  _ <  __/\\__ \\  __/ \n" +
"  \\____|\\___/|_| \\_\\___|\\___\\___/|_| |_| |_|_| \\_\\___||___/\\___| \n" +
"\033[0m"


func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go-recon <target>")
		return
	}

	target := os.Args[1]
	fmt.Println(banner)
	fmt.Printf("\033[1;33m[+] Starting recon on:\033[0m %s\n\n", target)

	modules.RunDNS(target)
	modules.RunPortScan(target)
	modules.RunWhois(target)
	modules.RunSubdomains(target)
}

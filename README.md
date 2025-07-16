
# go-recon

**go-recon** is a lightweight, terminal-based reconnaissance tool written in Go, designed for cybersecurity professionals and students. It runs on Kali Linux (or any OS with Go installed) and provides key information about a target domain or IP.

---

##  Features

-  **DNS Lookup** — A, CNAME, MX, NS, TXT records
-  **Subdomain Discovery** — Passive search using crt.sh
-  **WHOIS Lookup** — Basic domain registration info
-  **Port Scan** — Checks top common TCP ports
-  **Colorful CLI Output** — Easy to read, formatted results

---

## Installation

```bash
git clone https://github.com/Yousef-G0/go-recon.git
cd go-recon
go mod tidy
```

---

##  Usage

```bash
go run main.go <target>
```

Example:
```bash
go run main.go google.com
```

---

##  Project Structure

```
go-recon/
├── main.go               # CLI entrypoint
├── modules/              # Recon modules
│   ├── dns.go
│   ├── ports.go
│   ├── whois.go
│   └── subdomains.go
├── go.mod
├── go.sum
└── README.md             # This file
```

---


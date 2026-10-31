// Command euroton-bridge translates the Proton/Euroton API to IMAP + SMTP
// so that unmodified IMAP-capable mail clients can talk to a Euroton
// backend. See ../../README.md for the operations map and phase plan.
//
// Phase 0 (current): module + package skeleton boots, prints version, exits.
// Phase 1+ adds the Euroton API client, IMAP gateway, SMTP submission.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// Version is overwritten at build time via -ldflags "-X main.Version=…".
var Version = "0.0.0-dev"

func main() {
	imapAddr := flag.String("imap-addr", "127.0.0.1:1143", "IMAP gateway listen address")
	smtpAddr := flag.String("smtp-addr", "127.0.0.1:1025", "SMTP submission listen address")
	eurotonURL := flag.String("euroton-url", "http://127.0.0.1:4000", "euroton-server base URL")
	version := flag.Bool("version", false, "print version and exit")
	flag.Parse()

	if *version {
		fmt.Printf("euroton-bridge %s\n", Version)
		os.Exit(0)
	}

	log.Printf("euroton-bridge %s — scaffolding phase", Version)
	log.Printf("config: imap=%s smtp=%s euroton=%s", *imapAddr, *smtpAddr, *eurotonURL)
	log.Printf("no gateway loops implemented yet; exiting cleanly")
}

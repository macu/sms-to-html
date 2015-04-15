package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
)

var (
	outPath = flag.String("out", "messages.html", "output file path")
)

func main() {

	flag.Parse()

	// Check if output file already exists.
	outInfo, err := os.Stat(*outPath)
	if !os.IsNotExist(err) {
		// Must confirm overwrite.
		if outInfo.IsDir() {
			log.Fatal("cannot overwrite directory")
		}
		fmt.Printf("File already exists at: %s\n...Overwrite? [y/N] ", *outPath)
		var response string
		fmt.Scanln(&response)
		var yesPattern = regexp.MustCompile(`^[yY](es)?`)
		if !yesPattern.MatchString(response) {
			os.Exit(0)
		}
	}

	// Open input file.
	xmlPath := flag.Arg(0)
	xmlFile, err := os.Open(xmlPath)
	if err != nil {
		log.Fatal(err)
	}
	defer xmlFile.Close()

	// Parse input file.
	msgs := readMessages(xmlFile)
	log.Printf("Loaded %d SMS and %d MMS messages.\n", len(msgs.SMSList), len(msgs.MMSList))

	// Open output file.
	outFile, err := os.OpenFile(*outPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	// Write output file.
	writeMessages(msgs, outFile)
	log.Printf("Finished writing messages to file: %s\n", *outPath)

}

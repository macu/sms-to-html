package main

import (
	"encoding/xml"
	"io"
	"log"
)

// Messages represents a collection of messages.
type Messages struct {
	XMLName xml.Name `xml:"smses"`
	Count   int      `xml:"count,attr"`
	SMSList []SMS    `xml:"sms"`
	MMSList []MMS    `xml:"mms"`
}

// SMS represents a single message.
type SMS struct {
	XMLName      xml.Name `xml:"sms"`
	Date         uint64   `xml:"date,attr"`
	ReadableDate string   `xml:"readable_date,attr"`
	Address      string   `xml:"address,attr"`
	Type         int      `xml:"type,attr"`
	Body         string   `xml:"body,attr"`
	ContactName  string   `xml:"contact_name,attr"`
}

// IsIncoming returns whether the SMS represents a received message.
// Otherwise it was a sent message.
func (sms *SMS) IsIncoming() bool {
	// Type 1 is received
	// Type 2 is sent
	return sms.Type == 1
}

// MMS represents a multimedia message.
type MMS struct {
	XMLName      xml.Name  `xml:"mms"`
	Date         uint64    `xml:"date,attr"`
	ReadableDate string    `xml:"readable_date,attr"`
	Address      string    `xml:"address,attr"`
	ContactName  string    `xml:"contact_name,attr"`
	Parts        []MMSPart `xml:"parts>part"`
}

// IsIncoming returns whether this MMS represents a received message.
// Otherwise it was a sent message.
func (mms *MMS) IsIncoming() bool {
	// TODO
	return false
}

// MMSPart represents part of a multimedia message.
type MMSPart struct {
	XMLName     xml.Name `xml:"part"`
	ContentType string   `xml:"ct,attr"`
	Text        string   `xml:"text,attr"`
	Data        string   `xml:"data,attr"`
}

func readMessages(input io.Reader) *Messages {
	decoder := xml.NewDecoder(input)
	var messages Messages
	err := decoder.Decode(&messages)
	if err != nil {
		log.Fatal(err)
	}
	return &messages
}

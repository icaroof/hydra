package main

import (
	"encoding/xml"
	"log"
	"os"
)

//CrewMember struct
type CrewMember struct {
	XMLName           xml.Name `xml:"member"`
	ID                int      `xml:"id,omitempty"`
	Name              string   `xml:"name,attr"`
	SecurityClearance int      `xml:"clearance,attr"`
	AccessCodes       []string `xml:"codes>code"`
}

//ShipInfo struct
type ShipInfo struct {
	XMLName   xml.Name `xml:"ship"`
	ShipID    int      `xml:"ShipInfo>ShipID"`
	ShipClass string   `xml:"ShipInfo>ShipClass"`
	Captain   CrewMember
}

func main() {
	serializeStruct()
}

func serializeStruct() {
	file, err := os.Create("xmlfile.xml")
	if err != nil {
		log.Fatal("Could not create file", err)
	}
	defer file.Close()

	cm := CrewMember{Name: "Jaro", SecurityClearance: 10, AccessCodes: []string{"ADA", "LAL"}}
	si := ShipInfo{ShipID: 1, ShipClass: "Fighter", Captain: cm}

	enc := xml.NewEncoder(file)
	enc.Indent(" ", "  ")
	err = enc.Encode(si)

	if err != nil {
		log.Fatal("Could not encode xml file", err)
	}
}

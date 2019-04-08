package main

import (
	"encoding/json"
	"log"
	"os"
)

//CrewMember struct
type CrewMember struct {
	ID                int      `json:"id,omitempty"`
	Name              string   `json:"name"`
	SecurityClearance int      `json:"clearancelevel"`
	AccessCodes       []string `json:"accesscodes"`
}

//ShipInfo struct
type ShipInfo struct {
	ShipID    int
	ShipClass string
	Captain   CrewMember
}

func main() {
	SerializeStruct()
}

//SerializeStruct serializes a struct type into a json file
func SerializeStruct() {
	f, err := os.Create("jfile.json")
	PrintFatalError(err)
	defer f.Close()

	cm := CrewMember{Name: "Jaro", SecurityClearance: 10, AccessCodes: []string{"ADA", "LAL"}}
	si := ShipInfo{1, "Fighter", cm}

	err = json.NewEncoder(f).Encode(&si)
	PrintFatalError(err)
}

//PrintFatalError prints error if exists
func PrintFatalError(err error) {
	if err != nil {
		log.Fatal("Error happened while processing file", err)
	}
}

package main

import (
	"fmt"

	"github.com/icaroof/hydra/hydraconfigurator"
)

//ConfStruct default struct to file conversion
type ConfStruct struct {
	TS      string  `name:"testString" xml:"testString" json:"testString"`
	TB      bool    `name:"testBool" xml:"testBool" json:"testBool"`
	TF      float64 `name:"testFloat" xml:"testFloat" json:"testFloat"`
	TestInt int
}

func main() {
	testConfStructFromCustomFile()
	testConfStructFromJSONFile()
	testConfStructFromXMLFile()
}

func testConfStructFromCustomFile() {
	fmt.Println("Testing from Custom file")
	configStruct := new(ConfStruct)
	hydraconfigurator.GetConfiguration(hydraconfigurator.CUSTOM, configStruct, "configfile.conf")

	printStructData(configStruct)
}

func testConfStructFromJSONFile() {
	fmt.Println("Testing from JSON file")
	configStruct := new(ConfStruct)
	hydraconfigurator.GetConfiguration(hydraconfigurator.JSON, configStruct, "configfile.json")

	printStructData(configStruct)
}

func testConfStructFromXMLFile() {
	fmt.Println("Testing from XML file")
	configStruct := new(ConfStruct)
	hydraconfigurator.GetConfiguration(hydraconfigurator.XML, configStruct, "configfile.xml")

	printStructData(configStruct)
}

func printStructData(configStruct *ConfStruct) {
	fmt.Println(*configStruct)

	fmt.Println("bool is", configStruct.TB)

	fmt.Println(float64(4.8 * configStruct.TF))

	fmt.Println(5 * configStruct.TestInt)

	fmt.Println(configStruct.TS)
	fmt.Println("#####################################")
}

package main

import (
	"fmt"
	"hydra/hydraconfigurator"
)

type ConfS struct {
	TS      string  `name:"testString"`
	TB      bool    `name:"testBool"`
	TF      float64 `name:"testFloat"`
	TestInt int
}

func main() {
	configStruct := new(ConfS)
	hydraconfigurator.GetConfiguration(hydraconfigurator.CUSTOM, configStruct, "configfile.conf")
	fmt.Println(*configStruct)

	fmt.Println("bool is", configStruct.TB)

	fmt.Println(float64(4.8 * configStruct.TF))

	fmt.Println(5 * configStruct.TestInt)

	fmt.Println(configStruct.TS)
}

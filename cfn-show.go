package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type CloudFormationTemplate struct {
	AWSTemplateFormatVersion string
	Description              string
	Mappings                 map[string]interface{}
	Outputs                  map[string]interface{}
	Parameters               map[string]interface{}
	Resources                map[string]interface{}
}

func main() {

	if len(os.Args) != 3 {
		fmt.Println("You must provide a template file and an item name")
		os.Exit(1)
	}

	file, e := ioutil.ReadFile(os.Args[1])
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	var template CloudFormationTemplate
	json.Unmarshal(file, &template)

	itemPath := strings.Split(os.Args[2], "/")
	
	var itemRoot map[string]interface{}
	
	switch itemPath[0] {
	case "Mappings":
		itemRoot = template.Mappings
	case "Outputs":
		itemRoot = template.Outputs
	case "Parameters":
		itemRoot = template.Parameters
	case "Resources":
		itemRoot = template.Resources
	default:
		fmt.Printf("Invalid item: %s\n", os.Args[2])
	}
	
	var outputMap map[string]interface{}
	outputMap = make(map[string]interface{})
	outputMap[itemPath[1]] = itemRoot[itemPath[1]]

	output, _ := json.MarshalIndent(outputMap, "  ", "  ")

	fmt.Printf("%s\n", output)

}

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

type ScanConfig struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name string `yaml:"name"`
	} `yaml:"metadata"`
	Spec struct {
		ScanType   string   `yaml:"scanType"`
		Parameters []string `yaml:"parameters"`
	} `yaml:"spec"`
}

func main() {

	// WELCOME Banner
	fmt.Println(".............. WELCOME TO secureCodeBox CLI ..............")

	// Print Scan Options
	fmt.Println("Select Scan Type:")
	fmt.Println("1- Scanning network")
	fmt.Println("2- Repeat Scans on a Schedule")
	fmt.Println("3- Scanning Web Application")
	fmt.Println("4- Post-processing with Hooks")
	fmt.Println("5- Storing Findings with Persistence Hooks")
	fmt.Println("6- Automatically Scan your Cluster with Autodiscovery")
	fmt.Println("7- Enforce Engagement")

	reader := bufio.NewReader(os.Stdin)

	//Get Scan Option input from User
	fmt.Print("Enter your choice: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	//Choose Input Type
	scanType := strings.TrimSpace(input)

	//nmap Scanning and yaml file creation
	if scanType == "1" || scanType == "Scanning network" {
		var targetURL string
		fmt.Print("Enter target URL: ")
		targetURL, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading target URL:", err)
			os.Exit(1)
		}
		targetURL = strings.TrimSpace(targetURL)

		config := ScanConfig{
			APIVersion: "execution.securecodebox.io/v1",
			Kind:       "Scan",
			Metadata: struct {
				Name string `yaml:"name"`
			}{Name: fmt.Sprintf("nmap-%s", targetURL)},
			Spec: struct {
				ScanType   string   `yaml:"scanType"`
				Parameters []string `yaml:"parameters"`
			}{
				ScanType: scanType,
				Parameters: []string{
					targetURL,
				},
			},
		}

		output, err := yaml.Marshal(&config)
		if err != nil {
			fmt.Println("Error marshalling YAML:", err)
			os.Exit(1)
		}

		// Create the output_yaml directory if it doesnot already exist
		if _, err := os.Stat("output_yaml"); os.IsNotExist(err) {
			err := os.Mkdir("output_yaml", 0755)
			if err != nil {
				fmt.Println("Error creating output directory:", err)
				os.Exit(1)
			}
		}

		fileName := fmt.Sprintf("output_yaml/output_%s.yaml", time.Now().Format("2024-03-31-01-01-05"))
		err = ioutil.WriteFile(fileName, output, 0644)
		if err != nil {
			fmt.Println("Error writing YAML file:", err)
			os.Exit(1)
		}

		fmt.Println("YAML file generated successfully.")
	} else {

		//Other Scan options not created yet
		fmt.Println("THIS PART IS NOT CREATED YET!")
	}
}

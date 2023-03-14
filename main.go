package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const apiKey = "<API_KEY>"

type Organization struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Network struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Device struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Model  string `json:"model"`
	Mac    string `json:"mac"`
	Serial string `json:"serial"`
	IP     string `json:"lanIp"`
}

var org_ID string
var net_ID string

func main() {
	client := &http.Client{}

	url1 := "https://api.meraki.com/api/v0/organizations"

	req1, err := http.NewRequest("GET", url1, nil)
	if err != nil {
		panic(err)
	}

	req1.Header.Set("X-Cisco-Meraki-API-Key", apiKey)

	resp1, err := client.Do(req1)
	if err != nil {
		panic(err)
	}
	defer resp1.Body.Close()

	var orgs []Organization
	if err := json.NewDecoder(resp1.Body).Decode(&orgs); err != nil {
		panic(err)
	}

	fmt.Printf("\nFound %d organizations:\n", len(orgs))
	for _, org := range orgs {
		fmt.Printf("Orgaznization Name: %s       ID: %s\n", org.Name, org.ID)
	}

	fmt.Println("Enter ID of the organization you want to work with: ")
	fmt.Scan(&org_ID)

	url := fmt.Sprintf("https://api.meraki.com/api/v0/organizations/%s/networks", org_ID)
	req2, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req2.Header.Set("X-Cisco-Meraki-API-Key", apiKey)

	resp2, err := client.Do(req2)
	if err != nil {
		panic(err)
	}
	defer resp2.Body.Close()

	var networks []Network
	if err := json.NewDecoder(resp2.Body).Decode(&networks); err != nil {
		panic(err)
	}

	fmt.Printf("\nNetworks in the organization:\n")
	for _, network := range networks {
		fmt.Printf("ID: %s, Name: %s\n", network.ID, network.Name)
	}

	fmt.Println("\nEnter ID of network you want to work with: ")
	fmt.Scan(&net_ID)

	url3 := fmt.Sprintf("https://api.meraki.com/api/v0/networks/%s/devices", net_ID)
	req3, err := http.NewRequest("GET", url3, nil)
	if err != nil {
		panic(err)
	}

	req3.Header.Set("X-Cisco-Meraki-API-Key", apiKey)

	resp3, err := client.Do(req3)
	if err != nil {
		panic(err)
	}
	defer resp3.Body.Close()

	var devices []Device
	if err := json.NewDecoder(resp3.Body).Decode(&devices); err != nil {
		panic(err)
	}

	fmt.Printf("Found %d devices in network %s:\n", len(devices), net_ID)
	for _, device := range devices {
		fmt.Printf("ID: %s, Name: %s, Model: %s, MAC: %s, Serial: %s, IP: %s\n", device.ID, device.Name, device.Model, device.Mac, device.Serial, device.IP)
	}
}

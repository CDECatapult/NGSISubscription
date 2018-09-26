package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type config struct {
	Orion             string `json:"orion_url"`
	Cygnus            string `json:"cygnus_url"`
	FiwareService     string `json:"fiware_service"`
	FiwareServicePath string `json:"fiware_servicepath"`
}

type orionEntity struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	DateCreated string `json:"dateCreated"`
}

type entity struct {
	EntityID   string
	EntityType string
}

type subData struct {
	Config config
	Entity entity
}

// SubscriptionAPI : API endpoint
const SubscriptionAPI = "/v2/subscriptions"

// OrionAPI : Orion API get entities endpoint
const OrionAPI = "/v2/entities/?attrs=dateCreated&limit=500&type="

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No filename given")
		os.Exit(0)
	}

	fileName := os.Args[1]
	printBanner()

	// Loading configuration
	dat, err := ioutil.ReadFile("config.json")
	check(err)
	var c config
	err = json.Unmarshal(dat, &c)

	// Loading data sources file
	dat, err = ioutil.ReadFile(fileName)
	check(err)

	// Parsing data
	ds := parseData(dat)
	//check(err)

	// Showing data sources to import
	fmt.Println()
	fmt.Println("The following Types will be subscribed:")
	for _, el := range ds {
		fmt.Println(el)
	}
	// Ask for confirmation
	fmt.Println("WARNING: Are you sure? (yes/no)")
	if !askForConfirmation() {
		os.Exit(0)
	}

	// init
	aTmpl := template.Must(template.ParseFiles("templates/subscription"))
	var b bytes.Buffer

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	//client := &http.Client{}
	var arr []orionEntity
	var entities []entity
	for _, el := range ds {
		//pd := subData{c, el}
		//fmt.Println()
		//fmt.Println("Response for:", el)

		b.Reset()
		//aTmpl.Execute(&b, pd)
		en := sendGet(client, c, c.Orion+OrionAPI+el, b.Bytes())
		//fmt.Println(string(en))
		err = json.Unmarshal(en, &arr)
		//fmt.Println(arr)
		for _, item := range arr {
			ent := entity{item.ID, item.Type}
			entities = append(entities, ent)
		}
	}

	fmt.Println(len(entities))

	for _, el := range entities {
		pd := subData{c, el}
		fmt.Println()
		fmt.Println("Response for:", el.EntityID)

		b.Reset()
		aTmpl.Execute(&b, pd)
		//fmt.Println(string(b.Bytes()))
		fmt.Println("Subscription - ", sendPost(client, c, c.Orion+SubscriptionAPI, b.Bytes()))
	}
}

func sendPost(c *http.Client, conf config, url string, b []byte) string {
	req, err := http.NewRequest("POST", url, bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Fiware-Service", conf.FiwareService)
	req.Header.Add("Fiware-ServicePath", conf.FiwareServicePath)
	// Save a copy of this request for debugging.
	// requestDump, err := httputil.DumpRequest(req, true)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(requestDump))
	resp, err := c.Do(req)
	check(err)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return (resp.Status + " " + string(body))
}

func sendGet(c *http.Client, conf config, url string, b []byte) []byte {
	req, err := http.NewRequest("GET", url, bytes.NewReader(b))
	//req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Fiware-Service", conf.FiwareService)
	req.Header.Add("Fiware-ServicePath", conf.FiwareServicePath)
	// Save a copy of this request for debugging.
	// requestDump, err := httputil.DumpRequest(req, true)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(requestDump))
	resp, err := c.Do(req)
	check(err)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return (body)
}

func parseData(d []byte) []string {
	ds := []string{}

	lines := strings.Split(string(d), "\n")
	for _, line := range lines {
		if (line != "") && (line != "\n") {
			element := string(line)
			ds = append(ds, element)
		}
	}
	return ds
}

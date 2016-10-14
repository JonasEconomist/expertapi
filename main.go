package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/EconomistDigitalSolutions/expertapi/api"
)

var mode string
var nid string

func init() {
	flag.StringVar(&mode, "mode", "full", "analysis required")
	flag.StringVar(&nid, "nid", "", "node ID of content file")
}

func main() {
	flag.Parse()
	file, err := ioutil.ReadFile("samples/" + nid + ".sample")
	if err != nil {
		log.Fatal(err)
	}
	resp, err := api.RequestAnalysis(string(file), mode)
	if err != nil {
		log.Fatal(err)
	}
	writeResponseToFile(resp)
}

func writeResponseToFile(resp string) {
	t := time.Now()
	filename := fmt.Sprintf("output/%s_%s.xml", mode, t.Format("20060102150405"))
	target, _ := os.Create(filename)
	defer target.Close()
	_, err := target.WriteString(resp)
	if err != nil {
		log.Fatal(err)
	}
}

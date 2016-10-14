package api

import (
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	EXPERT_ENDPOINT_FULL              = "http://cogito-api.cloudapp.net:8091/cogito/v1/annotation/annotate/economist-plan.xml"
	EXPERT_ENDPOINT_CATEGORIZATION    = "http://cogito-api.cloudapp.net:8091/cogito/v1/annotation/annotate/economist-mediatopic.xml"
	EXPERT_ENDPOINT_EXTRACTION_MINING = "http://cogito-api.cloudapp.net:8091/cogito/v1/annotation/annotate/economist-cogito-synthesis.xml"
)

func RequestAnalysis(file, mode string) (output string, err error) {
	var endpoint string
	switch mode {
	case "categorization":
		// This endpoint produces the categorization (Media Topic taxonomy) results only.
		endpoint = EXPERT_ENDPOINT_CATEGORIZATION
	case "extraction_default_mining":
		// This endpoint produces results from:
		// 1. Extraction (relevant terms and concepts).
		// 2. Default mining (entities and tags), main sentences and topics.
		endpoint = EXPERT_ENDPOINT_EXTRACTION_MINING
	default:
		// This endpoint produces the full semantic analysis results:
		// 1. Categorization (Media Topic taxonomy).
		// 2. Extraction (relevant terms and concepts).
		// 3. Default mining (entities and tags), main sentences and topics.
		endpoint = EXPERT_ENDPOINT_FULL
	}
	client := &http.Client{}
	req, _ := http.NewRequest("POST", endpoint, strings.NewReader(file))

	req.Header.Set("Content-Type", "text/html;charset=ISO-8859-1")

	resp, err := client.Do(req)
	if err != nil {
		return output, err
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return string(b), err
}

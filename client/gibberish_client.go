package client

import (
	"encoding/json"
	"github.com/shitpostingio/analysis-commons/structs"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// PerformRequest performs a request to the gibberish service.
func PerformRequest(input, endpoint string) (data structs.GibberishResponse, err error) {

	client := http.Client{Timeout: time.Second * 30}
	request, err := http.NewRequest("POST", endpoint, strings.NewReader(input))
	if err != nil {
		log.Println("GibberishClient.PerformRequest: unable to create request:", err)
		return
	}

	response, err := client.Do(request)
	if err != nil {
		log.Println("GibberishClient.PerformRequest: unable to perform request:", err)
		return
	}

	defer func() {
		err := response.Body.Close()
		if err != nil {
			log.Println("GibberishClient.PerformRequest: unable to close request body:", err)
		}
	}()

	bodyResult, err := ioutil.ReadAll(response.Body)
	log.Debugln("GibberishClient.PerformRequest: body result: ", string(bodyResult))
	if err != nil {
		log.Println("GibberishClient.PerformRequest: unable to read request body:", err)
		return
	}

	err = json.Unmarshal(bodyResult, &data)
	if err != nil {
		log.Println("PerformRequest: error while unmarshaling ", err)
	}

	return

}

package handle

import (
	"bytes"
	"cf-pagerduty-service/pagerdutyapi/config"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// Payload body for the PagerDuty Events API requests
type Payload struct {
	EventType   string `json:"event_type"`
	Description string `json:"description"`
	ServiceKey  string `json:"service_key"`
}

// Headers headers for the PagerDuty Events API requests
type Headers struct {
	ContentType string
}

// Data is the expected structure of incoming requests
type Data struct {
	ServiceKey string `json:"service_key"`
}

// Trigger handle PagerDuty trigger events
func Trigger(w http.ResponseWriter, r *http.Request) {
	var serviceKey string
	var description string

	client := &http.Client{}
	apiConfigPath := configPath()

	parsedConfig, err := config.ParseConfig(apiConfigPath)
	if err != nil {
		log.Fatal("Error parsing config file\n")
		log.Fatal(err)
	}

	body, err := parseRequestBody(r)
	if err != nil {
		log.Println(err)
		// If error parsing body, use default values from env/config
		serviceKey = getEnv("SERVICE_KEY", parsedConfig.APIConfiguration.ServiceKey)
	} else {
		serviceKey = body.ServiceKey
	}

	// TODO: Allow users to send descriptions
	description = getEnv("DESCRIPTION", parsedConfig.APIConfiguration.Description)

	// TODO: Switch to POST /incidents when ready
	url := "https://events.pagerduty.com/generic/2010-04-15/create_event.json"
	data := &Payload{"trigger", description, serviceKey}
	payload, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(payload))
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)

	log.Println(res.StatusCode)
	log.Println("Successfully triggered PagerDuty incident")
}

func getEnv(env string, defaultValue string) string {
	var v string
	if v = os.Getenv(env); len(v) == 0 {
		log.Printf("Using default values for %v", env)
		return defaultValue
	}

	return env
}

func configPath() string {
	defaultConfigYamlPath := "config.yml"

	apiConfigYamlPath := os.Getenv("API_CONFIG_PATH")
	if apiConfigYamlPath == "" {
		log.Printf("API_CONFIG_PATH not set, using '%v'", defaultConfigYamlPath)
		return defaultConfigYamlPath
	}
	return apiConfigYamlPath
}

func parseRequestBody(req *http.Request) (Data, error) {
	decoder := json.NewDecoder(req.Body)
	var data Data
	err := decoder.Decode(&data)
	if err != nil {
		log.Println("Error parsing request body")
		return Data{}, err
	}

	defer req.Body.Close()
	return data, nil
}

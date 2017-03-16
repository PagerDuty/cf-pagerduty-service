package integrations

import (
	"os"

	"github.com/cloudfoundry-incubator/candiedyaml"
)

// Integrations is a list of integrations we have with Pivotal partners
type Integrations struct {
	IntegrationsInfo string        `yaml:"info"`
	Integrations     []Integration `yaml:"integrations"`
}

// Integration is a single integration's relevant data
type Integration struct {
	Name                string `yaml:"name"`
	IntegrationGuideURL string `yaml:"integration_guide_url"`
}

// ParseIntegrations is a function to parse the integrations yaml
func ParseIntegrations(path string) (Integrations, error) {
	file, err := os.Open(path)
	if err != nil {
		return Integrations{}, err
	}

	var decodedIntegrations Integrations
	if err := candiedyaml.NewDecoder(file).Decode(&decodedIntegrations); err != nil {
		return Integrations{}, err
	}
	return decodedIntegrations, nil
}

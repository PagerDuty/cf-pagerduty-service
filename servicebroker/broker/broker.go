package broker

import (
	"context"

	"cf-pagerduty-service/servicebroker/config"
	"cf-pagerduty-service/servicebroker/integrations"

	"github.com/pivotal-cf/brokerapi"
)

// PagerDutyBroker broker configuration
type PagerDutyBroker struct {
	Config       config.Config
	Integrations integrations.Integrations
}

// Services services broker generates
func (pagerDutyBroker *PagerDutyBroker) Services(context context.Context) []brokerapi.Service {
	return []brokerapi.Service{
		brokerapi.Service{
			ID:            "3ba74f7d-d755-42ab-bf7c-9d43dc1381c8",
			Name:          "p-pagerduty",
			Description:   "Trigger PagerDuty incidents from your Cloud Foundry applications",
			Bindable:      true,
			Tags:          []string{"pagerduty", "pd", "incident-management"},
			PlanUpdatable: false,
			Requires:      nil,
			Plans: []brokerapi.ServicePlan{
				brokerapi.ServicePlan{
					ID:          "58391dd3-3d79-447e-8343-1e06794b05b0",
					Name:        "standard",
					Description: "Integrate CF with your PagerDuty account",
					Metadata: &brokerapi.ServicePlanMetadata{
						DisplayName: "Standard",
						Bullets:     []string{"PagerDuty"},
					},
				},
			},
			Metadata: &brokerapi.ServiceMetadata{
				DisplayName:         "PagerDuty",
				SupportUrl:          "https://support.pagerduty.com/",
				DocumentationUrl:    "https://www.pagerduty.com/docs/guides/cloud-foundry-integration-guide/",
				ProviderDisplayName: "PagerDuty",
				LongDescription:     "Trigger PagerDuty incidents from your Cloud Foundry applications",
			},
		},
	}
}

// Provision provision services
func (pagerDutyBroker *PagerDutyBroker) Provision(context context.Context, instanceID string, details brokerapi.ProvisionDetails, asyncAllowed bool) (brokerapi.ProvisionedServiceSpec, error) {
	return brokerapi.ProvisionedServiceSpec{}, nil
}

// Deprovision deprovision services
func (pagerDutyBroker *PagerDutyBroker) Deprovision(context context.Context, instanceID string, details brokerapi.DeprovisionDetails, asyncAllowed bool) (brokerapi.DeprovisionServiceSpec, error) {
	return brokerapi.DeprovisionServiceSpec{}, nil
}

// Bind bind services to apps
func (pagerDutyBroker *PagerDutyBroker) Bind(context context.Context, instanceID string, bindingID string, details brokerapi.BindDetails) (brokerapi.Binding, error) {
	return brokerapi.Binding{
			Credentials: pagerDutyBroker},
		nil
}

// Unbind unbind services from apps
func (pagerDutyBroker *PagerDutyBroker) Unbind(context context.Context, instanceID string, bindingID string, details brokerapi.UnbindDetails) error {
	return nil
}

// LastOperation the last operation to watch for when polling API
func (pagerDutyBroker *PagerDutyBroker) LastOperation(context context.Context, instanceID string, operationData string) (brokerapi.LastOperation, error) {
	return brokerapi.LastOperation{}, nil
}

// Update update services
func (pagerDutyBroker *PagerDutyBroker) Update(context context.Context, instanceID string, details brokerapi.UpdateDetails, asyncAllowed bool) (brokerapi.UpdateServiceSpec, error) {
	return brokerapi.UpdateServiceSpec{}, nil
}

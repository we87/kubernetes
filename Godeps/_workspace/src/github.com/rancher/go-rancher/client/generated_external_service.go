package client

const (
	EXTERNAL_SERVICE_TYPE = "externalService"
)

type ExternalService struct {
	Resource

	AccountId string `json:"accountId,omitempty" yaml:"account_id,omitempty"`

	Created string `json:"created,omitempty" yaml:"created,omitempty"`

	Data map[string]interface{} `json:"data,omitempty" yaml:"data,omitempty"`

	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	EnvironmentId string `json:"environmentId,omitempty" yaml:"environment_id,omitempty"`

	ExternalIpAddresses []string `json:"externalIpAddresses,omitempty" yaml:"external_ip_addresses,omitempty"`

	HealthCheck *InstanceHealthCheck `json:"healthCheck,omitempty" yaml:"health_check,omitempty"`

	Hostname string `json:"hostname,omitempty" yaml:"hostname,omitempty"`

	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	RemoveTime string `json:"removeTime,omitempty" yaml:"remove_time,omitempty"`

	Removed string `json:"removed,omitempty" yaml:"removed,omitempty"`

	State string `json:"state,omitempty" yaml:"state,omitempty"`

	Transitioning string `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`

	TransitioningMessage string `json:"transitioningMessage,omitempty" yaml:"transitioning_message,omitempty"`

	TransitioningProgress int64 `json:"transitioningProgress,omitempty" yaml:"transitioning_progress,omitempty"`

	Upgrade ServiceUpgrade `json:"upgrade,omitempty" yaml:"upgrade,omitempty"`

	Uuid string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type ExternalServiceCollection struct {
	Collection
	Data []ExternalService `json:"data,omitempty"`
}

type ExternalServiceClient struct {
	rancherClient *RancherClient
}

type ExternalServiceOperations interface {
	List(opts *ListOpts) (*ExternalServiceCollection, error)
	Create(opts *ExternalService) (*ExternalService, error)
	Update(existing *ExternalService, updates interface{}) (*ExternalService, error)
	ById(id string) (*ExternalService, error)
	Delete(container *ExternalService) error

	ActionActivate(*ExternalService) (*Service, error)

	ActionAddservicelink(*ExternalService, *AddRemoveServiceLinkInput) (*Service, error)

	ActionCancelupgrade(*ExternalService) (*Service, error)

	ActionCreate(*ExternalService) (*Service, error)

	ActionDeactivate(*ExternalService) (*Service, error)

	ActionRemove(*ExternalService) (*Service, error)

	ActionRemoveservicelink(*ExternalService, *AddRemoveServiceLinkInput) (*Service, error)

	ActionSetservicelinks(*ExternalService, *SetServiceLinksInput) (*Service, error)

	ActionUpdate(*ExternalService) (*Service, error)

	ActionUpgrade(*ExternalService, *ServiceUpgrade) (*Service, error)
}

func newExternalServiceClient(rancherClient *RancherClient) *ExternalServiceClient {
	return &ExternalServiceClient{
		rancherClient: rancherClient,
	}
}

func (c *ExternalServiceClient) Create(container *ExternalService) (*ExternalService, error) {
	resp := &ExternalService{}
	err := c.rancherClient.doCreate(EXTERNAL_SERVICE_TYPE, container, resp)
	return resp, err
}

func (c *ExternalServiceClient) Update(existing *ExternalService, updates interface{}) (*ExternalService, error) {
	resp := &ExternalService{}
	err := c.rancherClient.doUpdate(EXTERNAL_SERVICE_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ExternalServiceClient) List(opts *ListOpts) (*ExternalServiceCollection, error) {
	resp := &ExternalServiceCollection{}
	err := c.rancherClient.doList(EXTERNAL_SERVICE_TYPE, opts, resp)
	return resp, err
}

func (c *ExternalServiceClient) ById(id string) (*ExternalService, error) {
	resp := &ExternalService{}
	err := c.rancherClient.doById(EXTERNAL_SERVICE_TYPE, id, resp)
	return resp, err
}

func (c *ExternalServiceClient) Delete(container *ExternalService) error {
	return c.rancherClient.doResourceDelete(EXTERNAL_SERVICE_TYPE, &container.Resource)
}

func (c *ExternalServiceClient) ActionActivate(resource *ExternalService) (*Service, error) {

	resp := &Service{}

	err := c.rancherClient.doAction(EXTERNAL_SERVICE_TYPE, "activate", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ExternalServiceClient) ActionAddservicelink(resource *ExternalService, input *AddRemoveServiceLinkInput) (*Service, error) {

	resp := &Service{}

	err := c.rancherClient.doAction(EXTERNAL_SERVICE_TYPE, "addservicelink", &resource.Resource, input, resp)

	return resp, err
}

func (c *ExternalServiceClient) ActionCancelupgrade(resource *ExternalService) (*Service, error) {

	resp := &Service{}

	err := c.rancherClient.doAction(EXTERNAL_SERVICE_TYPE, "cancelupgrade", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ExternalServiceClient) ActionCreate(resource *ExternalService) (*Service, error) {

	resp := &Service{}

	err := c.rancherClient.doAction(EXTERNAL_SERVICE_TYPE, "create", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ExternalServiceClient) ActionDeactivate(resource *ExternalService) (*Service, error) {

	resp := &Service{}

	err := c.rancherClient.doAction(EXTERNAL_SERVICE_TYPE, "deactivate", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ExternalServiceClient) ActionRemove(resource *ExternalService) (*Service, error) {

	resp := &Service{}

	err := c.rancherClient.doAction(EXTERNAL_SERVICE_TYPE, "remove", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ExternalServiceClient) ActionRemoveservicelink(resource *ExternalService, input *AddRemoveServiceLinkInput) (*Service, error) {

	resp := &Service{}

	err := c.rancherClient.doAction(EXTERNAL_SERVICE_TYPE, "removeservicelink", &resource.Resource, input, resp)

	return resp, err
}

func (c *ExternalServiceClient) ActionSetservicelinks(resource *ExternalService, input *SetServiceLinksInput) (*Service, error) {

	resp := &Service{}

	err := c.rancherClient.doAction(EXTERNAL_SERVICE_TYPE, "setservicelinks", &resource.Resource, input, resp)

	return resp, err
}

func (c *ExternalServiceClient) ActionUpdate(resource *ExternalService) (*Service, error) {

	resp := &Service{}

	err := c.rancherClient.doAction(EXTERNAL_SERVICE_TYPE, "update", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ExternalServiceClient) ActionUpgrade(resource *ExternalService, input *ServiceUpgrade) (*Service, error) {

	resp := &Service{}

	err := c.rancherClient.doAction(EXTERNAL_SERVICE_TYPE, "upgrade", &resource.Resource, input, resp)

	return resp, err
}

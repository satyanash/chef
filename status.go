package chef

type StatusService struct {
	client *Client
}

// Status represents the body of the returned information.
type Status struct {
	Status    string            `json:"status"`
	Upstreams map[string]string `json:"upstreams"`
	Keygen    map[string]int    `json:"keygen"`
}

// Status gets the frontend & backend server information.
//
// https://docs.chef.io/api_chef_server/#_status
func (e *StatusService) Get() (data Status, err error) {
	err = e.client.magicRequestDecoder("GET", "_status", UseOrg, nil, &data)
	return
}

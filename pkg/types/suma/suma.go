package suma

type Options struct {
	SumaServer					 string   `json:"suma-server,omitempty" yaml:"suma-server,omitempty"`
	SumaUser	                 string   `json:"suma-user,omitempty" yaml:"suma-user,omitempty"`
	SumaPassword		         string   `json:"suma-password,omitempty" yaml:"suma-password,omitempty"`
	Group                        string   `json:"group,omitempty" yaml:"group,omitempty"`
	MasterNodes                  string   `json:"master-nodes,omitempty" yaml:"master-nodes,omitempty"`
	WorkerNodes                  string   `json:"worker-nodes,omitempty" yaml:"worker-nodes,omitempty"`
	BurnThemAll                  bool     `json:"burn-them-all" yaml:"burn-them-all"`
}

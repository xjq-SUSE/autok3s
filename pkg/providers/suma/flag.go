package suma

import (
	"github.com/cnrancher/autok3s/pkg/types"
)

const createUsageExample = `  autok3s -d create \
    --provider suma \
    --suma-server <suma-server> \
    --suma-user <suma-user> \
    --suma-password <suma-password> \
	--group <group> \
	--ssh-key-path <ssh-key-path> \
	--master-nodes <master-nodes> \
	--burn-them-all <burn-them-all>
`

const joinUsageExample = `  autok3s -d join \
	--provider suma \
	--suma-server <suma-server> \
	--suma-user <suma-user> \
	--suma-password <suma-password> \
	--group <group> \
	--ssh-key-path <ssh-key-path> \
	--master-node <master-node> \
	--worker-nodes <worker-nodes>
`

func (p *Suma) GetUsageExample(action string) string {
	switch action {
	case "create":
		return createUsageExample
	case "join":
		return joinUsageExample
	default:
		return "not support"
	}
}

func (p *Suma) GetCreateFlags() []types.Flag {
	cSSH := p.GetSSHConfig()
	p.SSH = *cSSH
	fs := p.GetClusterOptions()
	fs = append(fs, p.GetCreateOptions()...)
	return fs
}

func (p *Suma) GetOptionFlags() []types.Flag {
	return p.sharedFlags()
}

func (p *Suma) GetJoinFlags() []types.Flag {
	fs := p.sharedFlags()
	fs = append(fs, p.GetClusterOptions()...)
	fs = append(fs, types.Flag{
		Name:  "ip",
		P:     &p.IP,
		V:     p.IP,
		Usage: "IP for an existing k3s server",
	})
	return fs
}

func (p *Suma) GetSSHFlags() []types.Flag {
	return []types.Flag{}
}

func (p *Suma) GetDeleteFlags() []types.Flag {
	return []types.Flag{}
}

func (p *Suma) GetCredentialFlags() []types.Flag {
	return []types.Flag{}
}

func (p *Suma) GetSSHConfig() *types.SSH {
	ssh := &types.SSH{
		SSHUser:    defaultUser,
		SSHPort:    "22",
		SSHKeyPath: defaultSSHKeyPath,
	}
	return ssh
}

func (p *Suma) BindCredential() error {
	return nil
}

func (p *Suma) MergeClusterOptions() error {
	return nil
}

func (p *Suma) sharedFlags() []types.Flag {
	fs := []types.Flag{
		{
			Name:  "suma-server",
			P:     &p.SumaServer,
			V:     p.SumaServer,
			Usage: "IP address of SUMA server",
		},
		{
			Name:  "suma-user",
			P:     &p.SumaUser,
			V:     p.SumaUser,
			Usage: "Username of SUMA server",
		},
		{
			Name:  "suma-password",
			P:     &p.SumaPassword,
			V:     p.SumaPassword,
			Usage: "Password of SUMA server",
		},
		{
			Name:  "group",
			P:     &p.Group,
			V:     p.Group,
			Usage: "SUMA group to create K3S cluster",
		},
		{
			Name:  "master-nodes",
			P:     &p.MasterNodes,
			V:     p.MasterNodes,
			Usage: "Hostnames of nodes to install K3S master,separated by commas",
		},
		{
			Name:  "worker-nodes",
			P:     &p.WorkerNodes,
			V:     p.WorkerNodes,
			Usage: "Hostnames of nodes to install K3S worker,separated by commas",
		},
		{
			Name:  "worker-ips",
			P:     &p.WorkerIps,
			V:     p.WorkerIps,
			Usage: "Public IPs of worker nodes on which to install agent, multiple IPs are separated by commas",
		},
		{
			Name:  "burn-them-all",
			P:     &p.BurnThemAll,
			V:     p.BurnThemAll,
			Usage: "Flag to use all the nodes in SUMA group",
		},
	}

	return fs
}

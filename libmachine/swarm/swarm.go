package swarm

const (
	DockerImage              = "dockerswarm/swarm:master"
	DiscoveryServiceEndpoint = "https://discovery-stage.hub.docker.com/v1"
)

type SwarmOptions struct {
	IsSwarm    bool
	Address    string
	Discovery  string
	Master     bool
	Host       string
	Strategy   string
	Heartbeat  int
	Overcommit float64
	TlsCaCert  string
	TlsCert    string
	TlsKey     string
	TlsVerify  bool
}

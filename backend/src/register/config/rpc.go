package config

var (
	RPCAddress string
	RPCMethod  string
)

func init() {
	RPCMethod = "tcp"
	RPCAddress = "127.0.0.1:8972"
}

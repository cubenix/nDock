package types

// Config represents the Json object in config.json
type Config struct {
	Hosts []Host
}

// // GetHosts returns Docker hosts read from the configuration
// func (c *Config) GetHosts() []Host {
// 	return c.hosts
// }

// // IntializeHosts returns Docker hosts read from the configuration
// func (c *Config) IntializeHosts(hosts []Host) {
// 	c.hosts = hosts
// }

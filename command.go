package dcli

type DiscoveryNode interface {
	Run([]string)
	Help()
	Name() string
	Description() string
	Usage() string
}

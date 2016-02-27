package model

import (
	"net"
	"os"
	"strings"
)

var containerInfo *ContainerInfo

// ContainerInfo holds information about the host executing the web server
type ContainerInfo struct {
	Hostname   string
	EnvVars    []ContainerEnv
	Interfaces []ContainerInterface
}

// ContainerEnv holds information about an environment variable
type ContainerEnv struct {
	Name      string
	Value     string
	IsService bool
}

// ContainerInterface holds container net interfaces info
type ContainerInterface struct {
	Name      string
	Addresses []InterfaceAddresses
}

// InterfaceAddresses holds addresses assigned to an interface
type InterfaceAddresses struct {
	IP      string
	Mask    string
	CIDR    string
	Network string
}

// GetContainerInfo return container information struct
func GetContainerInfo() (*ContainerInfo, error) {

	if containerInfo == nil {

		hostName, err := os.Hostname()
		if err != nil {
			return nil, err
		}

		containerInterfaces, err := getContainerInterfaces()
		if err != nil {
			return nil, err
		}

		containerEnvironment := getContainerEnvironment()

		containerInfo = &ContainerInfo{
			Hostname:   hostName,
			Interfaces: containerInterfaces,
			EnvVars:    containerEnvironment,
		}
	}

	return containerInfo, nil
}

// getContainerInterfaces return container interfaces struct
func getContainerInterfaces() ([]ContainerInterface, error) {

	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	containerInterfaces := make([]ContainerInterface, len(interfaces), len(interfaces))
	for i := range interfaces {
		containerInterfaces[i].Name = interfaces[i].Name
		addrs, err := interfaces[i].Addrs()
		if err != nil {
			return nil, err
		}

		interfaceAddresses := make([]InterfaceAddresses, len(addrs), len(addrs))
		for j := range addrs {
			switch addr := addrs[j].(type) {
			case *net.IPAddr:
				interfaceAddresses[j].Network = addr.Network()
				interfaceAddresses[j].IP = addr.IP.String()
				interfaceAddresses[j].Mask = addr.IP.DefaultMask().String()
				interfaceAddresses[j].CIDR = addr.String()
			case *net.IPNet:
				interfaceAddresses[j].Network = addr.Network()
				interfaceAddresses[j].IP = addr.IP.String()
				interfaceAddresses[j].Mask = addr.IP.DefaultMask().String()
				interfaceAddresses[j].CIDR = addr.String()
			default:
				interfaceAddresses[j].Network = addr.Network()
				interfaceAddresses[j].CIDR = addr.String()
			}
		}
		containerInterfaces[i].Addresses = interfaceAddresses
	}

	return containerInterfaces, nil
}

// getContainerEnvironment return container environment variables struct
func getContainerEnvironment() []ContainerEnv {
	envVars := os.Environ()
	containerEnvs := make([]ContainerEnv, len(envVars), len(envVars))

	for i := range envVars {
		env := strings.Split(envVars[i], "=")
		containerEnvs[i].Name = env[0]
		containerEnvs[i].Value = env[1]
	}

	return containerEnvs
}

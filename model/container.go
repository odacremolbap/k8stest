package model

import (
	"net"
	"os"
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
	Name string
}

// hostName, err := os.Hostname()
// if err != nil {
// 	return "", err
// }

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

		containerInfo = &ContainerInfo{
			Hostname:   hostName,
			Interfaces: containerInterfaces,
		}
	}

	return containerInfo, nil

	// for _, iface := range interfaces {
	// 	addrs, err := iface.Addrs()
	// 	if err != nil {
	// 		return "", err
	// 	}
	// 	for i := range addrs {
	// 		hostInfo = fmt.Sprintf("%s\n%s-address-%d: %s", hostInfo, iface.Name, i, addrs[i].String())
	// 	}
	// }
	//
	// return hostInfo, nil
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
	}

	return containerInterfaces, nil

	// for _, iface := range interfaces {
	// 	addrs, err := iface.Addrs()
	// 	if err != nil {
	// 		return "", err
	// 	}
	// 	for i := range addrs {
	// 		hostInfo = fmt.Sprintf("%s\n%s-address-%d: %s", hostInfo, iface.Name, i, addrs[i].String())
	// 	}
	// }
	//
	// return hostInfo, nil
}

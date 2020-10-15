package vpn

import (
	"encoding/base64"
	_ "encoding/base64"
	"errors"
	"fmt"
	"os"
)

type vpnConfiguration struct {
	pemPassword string
	username    string
	password    string
	email       string
}

type VpnConnectInterface interface {
	GetListOfEndpoints() ([]string, error)
	ConnectToEndpoint(matching string)
}

type VpnConnection struct {
	url             string
	listOfEndpoints []string
	configuration   *vpnConfiguration
}

func NewVpnConnection() *VpnConnection {
	fmt.Println("Starting factory method NewVpnConnection")
	newVpn := new(VpnConnection)
	newVpn.url = "google.com"
	newVpn.listOfEndpoints = make([]string, 0)
	newVpn.setupConfigurations()
	return newVpn
}

func (v *VpnConnection) setupConfigurations() {
	v.getConfigurationFromEnvironmentVaribales()
}

func (v *VpnConnection) getConfigurationFromEnvironmentVaribales() {
	v.configuration = new(vpnConfiguration)
	var err error
	v.configuration.pemPassword, err = UnDecodeBase64(GetEnvironmentVariable("PEM_PASSWORD"))
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	v.configuration.username, err = UnDecodeBase64(GetEnvironmentVariable("USERNAME"))
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	v.configuration.password, err = UnDecodeBase64(GetEnvironmentVariable("PASSWORD"))
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	v.configuration.email, err = UnDecodeBase64(GetEnvironmentVariable("EMAIL"))
	if err != nil {
		fmt.Print(err.Error())
		return
	}

}

func (v *VpnConnection) GetListOfEndpoints() ([]string, error) {
	fmt.Println("hhm")
	return nil, nil
}

func (v *VpnConnection) ConnectToEndpoint(matching string) {

}

func UnDecodeBase64(str string, prevErr error) (string, error) {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New(err.Error() + "\n" + prevErr.Error())
	} else {
		str = string(data)
		//fmt.Println(str)
		return str, nil
	}
}

func GetEnvironmentVariable(environmentName string) (string, error) {
	str := os.Getenv(environmentName)
	if len(str) == 0 {
		fmt.Printf("Error: Environment variable %s was not found \n", environmentName)
		fmt.Println(str)
		return "", errors.New("Environment variable " + environmentName + " was not found")
	} else {
		fmt.Printf("Environment variable %s equals to %s \n", environmentName, str)
		return str, nil
	}

}

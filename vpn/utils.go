package vpn

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type AnyConnect struct {
	XMLName    xml.Name     `xml:"AnyConnectProfile"`
	ServerList []ServerList `xml:"ServerList"`
}

type ServerList struct {
	XMLName xml.Name `xml:"ServerList"`
	Hosts   []Host   `xml:"HostEntry"`
}

type Host struct {
	XMLName     xml.Name `xml:"HostEntry"`
	HostName    string   `xml:"HostName"`
	HostAddress string   `xml:"HostAddress"`
}

const SERVER_FILE_LOCATION = "/opt/cisco/anyconnect/profile/anyconnect-cert-latestv1.0-linux.xml"

func FetchMapOfServers() (*map[string]string, error) {
	mapOfServers, err := convertToMap(openXmlFile())
	return mapOfServers, err
}

func openXmlFile() *ServerList {
	xmlFile, err := os.Open(SERVER_FILE_LOCATION)
	if err != nil {
		fmt.Printf("Could not open server file : %s \n", SERVER_FILE_LOCATION)
		return nil
	}

	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)

	hosts := AnyConnect{}
	err = xml.Unmarshal(byteValue, &hosts)

	if err != nil {
		fmt.Printf(err.Error())
		return nil
	}

	return &(hosts.ServerList[0])
}

func convertToMap(serverList *ServerList) (*map[string]string, error) {

	if serverList == nil {
		return nil, errors.New("Opening config file and trying to unmarshall it failed")
	}

	mapToReturn := make(map[string]string)

	for _, host := range serverList.Hosts {
		//fmt.Println(host.HostName + " : " + host.HostAddress)
		mapToReturn[host.HostName] = host.HostAddress
	}

	return &mapToReturn, nil
}

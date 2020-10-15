package main

import (
	"fmt"

	"github.com/salmin36/go-vpn-client/vpn"
)

func main() {
	fmt.Println("Starting vpn")

	var vpnCon vpn.VpnConnectInterface = vpn.NewVpnConnection()

	_, _ = vpnCon.GetListOfEndpoints()

}

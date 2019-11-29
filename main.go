package main

import "fmt"
import "net"

func GetBroadcastAddress(ip net.IP, mask net.IPMask) net.IP {
	if len(mask) == 4 {
		ip = ip.To4()
	} else {
		ip = ip.To16()
	}
	broadcast := net.IP(make([]byte, len(ip)))
	for i := range ip {
		broadcast[i] = ip[i] | ^mask[i]
	}
	return broadcast
}

func BroadcastQuery() {
	ifaces, err := net.Interfaces()
	if err != nil {
		return
	}
	for _, iface := range ifaces {
		if iface.Flags|net.FlagBroadcast == 0 ||
			iface.Flags|net.FlagUp == 0 {
			continue
		}
		fmt.Printf("if: %s %d\n", iface.Name, iface.Flags);
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			switch ip := addr.(type) {
			case *net.IPNet:
				if !ip.IP.IsLoopback() {
					fmt.Printf("addr: %s\n", addr.String())
					fmt.Printf("ip: %s %t\n", ip.String(), ip.IP.IsLoopback())
					broadcast := GetBroadcastAddress(ip.IP, ip.Mask)
					fmt.Printf("broadcast: %s\n", broadcast)
				}
			}
		}
	}
}

func main() {
	BroadcastQuery()
}

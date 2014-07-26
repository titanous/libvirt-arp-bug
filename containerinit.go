package main

import (
	"log"
	"net"
	"time"

	"github.com/docker/libcontainer/netlink"
	"github.com/j-keck/arping"
)

func main() {
	log.SetFlags(log.Lshortfile)

	iface, err := net.InterfaceByName("eth0")
	if err != nil {
		log.Fatal(err)
	}
	ip, ipNet, _ := net.ParseCIDR("192.168.122.128/24")
	if err := netlink.NetworkLinkAddIp(iface, ip, ipNet); err != nil {
		log.Fatal(err)
	}
	if err := netlink.AddDefaultGw("192.168.122.1", "eth0"); err != nil {
		log.Fatal(err)
	}
	if err := netlink.NetworkLinkUp(iface); err != nil {
		log.Fatal(err)
	}
	start := time.Now()
	arping.SetTimeout(100 * time.Millisecond)
	for {
		addr, t, err := arping.PingOverIface(net.ParseIP("192.168.122.1"), *iface)
		if err != nil {
			log.Println("ARP error:", err)
			continue
		}
		log.Printf("ARP success after %s: %v %v", time.Now().Sub(start), addr, t)
		return
	}
}

package main

import "fmt"
import "./nmapLib"

func main() {
	fmt.Println("scanning...")
	hosts, err := nmap.ScanOpenTcpPorts("192.168.2.0/24", "22,80,443")
	fmt.Println("scanned!")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hosts)
}

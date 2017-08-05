
package nmap

import (
	"os"
	"os/exec"
	"fmt"
	"io"
	"encoding/xml"
	"bytes"
)

func scanOpenTcpPorts(subnet string, ports string) (io.Reader, error) {
	c := exec.Command("nmap", "-oX", "-", "-sS", "-O", subnet, "-p", ports)
	cout, cerr := c.Output()
	if cerr != nil {
		return nil, fmt.Errorf("need root")
	}
	return bytes.NewReader(cout), nil
}

func openXml(filename string) (io.Reader, error) {
	return os.Open(filename)
}

type xnDoc struct {
	Hosts []xnHost `xml:"host"`
}

type xnHost struct {
	Addrs []xnAddr `xml:"address"`
	Os []xnOs `xml:"os"`
}

type xnAddr struct {
	Addr string `xml:"addr,attr"`
	Type string `xml:"addrtype,attr"`
	Os string `xml:"os"`
}

type xnOs struct {
	OsMatch []xnOsMatch `xml:"osmatch"`
}

type xnOsMatch struct {
	Name string `xml:"name,attr"`
	Accuracy string `xml:"accuracy,attr"`
}

func parseXml(reader io.Reader) (hosts []Host, err error) {
	doc := xnDoc{}

	dec := xml.NewDecoder(reader)
	derr := dec.Decode(&doc)
	if derr != nil {
		return hosts, derr
	}

	fmt.Println("!!!!!!!!!!!!!!!!")
	fmt.Println(doc)

	fmt.Println("$$$$$$$$$$$$$")
	fmt.Println(doc.Hosts)

	for _, xh := range doc.Hosts {
		h := Host{}
		for _, xa := range xh.Addrs {
			switch {
			case xa.Type == "ipv4":
				h.Ip = xa.Addr
			case xa.Type == "mac":
				h.Mac = xa.Addr
			}
		}
		/*
		for _, xo := range xh.Os {
			for _, xom := range xo.OsMatch {
				h.Os = xo.Name
			}
		}
		*/

		if h.Ip != "" && h.Mac != "" {
			hosts = append(hosts, h)
		}
	}

	return
}

type Host struct {
	Mac string
	Ip string
	Os string
}

func ScanOpenTcpPorts(subnet, ports string) (hosts []Host, err error) {
	reader, rerr := scanOpenTcpPorts(subnet, ports)
	if rerr != nil {
		return hosts, rerr
	}
	return parseXml(reader)
}

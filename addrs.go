package addrs

import (
	"net"
	"strings"

	"github.com/lovego/slice"
)

func IsLocalhost(addr string) (bool, error) {
	ips, err := net.LookupIP(addr)
	if err != nil {
		return false, err
	}
	for _, ip := range ips {
		if ip.IsLoopback() {
			return true, nil
		}
	}
	ipsOfLocal, err := IPsOfLocalhost()
	if err != nil {
		return false, err
	}
	for _, ip := range ips {
		if slice.ContainsString(ipsOfLocal, ip.String()) {
			return true, nil
		}
	}
	return false, nil
}

var ipsOfLocalhost []string

func IPsOfLocalhost() ([]string, error) {
	if ipsOfLocalhost != nil {
		return ipsOfLocalhost, nil
	}

	ifcAddrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	ipsOfLocalhost := make([]string, len(ifcAddrs))
	for i, ifcAddr := range ifcAddrs {
		addr := ifcAddr.String()
		if i := strings.IndexByte(addr, '/'); i >= 0 {
			addr = addr[:i]
		}
		ipsOfLocalhost[i] = addr
	}
	ipsOfLocalhost = ipsOfLocalhost
	return ipsOfLocalhost, nil
}

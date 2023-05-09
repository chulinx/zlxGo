package net

import (
	"errors"
	kubeNet "k8s.io/apimachinery/pkg/util/net"
	"net"
	"strings"
)

func HostIP() string {
	nodeIP, _ := kubeNet.ChooseHostInterface()
	return nodeIP.String()
}

// LookAddrIPFromDomain example: www.ccc.com/www.ccc.com:9090 return 100.22.14.1/100.22.14.1:9090
func LookAddrIPFromDomain(s string) (string, error) {
	var (
		domain string
		port   string
	)

	sSplit := strings.Split(s, ":")
	switch {
	case len(sSplit) == 0:
		return "", errors.New("domain format is error")
	case len(sSplit) == 1:
		address := net.ParseIP(s)
		if address == nil {
			ns, err := net.LookupHost(s)
			if err != nil || len(ns) < 1 {
				return "", errors.New("look up ip error1")
			}
			return ns[0], nil
		}
		return s, nil
	case len(sSplit) == 2:
		domain, port = sSplit[0], sSplit[1]
		address := net.ParseIP(domain)
		if address == nil {
			ns, err := net.LookupHost(domain)
			if err != nil || len(ns) < 1 {
				return "", errors.New("look up ip error2")
			}
			return strings.Join([]string{ns[0], port}, ":"), nil
		}
		return s, nil
	default:
		return "", errors.New("domain format is error")
	}
}

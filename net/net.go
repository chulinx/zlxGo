package net

import (
	kubeNet "k8s.io/apimachinery/pkg/util/net"
)

func HostIP() string {
	nodeIP, _ := kubeNet.ChooseHostInterface()
	return nodeIP.String()
}

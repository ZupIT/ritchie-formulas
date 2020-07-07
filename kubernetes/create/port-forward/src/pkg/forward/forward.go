package forward

import (
	"kubernetes/core/pkg/kube"
	"log"
)

type Inputs struct {
	Namespace   string
	PodPartName string
	Kubeconfig  string
	PortMap		string
}

func (in Inputs) Run() {
	client, config := kube.LoadClientSet(in.Kubeconfig)
	podSelect := kube.GetPod(client, in.PodPartName, in.Namespace)
	log.Println("Status of Pod : ", podSelect.Status.Phase)
	log.Println("Starting port forwarding.")
	kube.PortForward(in.PortMap, in.Namespace, podSelect.Name, config)
}
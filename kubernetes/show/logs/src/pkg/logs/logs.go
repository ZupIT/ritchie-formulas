package logs

import (
	"bufio"
	"fmt"
	"kubernetes/core/pkg/kube"
	"kubernetes/core/pkg/prompt"
	"log"

	v1 "k8s.io/api/core/v1"
)

type Inputs struct {
	Namespace   string
	PodPartName string
	Kubeconfig  string
}

func (in Inputs) Run() {
	client, _ := kube.LoadClientSet(in.Kubeconfig)
	podSelect := kube.GetPod(client, in.PodPartName, in.Namespace)
	log.Println("Status of Pod : ", podSelect.Status.Phase)

	var containersItems []string
	for _, container := range podSelect.Spec.Containers {
		containersItems = append(containersItems, container.Name)
	}
	containerSelect, _ := prompt.List("Select container: ", containersItems)

	podLogOpts := v1.PodLogOptions{Follow: true, Container: containerSelect}
	req := client.CoreV1().Pods(podSelect.Namespace).GetLogs(podSelect.Name, &podLogOpts)
	podLogs, _ := req.Stream()
	scanner := bufio.NewScanner(podLogs)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}
}

package health

import (
	"fmt"
	"io/ioutil"
	"kubernetes/core/pkg/kube"
	"kubernetes/core/pkg/prompt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Inputs struct {
	Namespace   string
	PodPartName string
	Kubeconfig  string
}

func (in Inputs) Run() {
	client, config := kube.LoadClientSet(in.Kubeconfig)
	podSelect := kube.GetPod(client, in.PodPartName, in.Namespace)
	log.Println("Status of Pod : ", podSelect.Status.Phase)

	var containersItems []string
	for i, container := range podSelect.Spec.Containers {
		containersItems = append(containersItems, fmt.Sprint(i, " - ", container.Name))
	}
	containerSelectString, _ := prompt.List("Select container: ", containersItems)
	ind, _ := strconv.Atoi(strings.Split(containerSelectString, " - ")[0])
	containerSelect := podSelect.Spec.Containers[ind]
	if containerSelect.LivenessProbe == nil || containerSelect.LivenessProbe.Handler.HTTPGet == nil {
		log.Fatal("LivenessProbe not configured to http.")
	}
	pathHealth := containerSelect.LivenessProbe.Handler.HTTPGet.Path
	portHealth := containerSelect.LivenessProbe.Handler.HTTPGet.Port

	log.Println("Starting port forwarding for container.")
	go kube.PortForward(strconv.Itoa(int(portHealth.IntVal)), in.Namespace, podSelect.Name, config)
	time.Sleep(time.Second)
	log.Println("Successful port forwarding for container.")

	log.Println("Call Health check.")
	url := fmt.Sprintf("%s:%d%s", "http://localhost", portHealth.IntVal, pathHealth)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		log.Fatalf("Failed to Health check: %v\n", resp)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Health Check response: ")
	log.Println(string(bodyBytes))
}

package kube

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"kubernetes/core/pkg/prompt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/portforward"
	"k8s.io/client-go/transport/spdy"

    "k8s.io/apimachinery/pkg/api/errors"
    "k8s.io/client-go/util/homedir"
    "k8s.io/client-go/plugin/pkg/client/auth"
)

func LoadClientSet(configBase64 string) (*kubernetes.Clientset, *rest.Config) {
	kubeConfigBytes, _ := base64.StdEncoding.DecodeString(configBase64)
	clientConfig, err := clientcmd.NewClientConfigFromBytes(kubeConfigBytes)
	if err != nil {
		log.Fatalln("Failed to load config. Verify if you set credential kubeconfig in base64 format.")
	}

	config, err := clientConfig.ClientConfig()
	if err != nil {
		log.Fatalln("Failed to load config. Verify if you set credential kubeconfig in base64 format.")
	}

	// create the client
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln("Failed to create kubernetes client.")
	}
	return client, config
}

func GetPod(client *kubernetes.Clientset, filter string, namespace string) v1.Pod {
	pods, err := client.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Failed to list pods for namespace: %s.\n", namespace)
	}
	if len(pods.Items) == 0 {
		log.Fatalf("No result pods to namespace: %s.\n", namespace)
	}
	log.Println("Pod list:")
	for _, pod := range pods.Items {
		fmt.Printf("%s - %s\n", pod.Name, pod.Status.Phase)
	}
	var podsFilter []v1.Pod
	for _, pod := range pods.Items {
		if strings.Contains(pod.Name, filter) {
			podsFilter = append(podsFilter, pod)
		}
	}
	if len(podsFilter) == 0 {
		log.Fatalf("No result pods to name: %s and namespace: %s.\n", filter, namespace)
	}
	var items []string
	for i, pod := range podsFilter {
		items = append(items, fmt.Sprint(i, " - ", pod.Name))
	}
	itemSelect, _ := prompt.List("Select Pod: ", items)
	ind, _ := strconv.Atoi(strings.Split(itemSelect, " - ")[0])
	return podsFilter[ind]
}

func PortForward(port, namespace, podName string, config *rest.Config) {
	ports := []string{fmt.Sprintf("%s:%s", port, port)}
	roundTripper, upgrader, err := spdy.RoundTripperFor(config)
	if err != nil {
		log.Fatal("Failed port forwarding for container.")
	}
	path := fmt.Sprintf("/api/v1/namespaces/%s/pods/%s/portforward", namespace, podName)
	hostIP := strings.TrimLeft(config.Host, "htps:/")
	serverURL := url.URL{Scheme: "https", Path: path, Host: hostIP}
	dialer := spdy.NewDialer(upgrader, &http.Client{Transport: roundTripper}, http.MethodPost, &serverURL)
	stopChan, readyChan := make(chan struct{}, 1), make(chan struct{}, 1)
	out, errOut := new(bytes.Buffer), new(bytes.Buffer)
	forwarder, err := portforward.New(dialer, ports, stopChan, readyChan, out, errOut)
	if err != nil {
		log.Fatal("Failed port forwarding for container.")
	}
	go func() {
		for range readyChan { // Kubernetes will close this channel when it has something to tell us.
		}
		if len(errOut.String()) != 0 {
			panic(errOut.String())
		} else if len(out.String()) != 0 {
			fmt.Println(out.String())
		}
	}()

	if err := forwarder.ForwardPorts(); err != nil { // Locks until stopChan is closed.
		log.Fatal("Failed port forwarding for container.")
	}
}

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// Initialize cmd flags
	kubeconfig := flag.String("kubeconfig", "kubeconfig", "absolute path to the kubeconfig file")
	nsTarget := flag.String("namespace", "workloads", "namespace to use")
	podsToKill := flag.Int("pods-to-kill", 1, "set maximum amount of pods to kill")
	dryRun := flag.Bool("dry-run", true, "run chaosnetes as dry-run.")
	flag.Parse()

	// Check kubeconfig exist
	if _, err := os.Stat(*kubeconfig); err != nil {
		log.Fatalf("error retrieving kubeconfig: %v\n", err)
	}

	// Create clientSet
	clientSet := initializeClient(kubeconfig)

	// Create podList
	podList, err := clientSet.CoreV1().Pods(*nsTarget).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("error retrieving podList: %v\n", err)
	}

	for i := 0; i < *podsToKill; i++ {
		// Pick random pod
		randomPod := podList.Items[rand.Intn(len(podList.Items))]
		fmt.Println("deleting", randomPod.Name)

		if !*dryRun {
			// Delete random pod
			err := clientSet.CoreV1().Pods(*nsTarget).Delete(context.TODO(), randomPod.Name, metav1.DeleteOptions{})
			if err != nil {
				log.Fatalf("error deleting pod: %v\n", err)
			}
		}
	}
}

func initializeClient(kubeconfigPath *string) *kubernetes.Clientset {
	// Create config
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfigPath)
	if err != nil {
		log.Fatalf("error in initializeClient.BuildConfigFromFlags: %v\n", err)
	}

	// Create clientSet
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("error in initializeClient.NewForConfig: %v\n", err)
	}
	return client
}

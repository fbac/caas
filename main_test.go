package main

import (
	"reflect"
	"testing"

	"k8s.io/client-go/kubernetes"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_initializeClient(t *testing.T) {
	type args struct {
		kubeconfigPath *string
	}
	tests := []struct {
		name string
		args args
		want *kubernetes.Clientset
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := initializeClient(tt.args.kubeconfigPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initializeClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

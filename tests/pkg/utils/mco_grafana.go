// Copyright (c) Red Hat, Inc.
// Copyright Contributors to the Open Cluster Management project
// Licensed under the Apache License 2.0

package utils

import (
	"os"
	"strings"

	"gopkg.in/yaml.v2"
	"k8s.io/klog"
)

var (

	BearerToken             string
	optionsFile             string

	testOptions          TestOptions
	testOptionsContainer TestOptionsContainer

)

func GetGrafanaURL(opt TestOptions) string {
	if optionsFile == "" {
		optionsFile = os.Getenv("OPTIONS")
		if optionsFile == "" {
			optionsFile = "resources/options.yaml"
		}
	}

	data, err := os.ReadFile(optionsFile)
	if err != nil {
		klog.Errorf("--options error: %v", err)
	}

	err = yaml.Unmarshal([]byte(data), &testOptionsContainer)
	if err != nil {
		klog.Errorf("--options error: %v", err)
	}

	testOptions = testOptionsContainer.Options

	cloudProvider := strings.ToLower(os.Getenv("CLOUD_PROVIDER"))
	substring1 := "rosa"
	substring2 := "hcp"
	if strings.Contains(cloudProvider, substring1) && strings.Contains(cloudProvider, substring2) {

		grafanaConsoleURL := "https://grafana-open-cluster-management-observability.apps.rosa." + opt.HubCluster.BaseDomain
		if opt.HubCluster.GrafanaURL != "" {
			grafanaConsoleURL = opt.HubCluster.GrafanaURL
		} else {
			opt.HubCluster.GrafanaHost = "grafana-open-cluster-management-observability.apps.rosa." + opt.HubCluster.BaseDomain
		}
		return grafanaConsoleURL
	} else {
		grafanaConsoleURL := "https://grafana-open-cluster-management-observability.apps." + opt.HubCluster.BaseDomain
		if opt.HubCluster.GrafanaURL != "" {
			grafanaConsoleURL = opt.HubCluster.GrafanaURL
		} else {
			opt.HubCluster.GrafanaHost = "grafana-open-cluster-management-observability.apps." + opt.HubCluster.BaseDomain
		}
		return grafanaConsoleURL
	}
}

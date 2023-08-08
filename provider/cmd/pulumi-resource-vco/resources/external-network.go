package resources

import (
	"bytes"
	"encoding/json"
	"fmt"
	p "github.com/pulumi/pulumi-go-provider"
	"net/http"
)

type ExternalNetwork struct{}

type ExternalNetworkState struct {
	ExternalNetworkArgs
	ExternalNetworkID string `pulumi:"external_network_id"`
	Metric            int    `pulumi:"metric"`
	ExternalNetworkIP string `pulumi:"external_network_ip"`
	Status            bool   `json:"success" pulumi:"success"`
}

type ExternalNetworkArgs struct {
	URL                 string `pulumi:"url"`
	Token               string `pulumi:"token"`
	CustomerID          string `pulumi:"customerID"`
	CloudSpaceID        string `pulumi:"cloudspace_id"`
	ExternalNetworkID   string `pulumi:"external_network_id"`
	ExternalNetworkType string `pulumi:"external_network_type"`
	Metric              int    `pulumi:"metric"`
	ExternalNetworkIP   string `pulumi:"external_network_ip"`
}

func (ExternalNetwork) Create(ctx p.Context, name string, input ExternalNetworkArgs, preview bool) (string, ExternalNetworkState, error) {
	state := ExternalNetworkState{ExternalNetworkArgs: input}
	state.ExternalNetworkID = input.ExternalNetworkID
	state.Metric = input.Metric
	state.ExternalNetworkIP = input.ExternalNetworkIP

	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/external-networks?external_network_id=%s&external_network_type=%s&metric=%d&external_network_ip=%s", input.URL, input.CustomerID, input.CloudSpaceID, input.ExternalNetworkID, input.ExternalNetworkType, input.Metric, input.ExternalNetworkIP)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(nil))
	if err != nil {
		return "", state, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", input.Token))

	resp, err := client.Do(req)
	if err != nil {
		return "", state, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", state, err
	}

	state.Status = result["success"].(bool)

	return name, state, nil
}

func (ex ExternalNetwork) Update(ctx p.Context, state ExternalNetworkState, input ExternalNetworkArgs) (bool, ExternalNetworkState, error) {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/external-networks?external_network_id=%s&external_network_type=%s&metric=%d&external_network_ip=%s", input.URL, input.CustomerID, input.CloudSpaceID, input.ExternalNetworkID, input.ExternalNetworkType, input.Metric, input.ExternalNetworkIP)

	client := &http.Client{}
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(nil))
	if err != nil {
		return false, state, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", input.Token))

	resp, err := client.Do(req)
	if err != nil {
		return false, state, err
	}
	defer resp.Body.Close()

	var status map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&status); err != nil {
		return false, state, err
	}

	success := status["success"].(bool)

	state.ExternalNetworkID = input.ExternalNetworkID
	state.Metric = input.Metric
	state.ExternalNetworkIP = input.ExternalNetworkIP

	return success, state, nil
}

func (ExternalNetwork) Delete(ctx p.Context, input ExternalNetworkArgs) (bool, error) {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/external-networks?external_network_id=%s&external_network_ip=%s", input.URL, input.CustomerID, input.CloudSpaceID, input.ExternalNetworkID, input.ExternalNetworkIP)

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(nil))
	if err != nil {
		return false, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", input.Token))

	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	status := result["success"].(bool)

	return status, nil
}

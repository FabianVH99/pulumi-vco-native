package cloudspace

import (
	"bytes"
	"fmt"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"io/ioutil"
	"net/http"
)

type ExternalNetwork struct{}

type ExternalNetworkState struct {
	ExternalNetworkArgs
	URL               string `pulumi:"url"`
	Token             string `pulumi:"token"`
	CustomerID        string `pulumi:"customerID"`
	CloudSpaceID      string `pulumi:"cloudspace_id"`
	ExternalNetworkID string `pulumi:"external_network_id"`
	Metric            int    `pulumi:"metric"`
	ExternalNetworkIP string `pulumi:"external_network_ip"`
}

type ExternalNetworkArgs struct {
	URL                 string `pulumi:"url" provider:"secret"`
	Token               string `pulumi:"token" provider:"secret"`
	CustomerID          string `pulumi:"customerID" provider:"secret"`
	CloudSpaceID        string `pulumi:"cloudspace_id"`
	ExternalNetworkID   string `pulumi:"external_network_id"`
	ExternalNetworkType string `pulumi:"external_network_type"`
	Metric              int    `pulumi:"metric"`
	ExternalNetworkIP   string `pulumi:"external_network_ip"`
}

func (ex ExternalNetwork) WireDependencies(f infer.FieldSelector, args *ExternalNetworkArgs, state *ExternalNetworkState) {
	f.OutputField(&state.URL).DependsOn(f.InputField(&args.URL))
	f.OutputField(&state.Token).DependsOn(f.InputField(&args.Token))
	f.OutputField(&state.CustomerID).DependsOn(f.InputField(&args.CustomerID))
	f.OutputField(&state.CloudSpaceID).DependsOn(f.InputField(&args.CloudSpaceID))
	f.OutputField(&state.ExternalNetworkID).DependsOn(f.InputField(&args.ExternalNetworkID))
	f.OutputField(&state.ExternalNetworkIP).DependsOn(f.InputField(&args.ExternalNetworkIP))
	f.OutputField(&state.Metric).DependsOn(f.InputField(&args.Metric))
}

func (ExternalNetwork) Create(ctx p.Context, name string, input ExternalNetworkArgs, preview bool) (string, ExternalNetworkState, error) {
	state := ExternalNetworkState{ExternalNetworkArgs: input}
	id, err := resource.NewUniqueHex(name, 8, 0)
	if err != nil {
		return "", state, err
	}
	if preview {
		return name, state, nil
	}

	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/external-networks?external_network_id=%s&external_network_type=%s&metric=%d&external_network_ip=%s", input.URL, input.CustomerID, input.CloudSpaceID, input.ExternalNetworkID, input.ExternalNetworkType, input.Metric, input.ExternalNetworkIP)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(nil))
	if err != nil {
		fmt.Printf("Error making API request %s: %v\n", id, err)
		return "", state, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", input.Token))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error creating resource %s: %v\n", id, err)
		return "", state, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading response body for %s: %v\n", id, err)
			return "", state, err
		}
		fmt.Printf("Error creating resource %s: received status code %d\n: %s\n", id, resp.StatusCode, string(body))
		return "", state, fmt.Errorf("received status code %d\n: %s\n", resp.StatusCode, string(body))
	}

	state.URL = input.URL
	state.CustomerID = input.CustomerID
	state.Token = input.Token
	state.CloudSpaceID = input.CloudSpaceID
	state.ExternalNetworkID = input.ExternalNetworkID
	state.Metric = input.Metric
	state.ExternalNetworkIP = input.ExternalNetworkIP

	return id, state, nil
}

func (ex ExternalNetwork) Update(ctx p.Context, id string, state ExternalNetworkState, input ExternalNetworkArgs, preview bool) (ExternalNetworkState, error) {
	if preview {
		return state, nil
	}
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/external-networks?external_network_id=%s&external_network_type=%s&metric=%d&external_network_ip=%s", state.URL, state.CustomerID, state.CloudSpaceID, input.ExternalNetworkID, state.ExternalNetworkType, input.Metric, input.ExternalNetworkIP)

	client := &http.Client{}
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(nil))
	if err != nil {
		return state, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", state.Token))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error updating resource %s: %v\n", id, err)
		return state, err
	}
	defer resp.Body.Close()

	state.ExternalNetworkID = input.ExternalNetworkID
	state.Metric = input.Metric
	state.ExternalNetworkIP = input.ExternalNetworkIP

	return state, nil
}

func (ExternalNetwork) Delete(ctx p.Context, id string, state ExternalNetworkState) error {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/external-networks?external_network_id=%s&external_network_ip=%s", state.URL, state.CustomerID, state.CloudSpaceID, state.ExternalNetworkID, state.ExternalNetworkIP)

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(nil))
	if err != nil {
		fmt.Printf("Error making API request for %s: %v", id, err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", state.Token))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error deleting resource %s: %v\n", id, err)
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading response body for %s: %v\n", id, err)
			return err
		}
		fmt.Printf("Error creating resource %s: received status code %d\n: %s\n", id, resp.StatusCode, string(body))
		return fmt.Errorf("received status code %d\n: %s\n", resp.StatusCode, string(body))
	}

	return nil
}

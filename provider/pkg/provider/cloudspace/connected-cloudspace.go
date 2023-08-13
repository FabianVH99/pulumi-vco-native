package cloudspace

import (
	"bytes"
	"fmt"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"net/http"
)

type ConnectedCloudspace struct{}

type ConnectedCloudspaceState struct {
	ConnectedCloudspaceArgs
	URL                   string `pulumi:"url"`
	Token                 string `pulumi:"token"`
	CustomerID            string `pulumi:"customerID"`
	CloudSpaceID          string `pulumi:"cloudspace_id"`
	ConnectedCloudSpaceID string `json:"connected_cloudspace_id" pulumi:"connected_cloudspace_id"`
	RemoteCloudSpaceIP    string `json:"remote_cloudspace_ip" pulumi:"remote_cloudspace_ip"`
	LocalCloudSpaceIP     string `json:"local_cloudspace_ip" pulumi:"local_cloudspace_ip"`
	Status                string `json:"status" pulumi:"status"`
}

type ConnectedCloudspaceArgs struct {
	URL                   string `pulumi:"url" provider:"secret"`
	Token                 string `pulumi:"token" provider:"secret"`
	CustomerID            string `pulumi:"customerID" provider:"secret"`
	CloudSpaceID          string `pulumi:"cloudspace_id"`
	ConnectedCloudSpaceID string `pulumi:"connected_cloudspace_id"`
	RemoteCloudSpaceIP    string `pulumi:"remote_cloudspace_ip "`
	LocalCloudSpaceIP     string `pulumi:"local_cloudspace_ip "`
}

func (cs ConnectedCloudspace) WireDependencies(f infer.FieldSelector, args *ConnectedCloudspaceArgs, state *ConnectedCloudspaceState) {
	f.OutputField(&state.URL).DependsOn(f.InputField(&args.URL))
	f.OutputField(&state.Token).DependsOn(f.InputField(&args.Token))
	f.OutputField(&state.CustomerID).DependsOn(f.InputField(&args.CustomerID))
	f.OutputField(&state.CloudSpaceID).DependsOn(f.InputField(&args.CloudSpaceID))
	f.OutputField(&state.ConnectedCloudSpaceID).DependsOn(f.InputField(&args.ConnectedCloudSpaceID))
	f.OutputField(&state.RemoteCloudSpaceIP).DependsOn(f.InputField(&args.RemoteCloudSpaceIP))
	f.OutputField(&state.LocalCloudSpaceIP).DependsOn(f.InputField(&args.LocalCloudSpaceIP))
}

func (cs ConnectedCloudspace) Create(ctx p.Context, name string, input ConnectedCloudspaceArgs, preview bool) (string, ConnectedCloudspaceState, error) {
	state := ConnectedCloudspaceState{ConnectedCloudspaceArgs: input}
	id, err := resource.NewUniqueHex(name, 8, 0)
	if err != nil {
		return "", state, err
	}
	if preview {
		return name, state, nil
	}

	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/connected-cloudspaces?connected_cloudspace_id=%s&remote_cloudspace_ip=%s&local_cloudspace_ip=%s", input.URL, input.CustomerID, input.CloudSpaceID, input.ConnectedCloudSpaceID, input.RemoteCloudSpaceIP, input.LocalCloudSpaceIP)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(nil))
	if err != nil {
		fmt.Printf("Error making API request for %s: %v", id, err)
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

	state.URL = input.URL
	state.CustomerID = input.CustomerID
	state.Token = input.Token
	state.CloudSpaceID = input.CloudSpaceID
	state.ConnectedCloudSpaceID = input.ConnectedCloudSpaceID
	state.RemoteCloudSpaceIP = input.RemoteCloudSpaceIP
	state.LocalCloudSpaceIP = input.LocalCloudSpaceIP

	return id, state, nil
}

func (ConnectedCloudspace) Update(ctx p.Context, id string, state ConnectedCloudspaceState, input ConnectedCloudspaceArgs, preview bool) (ConnectedCloudspaceState, error) {
	if preview {
		return state, nil
	}
	return state, nil
}

func (ConnectedCloudspace) Delete(ctx p.Context, id string, state ConnectedCloudspaceState) error {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/connected-cloudspaces/%s?remote_cloudspace_ip=%s&local_cloudspace_ip=%s", state.URL, state.CustomerID, state.CloudSpaceID, state.ConnectedCloudSpaceID, state.RemoteCloudSpaceIP, state.LocalCloudSpaceIP)

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

	return nil
}

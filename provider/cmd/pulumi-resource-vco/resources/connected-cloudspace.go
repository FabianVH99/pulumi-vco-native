package resources

import (
	"bytes"
	"fmt"
	p "github.com/pulumi/pulumi-go-provider"
	"net/http"
)

type ConnectedCloudspace struct{}

type ConnectedCloudspaceState struct {
	ConnectedCloudspaceArgs
	ConnectedCloudSpaceID string `json:"connected_cloudspace_id" pulumi:"connected_cloudspace_id"`
	RemoteCloudSpaceIP    string `json:"remote_cloudspace_ip" pulumi:"remote_cloudspace_ip"`
	LocalCloudSpaceIP     string `json:"local_cloudspace_ip" pulumi:"local_cloudspace_ip"`
	Status                string `json:"status" pulumi:"status"`
}

type ConnectedCloudspaceArgs struct {
	URL                   string `pulumi:"url"`
	Token                 string `pulumi:"token"`
	CustomerID            string `pulumi:"customerID"`
	CloudSpaceID          string `pulumi:"cloudspace_id"`
	ConnectedCloudSpaceID string `pulumi:"connected_cloudspace_id"`
	RemoteCloudSpaceIP    string `pulumi:"remote_cloudspace_ip "`
	LocalCloudSpaceIP     string `pulumi:"local_cloudspace_ip "`
}

func (cs ConnectedCloudspace) Create(ctx p.Context, name string, input ConnectedCloudspaceArgs, preview bool) (string, ConnectedCloudspaceState, error) {
	state := ConnectedCloudspaceState{ConnectedCloudspaceArgs: input}
	state.ConnectedCloudSpaceID = input.ConnectedCloudSpaceID
	state.RemoteCloudSpaceIP = input.RemoteCloudSpaceIP
	state.LocalCloudSpaceIP = input.LocalCloudSpaceIP

	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/connected-cloudspaces?connected_cloudspace_id=%s&remote_cloudspace_ip=%s&local_cloudspace_ip=%s", input.URL, input.CustomerID, input.CloudSpaceID, input.ConnectedCloudSpaceID, input.RemoteCloudSpaceIP, input.LocalCloudSpaceIP)

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

	return name, state, nil
}

func (ConnectedCloudspace) Delete(ctx p.Context, input ConnectedCloudspaceArgs) (int, error) {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/connected-cloudspaces/%s?remote_cloudspace_ip=%s&local_cloudspace_ip=%s", input.URL, input.CustomerID, input.CloudSpaceID, input.ConnectedCloudSpaceID, input.RemoteCloudSpaceIP, input.LocalCloudSpaceIP)

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(nil))
	if err != nil {
		return 0, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", input.Token))

	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}

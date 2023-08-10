package resources

import (
	"bytes"
	"encoding/json"
	"fmt"
	p "github.com/pulumi/pulumi-go-provider"
	"net/http"
)

type ObjectSpaceLink struct{}

type ObjectSpaceLinkState struct {
	ObjectSpaceLinkArgs
	ObjectSpaceID string `pulumi:"objectspace_id" json:"objectspace_id"`
	Status        bool   `pulumi:"success" json:"success"`
}

type ObjectSpaceLinkArgs struct {
	URL           string `pulumi:"url"`
	Token         string `pulumi:"token"`
	CustomerID    string `pulumi:"customerID"`
	CloudSpaceID  string `pulumi:"cloudspace_id"`
	ObjectSpaceID string `pulumi:"objectspace_id"`
}

func (ObjectSpaceLink) Create(ctx p.Context, name string, input ObjectSpaceLinkArgs, preview bool) (string, ObjectSpaceLinkState, error) {
	state := ObjectSpaceLinkState{ObjectSpaceLinkArgs: input}
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/objectspaces?objectspace_id=%s", input.URL, input.CustomerID, input.CloudSpaceID, input.ObjectSpaceID)

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

	state.ObjectSpaceID = input.ObjectSpaceID
	state.Status = result["success"].(bool)

	return name, state, nil
}

func (ObjectSpaceLink) Delete(ctx p.Context, state ObjectSpaceLinkState, input ObjectSpaceLinkArgs) (bool, error) {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/objectspaces?objectspace_id=%s", input.URL, input.CustomerID, input.CloudSpaceID, state.ObjectSpaceID)

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

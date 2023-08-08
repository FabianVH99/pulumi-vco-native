package resources

import (
	"bytes"
	"encoding/json"
	"fmt"
	p "github.com/pulumi/pulumi-go-provider"
	"net/http"
)

type AntiAffinityGroup struct{}

type AntiAffinityGroupState struct {
	AntiAffinityGroupArgs
	GroupID string `json:"group_id" pulumi:"group_id"`
	Spread  int    `json:"spread" pulumi:"spread"`
	Status  string `json:"status" pulumi:"status"`
}

type AntiAffinityGroupArgs struct {
	URL          string `pulumi:"url"`
	Token        string `pulumi:"token"`
	CustomerID   string `pulumi:"customerID"`
	CloudSpaceID string `pulumi:"cloudspace_id"`
	GroupID      string `pulumi:"group_id"`
	Spread       int    `pulumi:"spread"`
}

func (ag AntiAffinityGroup) Create(ctx p.Context, name string, input AntiAffinityGroupArgs, preview bool) (string, AntiAffinityGroupState, error) {
	state := AntiAffinityGroupState{AntiAffinityGroupArgs: input}

	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/anti-affinity-groups?spread=%d&group_id=%s", input.URL, input.CustomerID, input.CloudSpaceID, input.Spread, input.GroupID)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, nil)
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

	updatedState, err := ag.Read(ctx, state, input)
	if err != nil {
		return "", state, err
	}

	return name, updatedState, nil
}

func (AntiAffinityGroup) Read(ctx p.Context, state AntiAffinityGroupState, input AntiAffinityGroupArgs) (AntiAffinityGroupState, error) {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/anti-affinity-groups/%s", input.URL, input.CustomerID, input.CloudSpaceID, input.GroupID)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return AntiAffinityGroupState{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", input.Token))

	resp, err := client.Do(req)
	if err != nil {
		return AntiAffinityGroupState{}, err
	}
	defer resp.Body.Close()

	var result AntiAffinityGroupState
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return AntiAffinityGroupState{}, err
	}

	return result, nil
}

func (ag AntiAffinityGroup) Update(ctx p.Context, state AntiAffinityGroupState, input AntiAffinityGroupArgs) (bool, AntiAffinityGroupState, error) {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/anti-affinity-groups/%s?spread=%d", input.URL, input.CustomerID, input.CloudSpaceID, input.GroupID, input.Spread)

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

	updatedState, err := ag.Read(ctx, state, input)
	if err != nil {
		return false, state, err
	}

	return success, updatedState, nil
}

func (AntiAffinityGroup) Delete(ctx p.Context, state AntiAffinityGroupState, input AntiAffinityGroupArgs) (bool, error) {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/anti-affinity-groups/%s", input.URL, input.CustomerID, input.CloudSpaceID, state.GroupID)
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

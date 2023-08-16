package cloudspace

import (
	"bytes"
	"encoding/json"
	"fmt"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type AntiAffinityGroup struct{}

type AntiAffinityGroupState struct {
	AntiAffinityGroupArgs
	URL          string `pulumi:"url"`
	Token        string `pulumi:"token"`
	CustomerID   string `pulumi:"customerID"`
	CloudSpaceID string `json:"cloudspace_id" pulumi:"cloudspace_id"`
	GroupID      string `json:"group_id" pulumi:"group_id"`
	Spread       int    `json:"spread" pulumi:"spread"`
	Status       string `json:"status" pulumi:"status"`
}

type AntiAffinityGroupArgs struct {
	URL          string `pulumi:"url" provider:"secret"`
	Token        string `pulumi:"token" provider:"secret"`
	CustomerID   string `pulumi:"customerID" provider:"secret"`
	CloudSpaceID string `pulumi:"cloudspace_id"`
	GroupID      string `pulumi:"group_id"`
	Spread       int    `pulumi:"spread"`
}

func (ag AntiAffinityGroup) WireDependencies(f infer.FieldSelector, args *AntiAffinityGroupArgs, state *AntiAffinityGroupState) {
	f.OutputField(&state.URL).DependsOn(f.InputField(&args.URL))
	f.OutputField(&state.Token).DependsOn(f.InputField(&args.Token))
	f.OutputField(&state.CustomerID).DependsOn(f.InputField(&args.CustomerID))
	f.OutputField(&state.CloudSpaceID).DependsOn(f.InputField(&args.CloudSpaceID))
	f.OutputField(&state.GroupID).DependsOn(f.InputField(&args.GroupID))
	f.OutputField(&state.Spread).DependsOn(f.InputField(&args.Spread))
}

func (ag AntiAffinityGroup) Create(ctx p.Context, name string, input AntiAffinityGroupArgs, preview bool) (string, AntiAffinityGroupState, error) {
	state := AntiAffinityGroupState{AntiAffinityGroupArgs: input}
	id, err := resource.NewUniqueHex(name, 8, 0)
	if err != nil {
		return "", state, err
	}
	if preview {
		return name, state, nil
	}

	u, err := url.Parse(fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/anti-affinity-groups", input.URL, input.CustomerID, input.CloudSpaceID))
	if err != nil {
		fmt.Printf("Error making API request for %s: %v", id, err)
		return "", state, err
	}

	q := u.Query()
	q.Set("group_id", input.GroupID)
	q.Set("spread", strconv.Itoa(input.Spread))

	u.RawQuery = q.Encode()
	url := u.String()

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, nil)
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
	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading response body for %s: %v\n", id, err)
			return "", state, err
		}
		fmt.Printf("Error creating resource %s: received status code %d\n: %s\n", id, resp.StatusCode, string(body))
		return "", state, fmt.Errorf("received status code %d\n: %s\n", resp.StatusCode, string(body))
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Printf("Error decoding response body for %s: %v\n", id, err)
		return "", state, err
	}
	state.URL = input.URL
	state.CustomerID = input.CustomerID
	state.Token = input.Token
	state.CloudSpaceID = input.CloudSpaceID
	state.GroupID = input.GroupID

	updatedState, err := ag.Read(ctx, id, state)
	if err != nil {
		return "", state, err
	}

	return id, updatedState, nil
}

func (AntiAffinityGroup) Read(ctx p.Context, id string, state AntiAffinityGroupState) (AntiAffinityGroupState, error) {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/anti-affinity-groups/%s", state.URL, state.CustomerID, state.CloudSpaceID, state.GroupID)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error making API request for %s: %v", id, err)
		return AntiAffinityGroupState{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", state.Token))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error retrieving resource %s: %v\n", id, err)
		return AntiAffinityGroupState{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading response body for %s: %v\n", id, err)
			return state, err
		}
		fmt.Printf("Error creating resource %s: received status code %d\n: %s\n", id, resp.StatusCode, string(body))
		return state, fmt.Errorf("received status code %d\n: %s\n", resp.StatusCode, string(body))
	}

	var result AntiAffinityGroupState
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Printf("Error decoding response body for %s: %v\n", id, err)
		return AntiAffinityGroupState{}, err
	}

	result.URL = state.URL
	result.CustomerID = state.CustomerID
	result.Token = state.Token
	result.CloudSpaceID = state.CloudSpaceID

	return result, nil
}

func (ag AntiAffinityGroup) Update(ctx p.Context, id string, state AntiAffinityGroupState, input AntiAffinityGroupArgs, preview bool) (AntiAffinityGroupState, error) {
	if preview {
		return state, nil
	}
	u, err := url.Parse(fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/anti-affinity-groups/%s", state.URL, state.CustomerID, state.CloudSpaceID, state.GroupID))
	if err != nil {
		fmt.Printf("Error making API request for %s: %v", id, err)
		return state, err
	}
	q := u.Query()
	q.Set("spread", strconv.Itoa(input.Spread))
	u.RawQuery = q.Encode()
	url := u.String()

	client := &http.Client{}
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(nil))
	if err != nil {
		fmt.Printf("Error making API request for %s: %v", id, err)
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
	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading response body for %s: %v\n", id, err)
			return state, err
		}
		fmt.Printf("Error creating resource %s: received status code %d\n: %s\n", id, resp.StatusCode, string(body))
		return state, fmt.Errorf("received status code %d\n: %s\n", resp.StatusCode, string(body))
	}

	updatedState, err := ag.Read(ctx, id, state)
	if err != nil {
		return state, err
	}

	return updatedState, nil
}

func (AntiAffinityGroup) Delete(ctx p.Context, id string, state AntiAffinityGroupState) error {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/anti-affinity-groups/%s", state.URL, state.CustomerID, state.CloudSpaceID, state.GroupID)
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

package base

import (
	"bytes"
	"encoding/json"
	"fmt"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"io/ioutil"
	"net/http"
)

type ObjectSpace struct{}

type ObjectSpaceState struct {
	ObjectSpaceArgs
	URL             string `pulumi:"url"`
	Token           string `pulumi:"token"`
	CustomerID      string `pulumi:"customerID"`
	ObjectSpaceID   string `pulumi:"objectspace_id" json:"objectspace_id"`
	ObjectSpaceName string `pulumi:"objectspace_name" json:"objectspace_name"`
	Status          string `pulumi:"status" json:"status"`
	AccessKey       string `pulumi:"access_key" json:"access_key"`
	Secret          string `pulumi:"secret" json:"secret"`
	CreationTime    string `pulumi:"creation_time" json:"creation_time"`
	UpdateTime      string `pulumi:"update_time" json:"update_time"`
	Location        string `pulumi:"location" json:"location"`
	Domain          string `pulumi:"domain" json:"domain"`
}

type ObjectSpaceArgs struct {
	URL              string  `pulumi:"url" provider:"secret"`
	Token            string  `pulumi:"token" provider:"secret"`
	CustomerID       string  `pulumi:"customerID" provider:"secret"`
	Location         string  `pulumi:"location"`
	ObjectSpaceName  string  `pulumi:"objectspace_name"`
	Domain           *string `pulumi:"domain,optional"`
	CloudspaceID     *string `pulumi:"cloudspaceID,optional"`
	Subnet           *string `pulumi:"subnet,optional"`
	ExternalNetwork  *int    `pulumi:"externalNetwork,optional"`
	LetsencryptEmail *string `pulumi:"letsencryptEmail,optional"`
	Letsencrypt      *bool   `pulumi:"letsencrypt,optional"`
}

func (c ObjectSpace) WireDependencies(f infer.FieldSelector, args *ObjectSpaceArgs, state *ObjectSpaceState) {
	f.OutputField(&state.URL).DependsOn(f.InputField(&args.URL))
	f.OutputField(&state.Token).DependsOn(f.InputField(&args.Token))
	f.OutputField(&state.CustomerID).DependsOn(f.InputField(&args.CustomerID))
	f.OutputField(&state.ObjectSpaceName).DependsOn(f.InputField(&args.ObjectSpaceName))
	f.OutputField(&state.Location).DependsOn(f.InputField(&args.Location))
	f.OutputField(&state.Location).DependsOn(f.InputField(&args.Location))
}

func (obj ObjectSpace) Create(ctx p.Context, name string, input ObjectSpaceArgs, preview bool) (string, ObjectSpaceState, error) {
	state := ObjectSpaceState{ObjectSpaceArgs: input}
	id, err := resource.NewUniqueHex(name, 8, 0)
	if err != nil {
		return "", state, err
	}
	if preview {
		return name, state, nil
	}

	url := fmt.Sprintf("https://%s/api/1/customers/%s/objectspaces?name=%s&location=%s", input.URL, input.CustomerID, input.ObjectSpaceName, input.Location)
	payload := map[string]interface{}{}

	if input.Domain != nil {
		payload["domain"] = *input.Domain
	}

	if input.CloudspaceID != nil || input.Subnet != nil || input.ExternalNetwork != nil {
		cloudspace := map[string]interface{}{}
		if input.CloudspaceID != nil {
			cloudspace["cloudspace_id"] = *input.CloudspaceID
		}
		if input.Subnet != nil {
			cloudspace["subnet"] = *input.Subnet
		}
		if input.ExternalNetwork != nil {
			cloudspace["external_network"] = *input.ExternalNetwork
		}
		payload["cloudspace"] = cloudspace
	}

	if input.LetsencryptEmail != nil {
		payload["letsencrypt_email"] = *input.LetsencryptEmail
	}

	if input.Letsencrypt != nil {
		payload["letsencrypt"] = *input.Letsencrypt
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", state, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
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
		return "", state, err
	}

	ObjectSpaceID := result["objectspace_id"].(string)
	state.ObjectSpaceID = ObjectSpaceID
	state.URL = input.URL
	state.CustomerID = input.CustomerID
	state.Token = input.Token

	updatedState, err := obj.Read(ctx, id, state)
	if err != nil {
		return "", state, err
	}

	return id, updatedState, nil
}

func (ObjectSpace) Read(ctx p.Context, id string, state ObjectSpaceState) (ObjectSpaceState, error) {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/objectspaces/%s", state.URL, state.CustomerID, state.ObjectSpaceID)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error making API request for %s: %v", id, err)
		return ObjectSpaceState{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", state.Token))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error retrieving resource %s: %v\n", id, err)
		return ObjectSpaceState{}, err
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

	var result ObjectSpaceState
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Printf("Error decoding response body for %s: %v\n", id, err)
		return ObjectSpaceState{}, err
	}

	result.URL = state.URL
	result.CustomerID = state.CustomerID
	result.Token = state.Token
	result.ObjectSpaceID = state.ObjectSpaceID

	return result, nil
}

func (ObjectSpace) Update(ctx p.Context, id string, state ObjectSpaceState, input ObjectSpaceArgs, preview bool) (ObjectSpaceState, error) {
	if preview {
		return state, nil
	}
	return state, nil
}

func (ObjectSpace) Delete(ctx p.Context, id string, state ObjectSpaceState) error {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/objectspaces/%s", state.URL, state.CustomerID, state.ObjectSpaceID)

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

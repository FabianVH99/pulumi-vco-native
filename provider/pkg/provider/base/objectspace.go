package base

import (
	"bytes"
	"encoding/json"
	"fmt"
	p "github.com/pulumi/pulumi-go-provider"
	"net/http"
)

type ObjectSpace struct{}

type CS struct {
	Name              string `pulumi:"name" json:"name"`
	CloudspaceID      string `pulumi:"cloudspace_id" json:"cloudspace_id"`
	IPAddress         string `pulumi:"ip_address" json:"ip_address"`
	ExternalIPAddress string `pulumi:"external_ip_address" json:"external_ip_address"`
	Status            string `pulumi:"status" json:"status"`
	Mode              string `pulumi:"mode" json:"mode"`
}

type Limit struct {
	BucketName string `pulumi:"bucket_name" json:"bucket_name"`
	Limit      int    `pulumi:"limit" json:"limit"`
}

type ObjectSpaceState struct {
	ObjectSpaceArgs
	ObjectSpaceID   string  `pulumi:"objectspace_id" json:"objectspace_id"`
	ObjectSpaceName string  `pulumi:"objectspace_name" json:"objectspace_name"`
	Status          string  `pulumi:"status" json:"status"`
	AccessKey       string  `pulumi:"access_key" json:"access_key"`
	Secret          string  `pulumi:"secret" json:"secret"`
	CreationTime    string  `pulumi:"creation_time" json:"creation_time"`
	UpdateTime      string  `pulumi:"update_time" json:"update_time"`
	Cloudspaces     []CS    `pulumi:"cloudspaces" json:"cloudspaces"`
	Limits          []Limit `pulumi:"limits" json:"limits"`
	Location        string  `pulumi:"location" json:"location"`
	Domain          string  `pulumi:"domain" json:"domain"`
}

type ObjectSpaceArgs struct {
	URL              string  `pulumi:"url"`
	Token            string  `pulumi:"token"`
	CustomerID       string  `pulumi:"customerID"`
	Location         string  `pulumi:"location"`
	Name             string  `pulumi:"name"`
	Domain           *string `pulumi:"domain,optional"`
	CloudspaceID     *string `pulumi:"cloudspaceID,optional"`
	Subnet           *string `pulumi:"subnet,optional"`
	ExternalNetwork  *int    `pulumi:"externalNetwork,optional"`
	LetsencryptEmail *string `pulumi:"letsencryptEmail,optional"`
	Letsencrypt      *bool   `pulumi:"letsencrypt,optional"`
}

func (obj ObjectSpace) Create(ctx p.Context, name string, input ObjectSpaceArgs, preview bool) (string, ObjectSpaceState, error) {
	state := ObjectSpaceState{ObjectSpaceArgs: input}

	if preview {
		return name, state, nil
	}

	url := fmt.Sprintf("https://%s/api/1/customers/%s/objectspaces?name=%s&location=%s", input.URL, input.CustomerID, input.Name, input.Location)
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

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", state, err
	}

	ObjectSpaceID := result["objectspace_id"].(string)
	state.ObjectSpaceID = ObjectSpaceID

	updatedState, err := obj.Read(ctx, state, input)
	if err != nil {
		return "", state, err
	}

	return name, updatedState, nil
}

func (ObjectSpace) Read(ctx p.Context, state ObjectSpaceState, input ObjectSpaceArgs) (ObjectSpaceState, error) {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/objectspaces/%s", input.URL, input.CustomerID, state.ObjectSpaceID)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ObjectSpaceState{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", input.Token))

	resp, err := client.Do(req)
	if err != nil {
		return ObjectSpaceState{}, err
	}
	defer resp.Body.Close()

	var result ObjectSpaceState
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return ObjectSpaceState{}, err
	}

	return result, nil
}

func (ObjectSpace) Delete(ctx p.Context, state ObjectSpaceState, input ObjectSpaceArgs) (bool, error) {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/objectspaces/%s", input.URL, input.CustomerID, state.ObjectSpaceID)

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

package objectspace

import (
	"bytes"
	"fmt"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"io/ioutil"
	"net/http"
)

type Link struct{}

type LinkState struct {
	LinkArgs
	URL           string `pulumi:"url"`
	Token         string `pulumi:"token"`
	CustomerID    string `pulumi:"customerID"`
	CloudSpaceID  string `pulumi:"cloudspace_id"`
	ObjectSpaceID string `pulumi:"objectspace_id" json:"objectspace_id"`
}

type LinkArgs struct {
	URL           string `pulumi:"url" provider:"secret"`
	Token         string `pulumi:"token" provider:"secret"`
	CustomerID    string `pulumi:"customerID" provider:"secret"`
	CloudSpaceID  string `pulumi:"cloudspace_id"`
	ObjectSpaceID string `pulumi:"objectspace_id"`
}

func (sv Link) WireDependencies(f infer.FieldSelector, args *LinkArgs, state *LinkState) {
	f.OutputField(&state.URL).DependsOn(f.InputField(&args.URL))
	f.OutputField(&state.Token).DependsOn(f.InputField(&args.Token))
	f.OutputField(&state.CustomerID).DependsOn(f.InputField(&args.CustomerID))
	f.OutputField(&state.CloudSpaceID).DependsOn(f.InputField(&args.CloudSpaceID))
	f.OutputField(&state.ObjectSpaceID).DependsOn(f.InputField(&args.ObjectSpaceID))
}

func (Link) Create(ctx p.Context, name string, input LinkArgs, preview bool) (string, LinkState, error) {
	state := LinkState{LinkArgs: input}
	id, err := resource.NewUniqueHex(name, 8, 0)
	if err != nil {
		return "", state, err
	}
	if preview {
		return name, state, nil
	}

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
	state.ObjectSpaceID = input.ObjectSpaceID

	return id, state, nil
}

func (Link) Update(ctx p.Context, id string, state LinkState, input LinkArgs, preview bool) (LinkState, error) {
	if preview {
		return state, nil
	}
	return state, nil
}

func (Link) Delete(ctx p.Context, id string, state LinkState) error {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/objectspaces?objectspace_id=%s", state.URL, state.CustomerID, state.CloudSpaceID, state.ObjectSpaceID)

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

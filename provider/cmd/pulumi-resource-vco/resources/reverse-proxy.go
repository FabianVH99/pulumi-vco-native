package resources

import (
	"bytes"
	"encoding/json"
	"fmt"
	p "github.com/pulumi/pulumi-go-provider"
	"net/http"
)

type ReverseProxy struct{}

type ReverseProxyState struct {
	ReverseProxyArgs
	ReverseProxyID string               `pulumi:"reverseproxy_id" json:"reverseproxy_id"`
	Name           string               `pulumi:"name" json:"name"`
	Description    string               `pulumi:"description" json:"description"`
	Type           string               `pulumi:"type" json:"type"`
	FrontEnd       ReverseProxyFrontEnd `pulumi:"front_end" json:"front_end"`
	BackEnd        ReverseProxyBackend  `pulumi:"back_end" json:"back_end"`
}

type ReverseProxyArgs struct {
	URL                  string               `pulumi:"url"`
	Token                string               `pulumi:"token"`
	CustomerID           string               `pulumi:"customerID"`
	CloudSpaceID         string               `pulumi:"cloudspace_id"`
	Name                 string               `pulumi:"name"`
	Description          *string              `pulumi:"description,optional"`
	ReverseProxyFrontEnd ReverseProxyFrontEnd `pulumi:"front_end"`
	ReverseProxyBackend  ReverseProxyBackend  `pulumi:"back_end"`
}

type ReverseProxyFrontEnd struct {
	Domain      string      `pulumi:"domain" json:"domain"`
	HTTPPort    *int        `pulumi:"http_port,optional" json:"http_port"`
	HTTPSPort   *int        `pulumi:"https_port,optional" json:"https_port"`
	IPAddress   *string     `pulumi:"ip_address,optional" json:"ip_address"`
	Scheme      string      `pulumi:"scheme" json:"scheme"`
	LetsEncrypt LetsEncrypt `pulumi:"letsencrypt" json:"letsencrypt"`
}

type LetsEncrypt struct {
	Enabled bool    `pulumi:"enabled" json:"enabled"`
	Email   *string `pulumi:"email" json:"email"`
}

type ReverseProxyBackend struct {
	Scheme       string   `pulumi:"scheme" json:"scheme"`
	ServerpoolID string   `pulumi:"serverpool_id" json:"serverpool_id"`
	TargetPort   int      `pulumi:"target_port" json:"target_port"`
	Options      *Options `pulumi:"options,optional" json:"options"`
}

type Options struct {
	StickySessionCookie *StickySessionCookie `pulumi:"sticky_session_cookie,optional" json:"sticky_session_cookie"`
	HealthCheck         *HealthCheck         `pulumi:"health_check,optional" json:"health_check"`
}

type StickySessionCookie struct {
	Name     *string `pulumi:"name,optional" json:"name"`
	Secure   *bool   `pulumi:"secure,optional" json:"secure"`
	HttpOnly *bool   `pulumi:"http_only,optional" json:"http_only"`
	SameSite *string `pulumi:"same_site,optional" json:"same_site"`
}

type HealthCheck struct {
	Path     *string `pulumi:"path,optional" json:"path"`
	Scheme   *string `pulumi:"scheme,optional" json:"scheme"`
	Port     *int    `pulumi:"port,optional" json:"port"`
	Interval *int    `pulumi:"interval,optional" json:"interval"`
	Timeout  *int    `pulumi:"timeout,optional" json:"timeout"`
}

func (rp ReverseProxy) Create(ctx p.Context, name string, input ReverseProxyArgs, preview bool) (string, ReverseProxyState, error) {
	state := ReverseProxyState{ReverseProxyArgs: input}

	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/ingress/reverse-proxies", input.URL, input.CustomerID, input.CloudSpaceID)
	payload := map[string]interface{}{
		"name": input.Name,
		"front_end": map[string]interface{}{
			"domain": input.ReverseProxyFrontEnd.Domain,
			"scheme": input.ReverseProxyFrontEnd.Scheme,
		},
		"back_end": map[string]interface{}{
			"scheme":        input.ReverseProxyBackend.Scheme,
			"serverpool_id": input.ReverseProxyBackend.ServerpoolID,
			"target_port":   input.ReverseProxyBackend.TargetPort,
		},
	}

	if input.Description != nil {
		payload["description"] = *input.Description
	}

	if input.ReverseProxyFrontEnd.HTTPPort != nil {
		payload["front_end"].(map[string]interface{})["http_port"] = *input.ReverseProxyFrontEnd.HTTPPort
	}

	if input.ReverseProxyFrontEnd.HTTPSPort != nil {
		payload["front_end"].(map[string]interface{})["https_port"] = *input.ReverseProxyFrontEnd.HTTPSPort
	}

	if input.ReverseProxyFrontEnd.IPAddress != nil {
		payload["front_end"].(map[string]interface{})["ip_address"] = *input.ReverseProxyFrontEnd.IPAddress
	} else {
		payload["front_end"].(map[string]interface{})["ip_address"] = ""
	}

	letsecrypt := map[string]interface{}{
		"enabled": input.ReverseProxyFrontEnd.LetsEncrypt.Enabled,
	}
	if input.ReverseProxyFrontEnd.LetsEncrypt.Email != nil {
		letsecrypt["email"] = *input.ReverseProxyFrontEnd.LetsEncrypt.Email
	}
	payload["front_end"].(map[string]interface{})["letsencrypt"] = letsecrypt

	if input.ReverseProxyBackend.Options != nil {
		if input.ReverseProxyBackend.Options.StickySessionCookie.Name != nil || input.ReverseProxyBackend.Options.StickySessionCookie.Secure != nil || input.ReverseProxyBackend.Options.StickySessionCookie.HttpOnly != nil || input.ReverseProxyBackend.Options.StickySessionCookie.SameSite != nil {
			stickySessionCookie := map[string]interface{}{}
			if input.ReverseProxyBackend.Options.StickySessionCookie.Name != nil {
				stickySessionCookie["name"] = *input.ReverseProxyBackend.Options.StickySessionCookie.Name
			}
			if input.ReverseProxyBackend.Options.StickySessionCookie.Secure != nil {
				stickySessionCookie["secure"] = *input.ReverseProxyBackend.Options.StickySessionCookie.Secure
			}
			if input.ReverseProxyBackend.Options.StickySessionCookie.HttpOnly != nil {
				stickySessionCookie["http_only"] = *input.ReverseProxyBackend.Options.StickySessionCookie.HttpOnly
			}
			if input.ReverseProxyBackend.Options.StickySessionCookie.SameSite != nil {
				stickySessionCookie["same_site"] = *input.ReverseProxyBackend.Options.StickySessionCookie.SameSite
			}

			if len(stickySessionCookie) > 0 {
				options := map[string]interface{}{}
				options["sticky_session_cookie"] = stickySessionCookie
				payload["back_end"].(map[string]interface{})["options"] = options
			}
		}

		if payload["back_end"].(map[string]interface{})["options"] == nil && (input.ReverseProxyBackend.Options.HealthCheck.Path != nil ||
			input.ReverseProxyBackend.Options.HealthCheck.Scheme != nil ||
			input.ReverseProxyBackend.Options.HealthCheck.Port != nil ||
			input.ReverseProxyBackend.Options.HealthCheck.Interval != nil ||
			input.ReverseProxyBackend.Options.HealthCheck.Timeout != nil) {

			options := map[string]interface{}{}
			payload["back_end"].(map[string]interface{})["options"] = options
		}

		if payload["back_end"].(map[string]interface{})["options"] != nil && (input.ReverseProxyBackend.Options.HealthCheck.Path != nil ||
			input.ReverseProxyBackend.Options.HealthCheck.Scheme != nil ||
			input.ReverseProxyBackend.Options.HealthCheck.Port != nil ||
			input.ReverseProxyBackend.Options.HealthCheck.Interval != nil ||
			input.ReverseProxyBackend.Options.HealthCheck.Timeout != nil) {

			healthCheck := map[string]interface{}{}
			if input.ReverseProxyBackend.Options.HealthCheck.Path != nil {
				healthCheck["path"] = *input.ReverseProxyBackend.Options.HealthCheck.Path
			}
			if input.ReverseProxyBackend.Options.HealthCheck.Scheme != nil {
				healthCheck["scheme"] = *input.ReverseProxyBackend.Options.HealthCheck.Scheme
			}
			if input.ReverseProxyBackend.Options.HealthCheck.Port != nil {
				healthCheck["port"] = *input.ReverseProxyBackend.Options.HealthCheck.Port
			}
			if input.ReverseProxyBackend.Options.HealthCheck.Interval != nil {
				healthCheck["interval"] = *input.ReverseProxyBackend.Options.HealthCheck.Interval
			}
			if input.ReverseProxyBackend.Options.HealthCheck.Timeout != nil {
				healthCheck["timeout"] = *input.ReverseProxyBackend.Options.HealthCheck.Timeout
			}

			payload["back_end"].(map[string]interface{})["options"].(map[string]interface{})["health_check"] = healthCheck
		}
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

	state.ReverseProxyID = result["id"].(string)

	updatedState, err := rp.Read(ctx, state, input)
	if err != nil {
		return "", state, err
	}

	return name, updatedState, nil
}

func (ReverseProxy) Read(ctx p.Context, state ReverseProxyState, input ReverseProxyArgs) (ReverseProxyState, error) {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/ingress/reverse-proxies/%s", input.URL, input.CustomerID, input.CloudSpaceID, state.ReverseProxyID)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ReverseProxyState{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", input.Token))

	resp, err := client.Do(req)
	if err != nil {
		return ReverseProxyState{}, err
	}
	defer resp.Body.Close()

	var result ReverseProxyState
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return ReverseProxyState{}, err
	}

	return result, nil
}

func (rp ReverseProxy) Update(ctx p.Context, state ReverseProxyState, input ReverseProxyArgs) (bool, ReverseProxyState, error) {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/ingress/reverse-proxies/%s", input.URL, input.CustomerID, input.CloudSpaceID, state.ReverseProxyID)
	payload := map[string]interface{}{
		"name": input.Name,
		"front_end": map[string]interface{}{
			"domain": input.ReverseProxyFrontEnd.Domain,
			"scheme": input.ReverseProxyFrontEnd.Scheme,
		},
		"back_end": map[string]interface{}{
			"scheme":        input.ReverseProxyBackend.Scheme,
			"serverpool_id": input.ReverseProxyBackend.ServerpoolID,
			"target_port":   input.ReverseProxyBackend.TargetPort,
		},
	}

	if input.Description != nil {
		payload["description"] = *input.Description
	}

	if input.ReverseProxyFrontEnd.HTTPPort != nil {
		payload["front_end"].(map[string]interface{})["http_port"] = *input.ReverseProxyFrontEnd.HTTPPort
	}

	if input.ReverseProxyFrontEnd.HTTPSPort != nil {
		payload["front_end"].(map[string]interface{})["https_port"] = *input.ReverseProxyFrontEnd.HTTPSPort
	}

	if input.ReverseProxyFrontEnd.IPAddress != nil {
		payload["front_end"].(map[string]interface{})["ip_address"] = *input.ReverseProxyFrontEnd.IPAddress
	} else {
		payload["front_end"].(map[string]interface{})["ip_address"] = ""
	}

	letsecrypt := map[string]interface{}{
		"enabled": input.ReverseProxyFrontEnd.LetsEncrypt.Enabled,
	}
	if input.ReverseProxyFrontEnd.LetsEncrypt.Email != nil {
		letsecrypt["email"] = *input.ReverseProxyFrontEnd.LetsEncrypt.Email
	}
	payload["front_end"].(map[string]interface{})["letsencrypt"] = letsecrypt

	if input.ReverseProxyBackend.Options != nil {
		if input.ReverseProxyBackend.Options.StickySessionCookie.Name != nil || input.ReverseProxyBackend.Options.StickySessionCookie.Secure != nil || input.ReverseProxyBackend.Options.StickySessionCookie.HttpOnly != nil || input.ReverseProxyBackend.Options.StickySessionCookie.SameSite != nil {
			stickySessionCookie := map[string]interface{}{}
			if input.ReverseProxyBackend.Options.StickySessionCookie.Name != nil {
				stickySessionCookie["name"] = *input.ReverseProxyBackend.Options.StickySessionCookie.Name
			}
			if input.ReverseProxyBackend.Options.StickySessionCookie.Secure != nil {
				stickySessionCookie["secure"] = *input.ReverseProxyBackend.Options.StickySessionCookie.Secure
			}
			if input.ReverseProxyBackend.Options.StickySessionCookie.HttpOnly != nil {
				stickySessionCookie["http_only"] = *input.ReverseProxyBackend.Options.StickySessionCookie.HttpOnly
			}
			if input.ReverseProxyBackend.Options.StickySessionCookie.SameSite != nil {
				stickySessionCookie["same_site"] = *input.ReverseProxyBackend.Options.StickySessionCookie.SameSite
			}

			if len(stickySessionCookie) > 0 {
				options := map[string]interface{}{}
				options["sticky_session_cookie"] = stickySessionCookie
				payload["back_end"].(map[string]interface{})["options"] = options
			}
		}

		if payload["back_end"].(map[string]interface{})["options"] == nil && (input.ReverseProxyBackend.Options.HealthCheck.Path != nil ||
			input.ReverseProxyBackend.Options.HealthCheck.Scheme != nil ||
			input.ReverseProxyBackend.Options.HealthCheck.Port != nil ||
			input.ReverseProxyBackend.Options.HealthCheck.Interval != nil ||
			input.ReverseProxyBackend.Options.HealthCheck.Timeout != nil) {

			options := map[string]interface{}{}
			payload["back_end"].(map[string]interface{})["options"] = options
		}

		if payload["back_end"].(map[string]interface{})["options"] != nil && (input.ReverseProxyBackend.Options.HealthCheck.Path != nil ||
			input.ReverseProxyBackend.Options.HealthCheck.Scheme != nil ||
			input.ReverseProxyBackend.Options.HealthCheck.Port != nil ||
			input.ReverseProxyBackend.Options.HealthCheck.Interval != nil ||
			input.ReverseProxyBackend.Options.HealthCheck.Timeout != nil) {

			healthCheck := map[string]interface{}{}
			if input.ReverseProxyBackend.Options.HealthCheck.Path != nil {
				healthCheck["path"] = *input.ReverseProxyBackend.Options.HealthCheck.Path
			}
			if input.ReverseProxyBackend.Options.HealthCheck.Scheme != nil {
				healthCheck["scheme"] = *input.ReverseProxyBackend.Options.HealthCheck.Scheme
			}
			if input.ReverseProxyBackend.Options.HealthCheck.Port != nil {
				healthCheck["port"] = *input.ReverseProxyBackend.Options.HealthCheck.Port
			}
			if input.ReverseProxyBackend.Options.HealthCheck.Interval != nil {
				healthCheck["interval"] = *input.ReverseProxyBackend.Options.HealthCheck.Interval
			}
			if input.ReverseProxyBackend.Options.HealthCheck.Timeout != nil {
				healthCheck["timeout"] = *input.ReverseProxyBackend.Options.HealthCheck.Timeout
			}

			payload["back_end"].(map[string]interface{})["options"].(map[string]interface{})["health_check"] = healthCheck
		}
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return false, state, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonPayload))
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

	updatedState, err := rp.Read(ctx, state, input)
	if err != nil {
		return false, state, err
	}

	return success, updatedState, nil
}

func (ReverseProxy) Delete(ctx p.Context, state ReverseProxyState, input ReverseProxyArgs) (bool, error) {
	url := fmt.Sprintf("https://%s/api/1/customers/%s/cloudspaces/%s/ingress/reverse-proxies/%s", input.URL, input.CustomerID, input.CloudSpaceID, state.ReverseProxyID)
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

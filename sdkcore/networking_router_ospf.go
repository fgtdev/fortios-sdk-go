package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// JSONNetworkingRouterOSPF contains the parameters for Create and Update API function
type JSONNetworkingRouterOSPF struct {
	Routerid                    string                            `json:"router-id"`
	Defaultinformationoriginate string                            `json:"default-information-originate"`
	Area                        []JSONNetworkingRouterOSPFArea    `json:"area"`
	Network                     []JSONNetworkingRouterOSPFNetwork `json:"network"`
}

// JSONNetworkingRouterOSPFArea contains the parameters for Create and Update API function
type JSONNetworkingRouterOSPFArea struct {
	ID                                        string        `json:"id"`
	QOriginKey                                string        `json:"q_origin_key"`
	Shortcut                                  string        `json:"shortcut"`
	Authentication                            string        `json:"authentication"`
	DefaultCost                               int           `json:"default-cost"`
	NssaTranslatorRole                        string        `json:"nssa-translator-role"`
	StubType                                  string        `json:"stub-type"`
	Type                                      string        `json:"type"`
	NssaDefaultInformationOriginate           string        `json:"nssa-default-information-originate"`
	NssaDefaultInformationOriginateMetric     int           `json:"nssa-default-information-originate-metric"`
	NssaDefaultInformationOriginateMetricType string        `json:"nssa-default-information-originate-metric-type"`
	NssaRedistribution                        string        `json:"nssa-redistribution"`
	Range                                     []interface{} `json:"range"`
	VirtualLink                               []interface{} `json:"virtual-link"`
	FilterList                                []interface{} `json:"filter-list"`
}

// JSONNetworkingRouterOSPFNetwork contains the parameters for Create and Update API function
type JSONNetworkingRouterOSPFNetwork struct {
	ID         int    `json:"id"`
	QOriginKey int    `json:"q_origin_key"`
	Prefix     string `json:"prefix"`
	Area       string `json:"area"`
}

// JSONCreateNetworkingRouterOSPFOutput contains the output results for Create API function
type JSONCreateNetworkingRouterOSPFOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       float64 `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// JSONUpdateNetworkingRouterOSPFOutput contains the output results for Update API function
// Attention: The RESTful API changed the Mkey type from float64 in CREATE to string in UPDATE!
type JSONUpdateNetworkingRouterOSPFOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// CreateNetworkingRouterOSPF is currently empty, as the global OSPF config cannot be created.
// See the router - ospf chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateNetworkingRouterOSPF(params *JSONNetworkingRouterOSPF) (output *JSONCreateNetworkingRouterOSPFOutput, err error) {

	return
}

// UpdateNetworkingRouterOSPF API operation for FortiOS updates the global ospf configuration.
// Returns error for service API and SDK errors.
// See the router - ospf chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateNetworkingRouterOSPF(params *JSONNetworkingRouterOSPF, mkey string) (output *JSONUpdateNetworkingRouterOSPFOutput, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/ospf"
	// path += "/" + mkey
	output = &JSONUpdateNetworkingRouterOSPFOutput{}
	locJSON, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
		return
	}

	bytes := bytes.NewBuffer(locJSON)
	req := c.NewRequest(HTTPMethod, path, nil, bytes)
	err = req.Send()
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %s", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %s", err)
		return
	}
	log.Printf("FOS-fortios response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	req.HTTPResponse.Body.Close()

	if result != nil {
		if result["vdom"] != nil {
			output.Vdom = result["vdom"].(string)
		}
		if result["mkey"] != nil {
			output.Mkey = result["mkey"].(string)
		}
		if result["status"] != nil {
			if result["status"] != "success" {
				if result["error"] != nil {
					err = fmt.Errorf("status is %s and error no is %.0f", result["status"], result["error"])
				} else {
					err = fmt.Errorf("status is %s and error no is not found", result["status"])
				}

				if result["http_status"] != nil {
					err = fmt.Errorf("%s and http_status no is %.0f", err, result["http_status"])
				} else {
					err = fmt.Errorf("%s and and http_status no is not found", err)
				}

				return
			}
			output.Status = result["status"].(string)
		} else {
			err = fmt.Errorf("cannot get status from the response")
			return
		}
		if result["http_status"] != nil {
			output.HTTPStatus = result["http_status"].(float64)
		}
	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}

// DeleteNetworkingRouterOSPF API operation for FortiOS does not do anything, since global OSPF config cannot be deleted.
// Consider at some point resetting global values to default instead of doing nothing.
// See the router - ospf chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteNetworkingRouterOSPF(mkey string) (err error) {

	return
}

// ReadNetworkingRouterOSPF API operation for FortiOS reads global OSPF settings.
// Returns error for service API and SDK errors.
// See the router - ospf chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadNetworkingRouterOSPF(mkey string) (output *JSONNetworkingRouterOSPF, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/ospf"
	//path += "/" + mkey

	output = &JSONNetworkingRouterOSPF{}

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	err = req.Send()
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %s", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %s", err)
		return
	}
	log.Printf("FOS-fortios reading response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	req.HTTPResponse.Body.Close()

	if result != nil {
		if result["http_status"] == nil {
			err = fmt.Errorf("cannot get http_status from the response")
			return
		}

		if result["http_status"] == 404.0 {
			output = nil
			return
		}

		if result["status"] == nil {
			err = fmt.Errorf("cannot get status from the response")
			return
		}

		if result["status"] != "success" {
			if result["error"] != nil {
				err = fmt.Errorf("status is %s and error no is %.0f", result["status"], result["error"])
			} else {
				err = fmt.Errorf("status is %s and error no is not found", result["status"])
			}

			if result["http_status"] != nil {
				err = fmt.Errorf("%s and http_status no is %.0f", err, result["http_status"])
			} else {
				err = fmt.Errorf("%s and and http_status no is not found", err)
			}

			return
		}

		mapTmp := (result["results"].(interface{})).(map[string]interface{})

		if mapTmp == nil {
			err = fmt.Errorf("cannot get the results from the response")
			return
		}

	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}

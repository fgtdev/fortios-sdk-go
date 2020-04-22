package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fgtdev/fortios-sdk-go/util"
	"io/ioutil"
	"log"
	"strconv"
)

// JSONNetworkingInterfaceDHCPIPReserve contains the IP Reserve parameters for Create and Update API function
type JSONNetworkingInterfaceDHCPIpRes struct {
	ID          int    `json:"id,omitempty"`
	QOriginKey  int    `json:"q_origin_key,omitempty"`
	Type        string `json:"type,omitempty"`
	IP          string `json:"ip"`
	Mac         string `json:"mac"`
	Action      string `json:"action,omitempty"`
	Description string `json:"description,omitempty"`
	Vdom        string `json:"vdom,omitempty"`
	//	Skey        string `json:"skey,omitempty"`
	//	Mkey        string `json:"mkey,omitempty"`
}

type JSONNetworkingInterfaceDHCPResResult struct {
	Vdom            string `json:"vdom,omitempty"`
	Mkey            string `json:"mkey"`
	Skey            string `json:"skey"`
	Status          string `json:"status"`
	HTTPStatus      string `json:"http_status,omitempty"`
	Version         string `json:"version,omitempty"`
	RevisionChanged bool   `json:"revision_changed,omitempty"`
	Error           string `json:"error"`
}

// CreateUpdateNetworkInterfaceDHCIPPReservation API operation for FortiOS
func (c *FortiSDKClient) CreateNetworkInterfaceDHCPIpPRes(input *JSONNetworkingInterfaceDHCPIpRes, skey string) (output *JSONNetworkingInterfaceDHCPResResult, err error) {
	logPrefix := "CreateUpdateNetworkInterfaceDHCIPPReservation - "
	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system.dhcp/server"
	path += "/" + EscapeURLString(skey)
	path += "/reserved-address"

	output = &JSONNetworkingInterfaceDHCPResResult{}

	locJSON, err := json.Marshal(input)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Printf(logPrefix+"%s: %s", HTTPMethod, string(locJSON))

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

	log.Printf(logPrefix+"PATH: %s", path)
	log.Printf(logPrefix+"FortiOS Response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	req.HTTPResponse.Body.Close()

	if result != nil {
		if result["vdom"] != nil {
			output.Vdom = result["vdom"].(string)
		}
		if result["mkey"] != nil {
			// Note that Fortios is inconsistent in how mkey is reported
			if fmt.Sprintf("%T", result["mkey"]) == "float64" {
				output.Mkey = fmt.Sprintf("%.0f", result["mkey"].(float64))
			} else {
				output.Mkey = result["mkey"].(string)
			}

		}
		if result["status"] != nil {
			if result["status"] != "success" {
				if result["error"] != nil {
					errorCode := fmt.Sprintf("%.0f", result["error"])
					switch errorCode {
					case "-3":
						err = fmt.Errorf("Invalid Interface, no such interface")
					case "-8":
						err = fmt.Errorf("IP out of dhcp server range.")
					case "-526":
						err = fmt.Errorf("DHCP Server already attached to interface")
						// Even if create fail a server config is created in the FG and mkey is increased.
					case "-15":
						err = fmt.Errorf("IP or MAC already assigned to reservation, %s", errorCode)
					default:
						err = fmt.Errorf("status is %s and error no is %.0f", result["status"], result["error"])
					}
				} else {
					err = fmt.Errorf("status is %s and error no is not found", result["status"])
				}

				if result["http_status"] != nil {
					err = fmt.Errorf("%s, details: %s", err, util.HttpStatus2Str(int(result["http_status"].(float64))))
				} else {
					err = fmt.Errorf("%s, and http_status no is not found", err)
				}

				return
			}
			output.Status = result["status"].(string)
		} else {
			err = fmt.Errorf("cannot get status from the response")
			return
		}
		if result["http_status"] != nil {
			output.HTTPStatus = fmt.Sprintf("%.0f", result["http_status"].(float64))
		}
	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}

func (c *FortiSDKClient) UpdateNetworkInterfaceDHCPIpPRes(input *JSONNetworkingInterfaceDHCPIpRes, skey string, mkey string) (output *JSONNetworkingInterfaceDHCPResResult, err error) {
	logPrefix := "UpdateNetworkInterfaceDHCPIpPRes - "
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.dhcp/server"
	path += "/" + EscapeURLString(skey)
	path += "/reserved-address/" + EscapeURLString(mkey)

	output = &JSONNetworkingInterfaceDHCPResResult{}

	locJSON, err := json.Marshal(input)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Printf(logPrefix+"%s: %s", HTTPMethod, string(locJSON))

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

	log.Printf(logPrefix+"PATH: %s", path)
	log.Printf(logPrefix+"FortiOS Response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	req.HTTPResponse.Body.Close()

	if result != nil {
		if result["vdom"] != nil {
			output.Vdom = result["vdom"].(string)
		}
		if result["mkey"] != nil {
			// Note that Fortios is inconsistent in how mkey is reported
			if fmt.Sprintf("%T", result["mkey"]) == "float64" {
				output.Mkey = fmt.Sprintf("%.0f", result["mkey"].(float64))
			} else {
				output.Mkey = result["mkey"].(string)
			}
		}
		if result["status"] != nil {
			if result["status"] != "success" {
				if result["error"] != nil {
					errorCode := fmt.Sprintf("%.0f", result["error"])
					switch errorCode {
					case "-3":
						err = fmt.Errorf("Invalid Interface, no such interface")
					case "-8":
						err = fmt.Errorf("IP out of dhcp server range.")
					case "-15":
						err = fmt.Errorf("IP or MAC already assigned to reservation, %s", errorCode)
					case "-526":
						err = fmt.Errorf("DHCP Server already attached to interface")
						// Even if create fail a server config is created in the FG and mkey is increased.
					default:
						err = fmt.Errorf("status is %s and error no is %.0f", result["status"], result["error"])
					}
				} else {
					err = fmt.Errorf("status is %s and error no is not found", result["status"])
				}

				if result["http_status"] != nil {
					err = fmt.Errorf("%s, details: %s", err, util.HttpStatus2Str(int(result["http_status"].(float64))))
				} else {
					err = fmt.Errorf("%s, and http_status no is not found", err)
				}

				return
			}
			output.Status = result["status"].(string)
		} else {
			err = fmt.Errorf("cannot get status from the response")
			return
		}
		if result["http_status"] != nil {
			output.HTTPStatus = fmt.Sprintf("%.0f", result["http_status"].(float64))
		}
	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}

// DeleteNetworkingInterfaceDHCPIpRes API operation for FortiOS
func (c *FortiSDKClient) DeleteNetworkingInterfaceDHCPIpRes(skey string, mkey string) (output *JSONNetworkingInterfaceDHCPResult, err error) {
	logPrefix := "DeleteNetworkingInterfaceDHCPIpRes - "
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system.dhcp/server"
	path += "/" + EscapeURLString(skey)
	path += "/reserved-address/" + EscapeURLString(mkey)

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

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	req.HTTPResponse.Body.Close()

	log.Printf(logPrefix+"Path called %s", path)
	log.Printf(logPrefix+"FortiOS response: %s", string(body))

	if result != nil {
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
				err = fmt.Errorf("%s, details: %s", err, util.HttpStatus2Str(int(result["http_status"].(float64))))
			} else {
				err = fmt.Errorf("%s, and http_status no is not found", err)
			}

			return
		}

	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return

}

// ReadNetworkingInterfaceDHCPIpRes API operation for FortiOS
func (c *FortiSDKClient) ReadNetworkingInterfaceDHCPIpRes(skey string, mkey string) (output *JSONNetworkingInterfaceDHCPIpRes, err error) {
	logPrefix := "ReadNetworkingInterfaceDHCPIpRes - "
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.dhcp/server"
	path += "/" + EscapeURLString(skey)
	path += "/reserved-address/" + EscapeURLString(mkey)

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

	log.Printf(logPrefix+"PATH: %s", path)
	log.Printf(logPrefix+"FortiOS Response: %s", string(body))

	output = &JSONNetworkingInterfaceDHCPIpRes{}

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	req.HTTPResponse.Body.Close()

	if result != nil {
		if result["vdom"] != nil {
			output.Vdom = result["vdom"].(string)
		}
		if result["http_status"] == nil {
			err = fmt.Errorf("cannot get http_status from the response")
			return
		}

		if result["http_status"] == 404.0 {
			log.Printf(logPrefix+"DHCP Reservaton not found for mkey %s", mkey)
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
				err = fmt.Errorf("%s, details: %s", err, util.HttpStatus2Str(int(result["http_status"].(float64))))
			} else {
				err = fmt.Errorf("%s, and http_status no is not found", err)
			}

			return
		}

		mapTmp := (result["results"].([]interface{}))[0].(map[string]interface{})

		if mapTmp == nil {
			err = fmt.Errorf("cannot get the results from the response")
			return
		}

		// Result to json string
		out, err := json.Marshal(mapTmp)
		if err != nil {
			err = fmt.Errorf("cannot convert the results from the response")
			return nil, err
		}

		// json string to struct
		json.Unmarshal([]byte(out), output)

	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	//log.Printf(logPrefix+"Read output: %s", path)

	return
}

func (c *FortiSDKClient) NetworkingInterfaceDHCPSrvByName(intfName string) (id string, err error) {
	data, err := c.ReadNetworkingInterfaceDHCPServers()
	if err != nil {
		fmt.Println(err.Error())
	}
	servers := *data

	for _, server := range servers {
		if server.Interface == intfName {
			return strconv.Itoa(server.ID), nil
		}
	}
	err = fmt.Errorf("could not find dhcp server on interface, %s", intfName)
	return
}

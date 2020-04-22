package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fgtdev/fortios-sdk-go/util"
	"io/ioutil"
	"log"
)

// JSONNetworkInterfaceDHCP contains the parameters for Create and Update API function
type JSONNetworkingInterfaceDHCP struct {
	ID                        int                                  `json:"id,omitempty"`
	QOriginKey                int                                  `json:"q_origin_key,omitempty"`
	Status                    string                               `json:"status,omitempty"`
	LeaseTime                 int                                  `json:"lease-time,omitempty"`
	MacACLDefaultAction       string                               `json:"mac-acl-default-action,omitempty"`
	ForticlientOnNetStatus    string                               `json:"forticlient-on-net-status,omitempty"`
	DNSService                string                               `json:"dns-service,omitempty"`
	DNSServer1                string                               `json:"dns-server1,omitempty"`
	DNSServer2                string                               `json:"dns-server2,omitempty"`
	DNSServer3                string                               `json:"dns-server3,omitempty"`
	DNSServer4                string                               `json:"dns-server4,omitempty"`
	WifiAcService             string                               `json:"wifi-ac-service,omitempty"`
	WifiAc1                   string                               `json:"wifi-ac1,omitempty"`
	WifiAc2                   string                               `json:"wifi-ac2,omitempty"`
	WifiAc3                   string                               `json:"wifi-ac3,omitempty"`
	NtpService                string                               `json:"ntp-service,omitempty"`
	NtpServer1                string                               `json:"ntp-server1,omitempty"`
	NtpServer2                string                               `json:"ntp-server2,omitempty"`
	NtpServer3                string                               `json:"ntp-server3,omitempty"`
	Domain                    string                               `json:"domain,omitempty"`
	WinsServer1               string                               `json:"wins-server1,omitempty"`
	WinsServer2               string                               `json:"wins-server2,omitempty"`
	DefaultGateway            string                               `json:"default-gateway,omitempty"`
	NextServer                string                               `json:"next-server,omitempty"`
	Netmask                   string                               `json:"netmask,omitempty"`
	Interface                 string                               `json:"interface,omitempty"`
	IPRange                   []JSONNetworkingInterfaceDHCPIPRange `json:"ip-range,omitempty"`
	TimezoneOption            string                               `json:"timezone-option,omitempty"`
	Timezone                  string                               `json:"timezone,omitempty"`
	TftpServer                []interface{}                        `json:"tftp-server,omitempty"`
	Filename                  string                               `json:"filename,omitempty"`
	Options                   []JSONNetworkingInterfaceDHCPOptions `json:"options,omitempty"`
	ServerType                string                               `json:"server-type,omitempty"`
	IPMode                    string                               `json:"ip-mode,omitempty"`
	ConflictedIPTimeout       int                                  `json:"conflicted-ip-timeout,omitempty"`
	IpsecLeaseHold            int                                  `json:"ipsec-lease-hold,omitempty"`
	AutoConfiguration         string                               `json:"auto-configuration,omitempty"`
	DhcpSettingsFromFortiipam string                               `json:"dhcp-settings-from-fortiipam,omitempty"`
	AutoManagedStatus         string                               `json:"auto-managed-status,omitempty"`
	DdnsUpdate                string                               `json:"ddns-update,omitempty"`
	DdnsUpdateOverride        string                               `json:"ddns-update-override,omitempty"`
	DdnsServerIP              string                               `json:"ddns-server-ip,omitempty"`
	DdnsZone                  string                               `json:"ddns-zone,omitempty"`
	DdnsAuth                  string                               `json:"ddns-auth,omitempty"`
	DdnsKeyname               string                               `json:"ddns-keyname,omitempty"`
	DdnsKey                   string                               `json:"ddns-key,omitempty"`
	DdnsTTL                   int                                  `json:"ddns-ttl,omitempty"`
	VciMatch                  string                               `json:"vci-match,omitempty"`
	VciString                 []struct {
		VciString  string `json:"vci-string,omitempty"`
		QOriginKey string `json:"q_origin_key,omitempty"`
	} `json:"vci-string,omitempty"`
	ExcludeRange    []interface{}                      `json:"exclude-range,omitempty"`
	ReservedAddress []JSONNetworkingInterfaceDHCPIpRes `json:"reserved-address,omitempty"`
	Vdom            string                             `json:"vdom,omitempty"`
}

// JSONNetworkInterfaceDHCPIPRange contains the IP Range parameters for Create and Update API function
type JSONNetworkingInterfaceDHCPIPRange struct {
	ID         int    `json:"id,omitempty"`
	QOriginKey int    `json:"q_origin_key,omitempty"`
	StartIP    string `json:"start-ip"`
	EndIP      string `json:"end-ip"`
}

// JSONNetworkInterfaceDHCPOptions contains the parameters for Create and Update API function
type JSONNetworkingInterfaceDHCPOptions struct {
	ID         int    `json:"id,omitempty"`
	QOriginKey int    `json:"q_origin_key,omitempty"`
	Code       string `json:"code,omitempty"`
	Type       string `json:"type,omitempty"`
	Value      string `json:"value,omitempty"`
	IP         string `json:"ip,omitempty"`
}

// JSONCreateUpdateNetworkInterfaceDHCP contains the output results for Create API function
type JSONNetworkingInterfaceDHCPResult struct {
	Vdom            string `json:"vdom,omitempty"`
	Mkey            string `json:"mkey"`
	Status          string `json:"status"`
	HTTPStatus      string `json:"http_status,omitempty"`
	Version         string `json:"version,omitempty"`
	RevisionChanged bool   `json:"revision_changed,omitempty"`
	Error           string `json:"error"`
}

// CreateNetworkingInterfaceDHCP API operation for FortiOS
func (c *FortiSDKClient) CreateNetworkingInterfaceDHCP(params *JSONNetworkingInterfaceDHCP) (output *JSONNetworkingInterfaceDHCPResult, err error) {
	logPrefix := "CreateNetworkingInterfaceDHCP - "
	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system.dhcp/server"

	// Check if there is already a server attached to interface
	// Create will fail with error but still creat new config if interface already associated with a config., so we don't post the request if interface already has a config
	data, err := c.ReadNetworkingInterfaceDHCPServers()
	if err != nil {
		fmt.Println(err.Error())
	}
	servers := *data

	for _, server := range servers {
		if server.Interface == params.Interface {
			err = fmt.Errorf("DHCP Server already attached to interface!")
			return
		}
	}
	// End Check

	output = &JSONNetworkingInterfaceDHCPResult{}
	locJSON, err := json.Marshal(params)
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
					case "-526":
						err = fmt.Errorf("DHCP Server already attached to interface")
						// Even if create fail a server config is created. Check func added to avoid getting here
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

// UpdateNetworkingInterfaceDHCP API operation for FortiOS
// Returns the execution result when the request executes successfully.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) UpdateNetworkingInterfaceDHCP(input *JSONNetworkingInterfaceDHCP, mkey string) (output *JSONNetworkingInterfaceDHCPResult, err error) {
	logPrefix := "UpdateNetworkingInterfaceDHCP - "
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.dhcp/server"
	path += "/" + EscapeURLString(mkey)

	// Check if config exists
	_, err = c.ReadNetworkingInterfaceDHCP(mkey)
	if err != nil {
		return
	}

	output = &JSONNetworkingInterfaceDHCPResult{}
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
					case "-526":
						err = fmt.Errorf("DHCP Server already attached to interface")

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

// DeleteNetworkingInterfaceDHCP API operation for FortiOS
func (c *FortiSDKClient) DeleteNetworkingInterfaceDHCP(mkey string) (err error) {
	logPrefix := "DeleteNetworkingInterfaceDHCP - "
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system.dhcp/server"
	path += "/" + EscapeURLString(mkey)

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

// ReadNetworkingInterfaceDHCP API operation for FortiOS gets the dhcp server setting.
func (c *FortiSDKClient) ReadNetworkingInterfaceDHCP(mkey string) (output *JSONNetworkingInterfaceDHCP, err error) {
	logPrefix := "ReadNetworkingInterfaceDHCP - "

	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.dhcp/server"
	path += "/" + EscapeURLString(mkey)

	output = &JSONNetworkingInterfaceDHCP{}

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

	log.Printf(logPrefix+"Path called %s", path)
	log.Printf(logPrefix+"FortiOS response: %s", string(body))

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
			log.Printf(logPrefix+"DHCP Server config not found for mkey %s", mkey)
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

	log.Printf(logPrefix+"Read output: %s", path)

	return
}

func (c *FortiSDKClient) ReadNetworkingInterfaceDHCPServers() (output *[]JSONNetworkingInterfaceDHCP, err error) {
	logPrefix := "ReadNetworkingInterfaceDHCPServers - "
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.dhcp/server"

	var data []JSONNetworkingInterfaceDHCP

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

	log.Printf(logPrefix+"Path called %s", path)
	log.Printf(logPrefix+"FortiOS response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	req.HTTPResponse.Body.Close()

	if result != nil {
		if result["http_status"] == nil {
			err = fmt.Errorf("cannot get http_status from the response")
			return
		}

		if result["http_status"] == 404.0 {
			err = fmt.Errorf("DHCP Server configs not found")
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

		for _, srv := range result["results"].([]interface{}) {

			mapTmp := srv.(map[string]interface{})

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

			var o JSONNetworkingInterfaceDHCP
			// json string to struct
			json.Unmarshal(out, &o)
			if err != nil {
				err = fmt.Errorf("cannot convert the results to object")
				return nil, err
			}

			data = append(data, o)

		}

	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return &data, nil
}

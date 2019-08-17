package maidwhiteclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

const SCREEN_TYPE_SMART_PLUG ScreenType = "smartplug"
const SCREEN_TYPE_THERMOSTAT ScreenType = "thermostat"
const SCREEN_TYPE_ROLLING_DOOR ScreenType = "rolling_door"
const SCREEN_TYPE_HIDDEN ScreenType = "hidden"
const SCREEN_TYPE_PERMISSION ScreenType = "permission"

type ScreenType string

type Device struct {
	Id             int    `json:"id,omitempty"`
	DisplayName    string `json:"display_name,omitempty"`
	DisplayPicture string `json:"display_picture,omitempty"`
	RoomId         int    `json:"room_id,omitempty"`
	DriverName     string `json:"driver_name,omitempty"`
}

type DriverInfoItem struct {
	Parameters []string   `json:"parameters"`
	ScreenType ScreenType `json:"screen_type"`
}

var deviceInfoCache = map[string]*DriverInfoItem{}

func LoadDriverInfo() {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://www.tih.tw/2/driver-info"), nil)
	if err != nil {
		log.Println(err)
		return
	}
	dump, _ := httputil.DumpRequest(req, true)
	log.Println("req", string(dump))
	resp, err := http.DefaultClient.Do(req)
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println("received:", string(data))

	json.Unmarshal(data, &deviceInfoCache)

	// var v DeviceListResponse

	// json.Unmarshal(data, &v)
	// return v.Devices, nil
	// return nil
}

func GetScreenTypeByDriverName(driverName string) ScreenType {
	return deviceInfoCache[driverName].ScreenType

}

func RemoveHiddenDevice(list []*Device) []*Device {
	newList := []*Device{}
	for _, d := range list {
		st := deviceInfoCache[d.DriverName].ScreenType
		if st != SCREEN_TYPE_HIDDEN && st != SCREEN_TYPE_PERMISSION {
			nd := d
			newList = append(newList, nd)

		}
	}
	return newList
}

func FilterThermostatList(list []*Device) []*Device {
	newList := []*Device{}
	for _, d := range list {
		st := deviceInfoCache[d.DriverName].ScreenType
		if st == SCREEN_TYPE_THERMOSTAT {
			nd := d
			newList = append(newList, nd)
		}
	}
	return newList
}

type DeviceListResponse struct {
	Status  string    `json:"status,omitempty"`
	Devices []*Device `json:"devices,omitempty"`
}

func (c *Client) GetDevices() ([]*Device, error) {
	req, err := c.NewRequest("GET", "/devices", nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(req)
	data, err := ioutil.ReadAll(resp.Body)

	if Debug {
		fmt.Println("received:", string(data))
	}
	var v DeviceListResponse

	json.Unmarshal(data, &v)
	return v.Devices, nil

}

func (c *Client) SendDeviceControlRequest(d *Device, cmd map[string]string) error {
	u := url.Values{}
	for k, v := range cmd {
		u.Set(k, v)
	}
	if Debug {
		log.Println("xxx", u.Encode())
	}
	req, err := c.NewRequest("POST", fmt.Sprintf("/devices/%d", d.Id), strings.NewReader(u.Encode()))
	if err != nil {
		return err
	}
	dump, _ := httputil.DumpRequest(req, true)

	if Debug {
		log.Println("req", string(dump))
	}
	resp, err := c.client.Do(req)
	data, err := ioutil.ReadAll(resp.Body)

	if Debug {
		fmt.Println("received:", string(data))
	}
	// var v DeviceListResponse

	// json.Unmarshal(data, &v)
	// return v.Devices, nil
	return nil
}
func (c *Client) SendDeviceQueryRequest(d *Device) (map[string]interface{}, error) {
	req, err := c.NewRequest("GET", fmt.Sprintf("/devices/%d", d.Id), nil)
	if err != nil {
		return nil, err
	}
	dump, _ := httputil.DumpRequest(req, true)
	if Debug {
		log.Println("req", string(dump))
	}
	resp, err := c.client.Do(req)
	data, err := ioutil.ReadAll(resp.Body)
	if Debug {
		log.Println("received:", string(data))
	}
	response := map[string]interface{}{}
	json.Unmarshal(data, &response)
	// var v DeviceListResponse

	// json.Unmarshal(data, &v)
	// return v.Devices, nil
	return response, nil
}

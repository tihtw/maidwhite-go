package maidwhiteclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http/httputil"
	"net/url"
	"strings"
)

type Device struct {
	Id             int    `json:"id,omitempty"`
	DisplayName    string `json:"display_name,omitempty"`
	DisplayPicture string `json:"display_picture,omitempty"`
	RoomId         int    `json:"room_id,omitempty"`
	DriverName     string `json:"driver_name,omitempty"`
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
	fmt.Println("received:", string(data))
	var v DeviceListResponse

	json.Unmarshal(data, &v)
	return v.Devices, nil

}

func (c *Client) SendDeviceRequest(d *Device, cmd map[string]string) error {
	u := url.Values{}
	for k, v := range cmd {
		u.Set(k, v)
	}
	log.Println("xxx", u.Encode())
	req, err := c.NewRequest("POST", fmt.Sprintf("/devices/%d", d.Id), strings.NewReader(u.Encode()))
	if err != nil {
		return err
	}
	dump, _ := httputil.DumpRequest(req, true)
	log.Println("req", string(dump))
	resp, err := c.client.Do(req)
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println("received:", string(data))
	// var v DeviceListResponse

	// json.Unmarshal(data, &v)
	// return v.Devices, nil
	return nil
}

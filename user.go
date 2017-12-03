package maidwhiteclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type User struct {
	Id          int    `json:"id,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	Email       string `json:"email,omitempty"`
}

func (c *Client) GetUser() (*User, error) {
	req, err := c.NewRequest("GET", "/me", nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(req)
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println("received:", string(data))
	var v User

	json.Unmarshal(data, &v)
	return &v, nil

}

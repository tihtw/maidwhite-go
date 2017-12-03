package maidwhiteclient

import (
	"context"
	"golang.org/x/oauth2"
	"io"
	"net/http"
)

type Client struct {
	client       *http.Client
	accessToken  string
	oauth2Config oauth2.Config
}

type Token struct {
	*oauth2.Token
}

const (
	SCOPE_ALL = "*"
	BASE_URL  = "https://www.tih.tw/2"
)

//

func NewClient() *Client {
	return &Client{
		oauth2Config: oauth2.Config{
			ClientID:     "",
			ClientSecret: "",
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://www.tih.tw/o2/auth",
				TokenURL: "https://www.tih.tw/o2/token",
			},
			RedirectURL: "",
			Scopes:      []string{SCOPE_ALL},
		},
	}
}

func init() {
	oauth2.RegisterBrokenAuthHeaderProvider("https://www.tih.tw")

}

func (c *Client) SetClientID(clientID string) {
	c.oauth2Config.ClientID = clientID
}
func (c *Client) SetClientSecret(clientSecret string) {
	c.oauth2Config.ClientSecret = clientSecret
}
func (c *Client) SetRedirectURL(clientURL string) {
	c.oauth2Config.RedirectURL = clientURL
}

func (c *Client) SetAccessToken(t string) {
	c.accessToken = t
}

func (c *Client) AuthCodeURL(state string) string {
	return c.oauth2Config.AuthCodeURL(state)
	// oauth2.Config
}

func (c *Client) Exchange(ctx context.Context, code string) (*Token, error) {
	t, err := c.oauth2Config.Exchange(ctx, code)
	if t != nil {
		c.client = c.oauth2Config.Client(ctx, t)
	}
	return &Token{Token: t}, err
}
func (c *Client) NewRequest(method, urlStr string, body io.ReadWriter) (*http.Request, error) {
	req, err := http.NewRequest(method, BASE_URL+urlStr, body)
	if err != nil {
		return nil, err
	}

	// req.Header.Add("Authorization", "Bearer "+c.client.Client.AccessToken)
	return req, nil
}

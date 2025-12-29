package kiwoom

import (
	"fmt"
	"reflect"
)

type TokenResponse struct {
	ExpiresDt string `json:"expires_dt"`
	TokenType string `json:"token_type"`
	Token     string `json:"token"`
}

func (c *Client) FetchToken() error {
	res, err := c.CallEndpoint("au10001", nil, map[string]string{"grant_type": "client_credentials",
		"appkey":    c.AppKey,
		"secretkey": c.AppSecret,
	})

	if err != nil {
		return err
	}

	v := reflect.ValueOf(res).Elem()
	retMsg := v.FieldByName("ReturnMessage").String()
	retCode := v.FieldByName("ReturnCode").Int()
	token := v.FieldByName("Token").String()

	if retCode != 0 {
		return fmt.Errorf("failed to fetch token: %s", retMsg)
	}

	ttl, err := ParseKiwoomTime(v.FieldByName("ExpiresDt").String())
	if err != nil {
		return err
	}

	c.AccessToken = token
	c.TokenTTL = ttl

	return nil
}

func (c *Client) RevokeToken() error {
	res, err := c.CallEndpoint("au10002", nil, map[string]string{"grant_type": "client_credentials",
		"appkey":    c.AppKey,
		"secretkey": c.AppSecret,
		"token":     c.AccessToken,
	})

	if err != nil {
		return err
	}

	v := reflect.ValueOf(res).Elem()
	retMsg := v.FieldByName("ReturnMessage").String()
	retCode := v.FieldByName("ReturnCode").Int()

	if retCode != 0 {
		return fmt.Errorf("failed to revoke token: %s", retMsg)
	}

	return nil
}

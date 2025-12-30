// package kiwoom

// import "fmt"

// // It seems kiwoom only uses POST method for REST API
// // Nice job done

// // Note: using auto-parse of resty v3 (https://resty.dev/docs/response-auto-parse/)
// func (c *Client) POST(endpoint string, header map[string]string, reqBody any, respBody any, errBody any) error {
// 	res, err := c.RestClient.R().SetHeaders(header).SetBody(reqBody).SetResult(&respBody).SetError(&errBody).Post(endpoint)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println(res.Result().(*respBody))
// 	return nil
// }

package kiwoom

import (
	"fmt"
	"reflect"
)

type Endpoint struct {
	Url        string
	ResultType reflect.Type
}

var Endpoints = map[string]Endpoint{
	// 접근토큰발급
	"au10001": {
		Url:        "/oauth2/token",
		ResultType: reflect.TypeFor[resAU10001](),
	},
	// 접근토큰폐기
	"au10002": {
		Url:        "/oauth2/revoke",
		ResultType: reflect.TypeFor[resAU10002](),
	},
	// 당일거래량상위요청
	"ka10030 ": {
		Url:        "/api/dostk/rkinfo",
		ResultType: reflect.TypeFor[resKA10030](),
	},
}

func (c *Client) CallEndpoint(tr string, header map[string]string, body any) (any, error) {
	ep, ok := Endpoints[tr]
	if !ok {
		return nil, fmt.Errorf("endpoint key not found: %s", tr)
	}

	resultPtr := reflect.New(ep.ResultType).Interface()

	resp, err := c.RestClient.R().
		SetHeader("Content-Type", "application/json").
		SetHeaders(header).
		SetBody(body).
		SetResult(resultPtr).
		Post(ep.Url)

	if err != nil {
		return nil, fmt.Errorf("resty request error: %w", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("api response error: %s (status: %d)", resp.String(), resp.StatusCode())
	}

	return resultPtr, nil
}

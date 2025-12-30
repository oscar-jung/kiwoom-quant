package kiwoom

import (
	"fmt"
	"reflect"
)

func (c *Client) GetVolumeLeaders() ([]string, error) {
	res, err := c.CallEndpoint("ka10030", map[string]string{
		"api-id":        "ka10030",
		"authorization": fmt.Sprintf("Bearer %s", c.AccessToken),
	}, map[string]any{
		"mrkt_tp": "000",
		"sort_tp": "1",
		// "mang_stk_incls": "16",
		"mang_stk_incls": "0",
		"crd_tp":         "0",
		"trde_qty_tp":    "0",
		"pric_tp":        "0",
		"trde_prica_tp":  "0",
		"mrkt_open_tp":   "0",
		"stex_tp":        "1",
	})

	if err != nil {
		return nil, err
	}

	v := reflect.ValueOf(res).Elem()
	retMsg := v.FieldByName("ReturnMessage").String()
	retCode := v.FieldByName("ReturnCode").Int()

	if retCode != 0 {
		return nil, fmt.Errorf("failed to fetch result: %s", retMsg)
	}

	leaders := []string{}
	results := v.FieldByName("TodayTradeQtyUpper")
	slice, ok := results.Interface().([]resTdyTrdeQtyUpper)
	if !ok {
		return nil, fmt.Errorf("failed to parse result list")
	}
	for _, v := range slice {
		leaders = append(leaders, v.StockCode)
	}

	return leaders, nil
}

func (c *Client) GetTurnoverRatioLeaders() ([]string, error) {
	leaders := []string{}
	return leaders, nil
}

func (c *Client) GetvalueLeaders() ([]string, error) {
	leaders := []string{}
	return leaders, nil
}

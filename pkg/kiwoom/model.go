package kiwoom

type resAU10001 struct {
	ExpiresDt     string `json:"expires_dt"`
	TokenType     string `json:"token_type"`
	Token         string `json:"token"`
	ReturnCode    int    `json:"return_code"`
	ReturnMessage string `json:"return_msg"`
}

type resAU10002 struct {
	ReturnCode    int    `json:"return_code"`
	ReturnMessage string `json:"return_msg"`
}

type resKA10030 struct {
	TodayTradeQtyUpper []resTdyTrdeQtyUpper `json:"tdy_trde_qty_upper"`
}

type resTdyTrdeQtyUpper struct {
	StockCode string `json:"stk_cd"`
	StockName string `json:"stk_nm"`
	CurrentPrice string `json:"cur_prc"`
}
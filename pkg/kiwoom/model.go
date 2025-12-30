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
	StockCode                 string  `json:"stk_cd"`
	StockName                 string  `json:"stk_nm"`
	CurrentPrice              string  `json:"cur_prc"`
	PredPreSig                string  `json:"pred_pre_sig"`
	PredPre                   string  `json:"pred_pre"`
	FluRatio                  string  `json:"flu_rt"`
	TradeQuantity             int     `json:"trde_qty"`
	PredRatio                 float64 `json:"pred_rt"`
	TradeTernRatio            float64 `json:"trde_tern_rt"`
	TradeAmount               int     `json:"trde_amt"`
	OpenMarketTradeQuantity   int     `json:"opmk_trde_qty"`
	OpenMarketPredRatio       float64 `json:"opmk_pred_rt"`
	OpenMarketTradeRatio      float64 `json:"opmk_trde_rt"`
	OpenMarketTradeAmount     int     `json:"opmk_trde_amt"`
	AfterMarketTradeQuantity  int     `json:"af_mkrt_trde_qty"`
	AfterMarketPredRatio      float64 `json:"af_mkrt_prde_rt"`
	AfterMarketTradeRatio     float64 `json:"af_mkrt_trde_rt"`
	AfterMarketTradeAmount    int     `json:"af_mkrt_trde_amt"`
	BeforeMarketTradeQuantity int     `json:"bf_mkrt_trde_qty"`
	BeforeMarketPredRatio     float64 `json:"bf_mkrt_pred_rt"`
	BeforeMarketTradeRatio    float64 `json:"bf_mkrt_trde_rt"`
	BeforeMarketTradeAmount   int     `json:"bf_mkrt_trde_amt"`
}

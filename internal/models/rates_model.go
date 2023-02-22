package models

type Currency string

const (
	NGN Currency = "NGN"
	USD Currency = "USD"
)

type Exchange struct {
	Currency Currency `json:"currency" binding:"required"`
	Amount   float64  `json:"amount" binding:"required"`
}

type Rates struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		PublicKey string `json:"publicKey"`
		Signature string `json:"signature"`
		Rates     struct {
			Btcngn struct {
				Rate float64 `json:"rate"`
				Key  string  `json:"key"`
			} `json:"BTCNGN"`
			Btcbusd struct {
				Rate int    `json:"rate"`
				Key  string `json:"key"`
			} `json:"BTCBUSD"`
			Dashbusd struct {
				Rate int    `json:"rate"`
				Key  string `json:"key"`
			} `json:"DASHBUSD"`
			Dashngn struct {
				Rate int    `json:"rate"`
				Key  string `json:"key"`
			} `json:"DASHNGN"`
			Ethngn struct {
				Rate float64 `json:"rate"`
				Key  string  `json:"key"`
			} `json:"ETHNGN"`
			Usdtngn struct {
				Rate float64 `json:"rate"`
				Key  string  `json:"key"`
			} `json:"USDTNGN"`
			TronUsdtngn struct {
				Rate float64 `json:"rate"`
				Key  string  `json:"key"`
			} `json:"TRON_USDTNGN"`
			Busdngn struct {
				Rate float64 `json:"rate"`
				Key  string  `json:"key"`
			} `json:"BUSDNGN"`
			Bnbngn struct {
				Rate float64 `json:"rate"`
				Key  string  `json:"key"`
			} `json:"BNBNGN"`
			Cusdngn struct {
				Rate float64 `json:"rate"`
				Key  string  `json:"key"`
			} `json:"CUSDNGN"`
			Btcbtc struct {
				Rate int    `json:"rate"`
				Key  string `json:"key"`
			} `json:"BTCBTC"`
			Dashdash struct {
				Rate int    `json:"rate"`
				Key  string `json:"key"`
			} `json:"DASHDASH"`
			TronUsdttronUsdt struct {
				Rate int    `json:"rate"`
				Key  string `json:"key"`
			} `json:"TRON_USDTTRON_USDT"`
			Busdbusd struct {
				Rate int    `json:"rate"`
				Key  string `json:"key"`
			} `json:"BUSDBUSD"`
			Cusdcusd struct {
				Rate int    `json:"rate"`
				Key  string `json:"key"`
			} `json:"CUSDCUSD"`
			Etheth struct {
				Rate int    `json:"rate"`
				Key  string `json:"key"`
			} `json:"ETHETH"`
			Bnbbnb struct {
				Rate int    `json:"rate"`
				Key  string `json:"key"`
			} `json:"BNBBNB"`
			Usdtusdt struct {
				Rate int    `json:"rate"`
				Key  string `json:"key"`
			} `json:"USDTUSDT"`
			Btcngn0 struct {
				Rate float64 `json:"rate"`
				Key  string  `json:"key"`
			} `json:"BTCNGN_"`
			Btcusd struct {
				Rate float64 `json:"rate"`
				Key  string  `json:"key"`
			} `json:"BTCUSD"`
			Ethusd struct {
				Rate float64 `json:"rate"`
				Key  string  `json:"key"`
			} `json:"ETHUSD"`
			Bnbusd struct {
				Rate float64 `json:"rate"`
				Key  string  `json:"key"`
			} `json:"BNBUSD"`
			Usdtusd struct {
				Rate int    `json:"rate"`
				Key  string `json:"key"`
			} `json:"USDTUSD"`
			Cusdusd struct {
				Rate float64 `json:"rate"`
				Key  string  `json:"key"`
			} `json:"CUSDUSD"`
			Busdusd struct {
				Rate int    `json:"rate"`
				Key  string `json:"key"`
			} `json:"BUSDUSD"`
			Btcusd0 struct {
				Rate float64 `json:"rate"`
				Key  string  `json:"key"`
			} `json:"BTCUSD_"`
			Ethusd0 struct {
				Rate float64 `json:"rate"`
				Key  string  `json:"key"`
			} `json:"ETHUSD_"`
			Bnbusd0 struct {
				Rate float64 `json:"rate"`
				Key  string  `json:"key"`
			} `json:"BNBUSD_"`
			Usdtusd0 struct {
				Rate int    `json:"rate"`
				Key  string `json:"key"`
			} `json:"USDTUSD_"`
			Cusdusd0 struct {
				Rate float64 `json:"rate"`
				Key  string  `json:"key"`
			} `json:"CUSDUSD_"`
			Busdusd0 struct {
				Rate int    `json:"rate"`
				Key  string `json:"key"`
			} `json:"BUSDUSD_"`
			TronUsdtusd struct {
				Rate int    `json:"rate"`
				Key  string `json:"key"`
			} `json:"TRON_USDTUSD"`
			TronUsdtusd0 struct {
				Rate int    `json:"rate"`
				Key  string `json:"key"`
			} `json:"TRON_USDTUSD_"`
			Bnbngn0 struct {
				Rate float64 `json:"rate"`
				Key  string  `json:"key"`
			} `json:"BNBNGN_"`
			Ethngn0 struct {
				Rate float64 `json:"rate"`
				Key  string  `json:"key"`
			} `json:"ETHNGN_"`
			Usdtngn0 struct {
				Rate float64 `json:"rate"`
				Key  string  `json:"key"`
			} `json:"USDTNGN_"`
			Busdngn0 struct {
				Rate float64 `json:"rate"`
				Key  string  `json:"key"`
			} `json:"BUSDNGN_"`
			Cusdngn0 struct {
				Rate float64 `json:"rate"`
				Key  string  `json:"key"`
			} `json:"CUSDNGN_"`
			TronUsdtngn0 struct {
				Rate float64 `json:"rate"`
				Key  string  `json:"key"`
			} `json:"TRON_USDTNGN_"`
			Usdcusd struct {
				Rate int    `json:"rate"`
				Key  string `json:"key"`
			} `json:"USDCUSD"`
			Usdcusd0 struct {
				Rate int    `json:"rate"`
				Key  string `json:"key"`
			} `json:"USDCUSD_"`
			Usdcngn struct {
				Rate float64 `json:"rate"`
				Key  string  `json:"key"`
			} `json:"USDCNGN"`
			Usdcngn0 struct {
				Rate float64 `json:"rate"`
				Key  string  `json:"key"`
			} `json:"USDCNGN_"`
			Maticusd struct {
				Rate float64 `json:"rate"`
				Key  string  `json:"key"`
			} `json:"MATICUSD"`
			Maticusd0 struct {
				Rate float64 `json:"rate"`
				Key  string  `json:"key"`
			} `json:"MATICUSD_"`
			Maticngn struct {
				Rate string `json:"rate"`
				Key  string `json:"key"`
			} `json:"MATICNGN_"`
			Maticngn0 struct {
				Rate float64 `json:"rate"`
				Key  string  `json:"key"`
			} `json:"MATICNGN"`
			Celousd struct {
				Rate float64 `json:"rate"`
				Key  string  `json:"key"`
			} `json:"CELOUSD"`
			Celousd0 struct {
				Rate float64 `json:"rate"`
				Key  string  `json:"key"`
			} `json:"CELOUSD_"`
			Celongn struct {
				Rate string `json:"rate"`
				Key  string `json:"key"`
			} `json:"CELONGN_"`
			Celongn0 struct {
				Rate float64 `json:"rate"`
				Key  string  `json:"key"`
			} `json:"CELONGN"`
			Ethbtc struct {
				Rate float64 `json:"rate"`
				Key  string  `json:"key"`
			} `json:"ETHBTC"`
			Ethbtc0 struct {
				Rate float64 `json:"rate"`
				Key  string  `json:"key"`
			} `json:"ETHBTC_"`
			Bnbbtc struct {
				Rate float64 `json:"rate"`
				Key  string  `json:"key"`
			} `json:"BNBBTC"`
			Bnbbtc0 struct {
				Rate float64 `json:"rate"`
				Key  string  `json:"key"`
			} `json:"BNBBTC_"`
			Bnbeth struct {
				Rate float64 `json:"rate"`
				Key  string  `json:"key"`
			} `json:"BNBETH"`
			Bnbeth0 struct {
				Rate float64 `json:"rate"`
				Key  string  `json:"key"`
			} `json:"BNBETH_"`
		} `json:"rates"`
	} `json:"data"`
}

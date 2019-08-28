package protocol

type Basicdata struct {
	CapitalValues  float64 `json:"CapitalValues,omitempty"`  //总市值
	CapitalAmount  float64 `json:"CapitalAmount,omitempty"`  //总股本
	FloatingValues float64 `json:"FloatingValues,omitempty"` //流通值
	FloatingStocks float64 `json:"FloatingStocks,omitempty"` //流通股
	PER            float64 `json:"PER,omitempty"`            //市盈率
	PBR            float64 `json:"PBR,omitempty"`            //市净率
	EPS            float64 `json:"EPS,omitempty"`            //每股收益
}

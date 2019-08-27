package protocol

type Basicdata struct {
	CapitalValues  int64 `json:"CapitalValues,omitempty"`
	CapitalAmount  int64 `json:"CapitalAmount,omitempty"`
	FloatingValues int64 `json:"FloatingValues,omitempty"`
	FloatingStocks int64 `json:"FloatingStocks,omitempty"`
	PER            int64 `json:"PER,omitempty"`
	PBR            int64 `json:"PBR,omitempty"`
	EPS            int64 `json:"EPS,omitempty"`
}

package ctype

import "encoding/json"

type SignType int

const (
	SignQQ SignType = 1 //QQ
	SignWX SignType = 2 // 微信
	SignWB SignType = 3 // 微博

)

func (s SignType) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s SignType) String() string {
	var str string
	switch s {
	case SignQQ:
		str = "QQ"
	case SignWX:
		str = "WX"
	case SignWB:
		str = "WB"
	default:
		str = "Other"
	}
	return str
}

package types

type RespCmd struct {
	Cmd
	Code int `json:"code"`
}

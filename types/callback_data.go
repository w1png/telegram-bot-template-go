package types

import "encoding/json"

type StateBackData struct {
	Destination string `json:"dest"`
	Data        any    `json:"data"`
}

type CallbackData struct {
	Call string `json:"c"`
	Data any    `json:"d"`
}

func NewCallbackData(call string, data any) *CallbackData {
	return &CallbackData{
		Call: call,
		Data: data,
	}
}

func (c *CallbackData) Marshal() (string, error) {
	bytes, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func UnmarshalCallbackData(callback_data string) (*CallbackData, error) {
	var callback CallbackData
	err := json.Unmarshal([]byte(callback_data), &callback)
	if err != nil {
		return nil, err
	}
	return &callback, nil
}

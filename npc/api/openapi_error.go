package api

import (
	"encoding/json"
	"fmt"
)

type ApiError struct {
	Status  int    `json:"-"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Cause   error  `json:"-"`
}

func (e *ApiError) UnmarshalJSON(data []byte) error {
	target := make(map[string]interface{})
	if err := json.Unmarshal(data, &target); err != nil {
		return err
	}
	if code, ok := findInMap(target, "code"); ok {
		switch code := code.(type) {
		case float64:
			e.Code = fmt.Sprint(int64(code))
		default:
			e.Code = fmt.Sprint(code)
		}
	}
	if message, ok := findInMap(target, "message", "msg"); ok {
		e.Message = fmt.Sprint(message)
	}
	return nil
}

func findInMap(target map[string]interface{}, fields ...string) (interface{}, bool) {
	for _, f := range fields {
		if val, ok := target[f]; ok {
			return val, true
		}
	}
	return nil, false
}

func (e *ApiError) Error() string {
	if e == nil {
		return "nil/ApiError"
	}
	if e.Cause != nil {
		return fmt.Sprintf("HTTP-%v [code=%v]: %v\n- Cause by %v",
			e.Status, e.Code, e.Message,
			e.Cause.Error())
	}
	return fmt.Sprintf("HTTP-%v [code=%v]: %v", e.Status, e.Code, e.Message)
}

package api

import (
	"encoding/json"
	"log"
	"testing"
)

func TestUnmarshalApiError(t *testing.T) {
	e1, e2, e3 := ApiError{}, ApiError{}, ApiError{}
	jsonE1, jsonE2, jsonE3 :=
		[]byte(`{"message":"message 1.","code":4000001}`),
		[]byte(`{"code":4030002,"msg":"message 2."}`),
		[]byte(`{"message":"message 3.","code":"4000003"}`)

	if err := json.Unmarshal(jsonE1, &e1); err != nil || e1.Code != "4000001" || e1.Message != "message 1." {
		t.Error("[ERROR1]", err)
	}
	log.Print("[OK1]", e1.Error())
	if err := json.Unmarshal(jsonE2, &e2); err != nil || e2.Code != "4030002" || e2.Message != "message 2." {
		t.Error("[ERROR2]", err)
	}
	log.Print("[OK2]", e2.Error())
	if err := json.Unmarshal(jsonE3, &e3); err != nil || e3.Code != "4000003" || e3.Message != "message 3." {
		t.Error("[ERROR3]", err)
	}
	log.Print("[OK3]", e3.Error())
}

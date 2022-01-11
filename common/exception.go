package common

import (
    "encoding/json"
    "log"
)

type Exception struct {
    Code int    `json:"code"`
    Msg  string `json:"msg"`
}

func ErrorMsgException(code int, msg string) string {
    e := Exception{Code: code, Msg: msg}
    data, err := json.Marshal(e)
    if err != nil {
        log.Panic(err)
    }

    return string(data)
}

func ErrorCodeException(code int) string {
    e := Exception{Code: code, Msg: ""}
    data, err := json.Marshal(e)
    if err != nil {
        log.Panic(err)
    }

    return string(data)
}

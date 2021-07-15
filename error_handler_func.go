package binance_api_go

import (
	"encoding/json"
	"fake_buy_cli/src/cmn"
	"log"
	"net/http"
	"strings"
)

func TradeBotErrorCheck(body []byte, res *http.Response, resErr, bodyErr error) error {
	if resErr != nil {
		log.Print(resErr)
		return resErr

	} else if bodyErr != nil {
		log.Print(bodyErr)
		return bodyErr

	} else if res.StatusCode == 429 {
		return &cmn.FakeBuyError{
			Type:    cmn.HttpErr,
			Code:    res.StatusCode,
			Message: res.Status,
		}
	} else if body != nil {
		if strings.Contains(string(body), "msg") {
			var bodyAnswer BodyAnswer

			jsonErr := json.Unmarshal(body, &bodyAnswer)
			if jsonErr != nil {
				return jsonErr
			}

			if bodyAnswer.Code < 0 {
				return &cmn.FakeBuyError{
					Type:    cmn.BinanceErr,
					Code:    bodyAnswer.Code,
					Message: bodyAnswer.Msg,
				}
			}
		} else {
			return nil
		}

	} else if res != nil {
		if res.StatusCode != 200 {
			return &cmn.FakeBuyError{
				Type:    cmn.HttpErr,
				Code:    res.StatusCode,
				Message: res.Status,
			}
		}
	}
	return &cmn.FakeBuyError{Type: cmn.BinanceErr, Message: "Body is nil"}
}

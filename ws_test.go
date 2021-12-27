package bittrex

import (
	"errors"
	"testing"
	"time"
)

func TestBittrexSubscribeOrderBook(t *testing.T) {
	bt := New("a32fa86b865e451cac802596a8092bda", "01f10c85c06f4bffa9107a66a6c2ccaa")
	ch := make(chan ExchangeState, 16)
	errCh := make(chan error)
	go func() {
		var haveInit bool
		var msgNum int
		for st := range ch {
			haveInit = haveInit || st.Initial
			msgNum++
			if msgNum >= 3 {
				break
			}
		}
		if haveInit {
			errCh <- nil
		} else {
			errCh <- errors.New("no initial message")
		}
	}()
	go func() {
		errCh <- bt.SubscribeExchangeUpdate("USDT-BTC", ch, nil)
	}()
	select {
	case <-time.After(time.Second * 6):
		t.Error("timeout")
	case err := <-errCh:
		if err != nil {
			t.Error(err)
		}
	}
}

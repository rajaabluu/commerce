package config

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func (c *Config) NewPaymentLib() snap.Client {
	var s snap.Client
	s.New(c.PaymentLib.ServerKey, midtrans.Sandbox)
	return s
}

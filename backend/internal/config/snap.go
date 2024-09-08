package config

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"github.com/spf13/viper"
)

func NewSnapClient(viper *viper.Viper) *snap.Client {
	client := new(snap.Client)
	client.New(viper.GetString("midtrans.key.server"), midtrans.Sandbox)
	return client
}

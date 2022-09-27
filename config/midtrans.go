package config

import (
	midtrans "github.com/midtrans/midtrans-go"
	coreMidtrans "github.com/midtrans/midtrans-go/coreapi"
  snapMidtrans  "github.com/midtrans/midtrans-go/snap"
  irisMidtrans  "github.com/midtrans/midtrans-go/iris"
)


type MidtransConfig struct {
	Snap snapMidtrans.Client
	Core coreMidtrans.Client
	Iris irisMidtrans.Client
	ServerKey string
	ClientKey string
}


func InitMidtrans(conf Config) MidtransConfig {
	midtrans.ServerKey = conf.Midtrans.ServerKey
	midtrans.Environment = midtrans.Sandbox
	
	// var midtrans.ClientKey = conf.Midtrans.ClientKey
	return MidtransConfig{
		ServerKey: conf.Midtrans.ServerKey,
		ClientKey: conf.Midtrans.ClientKey,
		Snap: snapMidtrans.Client{ServerKey: conf.Midtrans.ServerKey, Env: midtrans.Sandbox},
		Core: coreMidtrans.Client{ServerKey: conf.Midtrans.ServerKey, Env: midtrans.Sandbox},
	}
}
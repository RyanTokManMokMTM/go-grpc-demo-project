package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	MySql struct {
		Datasource string
	}
	Salt string
}
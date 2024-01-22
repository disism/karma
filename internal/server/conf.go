package server

import (
	"github.com/spf13/viper"
	"strings"
)

const (
	DefaultRedisAddr     = "localhost:6379"
	DefaultRedisPassword = ""
)

func RedisAddr() string {
	addr := viper.GetString("redis.addr")
	if strings.TrimSpace(addr) != "" {
		return addr
	}
	return DefaultRedisAddr
}

func RedisPassword() string {
	password := viper.GetString("redis.password")
	if strings.TrimSpace(password) != "" {
		return password
	}
	return DefaultRedisPassword
}

func JWTSecret() string {
	return viper.GetString("jwt.secret")
}

func JWTIssues() string {
	return viper.GetString("server.addr")
}

func IPFSAddr() string {
	return viper.GetString("ipfs.addr")
}

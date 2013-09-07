package goredis

import (
	"./storage"
)

// GoRedisServer
type GoRedisServer struct {
	RedisServer
	Storages storage.RedisStorages
}

func NewGoRedisServer() (server *GoRedisServer) {
	server = &GoRedisServer{}
	server.SetHandler(server)
	server.Storages = storage.RedisStorages{}
	server.Storages.StringStorage = storage.NewMemoryStringStorage()
	server.Storages.KeyTypeStorage = storage.NewMemoryKeyTypeStorage()
	return
}

func (server *GoRedisServer) OnPETER(cmd *Command) (reply *Reply) {
	reply = StatusReply("Hello, Peter")
	return
}
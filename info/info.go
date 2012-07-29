/*
   简单的获取cpu利用率，系统内存使用情况，磁盘读写情况.
   查看cpu温度命令sensors输出内容，判断是否高于（70）等.
   获取硬盘温度.
   使用package net检查tcp/udp等服务是否在线.
   使用shell脚本检查特定进程。
*/
package info

import (
	"log"
	"os"
)

func Logger(name string) *log.Logger {
	file, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Open log:", err)
	}
	return log.New(file, "", log.LstdFlags)
}

var ErrorLog *log.Logger

type Server struct {
	Address string
	Host    *Host
	Load    *Load
	Temp    *Temp
}

func (s *Server) Get(address *string, reply *Server) error {
	reply.Address = *address
	host, err := GetHost()
	if err != nil {
		return err
	}
	reply.Host = host
	load, err := GetLoad()
	if err != nil {
		return err
	}
	reply.Load = load
	temp, err := GetTemp()
	if err != nil {
		return err
	}
	reply.Temp = temp
	return nil
}
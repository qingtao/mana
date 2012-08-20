package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"time"
)

var (
	sys_time  = 311 * time.Second
	srv_time  = 23 * time.Second
	proc_time = 29 * time.Second
	//sh_time      = 181 * time.Second
	sh_time      = 31 * time.Second
	running_time = 5 * time.Second
)

var help = flag.Bool("h", false, "help")

type Group struct {
	Name     string      `json:"name"`
	Computer []*Computer `json:"computer"`
}

func main() {
	if *help {
		flag.PrintDefaults()
		return
	}
	group_bytes, err := ioutil.ReadFile("etc/group")
	if err != nil {
		fmt.Println(err)
		return
	}
	var group Group
	err = json.Unmarshal(group_bytes, &group)
	if err != nil {
		fmt.Println(err)
		return
	}

	var (
		sys     = time.Tick(sys_time)
		srv     = time.Tick(srv_time)
		proc    = time.Tick(proc_time)
		sh      = time.Tick(sh_time)
		running = time.Tick(running_time)
	)

	go check_if_warn(notify.Retry, notify.Warn)

	go func() {
		for {
			select {
			case warn := <-notify.Warn:
				fmt.Println(warn)
			case err := <-notify.err:
				fmt.Println(err)
			}
		}
	}()

	for {
		select {
		case <-running:
			for _, co := range group.Computer {
				co.Status()
			}
		case <-sys:
			for _, co := range group.Computer {
				co.System()
			}
		case <-srv:
			for _, co := range group.Computer {
				co.Tcp()
				co.Udp()
			}
		case <-proc:
			for _, co := range group.Computer {
				co.Process()
			}
		case <-sh:
			for _, co := range group.Computer {
				co.Shell()
			}
		}
	}
}

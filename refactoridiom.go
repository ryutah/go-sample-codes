package main

import (
	"fmt"
	"strings"
	"time"
)

// for-selectイディオムを1つの関数にラッピングする
func ForSelectBefore() {
L:
	for {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("hello")
			break L
		}
	}

	fmt.Println("ending")
}

func ForSelectAfter() {
	f := func() {
		for {
			select {
			case <-time.After(time.Second):
				fmt.Println("hello")
				return
			}
		}
	}

	f()
	fmt.Println("ending")
}

// 整数型定数値にString()メソッドを追加する
type State int

const (
	Running State = iota + 1 // iotaはインクリメントしてから使用する
	Stopped
	Rebooting
	Terminate
)

func (s State) String() string {
	switch s {
	case Running:
		return "Running"
	case Stopped:
		return "Stopped"
	case Terminate:
		return "Terminate"
	default:
		return "Unknown"
	}
}

func PrintState() {
	state := Running
	fmt.Println("state ", state)
}

// Sliceやmapなどをカスタムな型に変換する
type Server struct {
	Name string
}

func ListServersBefore(name string) []Server {
	servers := []Server{
		{Name: "Server1"},
		{Name: "Server2"},
		{Name: "Foo1"},
		{Name: "Foo2"},
	}

	if name == "" {
		return servers
	}

	filtered := make([]Server, 0)

	for _, server := range servers {
		if strings.Contains(server.Name, name) {
			filtered = append(filtered, server)
		}
	}

	return filtered
}

type Servers []Server

func (s Servers) Filter(name string) Servers {
	filtered := make(Servers, 0)

	for _, server := range s {
		if strings.Contains(server.Name, name) {
			filtered = append(filtered, server)
		}
	}

	return filtered
}

func ListServersAfter() Servers {
	return []Server{
		{Name: "Server1"},
		{Name: "Server2"},
		{Name: "Foo1"},
		{Name: "Foo2"},
	}
}

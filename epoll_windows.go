// +build windows,cgo

package epoller

import (
	"github.com/TradingStartup/epoller/wepoll"
)

type epoll = wepoll.Epoll

func NewPoller() (Poller, error) {
	return wepoll.NewPoller()
}

func NewPollerWithBuffer(count int) (Poller, error) {
	return wepoll.NewPollerWithBuffer(count)
}

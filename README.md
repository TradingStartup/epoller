# epoller
epoll implementation for connections in Linux, MacOS and windows.

[![GoDoc](https://godoc.org/github.com/TradingStartup/epoller?status.png)](http://godoc.org/github.com/smallnest/epoller)

Its target is implementing a simple epoll for connection, so you should see it only contains few methods:

```go
type Poller interface {
	Add(conn net.Conn) error
	Remove(conn net.Conn) error
	Wait() ([]net.Conn, error)
	Close() error
}
```

Welcome any PRs for windows IOCompletePort.

Inspired by [1m-go-websockets](https://github.com/eranyanay/1m-go-websockets).

Thanks @sunnyboy00 for providing windows implementation.
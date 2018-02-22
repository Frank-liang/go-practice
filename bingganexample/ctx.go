package main

import (
	"context"
	"net"
	"net/http"
	"time"
)

func callRedis(ctx context.Context) {
	dialer := new(net.Dialer)
	conn, err := dialer.DialContext(ctx, "tcp", "")
	deadline, ok := ctx.Deadline()
	conn.SetDeadline(deadline)

	req := http.NewRequest("GET", "", nil)
	req = req.WithContext(ctx)

}

func callMySql(ctx context.Context) {

}

func callApi(ctx context.Context) {

}

func main() {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	go callRedis(ctx)
	go callMySql(ctx)
	go callApi(ctx)

	user := make(chan bool)
	select {
	case <-ctx.Done():
	case <-user:
		cancel()
	}

}

package main

import (
	"context"
	"fmt"
)

type database map[string]bool
type userIdKeyType string

var db database = database{
	"jane": true,
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	processRequest(ctx, "jane")
}

func processRequest(ctx context.Context, userid string) {
	vctx := context.WithValue(ctx, userIdKeyType("userIDKey"), "jane")
	ch := checkMembership(vctx)
	status := <-ch
	fmt.Println("membership status of userid : %s : %v\n", userid, status)
}

func checkMembership(ctx context.Context) <-chan bool {
	ch := make(chan bool)
	go func() {
		defer close(ch)
		userid := ctx.Value(userIdKeyType("userIDKey")).(string)
		status := db[userid]
		ch <- status
	}()
	return ch
}
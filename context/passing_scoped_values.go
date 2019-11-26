package main

import "context"
import "fmt"

type KeyType struct{}

var key = &KeyType{}

func WithKey(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, key, value)
}

func GetKey(ctx context.Context) (string, bool) {
	v := ctx.Value(key)
	if v == nil {
		return "", false
	}
	return v.(string), true
}

func main() {
	fmt.Println("vim-go")
}

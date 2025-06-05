package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "key1", "value1")
	bookHotel(&ctx)
}

func bookHotel(ctx *context.Context) {
	token := (*ctx).Value("key1")
	fmt.Println(token)
}

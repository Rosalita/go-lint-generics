package main

import (
	"context"
	"fmt"
	"time"

	catfacts "github.com/Rosalita/go-lint-generics/pkg/catfacts/client"
)

var timeout = time.Second * 30

func main() {
	client, err := catfacts.NewClient("https://catfact.ninja")
	if err != nil {
		fmt.Println(err)
		return
	}
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	catFact, err := client.Get(ctx, "/fact")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("here is your cat fact:")
	fmt.Println(catFact.Fact)
}

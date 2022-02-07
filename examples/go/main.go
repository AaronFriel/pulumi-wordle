package main

import (
	"fmt"

	"github.com/aaronfriel/pulumi-wordle/sdk/go/wordle"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		Wordle, err := wordle.NewWordle(ctx, "wordle", &wordle.WordleArgs{
			Word: pulumi.String("cloud"),
		})
		if err != nil {
			return fmt.Errorf("error creating Wordle: %v", err)
		}

		ctx.Export("result", Wordle.Result)

		return nil
	})
}

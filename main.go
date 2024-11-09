package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/satishcg12/dotnet-me/internal"
)

func main() {
	app := internal.NewApp()

	ctx, cancle := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancle()

	err := app.Start(ctx)
	if err != nil {
		panic(err)
	}

}

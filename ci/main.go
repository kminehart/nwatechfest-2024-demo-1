package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"dagger.io/dagger"
)

func main() {
	ctx := context.Background()
	c, err := dagger.Connect(ctx)
	if err != nil {
		panic(err)
	}

	repo := c.Git("https://github.com/kminehart/nwatechfest-2024-demo-1.git")
	source := repo.Branch("main").Tree()

	container := c.Container().From("golang:1.22").
		WithMountedDirectory("/src", source).
		WithWorkdir("/src").
		WithExec([]string{"go", "test", "./..."})

	container, err = container.Sync(ctx)
	if err != nil {
		log.Fatalln("error running containers", err)
	}

	if stdout, err := container.Stdout(ctx); err == nil {
		fmt.Fprint(os.Stdout, stdout)
	}

	if stderr, err := container.Stdout(ctx); err == nil {
		fmt.Fprint(os.Stderr, stderr)
	}
}

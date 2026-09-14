package main

import (
	"context"
	"fmt"
	"os"
)

func main() {
	cfg, err := configFromEnv()
	if err != nil {
		fmt.Printf("failed to configure databricks-model-provider: %v\n", err)
		os.Exit(1)
	}

	if len(os.Args) > 1 && os.Args[1] == "validate" {
		if _, err := cfg.listModels(context.Background()); err != nil {
			fmt.Printf("failed to validate databricks-model-provider: %v\n", err)
			os.Exit(1)
		}
		return
	}

	if err := cfg.run(); err != nil {
		fmt.Printf("failed to run databricks-model-provider: %v\n", err)
		os.Exit(1)
	}
}

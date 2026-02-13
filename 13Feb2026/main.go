package main

import (
	"context"
	"fmt"
)

type jobIdKeyType struct{}

var jobIdKey = jobIdKeyType{}

func validate(ctx context.Context) {
	jobId, _ := ctx.Value(jobIdKey).(string)
	fmt.Println("Validating job:", jobId)
}

func process(ctx context.Context) {
	jobId, _ := ctx.Value(jobIdKey).(string)
	fmt.Println("Processing job:", jobId)
}

func store(ctx context.Context) {
	jobId, _ := ctx.Value(jobIdKey).(string)
	fmt.Println("Storing the results for the job:", jobId)
}

func main() {
	ctx := context.WithValue(context.Background(), jobIdKey, "job-1234")

	validate(ctx)
	process(ctx)
	store(ctx)
}

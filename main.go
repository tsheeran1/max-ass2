package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	// Age    int `json:"age"`
	// Height int `json:"height"`
	Income int `json:"income"`
}

func handler(ctx context.Context, e Event) (int, error) {

	fmt.Println("Event:", e)
	// fmt.Println("Age: ", e.Age)
	// fmt.Println("Height", e.Height)
	fmt.Println("Income", e.Income)

	einc := e.Income
	return einc * 12, nil
}

func main() {
	lambda.Start(handler)
}

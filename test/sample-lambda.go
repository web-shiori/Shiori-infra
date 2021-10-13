package main

import (
    "fmt"

    "github.com/aws/aws-lambda-go/lambda"
)

func sampleHandler() {
    fmt.Println("Hello, lambda")
}

func main() {
    lambda.Start(sampleHandler)
}

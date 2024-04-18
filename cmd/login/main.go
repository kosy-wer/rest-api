package main

import (
    "fmt"
    "rest_api/internal/app/register" // Import package rest_api_register dari folder internal/app/rest_api_register
)

func main() {
    fmt.Println("Hello from main package!")
    register.SayHello() // Panggil fungsi dari package register
}


package gotrack

import (
    "fmt"
    "github.com/stvp/rollbar"
)

func ReportInfo(message string) {
    fmt.Println("Message:", message)
    rollbar.Message("info", message)
}

func ReportDebug(message string) {
    fmt.Println("Debug:", message)
    rollbar.Message("debug", message)
}

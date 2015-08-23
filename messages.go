package gotrack

import (
    "github.com/stvp/rollbar"
)

func ReportInfo(message string) {
    rollbar.Message("info", message)
}

func ReportDebug(message string) {
    rollbar.Message("debug", message)
}

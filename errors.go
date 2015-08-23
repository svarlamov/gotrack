package gotrack

import (
    "github.com/stvp/rollbar"
)

func ReportError(err error, message string) {
    m := make(map[string]string)
    m["message"] = message
    ReportCriticalVerbose(err, m)
}

func ReportCritical(err error, message string) {
    m := make(map[string]string)
    m["message"] = message
    ReportCriticalVerbose(err, m)
}

func ReportErrorVerbose(err error, fields map[string]string) {
    rollbarFields := mapToRollbarFields(fields)
    rollbar.Error("error", err, rollbarFields...)
}

func ReportCriticalVerbose(err error, fields map[string]string) {
    rollbarFields := mapToRollbarFields(fields)
    rollbar.Error("critical", err, rollbarFields...)
}

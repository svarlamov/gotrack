package gotrack

import (
    "github.com/stvp/rollbar"
)

func ReportError(err error, message string) {
    m := make(map[string]string)
    ReportErrorVerbose(err, message, m)
}

func ReportCritical(err error, message string) {
    m := make(map[string]string)
    ReportCriticalVerbose(err, message, m)
}

func ReportErrorVerbose(err error, message string, fields map[string]string) {
    fields["message"] = message
    rollbarFields := mapToRollbarFields(fields)
    rollbar.Error("error", err, rollbarFields...)
}

func ReportCriticalVerbose(err error, message string, fields map[string]string) {
    fields["message"] = message
    rollbarFields := mapToRollbarFields(fields)
    rollbar.Error("critical", err, rollbarFields...)
}

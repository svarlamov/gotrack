package gotrack

import (
    "github.com/stvp/rollbar"

    "errors"
)

func SetupTracking(tool string, tkn string, env string) error {
    if tool == "rollbar" {
        rollbar.Token = tkn
        rollbar.Environment = env
        ReportInfo("Node launched")
        return nil
    } else {
        return errors.New("Unrecognized tracking service specified")
    }
}

func StopTracking() {
    ReportInfo("Node stopped")
    wait()
}

func wait() {
    rollbar.Wait()
}

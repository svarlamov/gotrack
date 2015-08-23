package gotrack

import (
    "github.com/stvp/rollbar"
)

func mapToRollbarFields(input map[string]string) []*rollbar.Field {
    result := make([]*rollbar.Field, len(input))
    current := 0
    for ind, item := range input {
        result[current] = &rollbar.Field{
            Name: ind,
            Data: item,
        }
        current++
    }
    return result
}

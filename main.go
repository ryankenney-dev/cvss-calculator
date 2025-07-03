package main

import (
    "fmt"
    "os"

    "github.com/goark/go-cvss/v3/metric"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Fprintln(os.Stderr, "Usage: <program> <CVSS string>")
        os.Exit(1)
    }
    cvssString := os.Args[1]
    em, err := metric.NewEnvironmental().Decode(cvssString)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
    fmt.Printf("Base Severity: %v (%v)\n", em.BaseMetrics().Severity(), em.BaseMetrics().Score())
    fmt.Printf("Temporal Severity: %v (%v)\n", em.TemporalMetrics().Severity(), em.TemporalMetrics().Score())
    fmt.Printf("Environmental Severity: %v (%v)\n", em.Severity(), em.Score())
}

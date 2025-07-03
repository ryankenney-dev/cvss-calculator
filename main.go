package main

import (
    "fmt"
    "io"
    "os"
    "strings"

    "github.com/goark/go-cvss/v3/metric"
    "github.com/goark/go-cvss/v3/report"
)

var template = `- CVSS Version {{ .Version }}
- Vector: {{ .Vector }}

## Base Metrics

- Base Score: {{ .BaseScore }}

| {{ .BaseMetrics }} | {{ .BaseMetricValue }} |
|--------|-------|
| {{ .AVName }} | {{ .AVValue }} |
| {{ .ACName }} | {{ .ACValue }} |
| {{ .PRName }} | {{ .PRValue }} |
| {{ .UIName }} | {{ .UIValue }} |
| {{ .SName }} | {{ .SValue }} |
| {{ .CName }} | {{ .CValue }} |
| {{ .IName }} | {{ .IValue }} |
| {{ .AName }} | {{ .AValue }} |

## Temporal Metrics

- Temporal Score: {{ .TemporalScore }}
- {{ .SeverityName }}: {{ .SeverityValue }}

| {{ .TemporalMetrics }} | {{ .TemporalMetricValue }} |
|--------|-------|
| {{ .EName }} | {{ .EValue }} |
| {{ .RLName }} | {{ .RLValue }} |
| {{ .RCName }} | {{ .RCValue }} |

## Environmental Metrics

- {{ .SeverityName }}: {{ .SeverityValue }} ({{ .EnvironmentalScore }})

| {{ .EnvironmentalMetrics }} | {{ .EnvironmentalMetricValue }} |
|--------|-------|
| {{ .CRName }} | {{ .CRValue }} |
| {{ .IRName }} | {{ .IRValue }} |
| {{ .ARName }} | {{ .ARValue }} |
| {{ .MAVName }} | {{ .MAVValue }} |
| {{ .MACName }} | {{ .MACValue }} |
| {{ .MPRName }} | {{ .MPRValue }} |
| {{ .MUIName }} | {{ .MUIValue }} |
| {{ .MSName }}  | {{ .MSValue }} |
| {{ .MCName }}  | {{ .MCValue }} |
| {{ .MIName }}  | {{ .MIValue }} |
| {{ .MAName }}  | {{ .MAValue }} |
`

func main() {
    if len(os.Args) < 2 {
        fmt.Fprintln(os.Stderr, "Usage: <program> <CVSS string>")
        os.Exit(1)
    }
    cvssString := os.Args[1]

    // Simpler example
    //
    // em, err := metric.NewEnvironmental().Decode(cvssString)
    // if err != nil {
    //     fmt.Fprintln(os.Stderr, err)
    //     os.Exit(1)
    // }
    // fmt.Printf("Base Severity: %v (%v)\n", em.BaseMetrics().Severity(), em.BaseMetrics().Score())
    // fmt.Printf("Temporal Severity: %v (%v)\n", em.TemporalMetrics().Severity(), em.TemporalMetrics().Score())
    // fmt.Printf("Environmental Severity: %v (%v)\n", em.Severity(), em.Score())

    r, err := report.NewEnvironmental(em).ExportWith(strings.NewReader(template))
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    if _, err := io.Copy(os.Stdout, r); err != nil {
        fmt.Fprintln(os.Stderr, err)
    }
}

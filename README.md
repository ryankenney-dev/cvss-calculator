# cvss-calculator

## Overview

I needed a tool similar to
[Common Vulnerability Scoring System Calculator](https://nvd.nist.gov/vuln-metrics/cvss/v3-calculator),
but that could easily convert from a CVSS string into a score.
(For some strange reason that website cannot do it...)

This repo provides a command line tool that can.
I tool from a docker/podman container with no network
just because I did nothing to vette the emedded libraries for security.

## Setup

Create a config file:

    cp config.sh.template config.sh

... and customize it appropriately

## Usage

Convert a base v3.1 CVSS to scores:

    ./build-run.sh "CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:C/C:L/I:L/A:L"

    Base Severity: High (8.3)
    Temporal Severity: High (8.3)
    Environmental Severity: Medium (8.3)

Convert a full (base/temporal/environmental) v3.1 CVSS (with `MAV:L` as an environmental modifier) to scores:

    ./build-run.sh "CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:C/C:L/I:L/A:L/CR:X/IR:X/AR:X/MAV:L/MAC:X/MPR:X/MUI:X/MS:X/MC:X/MI:X/MA:X"

    Base Severity: High (8.3)
    Temporal Severity: High (8.3)
    Environmental Severity: Medium (6.8)


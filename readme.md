# hourbucket

[![Go Reference](https://pkg.go.dev/badge/github.com/Aaron2599/hourbucket.svg)](https://pkg.go.dev/github.com/Aaron2599/hourbucket)
A simple Go utility package for converting time values to and from an "hourly epoch" representation. The hourly epoch is defined here as the total number of full hours elapsed since the Unix epoch (January 1, 1970, 00:00:00 UTC).

This can be useful for creating time-based identifiers or partition keys that group data by the hour.

## Installation
```bash
go get github.com/Aaron2599/hourbucket
```
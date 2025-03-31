# Hour Bucket

The `hourbucket` Package offers utilities to easily convert Go's standard `time.Time` into hourly intervals. It represents each hour with a single `int64` value: the total number of hours passed since the Unix epoch (Jan 1, 1970 UTC).

**Key Features:**

* Convert any `time.Time` to its unique "hourly epoch".
* Get the current hourly epoch with `hourbucket.Now()`.
* Convert an hourly epoch time back into standard Unix timestamps (seconds, ms, µs, ns).

**Common Use Cases:**

* Creating efficient, time-based partition keys for databases (like Cassandra, InfluxDB) to group data hourly.
* Generating consistent, time-bucketed identifiers or keys.
* Aggregating or analyzing time-series data within distinct hourly intervals.

## Installation
```bash
go get github.com/Aaron2599/hourbucket
```

## Usage
```go
package main

import (
	"fmt"
	"github.com/Aaron2599/hourbucket"
	"time"
)

func main() {
	var currentHour = hourbucket.Now()
	fmt.Printf("Current Hourly Epoch: %d\n", currentHour)

	fmt.Println("---")

	var someDate = time.Date(2025, time.April, 15, 9, 45, 30, 0, time.UTC)
	var timeToHourly = hourbucket.ToHourly(someDate)
	fmt.Printf("Hourly Epoch for %s: %d\n", someDate.Format(time.RFC3339), timeToHourly)

	fmt.Println("--- Converting Current Hourly Epoch Back ---")

	var hourlyToUnix = hourbucket.ToUnix(currentHour)
	fmt.Printf("%d -> Unix Seconds: %d\n", currentHour, hourlyToUnix)

	var hourlyToUnixMilli = hourbucket.ToUnixMilli(currentHour)
	fmt.Printf("%d -> Unix Millis: %d\n", currentHour, hourlyToUnixMilli)

	var hourlyToUnixMicro = hourbucket.ToUnixMicro(currentHour)
	fmt.Printf("%d -> Unix Micros: %d\n", currentHour, hourlyToUnixMicro)

	var hourlyToUnixNano = hourbucket.ToUnixNano(currentHour)
	fmt.Printf("%d -> Unix Nanos:  %d\n", currentHour, hourlyToUnixNano)
}
```



package test

import (
	"time"
)

const secondsPerHour int64 = 3600

const millisPerHour int64 = 3600000

const microsPerHour int64 = 3600000000

const nanosPerHour int64 = 3600000000000

func Now() int64 {
	return time.Now().UTC().Unix() / secondsPerHour
}

func ToHourly(t time.Time) int64 {
	return t.UTC().Unix() / secondsPerHour
}

func ToUnix(hourly int64) int64 {
	return hourly * secondsPerHour
}

func ToUnixMilli(hourly int64) int64 {
	return hourly * millisPerHour
}

func ToUnixMicro(hourly int64) int64 {
	return hourly * microsPerHour
}

func ToUnixNano(hourly int64) int64 {
	return hourly * nanosPerHour
}

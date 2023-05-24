package NTPPrint

import (
	"testing"
	"time"
)

func TestTimePrint(t *testing.T) {
	curTimeNTP := printTime()
	curTime := time.Now()
	if curTimeNTP.Year() != curTime.Year() || 
	curTimeNTP.Month() != curTime.Month() || 
	curTimeNTP.Day() != curTime.Day() || 
	curTimeNTP.Hour() != curTime.Hour() || 
	curTimeNTP.Minute() != curTime.Minute() {
		t.Errorf("NTP time nearly not as time-package time\nNTP time: %v\ntime-pkg time: %v\n", curTimeNTP, curTime)
	}
}

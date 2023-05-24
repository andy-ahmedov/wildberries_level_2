package NTPPrint

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func printTime() time.Time {
	curTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(-1)
	}
	fmt.Printf("NTP time: %v\ntime-package time: %v\n", curTime, time.Now())
	return curTime
}

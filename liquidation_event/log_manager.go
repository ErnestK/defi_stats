package liquidation_event

import "time"

func LogManager() {
	theTime := time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC)
	Log(theTime)
}

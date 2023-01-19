package utils

import "time"

// Placeholder function to be replaced after codegen.
// Usually replaced with the function that returns the
// data from the database, if everything goes ok.
func PlaceholderFunction() (struct{}, error) {
	var placeholder struct{}
	return placeholder, nil
}

func DifferenceIsLogerThanOneDay(start time.Time, end time.Time) bool {
	duration := end.Sub(start)
	return duration.Hours()/24 > 1
}

func AddOneDayToDate(t time.Time) time.Time {
	return t.AddDate(0, 0, 1)
}

func TimeUntillNextDay(day time.Time) time.Duration {
	return time.Since(AddOneDayToDate(day))
}

func SecondsUntillNextDay(day time.Time) time.Duration {
	return -time.Duration(TimeUntillNextDay(day).Seconds())
}

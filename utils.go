package gonba

import (
	"fmt"
	"strings"
	"time"
)

const dateLayout = "2006-01-02"
const timeLayout = "2006-01-02 3:04 pm MST"

func ParseDateFromStrings(dateString string, timeString string) time.Time {
	timeZone, _ := time.LoadLocation("America/New_York")
	timeZoneCode, _ := time.Now().In(timeZone).Zone()
	dateString = strings.TrimSuffix(dateString, "T00:00:00")
	parseLayout := dateLayout
	if timeString != "" {
		timeString = strings.TrimSuffix(timeString, " ET")
		dateString = fmt.Sprintf("%s %s %s", dateString, timeString, timeZoneCode)
		parseLayout = timeLayout
	}
	date, err := time.Parse(parseLayout, dateString)
	if err != nil {
		fmt.Printf("Error when parsing dateString %s with layout %s\n", dateString, parseLayout)
	}
	return date
}

func FormatGameIdString(gameId int) string {
	return fmt.Sprintf("00%d", gameId)
}
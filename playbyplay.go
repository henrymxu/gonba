package gonba

import (
	"strconv"
)

const endpointPlayByPlay = "playbyplayv2"

// first period is startPeriod, second period is endPeriod
func (c *Client) GetPlayByPlay(gameId int, periods ...int) {
	params := make(map[string]string)
	if len(periods) > 0 {
		params["StartPeriod"] = strconv.Itoa(periods[0])
		if len(periods) > 1 {
			params["EndPeriod"] = strconv.Itoa(periods[1])
		}
	}
	params["GameID"] = FormatGameIdString(gameId)
	c.makeRequestWithoutJson(endpointPlayByPlay, params)
}

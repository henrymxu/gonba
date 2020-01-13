package gonba

import (
	"fmt"
	"time"
)

const endpointPlayByPlayV2 = "v1/%s/%s_pbp_%d.json"
const playbyplayV2DateFormat = "20060102"

func (c *Client) GetPlayByPlayV2(gameDate time.Time, gameId string, periods ...int) (PlayByPlayV2, int) {
	params := map[string]string {"version": "3"}
	dateString := gameDate.Format(playbyplayV2DateFormat)
	result := PlayByPlayV2{
		Plays: make([]Play, 0),
	}
	var status int
	for _, period := range periods {
		var playByPlay PlayByPlayV2
		endpoint := fmt.Sprintf(endpointPlayByPlayV2, dateString, gameId, period)
		status = c.makeRequest(endpoint, params, &playByPlay)
		result.Internal = playByPlay.Internal
		result.Plays = append(result.Plays, playByPlay.Plays...)
	}
	return result, status
}

func (c *Client) GetPlayByPlayV2All(gameDate time.Time, gameId string) (PlayByPlayV2, int) {
	params := map[string]string {"version": "3"}
	dateString := gameDate.Format(playbyplayV2DateFormat)
	result := PlayByPlayV2{
		Plays: make([]Play, 0),
	}
	var status int
	currentQuarterHasPlays := true
	period := 1
	for currentQuarterHasPlays {
		var playByPlay PlayByPlayV2
		endpoint := fmt.Sprintf(endpointPlayByPlayV2, dateString, gameId, period)
		status = c.makeRequest(endpoint, params, &playByPlay)
		result.Internal = playByPlay.Internal
		result.Plays = append(result.Plays, playByPlay.Plays...)
		currentQuarterHasPlays = len(playByPlay.Plays) != 0
		period++
	}
	return result, status
}

type PlayByPlayV2 struct {
	Internal struct {
		PubDateTime string `json:"pubDateTime"`
		Xslt        string `json:"xslt"`
		EventName   string `json:"eventName"`
	} `json:"_internal"`
	Plays []Play `json:"plays"`
}

type Play struct {
	Clock            string `json:"clock"`
	EventMsgType     int `json:"eventMsgType,string"`
	Description      string `json:"description"`
	PersonID         int `json:"personId,string"`
	TeamID           int `json:"teamId,string"`
	VTeamScore       int `json:"vTeamScore,string"`
	HTeamScore       int `json:"hTeamScore,string"`
	IsScoreChange    bool   `json:"isScoreChange"`
	IsVideoAvailable bool   `json:"isVideoAvailable"`
	Formatted        struct {
		Description string `json:"description"`
	} `json:"formatted"`
}
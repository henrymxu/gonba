package gonba

import (
	"fmt"
	"time"
)

const endpointScheduleV3 = "v2/%s/scoreboard.json"
const scheduleV3DateFormat = "20060102"

func (c *Client) GetScheduleV3(date time.Time) (ScheduleV3, int) {
	var schedule ScheduleV3
	params := map[string]string{"version": "3"}
	dateString := date.Format(scheduleV3DateFormat)
	endpoint := fmt.Sprintf(endpointScheduleV3, dateString)
	status := c.makeRequest(endpoint, params, &schedule)
	return schedule, status
}

type ScheduleV3 struct {
	Internal struct {
		PubDateTime string `json:"pubDateTime"`
		Xslt        string `json:"xslt"`
		EventName   string `json:"eventName"`
	} `json:"_internal"`
	NumGames int      `json:"numGames"`
	Games    []GameV3 `json:"games"`
}

type GameV3 struct {
	SeasonStageID         int       `json:"seasonStageId"`
	SeasonYear            int    `json:"seasonYear,string"`
	GameID                string    `json:"gameId"`
	Arena                 Arena     `json:"arena"`
	IsGameActivated       bool      `json:"isGameActivated"`
	StatusNum             int       `json:"statusNum"`
	ExtendedStatusNum     int       `json:"extendedStatusNum"`
	StartTimeEastern      string    `json:"startTimeEastern"`
	StartTimeUTC          time.Time `json:"startTimeUTC"`
	EndTimeUTC            time.Time `json:"endTimeUTC"`
	StartDateEastern      string    `json:"startDateEastern"`
	Clock                 string    `json:"clock"`
	IsBuzzerBeater        bool      `json:"isBuzzerBeater"`
	IsPreviewArticleAvail bool      `json:"isPreviewArticleAvail"`
	IsRecapArticleAvail   bool      `json:"isRecapArticleAvail"`
	Tickets               Ticket    `json:"tickets"`
	HasGameBookPdf        bool      `json:"hasGameBookPdf"`
	IsStartTimeTBD        bool      `json:"isStartTimeTBD"`
	Nugget                struct {
		Text string `json:"text"`
	} `json:"nugget"`
	Attendance   string `json:"attendance"`
	GameDuration struct {
		Hours   string `json:"hours"`
		Minutes string `json:"minutes"`
	} `json:"gameDuration"`
	Period Period `json:"period"`
	VTeam  TeamV3 `json:"vTeam"`
	HTeam  TeamV3 `json:"hTeam"`
	Watch  struct {
		Broadcast struct {
			Broadcasters TypeOfBroadcastersV3 `json:"broadcasters"`
			Video        Video                `json:"video"`
			Audio        struct {
				National BroadcastersV3 `json:"national"`
				VTeam    BroadcastersV3 `json:"vTeam"`
				HTeam    BroadcastersV3 `json:"hTeam"`
			} `json:"audio"`
		} `json:"broadcast"`
	} `json:"watch"`
}

type Arena struct {
	Name       string `json:"name"`
	IsDomestic bool   `json:"isDomestic"`
	City       string `json:"city"`
	StateAbbr  string `json:"stateAbbr"`
	Country    string `json:"country"`
}

type Ticket struct {
	MobileApp  string `json:"mobileApp"`
	DesktopWeb string `json:"desktopWeb"`
	MobileWeb  string `json:"mobileWeb"`
}

type Period struct {
	Current       int  `json:"current"`
	Type          int  `json:"type"`
	MaxRegular    int  `json:"maxRegular"`
	IsHalftime    bool `json:"isHalftime"`
	IsEndOfPeriod bool `json:"isEndOfPeriod"`
}

type TeamV3 struct {
	TeamID     int `json:"teamId,string"`
	TriCode    string `json:"triCode"`
	Win        int `json:"win,string"`
	Loss       int `json:"loss,string"`
	SeriesWin  int `json:"seriesWin,string"`
	SeriesLoss int `json:"seriesLoss,string"`
	Score      int `json:"score,string"`
	Linescore  []struct {
		Score int `json:"score,string"`
	} `json:"linescore"`
}

type Video struct {
	RegionalBlackoutCodes string        `json:"regionalBlackoutCodes"`
	CanPurchase           bool          `json:"canPurchase"`
	IsLeaguePass          bool          `json:"isLeaguePass"`
	IsNationalBlackout    bool          `json:"isNationalBlackout"`
	IsTNTOT               bool          `json:"isTNTOT"`
	IsVR                  bool          `json:"isVR"`
	TntotIsOnAir          bool          `json:"tntotIsOnAir"`
	IsNextVR              bool          `json:"isNextVR"`
	IsNBAOnTNTVR          bool          `json:"isNBAOnTNTVR"`
	IsMagicLeap           bool          `json:"isMagicLeap"`
	IsOculusVenues        bool          `json:"isOculusVenues"`
	Streams               []VideoStream `json:"streams"`
	DeepLink              []DeepLink    `json:"deepLink"`
}

type VideoStream struct {
	StreamType            string `json:"streamType"`
	IsOnAir               bool   `json:"isOnAir"`
	DoesArchiveExist      bool   `json:"doesArchiveExist"`
	IsArchiveAvailToWatch bool   `json:"isArchiveAvailToWatch"`
	StreamID              string `json:"streamId"`
	Duration              int    `json:"duration"`
}

type DeepLink struct {
	Broadcaster         string `json:"broadcaster"`
	RegionalMarketCodes string `json:"regionalMarketCodes"`
	IosApp              string `json:"iosApp"`
	AndroidApp          string `json:"androidApp"`
	DesktopWeb          string `json:"desktopWeb"`
	MobileWeb           string `json:"mobileWeb"`
}

type TypeOfBroadcastersV3 struct {
	National        []BroadcasterV3 `json:"national"`
	Canadian        []BroadcasterV3 `json:"canadian"`
	VTeam           []BroadcasterV3 `json:"vTeam"`
	HTeam           []BroadcasterV3 `json:"hTeam"`
	SpanishHTeam    []BroadcasterV3 `json:"spanish_hTeam"`
	SpanishVTeam    []BroadcasterV3 `json:"spanish_vTeam"`
	SpanishNational []BroadcasterV3 `json:"spanish_national"`
}

type BroadcastersV3 struct {
	Streams      []BroadcasterStream `json:"streams"`
	Broadcasters []BroadcasterV3     `json:"broadcasters"`
}

type BroadcasterStream struct {
	Language string `json:"language"`
	IsOnAir  bool   `json:"isOnAir"`
	StreamID string `json:"streamId"`
}

type BroadcasterV3 struct {
	ShortName string `json:"shortName"`
	LongName  string `json:"longName"`
}

package gonba

import (
	"fmt"
	"time"
)

const endpointScheduleV2 = "scoreboard/%s/games.json"
const scheduleV2DateFormat = "20060102"

func (c *Client) GetScheduleV2(date time.Time) (ScheduleV2, int) {
	var schedule scheduleResponseV2
	params := map[string]string{"version": "2"}
	dateString := date.Format(scheduleV2DateFormat)
	endpoint := fmt.Sprintf(endpointScheduleV2, dateString)
	status := c.makeRequest(endpoint, params, &schedule)
	return schedule.SportsContent, status
}

type scheduleResponseV2 struct {
	SportsContent ScheduleV2  `json:"sports_content"`
}

type ScheduleV2 struct {
	SportsMeta SportsMeta `json:"sports_meta"`
	Games      struct {
		Game [] GameV2 `json:"game"`
	} `json:"games"`
}

type SportsMeta struct {
	DateTime           string     `json:"date_time"`
	EndToEndTimeMillis string     `json:"end_to_end_time_millis"`
	ConsolidatedDomKey string     `json:"consolidatedDomKey"`
	SeasonMeta         SeasonMeta `json:"season_meta"`
	Next               struct {
		URL string `json:"url"`
	} `json:"next"`
}

type SeasonMeta struct {
	CalendarDate        string `json:"calendar_date"`
	SeasonYear          int `json:"season_year,string"`
	StatsSeasonYear     int `json:"stats_season_year,string"`
	StatsSeasonID       int `json:"stats_season_id,string"`
	StatsSeasonStage    int `json:"stats_season_stage,string"`
	RosterSeasonYear    int `json:"roster_season_year,string"`
	ScheduleSeasonYear  int `json:"schedule_season_year,string"`
	StandingsSeasonYear int `json:"standings_season_year,string"`
	SeasonID            int `json:"season_id,string"`
	DisplayYear         string `json:"display_year"`
	DisplaySeason       string `json:"display_season"`
	SeasonStage         int `json:"season_stage,string"`
	LeagueID            int `json:"league_id,string"`
}

type GameV2 struct {
	ID                string `json:"id"`
	GameURL           string `json:"game_url"`
	SeasonID          int `json:"season_id,string"`
	Date              string `json:"date"`
	Time              string `json:"time"`
	Arena             string `json:"arena"`
	City              string `json:"city"`
	State             string `json:"state"`
	Country           string `json:"country"`
	HomeStartDate     string `json:"home_start_date"`
	HomeStartTime     string `json:"home_start_time"`
	VisitorStartDate  string `json:"visitor_start_date"`
	VisitorStartTime  string `json:"visitor_start_time"`
	PreviewAvailable  int `json:"previewAvailable,string"`
	RecapAvailable    int `json:"recapAvailable,string"`
	NotebookAvailable int `json:"notebookAvailable,string"`
	TntOt             int `json:"tnt_ot,string"`
	BuzzerBeater      int `json:"buzzerBeater,string"`
	Ticket            struct {
		TicketLink string `json:"ticket_link"`
	} `json:"ticket"`
	PeriodTime PeriodTime `json:"period_time"`
	Lp         Lp         `json:"lp"`
	Dl         struct {
		Link []Link `json:"link"`
	} `json:"dl"`
	Broadcasters Broadcasters `json:"broadcasters"`
	Visitor      TeamV2       `json:"visitor"`
	Home         TeamV2       `json:"home"`
}

type PeriodTime struct {
	PeriodValue  int `json:"period_value,string"`
	PeriodStatus string `json:"period_status"`
	GameStatus   int `json:"game_status,string"`
	GameClock    string `json:"game_clock"`
	TotalPeriods int `json:"total_periods,string"`
	PeriodName   string `json:"period_name"`
}

type Lp struct {
	LpVideo     bool`json:"lp_video,string"`
	CondensedBb bool `json:"condensed_bb,string"`
	Visitor     LpTeam `json:"visitor"`
	Home        LpTeam `json:"home"`
}

type LpTeam struct {
	Audio struct {
		ENG bool `json:"ENG,string"`
		SPA bool `json:"SPA,string"`
	} `json:"audio"`
	Video struct {
		Avl    bool `json:"avl,string"`
		OnAir  bool `json:"onAir,string"`
		ArchBB bool `json:"archBB,string"`
	} `json:"video"`
}

type Link struct {
	Name        string `json:"name"`
	LongNm      string `json:"long_nm"`
	Code        string `json:"code"`
	URL         string `json:"url"`
	MobileURL   string `json:"mobile_url"`
	HomeVisitor string `json:"home_visitor"`
}

type Broadcasters struct {
	Radio struct {
		Broadcaster []Broadcaster `json:"broadcaster"`
	} `json:"radio"`
	Tv struct {
		Broadcaster []Broadcaster `json:"broadcaster"`
	} `json:"tv"`
}

type Broadcaster struct {
	Scope       string `json:"scope"`
	HomeVisitor string `json:"home_visitor"`
	DisplayName string `json:"display_name"`
}

type TeamV2 struct {
	ID           string `json:"id"`
	TeamKey      string `json:"team_key"`
	City         string `json:"city"`
	Abbreviation string `json:"abbreviation"`
	Nickname     string `json:"nickname"`
	URLName      string `json:"url_name"`
	TeamCode     string `json:"team_code"`
	Score        int `json:"score,string"`
}

package gonba

import (
	"time"
)

const endpointSchedule = "scoreboardv2"
const scheduleDateFormat = "2006-01-02"

func (c *Client) GetSchedule(date time.Time) (Schedule, int) {
	var scheduleResponse scheduleResponse
	params := map[string]string {
		"GameDate": date.Format(scheduleDateFormat),
		"LeagueId": "00",
		"DayOffset": "0",
	}
	status := c.makeRequest(endpointSchedule, params, &scheduleResponse)
	return parseScheduleResponse(scheduleResponse), status
}

func parseScheduleResponse(response scheduleResponse) Schedule {
	gameHeaders := response.ResultSets[0]
	lineScores := response.ResultSets[1]
	eastStandings := response.ResultSets[4]
	westStandings := response.ResultSets[5]
	games := make([]Game, 0, len(gameHeaders.RowSet))
	for _, gameHeader := range gameHeaders.RowSet {
		game := Game{}
		game.GameId = gameHeader[2].(string) // GAME_ID
		game.Venue = gameHeader[15].(string) // ARENA_NAME
		game.Quarter = int(gameHeader[9].(float64)) // LIVE_PERIOD
		game.QuarterTime = gameHeader[10].(string) // LIVE_PC_TIME
		status := GameStatus{
			StatusCode: int(gameHeader[3].(float64)), // GAME_STATUS_ID (1 for pregame, 2 for live)
			StatusString: gameHeader[4].(string), // GAME_STATUS_TEXT
		}
		game.Status = status
		timeString := ""
		if status.StatusCode == 1 {
			timeString = status.StatusString
		}
		game.GameDate = ParseDateFromStrings(gameHeader[0].(string), timeString) // GAME_DATE_EST, GAME_STATUS_TEXT
		teamIds := []int{int(gameHeader[6].(float64)), int(gameHeader[7].(float64))} // [HOME_TEAM_ID, AWAY_TEAM_ID]
		conferences := []resultSet{eastStandings, westStandings}
		for index, teamId := range teamIds {
			team := Team{}
			team.Id = teamId
			for _, linescore := range lineScores.RowSet {
				if int(linescore[3].(float64)) == teamId { // TEAM_ID
					team.Name = linescore[6].(string) // TEAM_NAME
					team.Abbr = linescore[4].(string) // TEAM_ABBREVIATION
					team.City = linescore[5].(string) // TEAM_CITY_NAME
					if score, ok := linescore[22].(float64); ok { // PTS
						team.Score = int(score)
					}
					break
				}
			}
			for _, conference := range conferences {
				teamFound := false
				for index, standing := range conference.RowSet {
					if int(standing[0].(float64)) == teamId { // TEAM_ID
						team.GamesPlayed = int(standing[6].(float64)) // G
						team.Wins = int(standing[7].(float64)) // W
						team.Losses = int(standing[8].(float64)) // L
						team.ConferenceStanding = index
						teamFound = true
						break
					}
				}
				if teamFound {
					break
				}
			}
			if index == 0 {
				game.Teams.Home = team
			} else {
				game.Teams.Away = team
			}
		}
		games = append(games, game)
	}
	schedule := Schedule{
		Games: games,
	}
	return schedule
}

type Schedule struct {
	Games []Game
}

type Game struct {
	GameId string
	Teams struct {
		Home Team
		Away Team
	}
	GameDate time.Time
	Quarter int
	QuarterTime string
	Venue string
	Status GameStatus
}

type Team struct {
	Id int
	City string
	Name string
	Abbr string
	GamesPlayed int
	Wins int
	Losses int
	ConferenceStanding int
	Score int
}

type GameStatus struct {
	StatusCode int
	StatusString string
}

type scheduleResponse struct {
	Resource   string `json:"resource"`
	Parameters struct {
		GameDate  string `json:"GameDate"`
		LeagueID  string `json:"LeagueID"`
		DayOffset string `json:"DayOffset"`
	} `json:"parameters"`
	ResultSets []resultSet `json:"resultSets"`
}

type resultSet struct {
	Name    string          `json:"name"`
	Headers []string        `json:"headers"`
	RowSet  [][]interface{} `json:"rowSet"`
}
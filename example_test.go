package gonba

import (
	"fmt"
	"time"
)

func ExampleNewClient() {
	client := NewClient()
	fmt.Println(client)
}

func getTestTime() time.Time {
	location, _ := time.LoadLocation("America/Toronto")
	date := time.Date(2018, 10, 27, 16, 0, 0, 0, location)
	return date
}

func getTestGame() string {
	return "0021800075"
}

func ExampleClient_GetSchedule() {
	client := NewClient()
	schedule, _ := client.GetSchedule(getTestTime())
	home := schedule.Games[0].Teams.Home
	game := schedule.Games[1]
	fmt.Printf("%s, %d, %d, %d\n", home.Name, home.Score, home.GamesPlayed, home.Wins)
	fmt.Printf("%d, %s\n", game.Quarter, game.GameId)
	fmt.Printf("%s, %d", game.Status.StatusString, game.Status.StatusCode)
	// Output:
	// Pistons, 89, 5, 4
	// 4, 0021800075
	// Final, 3
}

func ExampleClient_GetScheduleV2() {
	client := NewClient()
	schedule, _ := client.GetScheduleV2(getTestTime())
	home := schedule.Games.Game[0].Home
	game := schedule.Games.Game[1]
	fmt.Printf("%s, %d, %s\n", home.Nickname, home.Score, home.Abbreviation)
	fmt.Printf("%d, %s\n", game.PeriodTime.PeriodValue, game.ID)
	// Output:
	// Pistons, 89, DET
	// 4, 0021800075
}

func ExampleClient_GetScheduleV3() {
	client := NewClient()
	schedule, _ := client.GetScheduleV3(getTestTime())
	home := schedule.Games[0].HTeam
	game := schedule.Games[1]
	fmt.Printf("%s, %d\n", home.TriCode, home.Score)
	fmt.Printf("%d, %s\n", game.Period.Current, game.GameID)
	// Output:
	// DET, 89
	// 4, 0021800075
}

func ExampleClient_GetPlayByPlay() {
	client := NewClient()
	client.GetPlayByPlay(21900193)
	// Output:
	// Pistons, 89, 5, 4
	// 4, 0021800075
}

func ExampleClient_GetPlayByPlayV2() {
	client := NewClient()
	playByPlay, _ := client.GetPlayByPlayV2(getTestTime(), getTestGame(), 1)
	fmt.Printf("%s, %d\n", playByPlay.Plays[0].Formatted.Description, playByPlay.Plays[0].PersonID)
	fmt.Printf("%s, %d\n", playByPlay.Plays[15].Clock, playByPlay.Plays[15].HTeamScore)
	// Output:
	// START OF 1ST QUARTER, 0
	// 10:26, 2
}

func ExampleClient_GetPlayByPlayV2Smart() {
	client := NewClient()
	playByPlay, _ := client.GetPlayByPlayV2All(getTestTime(), getTestGame())
	fmt.Printf("%s, %d\n", playByPlay.Plays[0].Formatted.Description, playByPlay.Plays[0].PersonID)
	fmt.Printf("%s, %d\n", playByPlay.Plays[15].Clock, playByPlay.Plays[15].HTeamScore)
	// Output:
	// START OF 1ST QUARTER, 0
	// 10:26, 2
}

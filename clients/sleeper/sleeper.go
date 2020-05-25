package sleeper

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strconv"

	"github.com/alexyou8021/sleeper-wrapper.git/entities"
)

func GetUserByUsername(username string) entities.User {
	url := "https://api.sleeper.app/v1/user/" + username
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var user entities.User
	json.Unmarshal(bodyBytes, &user)

	return user
}

func GetAllLeagues(user entities.User) []entities.League {
	user_id := user.UserId
	url := "https://api.sleeper.app/v1/user/" + user_id + "/leagues/nfl/2019"
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var leagues []entities.League
	json.Unmarshal(bodyBytes, &leagues)

	return leagues
}

func GetLeagueRosters(league entities.League) []entities.Roster {
	league_id := league.LeagueId
	url := "https://api.sleeper.app/v1/league/" + league_id + "/rosters"
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var rosters []entities.Roster
	json.Unmarshal(bodyBytes, &rosters)

	return rosters
}

func GetAllTransactions(league entities.League) []entities.Transaction {
	league_id := league.LeagueId
	totalTransactions := []entities.Transaction{}

	for week := 1; week <= 16; week++ {
		url := "https://api.sleeper.app/v1/league/" + league_id + "/transactions/" + strconv.Itoa(week)
		resp, _ := http.Get(url)
		defer resp.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		var transactions []entities.Transaction
		json.Unmarshal(bodyBytes, &transactions)
		totalTransactions = append(totalTransactions, transactions...)
	}

	return totalTransactions
}

func GetPlayers() map[string]map[string]interface{} {
	url := "https://api.sleeper.app/v1/players/nfl/"
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var players map[string]map[string]interface{}
	json.Unmarshal(bodyBytes, &players)

	return players
}
/*func GetStats(season string, week string) map[int]map[string]float32 {
	url := "https://api.sleeper.app/v1/stats/nfl/regular/" + season + "/" + week
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var stats map[int]map[string]float32
	json.Unmarshal(bodyBytes, &stats)

	return stats
}

func GetLeague(groupid string) League {
	url := "https://api.groupme.com/v3/groups/" + groupid + "?token="
	url = url + os.Getenv("token")
	resp, _ := http.Get(url)

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var league League
	json.Unmarshal(bodyBytes, &league)

	return league
}

func GetRosters(league string) []map[string]interface{} {
	url := "https://api.sleeper.app/v1/league/" + league + "/rosters"
	resp, _ := http.Get(url)

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var rosters []map[string]interface{}
	json.Unmarshal(bodyBytes, &rosters)

	return rosters
}*/

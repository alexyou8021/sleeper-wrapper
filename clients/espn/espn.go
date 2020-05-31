package espn

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/alexyou8021/sleeper-wrapper.git/entities"
)

func GetLeagueTransactions(leagueId string) []entities.ESPNTransaction {
	url := "https://fantasy.espn.com/apis/v3/games/ffl/seasons/2019/segments/0/leagues/" + leagueId + "/?view=kona_league_communication"
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var response entities.ESPNTransactionResponse
	json.Unmarshal(bodyBytes, &response)
	comms := response.Communication
	transactions := comms.Topics

	return transactions
}

func GetPlayers() []entities.ESPNPlayer {
	url := "http://fantasy.espn.com/apis/v3/games/ffl/seasons/2019/players?scoringPeriodId=0&view=players_wl"
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var players []entities.ESPNPlayer
	json.Unmarshal(bodyBytes, &players)

	return players
}
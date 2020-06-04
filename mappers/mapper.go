package mappers

import (
	"math"
	"strconv"

	"github.com/alexyou8021/sleeper-wrapper.git/entities"
	"github.com/alexyou8021/sleeper-wrapper.git/clients/db"
)

func ToTransactionResponse(transactions []entities.Transaction, league entities.League, rosters []entities.Roster) entities.TransactionResponse {
	details := []entities.TransactionDetails{}
	for _, transaction := range transactions {
		score := 0.0
		adds := []entities.Player{}
		drops := []entities.Player{}
		for key, _ := range transaction.Adds {
			if key == "OAK" {
				name := "Oakland Raiders"
				position := "DEF"
				adds = append(adds, entities.Player{
					Id:   "OAK",
					Name: name,
					Position: position,
					ImageURL: "https://sleepercdn.com/images/team_logos/nfl/oak.png",
					Hyperlink: "https://www.nfl.com/teams/las-vegas-raiders/stats",
				})
				playerScore, _ := db.QueryStats(name, strconv.Itoa(transaction.Week), position)
				score = score + playerScore.HalfPPR
				continue
			}
			player, _ := db.QueryPlayer(key)
			adds = append(adds, player)
			playerScore, _ := db.QueryStats(player.Name, strconv.Itoa(transaction.Week), player.Position)
			score = score + playerScore.HalfPPR
		}
		for key, _ := range transaction.Drops {
			if key == "OAK" {
				name := "Oakland Raiders"
				position := "DEF"
				drops = append(drops, entities.Player{
					Id:   "OAK",
					Name: name,
					Position: position,
					ImageURL: "https://sleepercdn.com/images/team_logos/nfl/oak.png",
					Hyperlink: "https://www.nfl.com/teams/las-vegas-raiders/stats",
				})
				playerScore, _ := db.QueryStats(name, strconv.Itoa(transaction.Week), position)
				score = score - playerScore.HalfPPR
				continue
			}
			player, _ := db.QueryPlayer(key)
			drops = append(drops, player)
			playerScore, _ := db.QueryStats(player.Name, strconv.Itoa(transaction.Week), player.Position)
			score = score - playerScore.HalfPPR
		}
		details = append(details, entities.TransactionDetails{
			Type: transaction.Type,
			Week: transaction.Week,
			Adds: adds,
			Drops: drops,
			Score: math.Round(score*100)/100,
		})
	}
	response := entities.TransactionResponse{
		Transactions: details,
		LeagueId: league.LeagueId,
		LeagueName: league.Name,
		LeagueMembers: rosters,
	}
	return response
}
package mappers

import (
	"math"
	"strconv"

	"github.com/alexyou8021/sleeper-wrapper.git/entities"
	"github.com/alexyou8021/sleeper-wrapper.git/clients/db"
)

func ToTransactionResponse(transactions []entities.Transaction) []entities.TransactionResponse {
	response := []entities.TransactionResponse{}
	for _, transaction := range transactions {
		score := 0.0
		adds := map[string]string{}
		drops := map[string]string{}
		for key, _ := range transaction.Adds {
			player, _ := db.QueryPlayer(key)
			adds[player.Name] = player.Id
			playerScore, _ := db.QueryStats(player.Name, strconv.Itoa(transaction.Week), player.Position)
			score = score + playerScore.HalfPPR
		}
		for key, _ := range transaction.Drops {
			player, _ := db.QueryPlayer(key)
			drops[player.Name] = player.Id
			playerScore, _ := db.QueryStats(player.Name, strconv.Itoa(transaction.Week), player.Position)
			score = score - playerScore.HalfPPR
		}
		response = append(response, entities.TransactionResponse{
			Type: transaction.Type,
			Week: transaction.Week,
			Adds: adds,
			Drops: drops,
			Score: math.Round(score*100)/100,
		})
	}
	return response
}

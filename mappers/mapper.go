package mappers

import (
	"log"
	"strconv"

	"github.com/alexyou8021/sleeper-wrapper.git/entities"
	"github.com/alexyou8021/sleeper-wrapper.git/clients/db"
)

func ToTransactionResponse(transactions []entities.Transaction) []entities.TransactionResponse {
	score := 1.4
	response := []entities.TransactionResponse{}
	for _, transaction := range transactions {
		adds := []string{}
		drops := []string{}
		for key, _ := range transaction.Adds {
			log.Println(key)
			player, _ := db.QueryPlayer(key)
			adds = append(adds, player.Name)
			playerScore, _ := db.QueryStats(player.Name, strconv.Itoa(transaction.Week))
			score = score + playerScore.HalfPPR
		}
		for key, _ := range transaction.Drops {
			log.Println(key)
			player, _ := db.QueryPlayer(key)
			drops = append(drops, player.Name)
			playerScore, _ := db.QueryStats(player.Name, strconv.Itoa(transaction.Week))
			score = score - playerScore.HalfPPR
		}
		response = append(response, entities.TransactionResponse{
			Type: transaction.Type,
			Week: transaction.Week,
			Adds: adds,
			Drops: drops,
			Score: score,
		})
	}
	return response
}

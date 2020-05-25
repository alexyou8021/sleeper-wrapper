package mappers

import (
	"log"

	"github.com/alexyou8021/sleeper-wrapper.git/entities"
	"github.com/alexyou8021/sleeper-wrapper.git/clients/db"
)

func ToTransactionResponse(transactions []entities.Transaction) []entities.TransactionResponse {
	log.Println("test")
	response := []entities.TransactionResponse{}
	for _, transaction := range transactions {
		adds := []string{}
		drops := []string{}
		for key, _ := range transaction.Adds {
			log.Println(key)
			player, _ := db.QueryPlayer(key)
			adds = append(adds, player.Name)
		}
		for key, _ := range transaction.Drops {
			log.Println(key)
			player, _ := db.QueryPlayer(key)
			drops = append(drops, player.Name)
		}
		response = append(response, entities.TransactionResponse{
			Type: transaction.Type,
			Week: transaction.Week,
			Adds: adds,
			Drops: drops,
			Score: transaction.Score,
		})
	}
	return response
}

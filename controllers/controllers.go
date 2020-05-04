package controllers

import (
	"github.com/alexyou8021/sleeper-wrapper.git/clients/sleeper"
	"github.com/alexyou8021/sleeper-wrapper.git/entities"
	"github.com/alexyou8021/sleeper-wrapper.git/mappers"
)

func Controller(username string) []entities.TransactionResponse {
	user := sleeper.GetUserByUsername(username)
	leagues := sleeper.GetAllLeagues(user)
	transactions := getTransactionsForUser(user, leagues[0])
	response := mappers.ToTransactionResponse(transactions)
	return response
}

func getTransactionsForUser(user entities.User, league entities.League) []entities.Transaction {
	transactions := sleeper.GetAllTransactions(league)
	rosterId := getRosterIdFromLeagueForUser(user, league)
	result := getTransactionsFromRosterId(transactions, rosterId)

	return result
}

func getRosterIdFromLeagueForUser(user entities.User, league entities.League) int {
	rosters := sleeper.GetLeagueRosters(league)
	for _, roster := range rosters {
		if roster.OwnerId == user.UserId {
			return roster.RosterId
		}
	}
	return -1
}

func getTransactionsFromRosterId(transactions []entities.Transaction, rosterId int) (result []entities.Transaction) {
	for _, transaction := range transactions {
		if transaction.Status != "complete" {
			continue
		}
		if transaction.Type == "commissioner" {
			continue
		}
		for _, tRosterId := range transaction.RosterIds {
			if tRosterId == rosterId {
				result = append(result, transaction)
				continue
			}
		}
	}
	return result
}

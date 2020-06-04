package controllers

import (
	"github.com/alexyou8021/sleeper-wrapper.git/clients/espn"
	"github.com/alexyou8021/sleeper-wrapper.git/clients/sleeper"
	"github.com/alexyou8021/sleeper-wrapper.git/entities"
	"github.com/alexyou8021/sleeper-wrapper.git/mappers"
	"github.com/alexyou8021/sleeper-wrapper.git/clients/db"

	"log"
	"strconv"
)

func Controller(username string) entities.TransactionResponse {
	user := sleeper.GetUserByUsername(username)
	leagues := sleeper.GetAllLeagues(user)
	league := leagues[0]
	rosters := sleeper.GetLeagueRosters(league)
	transactions := getTransactionsForUser(user, league, rosters)
	response := mappers.ToTransactionResponse(transactions, league, rosters)

	return response
}

func GetESPNTransactions(id string) entities.TransactionResponse {
	var details []entities.TransactionDetails

	transactions := espn.GetLeagueTransactions(id)
	for _, transaction := range transactions {
		rosterId := transaction.RosterId
		atype := transaction.Type
		messages := transaction.Messages
		totalScore := 0.0
		var ttype string
		var adds []entities.Player
		var drops []entities.Player

		if atype != "ACTIVITY_TRANSACTIONS" {
			continue
		}
		if rosterId != 0 {
			continue
		}
		for _, message := range messages {
			targetId := message.TargetId
			typeId := message.TypeId
			typeString := entities.ActivityMap[typeId]
			if typeString == "" {
				continue
			}
			player, _ := db.QueryESPNPlayer(strconv.Itoa(targetId))
			if typeString == "dropped" {
				drops = append(drops, player)
			} else {
				adds = append(adds, player)
				ttype = typeString
			}
		}
		obj := entities.TransactionDetails {
			Type: ttype,
			Week: 1,
			Adds: adds,
			Drops: drops,
			Score: totalScore,
		}
		log.Println(obj)
		details = append(details, obj)
	}
	response := entities.TransactionResponse{
		Transactions: details,
		LeagueId: id,
	}
	return response
}

func getTransactionsForUser(user entities.User, league entities.League, rosters []entities.Roster) []entities.Transaction {
	transactions := sleeper.GetAllTransactions(league)
	rosterId := getRosterIdForUser(user, rosters)
	result := getTransactionsFromRosterId(transactions, rosterId)

	return result
}

func getRosterIdForUser(user entities.User, rosters []entities.Roster) int {
	for _, roster := range rosters {
		if roster.OwnerId == user.UserId {
			return roster.RosterId
		}
		for _, coownerId := range roster.CoOwnersIds {
			if coownerId == user.UserId {
				return roster.RosterId
			}
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
		for key, tRosterId := range transaction.Adds {
			if tRosterId != rosterId {
				delete(transaction.Adds, key)
			}
		}
		for key, tRosterId := range transaction.Drops {
			if tRosterId != rosterId {
				delete(transaction.Drops, key)
			}
		}
		for _, tRosterId := range transaction.RosterIds {
			if tRosterId == rosterId {
				transaction.Score = calculateTransactionScore(transaction)
				result = append(result, transaction)
				continue
			}
		}
	}

	return result
}

func calculateTransactionScore(transaction entities.Transaction) float64 {
	return 1.24
}

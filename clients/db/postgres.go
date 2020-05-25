package db

import (
	"log"
	"os"
	"strings"
	"strconv"

	"database/sql"
	_ "github.com/lib/pq"

	"github.com/alexyou8021/sleeper-wrapper.git/entities"
	"github.com/alexyou8021/sleeper-wrapper.git/clients/sleeper"
	"github.com/alexyou8021/sleeper-wrapper.git/clients/datapros"
)

var (
	db  *sql.DB
)

func RemakePlayersTable() {
	CreatePlayersTable()
	StorePlayers()
}

func RemakeStatsTable() {
	CreateStatsTable()
	StoreStats()
}

func CreatePlayersTable() {
	if db == nil {
		db, _ = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	}

	result, err := db.Exec("DROP TABLE players;")
	result, err = db.Exec("CREATE TABLE players (id varchar(255), name varchar(255), position varchar(255));")
	log.Println(result)
	log.Println(err)
}

func CreateStatsTable() {
	if db == nil {
		db, _ = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	}

	result, err := db.Exec("DROP TABLE stats;")
	result, err = db.Exec("CREATE TABLE stats (name varchar(255), week int, position varchar(255), team varchar(255), halfppr float, ppr float, standard float);")
	log.Println(result)
	log.Println(err)
}

func StorePlayers() {
	if db == nil {
		db, _ = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	}
	players := sleeper.GetPlayers()

	for _, value := range players {
		id, _ := value["player_id"].(string)
		name, _ := value["full_name"].(string)
		name = strings.Replace(name, "'", "''", 1)
		if name == "" {
			name = value["first_name"].(string) + " " + value["last_name"].(string)
		}
		position, _ := value["position"].(string)
		execCmd := "INSERT INTO players VALUES ('" + id + "', '" + name + "', '" + position + "');"
		log.Println(execCmd)
		_, err := db.Exec(execCmd)
		if err != nil {
			log.Fatal(err)
			break
		}
	}
}

func StoreStats() {
	if db == nil {
		db, _ = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	}


	result := datapros.GetStatsFrom("2019", "1")
	for _, stats := range result {
		name := strings.Replace(stats.Name, "'", "''", 1)
		week := "1"
		position := stats.Position
		team := stats.Team
		halfppr := strconv.FormatFloat(stats.FantasyPoints["half_ppr"], 'f', 2, 64)
		ppr := strconv.FormatFloat(stats.FantasyPoints["ppr"], 'f', 2, 64)
		standard := strconv.FormatFloat(stats.FantasyPoints["standard"], 'f', 2, 64)
		execCmd := "INSERT INTO stats VALUES ('" + name + "', " + week + ", '" + position + "', '" + team + "', " + halfppr + ", " + ppr + ", " + standard + ");"
		log.Println(execCmd)
		_, err := db.Exec(execCmd)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func QueryPlayer(id string) (entities.Player, error) {
	var player entities.Player

	if db == nil {
		db, _ = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	}

	log.Println("SELECT * FROM players WHERE id='" + id+ "';")
	result, err := db.Query("SELECT * FROM players WHERE id='" + id+ "';")
	if err != nil {
		log.Fatal(err)
		return player, err
	}

	for result.Next() {
		err = result.Scan(&player.Id, &player.Name, &player.Position)
		if err != nil {
			log.Fatal(err)
			return player, err
		}
	}

	return player, nil
}

func QueryStats(name string) (entities.Stats, error) {
	var stats entities.Stats

	if db == nil {
		db, _ = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	}

	result, err := db.Query("SELECT * FROM stats WHERE name='" + name + "';")
	if err != nil {
		log.Fatal(err)
	}

	for result.Next() {
		err = result.Scan(&stats.Name, &stats.Week, &stats.Position, &stats.Team, &stats.HalfPPR, &stats.PPR, &stats.Standard)
		if err != nil {
			log.Fatal(err)
		}
	}

	return stats, nil
}

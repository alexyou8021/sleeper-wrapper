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


	for i := 1; i <= 16; i++ {
		weekString := strconv.Itoa(i)
		result := datapros.GetStatsFrom("2019", weekString)
		for _, stats := range result {
			name := strings.Replace(stats.Name, "'", "''", 1)
			week := weekString
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
}

func QueryPlayer(id string) (entities.Player, error) {
	var player entities.Player

	if db == nil {
		db, _ = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	}

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

func QueryStats(name string, week string, position string) (entities.Stats, error) {
	var stats entities.Stats

	if db == nil {
		db, _ = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	}

	queryName := strings.Replace(name, " Jr.", "", 1)
	queryName = strings.Replace(queryName, " Sr.", "", 1)
	queryName = strings.Replace(queryName, " IV", "", 1)
	queryName = strings.Replace(queryName, " V", "", 1)
	queryName = strings.Replace(queryName, " III", "", 1)
	queryName = strings.Replace(queryName, " II", "", 1)
	queryName = strings.Replace(queryName, "'", "''", 1)
	log.Println(queryName + " " + week)

	result, err := db.Query("SELECT name, position, team, sum(halfppr), sum(ppr), sum(standard) FROM stats WHERE name='" + queryName + "' AND week > " + week + " GROUP BY name, position, team;")
	if err != nil {
		log.Fatal(err)
	}

	for result.Next() {
		err = result.Scan(&stats.Name, &stats.Position, &stats.Team, &stats.HalfPPR, &stats.PPR, &stats.Standard)
		if err != nil {
			log.Fatal(err)
		}
	}
	stats.Week = week
	stats.Name = name
	if position == "DEF" || position == "K" {
		weekNum, _ := strconv.ParseFloat(week, 64)
		remainingWeeks := 16.0 - weekNum
		stats.HalfPPR = 8.0 * remainingWeeks
		stats.PPR = 8.0 * remainingWeeks
		stats.Standard = 8.0 * remainingWeeks
	}

	log.Println(stats)
	return stats, nil
}

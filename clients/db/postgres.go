package db

import (
	"log"
	"os"
	"strings"
	"strconv"

	"database/sql"
	_ "github.com/lib/pq"

	"github.com/alexyou8021/sleeper-wrapper.git/entities"
	"github.com/alexyou8021/sleeper-wrapper.git/clients/espn"
	"github.com/alexyou8021/sleeper-wrapper.git/clients/sleeper"
	"github.com/alexyou8021/sleeper-wrapper.git/clients/datapros"
)

var (
	db  *sql.DB
)

func RemakeSleeperPlayersTable() {
	CreateSleeperPlayersTable()
	StoreSleeperPlayers()
}

func RemakeESPNPlayersTable() {
	CreateESPNPlayersTable()
	StoreESPNPlayers()
}

func RemakeStatsTable() {
	CreateStatsTable()
	StoreStats()
}

func CreateSleeperPlayersTable() {
	if db == nil {
		db, _ = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	}

	result, err := db.Exec("DROP TABLE sleeper;")
	result, err = db.Exec("CREATE TABLE sleeper (id varchar(255), name varchar(255), position varchar(255));")
	log.Println(result)
	log.Println(err)
}

func CreateESPNPlayersTable() {
	if db == nil {
		db, _ = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	}

	result, err := db.Exec("DROP TABLE espn;")
	result, err = db.Exec("CREATE TABLE espn (id varchar(255), name varchar(255), position varchar(255));")
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

func StoreSleeperPlayers() {
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
		execCmd := "INSERT INTO sleeper VALUES ('" + id + "', '" + name + "', '" + position + "');"
		log.Println(execCmd)
		_, err := db.Exec(execCmd)
		if err != nil {
			log.Fatal(err)
			break
		}
	}
	log.Println(len(players))
}

func StoreESPNPlayers() {
	if db == nil {
		db, _ = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	}
	players := espn.GetPlayers()

	for _, value := range players {
		id := strconv.Itoa(value.Id)
		name := value.FullName
		name = strings.Replace(name, "'", "''", 1)
		if name == "" {
			name = value.FirstName + " " + value.LastName
		}
		position := entities.PositionMap[value.PositionId]
		execCmd := "INSERT INTO espn VALUES ('" + id + "', '" + name + "', '" + position + "');"
		log.Println(execCmd)
		_, err := db.Exec(execCmd)
		if err != nil {
			log.Fatal(err)
			break
		}
	}
	log.Println(len(players))
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
			name = strings.ReplaceAll(name, ".", "")
			week := weekString
			position := stats.Position
			team := stats.Team
			halfppr := strconv.FormatFloat(stats.FantasyPoints["half_ppr"], 'f', 2, 64)
			ppr := strconv.FormatFloat(stats.FantasyPoints["ppr"], 'f', 2, 64)
			standard := strconv.FormatFloat(stats.FantasyPoints["standard"], 'f', 2, 64)
			if halfppr == "0.00" && ppr == "0.00" && standard == "0.00" {
				continue
			}
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
	log.Println("QuerySleeperPlayer: " + id)
	var player entities.Player

	if db == nil {
		db, _ = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	}

	result, err := db.Query("SELECT * FROM sleeper WHERE id='" + id+ "';")
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
	if player.Position == "DEF" {
		formattedName := strings.ReplaceAll(player.Name, " ", "-")
		player.ImageURL = "https://sleepercdn.com/images/team_logos/nfl/" + strings.ToLower(player.Id) + ".png"
		player.Hyperlink = "https://www.nfl.com/teams/" + formattedName + "/stats"
	} else {
		replacedName := strings.ReplaceAll(player.Name, ".", "-")
		replacedName = strings.ReplaceAll(replacedName, "'", "-")
		splitName := strings.Split(replacedName, " ")
		firstName := splitName[0]
		if len(firstName) == 2 && firstName != "Bo" && firstName != "Ty" {
			firstName = firstName[:1] + "-" + firstName[1:]
		}
		lastName := splitName[1]
		formattedName := strings.Replace(firstName + "-" + lastName, "--", "-", 1)
		player.ImageURL = "https://sleepercdn.com/content/nfl/players/" + player.Id + ".jpg";
		player.Hyperlink = "https://www.nfl.com/players/" + formattedName + "/stats/logs"
	}

	return player, nil
}


func QueryESPNPlayer(id string) (entities.Player, error) {
	log.Println("QueryESPNPlayer: " + id)
	var player entities.Player

	if db == nil {
		db, _ = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	}

	result, err := db.Query("SELECT * FROM espn WHERE id='" + id+ "';")
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
	if player.Position == "D/ST" {
		dstId := entities.DstMap[player.Id]
		//formattedName := strings.ReplaceAll(player.Name, " ", "-")
		player.ImageURL = "https://a.espncdn.com/i/teamlogos/nfl/500/" + dstId + ".png"
		//player.Hyperlink = "https://www.nfl.com/teams/" + formattedName + "/stats"
	} else {
		replacedName := strings.ReplaceAll(player.Name, ".", "-")
		replacedName = strings.ReplaceAll(replacedName, "'", "-")
		splitName := strings.Split(replacedName, " ")
		firstName := splitName[0]
		if len(firstName) == 2 && firstName != "Bo" && firstName != "Ty" {
			firstName = firstName[:1] + "-" + firstName[1:]
		}
		lastName := splitName[1]
		formattedName := strings.Replace(firstName + "-" + lastName, "--", "-", 1)
		player.ImageURL = "https://a.espncdn.com/combiner/i?img=/i/headshots/nfl/players/full/" + player.Id + ".png";
		player.Hyperlink = "https://www.nfl.com/players/" + formattedName + "/stats/logs"
	}

	return player, nil
}

func QueryStats(name string, week string, position string) (entities.Stats, error) {
	log.Println("QueryStats: " + name + " " + week + " " + position)
	var stats entities.Stats

	if db == nil {
		db, _ = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	}

	splitName := strings.Split(name, " ")
	queryName := splitName[0] + " " + splitName[1]
	queryName = strings.Replace(queryName, "'", "''", 1)
	queryName = strings.ReplaceAll(queryName, ".", "")
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

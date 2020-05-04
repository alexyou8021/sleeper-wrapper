package db

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"database/sql"
	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
)

type Player struct {
	Id       int    `json: id`
	Name     string `json: name`
	Position string `json: position`
}

func RemakePlayersTable() {
	CreatePlayersTable()
	StorePlayers()
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

func StorePlayers() {
	if db == nil {
		db, _ = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	}
	url := "https://api.sleeper.app/v1/players/nfl/"
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var players map[int]map[string]interface{}
	json.Unmarshal(bodyBytes, &players)

	for _, value := range players {
		id, _ := value["player_id"].(string)
		name, _ := value["full_name"].(string)
		name = strings.Replace(name, "'", "''", 1)
		position, _ := value["position"].(string)
		log.Println(name + " " + position + " " + id)
		_, err := db.Exec("INSERT INTO players VALUES (" + id + ", '" + name + "', '" + position + "');")
		if err != nil {
			log.Fatal(err)
			break
		}
	}
}

func QueryPlayer(id string) (_ Player, _ error) {
	var player Player

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

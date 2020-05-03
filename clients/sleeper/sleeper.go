package sleeper

import (
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"

	"github.com/alexyou8021/sleeper-wrapper.git/entities"
)

func GetUserByUsername(username string) entities.User {
        username = "376687828232613888"
	url := "https://api.sleeper.app/v1/user/" + username
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var user entities.User
	json.Unmarshal(bodyBytes, &user)
	log.Println(user)

	return user
}

/*func GetStats(season string, week string) map[int]map[string]float32 {
	url := "https://api.sleeper.app/v1/stats/nfl/regular/" + season + "/" + week
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var stats map[int]map[string]float32
	json.Unmarshal(bodyBytes, &stats)

	return stats
}

func GetLeague(groupid string) League {
	url := "https://api.groupme.com/v3/groups/" + groupid + "?token="
	url = url + os.Getenv("token")
	resp, _ := http.Get(url)

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var league League
	json.Unmarshal(bodyBytes, &league)

	return league
}

func GetRosters(league string) []map[string]interface{} {
	url := "https://api.sleeper.app/v1/league/" + league + "/rosters"
	resp, _ := http.Get(url)

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var rosters []map[string]interface{}
	json.Unmarshal(bodyBytes, &rosters)

	return rosters
}*/

package datapros

import (
	"net/http"
	"encoding/json"
	"io/ioutil"

	"github.com/alexyou8021/sleeper-wrapper.git/entities"
)

func GetStatsFrom(year string, week string) []entities.Stats {
	url := "https://www.fantasyfootballdatapros.com/api/players/" + year + "/" + week
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var stats []entities.Stats
	json.Unmarshal(bodyBytes, &stats)

	return stats
}


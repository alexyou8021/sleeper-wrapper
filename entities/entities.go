package entities

type User struct {
	Username    string `json:"username"`
	UserId      string `json:"user_id"`
	DisplayName string `json:"display_name"`
}

type League struct {
	TotalRosters	int	`json:"total_rosters"`
	LeagueId	string	`json:"league_id"`
	Name		string	`json:"name"`
}

type Roster struct {
	RosterId	int	`json:"roster_id"`
	LeagueId	string	`json:"league_id"`
	OwnerId		string	`json:"owner_id"`
}

type Transaction struct {
	Type		string		`json:"type"`
	Status		string		`json:"status"`
	Week		int		`json:"leg"`
	RosterIds	[]int		`json:"roster_ids"`
	Drops		map[string]int	`json:"drops"`
	Adds		map[string]int	`json:"adds"`
}

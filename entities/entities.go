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
	RosterId	int	 `json:"roster_id"`
	LeagueId	string	 `json:"league_id"`
	OwnerId		string	 `json:"owner_id"`
	CoOwnersIds     []string `json:"co_owners"`
}

type Transaction struct {
	Type		string		`json:"type"`
	Status		string		`json:"status"`
	Week		int		`json:"leg"`
	RosterIds	[]int		`json:"roster_ids"`
	Adds		map[string]int	`json:"adds"`
	Drops		map[string]int	`json:"drops"`
	Score		float64		`json:"score"`
}

type TransactionDetails struct {
	Type		string		`json:"type"`
	Week		int		`json:"week"`
	Adds		[]Player	`json:"adds"`
	Drops		[]Player	`json:"drops"`
	Score		float64		`json:"score"`
}

type TransactionResponse struct {
	Transactions	[]TransactionDetails	`json:"transaction_details"`
	LeagueId	string			`json:"league_id"`
	LeagueName	string			`json:"league_name"`
	LeagueMembers	[]Roster		`json:"league_members"`
}
type Player struct {
        Id		string `json:"id"`
        Name     	string `json:"name"`
        Position 	string `json:"position"`
        ImageURL  	string `json:"image_url"`
        Hyperlink 	string `json:"hyperlink"`
}

type Stats struct {
	Name		string `json:"player_name"`
	Week		string
	Position	string `json:"position"`
	Team		string `json:"team"`
	FantasyPoints	map[string]float64 `json:"fantasy_points"`
	HalfPPR		float64
	PPR		float64
	Standard	float64
}

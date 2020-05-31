package entities

var ActivityMap = map[int]string{
	178: "free_agent",
	180: "waiver",
	179: "dropped",
	181: "dropped",
	239: "dropped",
	244: "traded",
}

var PositionMap = map[int]string{
    0: "QB",
    1: "TQB",
    2: "RB",
    3: "RB/WR",
    4: "WR",
    5: "WR/TE",
    6: "TE",
    7: "OP",
    8: "DT",
    9: "DE",
    10: "LB",
    11: "DL",
    12: "CB",
    13: "S",
    14: "DB",
    15: "DP",
    16: "D/ST",
    17: "K",
    18: "P",
    19: "HC",
    20: "BE",
    21: "IR",
    22: "",
    23: "RB/WR/TE",
    24: "ER",
    25: "Rookie",
}

var DstMap = map[string]string{
    "-16001" : "ATL",
    "-16002" : "BUF",
    "-16003" : "CHI",
    "-16004" : "CIN",
    "-16005" : "CLE",
    "-16006" : "DAL",
    "-16007" : "DEN",
    "-16008" : "DET",
    "-16009" : "GB",
    "-16010" : "TEN",
    "-16011" : "IND",
    "-16012" : "KC",
    "-16013" : "OAK",
    "-16014" : "LAR",
    "-16015" : "MIA",
    "-16016" : "MIN",
    "-16017" : "NE",
    "-16018" : "NO",
    "-16019" : "NYG",
    "-16020" : "NYJ",
    "-16021" : "PHI",
    "-16022" : "ARI",
    "-16023" : "PIT",
    "-16024" : "LAC",
    "-16025" : "SF",
    "-16026" : "SEA",
    "-16027" : "TB",
    "-16028" : "WSH",
    "-16029" : "CAR",
    "-16030" : "JAX",
    "-16033" : "BAL",
    "-16034" : "HOU",
}

type ESPNTransactionResponse struct {
	Communication	Topics	`json:"communication"`
}

type Topics struct {
	Topics	[]ESPNTransaction	`json:"topics"`
}

type ESPNTransaction struct {
	RosterId	int		`json:"targetId"`
	Messages	[]Message	`json:"messages"`
	Type		string		`json:"type"`
}

type Message struct {
	TypeId		int	`json:"messageTypeId"`
	TargetId	int	`json:"targetId"`
}

type ESPNPlayer struct {
	Id		int	`json:"id"`
	FirstName	string	`json:"firstName"`
	LastName	string	`json:"lastName"`
	FullName	string	`json:"fullName"`
	PositionId	int	`json:"defaultPositionId"`
}
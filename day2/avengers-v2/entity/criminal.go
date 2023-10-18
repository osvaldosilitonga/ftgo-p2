package entity

type Criminal struct {
	ID          int    `json:"id"`
	HeroId      int    `json:"hero"`
	VillainId   int    `json:"villain"`
	Location    string `json:"location"`
	Date        string `json:"date"`
	Description string `json:"desc"`
	Status      string `json:"status"`
}

type CriminalReports struct {
	ID          int    `json:"id"`
	Hero        string `json:"hero"`
	Villain     string `json:"villain"`
	Location    string `json:"location"`
	Date        string `json:"date"`
	Description string `json:"desc"`
	Status      string `json:"status"`
}

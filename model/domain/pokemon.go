package domain

type Pokemon struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	About About  `json:"about"`
	Types string `json:"types"`
	Stats Stats  `json:"stats"`
}

type About struct {
	Id_about  int    `json:"id_about"`
	Species   string `json:"species"`
	Height    int    `json:"height"`
	Weight    int    `json:"weight"`
	Abilities string `json:"abilities"`
}

type Stats struct {
	Id_stats int `json:"id_stats"`
	HP       int `json:"hp"`
	Attack   int `json:"attack"`
	Defense  int `json:"defense"`
	Speed    int `json:"speed"`
}
package data

type Hero struct {
	Id int `json:"id"`
	Name string `json:"name"`
	RealName string `json:"real_name"`
	Health int `json:"health"`
	Armour int `json:"armour"`
	Shield int 	`json:"shield"`
}

type HeroesResponse struct {
	Total int `json:"total"`
	Data []Hero `json:"data"`
}

type HeroResponse struct {
	Hero
	Abilities []Ability
}

type Ability struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	IsUltimate bool `json:"is_ultimate"`
	HeroObject Hero `json:"hero"`
}

type AbilitiesResponse struct {
	Total int `json:"total"`
	Data []Ability `json:"data"`
}

type AbilityResponse struct {
	Ability
}
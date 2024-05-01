package dto

type StoreInputDTO struct {
	Group Group `json:"group"`
}

type Group struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

package models

type Track struct {
	Name   string `json:"name"`
	Artist struct {
		Name string `json:"name"`
	} `json:"artist"`
}

type LastfmResponse struct {
	Tracks struct {
		Track []Track `json:"track"`
	} `json:"tracks"`
}

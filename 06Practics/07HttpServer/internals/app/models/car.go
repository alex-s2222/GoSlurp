package models

type Car struct { 
	Id int64 `json:"id"`
	Colour string `json:"colour"`
	Brand string `json:"brand"`
	LicensePlate string `json:"locense_plate"`
	Owner User `json:"owner"`
}
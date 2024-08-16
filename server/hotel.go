package main

type Hotel struct {
	Id          string `json:"id" binding:"required"`
	DisplayName string `json:"display_name"`
	StarRating  int    `json:"star_rating"`
	NoRooms     int    `json:"no_rooms"`
	Links       []Link `json:"links"`
}

type Link struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
	Type string `json:"type"`
}

var (
	repository map[string]*Hotel
)

func init() {
	repository = make(map[string]*Hotel)
}

package store

type Graffity struct {
	Image     string  `json:"image"`
	Id        string  `json:"id"`
	UserID    string  `json:"userid"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Height    float64 `json:"height"`
	Message   string  `json:"message"`
	Gang      string  `json:"gang"`
}

type User struct {
	UserID     string   `bson:"userid"`
	Gang       string   `bson:"gang"`
	Graffities []string `bson:"graffities"`
}

type Login struct {
	UserID string `json:"userid"`
	Gang   string `json:"gang"`
}

type Coordinates struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}
type MapZone struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Gang      string  `json:"gang"`
}

package model

// omit is the bool type for omitting a field of struct.
type omit bool

// Verify struct.
type Verify struct {
	IsVerified bool   `json:"isVerified" bson:"isVerified"`
	Code       string `json:"code"`
	CreatedAt  int64  `json:"createdAt" bson:"createdAt"`
}

// GeoJSON  struct
type GeoJSON struct {
	Type        string    `json:"-"`
	Coordinates []float64 `json:"coordinates"`
}

// GeoLocation is geo struct of simple location
type GeoLocation struct {
	PlaceID string   `json:"placeId" bson:"placeId" `
	Address string   `json:"address"`
	GeoJSON *GeoJSON `json:"geoJson"`
}

// Status is alias for status of every models
type Status int

package model

type Service struct {
	ID              string  `json:"id" bson:"id"`
	Name            string  `json:"name" bson:"name"`
	Description     string  `json:"description" bson:"description"`
	Price           float64 `json:"price" bson:"price"`
	ProviderID      string  `json:"provider_id" bson:"provider_id"`
	Category        string  `json:"category" bson:"category"`
	CategoryId      string  `json:"category_id" bson:"category_id"`
	AvgRating       float64 `json:"avg_rating" bson:"avg_rating"`
	RatingCount     int64   `json:"rating_count" bson:"rating_count"`
	ProviderName    string  `json:"provider_name,omitempty" bson:"provider_name"`
	ProviderContact string  `json:"provider_contact,omitempty" bson:"provider_contact"`
	ProviderAddress string  `json:"provider_address,omitempty" bson:"provider_address"`
	ProviderRating  float64 `json:"provider_rating,omitempty" bson:"provider_rating"`
}

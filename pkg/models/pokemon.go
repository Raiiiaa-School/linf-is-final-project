package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Pokemon struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string             `bson:"name" json:"name"`
	Type      []string           `bson:"type" json:"type"`
	Abilities []string           `bson:"abilities" json:"abilities"`
	Height    float64            `bson:"height" json:"height"`
	Weight    float64            `bson:"weight" json:"weight"`
	PokedexID int                `bson:"pokedex_id" json:"pokedex_id"`
}

package graph

type CreatePokemonInput struct {
	Name      string   `json:"name"`
	Type      []string `json:"type"`
	Abilities []string `json:"abilities"`
	Height    float64  `json:"height"`
	Weight    float64  `json:"weight"`
	PokedexID int      `json:"pokedexId"`
}

type UpdatePokemonInput struct {
	Name      *string   `json:"name,omitempty"`
	Type      []string  `json:"type,omitempty"`
	Abilities []string  `json:"abilities,omitempty"`
	Height    *float64  `json:"height,omitempty"`
	Weight    *float64  `json:"weight,omitempty"`
	PokedexID *int      `json:"pokedexId,omitempty"`
}

type CreateUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string `json:"token"`
}
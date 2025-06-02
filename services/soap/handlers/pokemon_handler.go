package handlers

import (
	"context"
	"encoding/xml"
	"log"
	"net/http"
	"project-is/pkg/db"
	"project-is/pkg/messaging"
	"project-is/pkg/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SoapEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    SoapBody
}

type SoapBody struct {
	XMLName  xml.Name     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	Request  *PokemonRequest `xml:",omitempty"`
	Response interface{}  `xml:",omitempty"`
}

type PokemonRequest struct {
	XMLName xml.Name `xml:"PokemonRequest"`
	Action  string   `xml:"action"`
	Pokemon struct {
		ID        string   `xml:"id,omitempty"`
		Name      string   `xml:"name"`
		Type      []string `xml:"type"`
		Abilities []string `xml:"abilities"`
		Height    float64  `xml:"height"`
		Weight    float64  `xml:"weight"`
		PokedexID int      `xml:"pokedex_id"`
	} `xml:"pokemon"`
}

type PokemonResponse struct {
	XMLName  xml.Name        `xml:"PokemonResponse"`
	Status   string          `xml:"status"`
	Message  string          `xml:"message"`
	Pokemon  *models.Pokemon `xml:"pokemon,omitempty"`
	Pokemons []models.Pokemon `xml:"pokemons,omitempty"`
}

func HandlePokemonService(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendSoapFault(w, "Invalid HTTP method", "Only POST method is supported")
		return
	}

	var envelope SoapEnvelope
	if err := xml.NewDecoder(r.Body).Decode(&envelope); err != nil {
		sendSoapFault(w, "Invalid request", err.Error())
		return
	}

	if envelope.Body.Request == nil {
		sendSoapFault(w, "Invalid request", "Missing PokemonRequest in SOAP Body")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var response PokemonResponse

	switch envelope.Body.Request.Action {
	case "create":
		pokemon := &models.Pokemon{
			Name:      envelope.Body.Request.Pokemon.Name,
			Type:      envelope.Body.Request.Pokemon.Type,
			Abilities: envelope.Body.Request.Pokemon.Abilities,
			Height:    envelope.Body.Request.Pokemon.Height,
			Weight:    envelope.Body.Request.Pokemon.Weight,
			PokedexID: envelope.Body.Request.Pokemon.PokedexID,
		}

		res, err := db.CreatePokemon(ctx, pokemon)
		if err != nil {
			sendSoapFault(w, "Database error", err.Error())
			return
		}

		if err := messaging.PublishPokemonEvent("created", pokemon); err != nil {
			log.Printf("Error publishing pokemon created event: %v", err)
		}

		response.Status = "success"
		response.Message = "Pokemon created successfully"
		response.Pokemon = pokemon
		pokemon.ID = res.InsertedID.(primitive.ObjectID)

	case "get":
		if envelope.Body.Request.Pokemon.ID == "" {
			pokemons, err := db.GetAllPokemons(ctx)
			if err != nil {
				sendSoapFault(w, "Database error", err.Error())
				return
			}
			response.Status = "success"
			response.Pokemons = pokemons
		} else {
			id, err := primitive.ObjectIDFromHex(envelope.Body.Request.Pokemon.ID)
			if err != nil {
				sendSoapFault(w, "Invalid ID", "Invalid Pokemon ID format")
				return
			}

			pokemon, err := db.GetPokemonByID(ctx, id)
			if err != nil {
				sendSoapFault(w, "Database error", err.Error())
				return
			}

			if pokemon == nil {
				response.Status = "error"
				response.Message = "Pokemon not found"
			} else {
				response.Status = "success"
				response.Pokemon = pokemon
			}
		}

	case "update":
		id, err := primitive.ObjectIDFromHex(envelope.Body.Request.Pokemon.ID)
		if err != nil {
			sendSoapFault(w, "Invalid ID", "Invalid Pokemon ID format")
			return
		}

		pokemon := &models.Pokemon{
			ID:        id,
			Name:      envelope.Body.Request.Pokemon.Name,
			Type:      envelope.Body.Request.Pokemon.Type,
			Abilities: envelope.Body.Request.Pokemon.Abilities,
			Height:    envelope.Body.Request.Pokemon.Height,
			Weight:    envelope.Body.Request.Pokemon.Weight,
			PokedexID: envelope.Body.Request.Pokemon.PokedexID,
		}

		res, err := db.UpdatePokemon(ctx, id, pokemon)
		if err != nil {
			sendSoapFault(w, "Database error", err.Error())
			return
		}

		if res.ModifiedCount == 0 {
			response.Status = "error"
			response.Message = "Pokemon not found or no changes made"
		} else {
			if err := messaging.PublishPokemonEvent("updated", pokemon); err != nil {
				log.Printf("Error publishing pokemon updated event: %v", err)
			}
			response.Status = "success"
			response.Message = "Pokemon updated successfully"
			response.Pokemon = pokemon
		}

	case "delete":
		id, err := primitive.ObjectIDFromHex(envelope.Body.Request.Pokemon.ID)
		if err != nil {
			sendSoapFault(w, "Invalid ID", "Invalid Pokemon ID format")
			return
		}

		res, err := db.DeletePokemon(ctx, id)
		if err != nil {
			sendSoapFault(w, "Database error", err.Error())
			return
		}

		if res.DeletedCount == 0 {
			response.Status = "error"
			response.Message = "Pokemon not found"
		} else {
			if err := messaging.PublishPokemonEvent("deleted", nil, id); err != nil {
				log.Printf("Error publishing pokemon deleted event: %v", err)
			}
			response.Status = "success"
			response.Message = "Pokemon deleted successfully"
		}

	default:
		sendSoapFault(w, "Invalid action", "Unsupported action requested")
		return
	}

	sendSoapResponse(w, response)
}

func sendSoapFault(w http.ResponseWriter, faultString, detail string) {
	w.Header().Set("Content-Type", "application/xml")
	fault := SoapFault{
		Code:   "Server",
		String: faultString,
		Detail: detail,
	}
	envelope := SoapEnvelope{
		Body: SoapBody{
			Response: fault,
		},
	}
	xml.NewEncoder(w).Encode(envelope)
}

func sendSoapResponse(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/xml")
	envelope := SoapEnvelope{
		Body: SoapBody{
			Response: response,
		},
	}
	xml.NewEncoder(w).Encode(envelope)
}

type SoapFault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`
	Code    string   `xml:"faultcode"`
	String  string   `xml:"faultstring"`
	Detail  string   `xml:"detail"`
}
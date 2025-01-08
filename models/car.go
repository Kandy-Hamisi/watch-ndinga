package models

import (
	"errors"
	"github.com/google/uuid"
	"strconv"
	"time"
)

type Car struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Year      string    `json:"year"`
	Brand     string    `json:"brand"`
	FuelType  string    `json:"fuel_type"`
	Engine    Engine    `json:"engine"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CarRequest struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Year     string    `json:"year"`
	Brand    string    `json:"brand"`
	FuelType string    `json:"fuel_type"`
	Engine   Engine    `json:"engine"`
	Price    float64   `json:"price"`
}

func validateName(name string) error {
	if name == "" {
		return errors.New("name is required")
	}
	return nil
}

func validateYear(year string) error {
	if year == "" {
		return errors.New("year is required")
	}

	_, err := strconv.Atoi(year)

	if err != nil {
		return errors.New("year must be a valid number")
	}

	currentYear := time.Now().Year()
	yearInt, _ := strconv.Atoi(year)
	if yearInt < 1886 || yearInt > currentYear {
		return errors.New("year must be between 1886 and current year")
	}
	return nil
}

func validatedbrand(brand string) error {
	if brand == "" {
		return errors.New("brand is required")
	}
	return nil
}

package repository

import "github.com/markledger/bookings/internal/models"

type DatabaseRepository interface {
	AllUsers() bool

	InsertReservation(reservation models.Reservation) (int, error)

	GetRooms() ([]models.Room, error)
}

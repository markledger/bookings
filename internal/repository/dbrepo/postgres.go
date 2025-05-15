package dbrepo

import (
	"github.com/markledger/bookings/internal/models"
	"time"
)

func (pg *postgresDBRepo) AllUsers() bool {
	return true
}

func (pg *postgresDBRepo) InsertReservation(reservation models.Reservation) error {
	sql := `INSERT INTO reservations 
    			(user_id, phone, start_date, end_date, room_id, created_at, updated_at)
    		VALUES 
    		    ($1, $2, $3, $4, $5, $6, $7)`

	pg.DB.Exec(sql,
		reservation.UserID,
		reservation.Phone,
		reservation.StartDate,
		reservation.EndDate,
		reservation.RoomID,
		reservation.CreatedAt,
		reservation.UpdatedAt,
		time.Now(),
		time.Now())

	return nil
}

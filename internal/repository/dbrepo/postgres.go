package dbrepo

import (
	"context"
	"github.com/markledger/bookings/internal/models"
	"time"
)

func (pg *postgresDBRepo) AllUsers() bool {
	return true
}

func (pg *postgresDBRepo) GetRooms() ([]models.Room, error) {
	var rooms []models.Room

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sql := `Select id, room_name from rooms`
	rows, err := pg.DB.QueryContext(ctx, sql)

	if err != nil {
		return rooms, err
	}

	for rows.Next() {
		var room models.Room
		if err := rows.Scan(&room.ID, &room.RoomName); err != nil {
			return rooms, err
		}
		rooms = append(rooms, room)

	}

	return rooms, nil
}

func (pg *postgresDBRepo) InsertReservation(reservation models.Reservation) (int, error) {
	var newID int
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sql := `INSERT INTO reservations 
    			(user_id, phone, start_date, end_date, room_id, created_at, updated_at)
    		VALUES 
    		    ($1, $2, $3, $4, $5, $6, $7)`

	err := pg.DB.QueryRowContext(ctx, sql,
		reservation.UserID,
		reservation.Phone,
		reservation.StartDate,
		reservation.EndDate,
		reservation.RoomID,
		time.Now(),
		time.Now()).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}

package handlers

import (
	"github.com/markledger/bookings/internal/forms"
	"github.com/markledger/bookings/internal/helpers"
	"github.com/markledger/bookings/internal/models"
	"github.com/markledger/bookings/internal/render"
	"net/http"
	"strconv"
	"time"
)

// PostReservation handles the posting of a reservation form
func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	// 01/02 03:04:05PM '06 -0700
	dateLayout := "2006-01-02"

	startDate, err := time.Parse(dateLayout, r.Form.Get("start_date"))
	if err != nil {
		helpers.ServerError(w, err)
	}

	endDate, err := time.Parse(dateLayout, r.Form.Get("end_date"))
	if err != nil {
		helpers.ServerError(w, err)
	}

	roomID, err := strconv.Atoi(r.Form.Get("room_id"))
	if err != nil {
		helpers.ServerError(w, err)
	}

	reservation := models.Reservation{
		UserID:    1,
		Phone:     r.Form.Get("phone"),
		StartDate: startDate,
		EndDate:   endDate,
		RoomID:    roomID,
	}

	form := forms.New(r.PostForm)
	form.MinLength("phone", 11)

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation
		render.Template(w, r, "make-reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	_, err = m.DB.InsertReservation(reservation)
	if err != nil {
		helpers.ServerError(w, err)
	}

	rooms, err := m.DB.GetRooms()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	var room models.Room
	for _, r := range rooms {
		if r.ID == roomID {
			room = r
		}
	}

	//restriction := models.RoomRestriction{
	//	StartDate:     startDate,
	//	EndDate:       endDate,
	//	RoomID:        roomID,
	//	ReservationID: reservationID,
	//	RestrictionID: 1,
	//}
	//err = m.DB.InsertRoomRestriction(restriction)
	//if err != nil {
	//	helpers.ServerError(w, err)
	//	return
	//}
	m.App.Session.Put(r.Context(), "reservation", reservation)

	http.Redirect(w, r, "reservation-summary", http.StatusSeeOther)
}

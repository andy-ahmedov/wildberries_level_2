package handlers

import (
	"errors"
	"fmt"
	"github.com/cothromachd/wb-internship/l2/develop/dev11/calendar"
	"github.com/cothromachd/wb-internship/l2/develop/dev11/utils"
	"net/http"
	"strconv"
)

type CalendarHandler struct {
	calendar *calendar.Calendar
}

func NewCalendarHandler() *CalendarHandler {
	return &CalendarHandler{calendar.NewCalendar()}
}

func (c *CalendarHandler) CreateEventRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		utils.SendResult(w, []byte("method not allowed"))
	}

	newEvent, err := utils.ParseCreateRequest(r)
	if err != nil {
		utils.SendError(w, err, http.StatusBadRequest)
		return
	}

	c.calendar.CreateEvent(newEvent)
	utils.SendResult(w, []byte("new event created"))
}

func (c *CalendarHandler) UpdateEventRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		utils.SendResult(w, []byte("method not allowed"))
	}

	ID, Time, Name, err := utils.ParseUpdateRequest(r)
	if err != nil {
		utils.SendError(w, err, http.StatusBadRequest)
		return
	}

	err = c.calendar.UpdateEvent(ID, Time, Name)
	if err != nil {
		utils.SendError(w, err, http.StatusBadRequest)
		return
	}

	utils.SendResult(w, []byte(fmt.Sprintf("event #%d updated", ID)))
}

func (c *CalendarHandler) DeleteEventRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		utils.SendResult(w, []byte("method not allowed"))
	}

	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "application/x-www-form-urlencoded" {
		utils.SendError(w, errors.New("invalid data"), http.StatusBadRequest)
		return
	}

	err := r.ParseForm()
	if err != nil {
		utils.SendError(w, err, http.StatusBadRequest)
		return
	}

	ID, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		utils.SendError(w, err, http.StatusBadRequest)
		return
	}

	deleted, err := c.calendar.DeleteEvent(ID)
	if err != nil {
		utils.SendError(w, err, http.StatusBadRequest)
		return
	}

	utils.SendResult(w, []byte(fmt.Sprintf("event #%d (%s, %v) removed", deleted.ID, deleted.Name, deleted.Time)))
}

func (c *CalendarHandler) EventsForDayRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		utils.SendResult(w, []byte("method not allowed"))
	}

	data, err := calendar.SerializeEventSlice(c.calendar.EventsForDay())
	if err != nil {
		utils.SendError(w, err, http.StatusServiceUnavailable)
		return
	}

	utils.SendResult(w, data)
}

func (c *CalendarHandler) EventsForWeekRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		utils.SendResult(w, []byte("method not allowed"))
	}

	data, err := calendar.SerializeEventSlice(c.calendar.EventsForWeek())
	if err != nil {
		utils.SendError(w, err, http.StatusServiceUnavailable)
		return
	}

	utils.SendResult(w, data)
}

func (c *CalendarHandler) EventsForMonthRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		utils.SendResult(w, []byte("method not allowed"))
	}

	data, err := calendar.SerializeEventSlice(c.calendar.EventsForMonth())
	if err != nil {
		utils.SendError(w, err, http.StatusServiceUnavailable)
		return
	}

	utils.SendResult(w, data)
}
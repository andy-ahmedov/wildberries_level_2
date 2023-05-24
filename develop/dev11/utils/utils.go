package utils

import (
	"encoding/json"
	"errors"
	"github.com/cothromachd/wb-internship/l2/develop/dev11/calendar"
	"net/http"
	"strconv"
	"time"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type ResultResponse struct {
	Result []byte `json:"result"`
}

func ParseCreateRequest(r *http.Request) (*calendar.Event, error) {
	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "application/x-www-form-urlencoded" {
		return nil, errors.New("invalid data")
	}
	err := r.ParseForm()
	if err != nil {
		return nil, err
	}

	t, err := time.Parse("2006-01-02 15:04", r.FormValue("time"))
	if err != nil {
		return nil, err
	}

	name := r.FormValue("name")
	if name == "" {
		return nil, errors.New("name can't be blank")
	}

	newEvent := calendar.NewEvent(t, name)
	return newEvent, nil
}

func ParseUpdateRequest(r *http.Request) (int, time.Time, string, error) {
	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "application/x-www-form-urlencoded" {
		return -1, time.Time{}, "", errors.New("invalid data")
	}
	err := r.ParseForm()
	if err != nil {
		return -1, time.Time{}, "", err
	}

	ID, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		return -1, time.Time{}, "", err
	}

	Time := r.FormValue("time")
	parsedTime, err := time.Parse("2006-01-02 15:04", r.FormValue("time"))
	if !(Time == "" || err == nil) {
		return -1, time.Time{}, "", err
	}

	Name := r.FormValue("name")

	return ID, parsedTime, Name, nil
}

func SendError(w http.ResponseWriter, err error, statusCode int) {
	data := ErrorResponse{err.Error()}
	result, _ := json.Marshal(data)
	w.WriteHeader(statusCode)
	w.Write(result)
}

func SendResult(w http.ResponseWriter, response []byte) {
	data := ResultResponse{response}
	result, _ := json.Marshal(data)
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
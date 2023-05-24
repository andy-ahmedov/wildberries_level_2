package calendar

import (
	"encoding/json"
	"errors"
	"sync"
	"time"
)

var LastID = 0
var LastIdMutex = sync.Mutex{}

type Calendar struct {
	events map[int]Event
	mu     sync.RWMutex
}

type Event struct {
	ID   int       `json:"id"`
	Name string    `json:"name"`
	Time time.Time `json:"time"`
}

type Result struct {
	Result []Event `json:"result"`
}

func SerializeEventSlice(events []Event) ([]byte, error) {
	data := Result{events}
	result, err := json.Marshal(data)
	return result, err
}

func NewCalendar() *Calendar {
	return &Calendar{make(map[int]Event), sync.RWMutex{}}
}

func NewEvent(Time time.Time, Name string) *Event { // Create event struct
	LastIdMutex.Lock()
	LastID++
	LastIdMutex.Unlock()
	return &Event{LastID, Name, Time}
}

func (c *Calendar) CreateEvent(event *Event) { // Add event to calendar
	c.mu.Lock()
	c.events[event.ID] = *event
	c.mu.Unlock()
}

func (c *Calendar) UpdateEvent(ID int, Time time.Time, Name string) error { // Add new event instead of old one
	c.mu.RLock()
	event, ok := c.events[ID]
	if !ok {
		c.mu.RUnlock()
		return errors.New("no such event")
	}
	c.mu.RUnlock()

	if !Time.IsZero() {
		event.Time = Time
	}
	if Name != "" {
		event.Name = Name
	}
	c.mu.Lock()
	c.events[ID] = event
	c.mu.Unlock()
	return nil
}

func (c *Calendar) DeleteEvent(ID int) (*Event, error) {
	c.mu.RLock()
	if _, ok := c.events[ID]; !ok {
		c.mu.RUnlock()
		return nil, errors.New("no such event")
	}
	c.mu.RUnlock()
	c.mu.Lock()
	deleted := c.events[ID]
	delete(c.events, ID)
	c.mu.Unlock()
	return &deleted, nil
}

func (c *Calendar) EventsForDay() []Event {
	var result []Event
	tYear, tMonth, tDay := time.Now().Date() // today
	c.mu.RLock()
	for _, v := range c.events {
		year, month, day := v.Time.Date()
		if tYear == year && tMonth == month && tDay == day {
			result = append(result, v)
		}
	}
	c.mu.RUnlock()
	return result
}

func (c *Calendar) EventsForWeek() []Event {
	var result []Event
	tYear, tWeek := time.Now().ISOWeek()
	c.mu.RLock()
	for _, v := range c.events {
		year, week := v.Time.ISOWeek()
		if tYear == year && tWeek == week {
			result = append(result, v)
		}
	}
	c.mu.RUnlock()
	return result
}

func (c *Calendar) EventsForMonth() []Event {
	var result []Event
	tYear, tMonth, _ := time.Now().Date() // today
	c.mu.RLock()
	for _, v := range c.events {
		year, month, _ := v.Time.Date()
		if tYear == year && tMonth == month {
			result = append(result, v)
		}
	}
	c.mu.RUnlock()
	return result
}
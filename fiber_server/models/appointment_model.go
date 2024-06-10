package model

import "time"

type Appointment struct {
    ID        int       `json:"id"`
    Title     string    `json:"title"`
    StartTime time.Time `json:"start_time"`
    EndTime   time.Time `json:"end_time"`
}
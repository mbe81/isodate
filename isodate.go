package isodate

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type IsoDate struct {
	Time time.Time
}

func New(time time.Time) IsoDate {
	i := IsoDate{
		Time: time,
	}
	return i
}

// UnmarshalJSON to read a JSON
func (d *IsoDate) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	t, _ := time.Parse("2006-01-02", s)
	d.Time = t
	return nil
}

// MarshalJSON to create a JSON
func (d IsoDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Time.Format("2006-01-02"))
}

// Scan to read from SQL into variable, implements sql.Scan interface
func (d *IsoDate) Scan(value interface{}) error {
	d.Time = value.(time.Time)
	return nil
}

// Value to write IsoDate value to SQL, implements sql.driver.Value interface
func (d IsoDate) Value() (driver.Value, error) {
	return d.Time.Format("2006-01-02"), nil
}
package coinspaid

import (
	"strconv"
	"time"
)

// Time is time.Time wrapper implementing custom JSON marshaller/unmarshaler.
type Time time.Time

// Time returns time.Time value.
func (t Time) Time() time.Time {
	return time.Time(t)
}

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t).Unix(), 10)), nil
}

func (t *Time) UnmarshalJSON(data []byte) error {
	sec, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}

	*(*time.Time)(t) = time.Unix(sec, 0)

	return nil
}

func (t Time) Equal(u Time) bool {
	return time.Time(t).Equal(time.Time(u))
}

// Duration is time.Duration wrapper implementing custom JSON marshaller/unmarshaler.
type Duration time.Duration

// Duration returns time.Duration value.
func (d Duration) Duration() time.Duration {
	return time.Duration(d)
}

func (d Duration) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(int64(time.Duration(d)/time.Second), 10)), nil
}

func (d *Duration) UnmarshalJSON(data []byte) error {
	sec, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}

	*(*time.Duration)(d) = time.Second * time.Duration(sec)

	return nil
}

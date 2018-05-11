package hxconnect

import (
  "time"
  "strconv"
)

func DiffString(bootTime int64) string {
  a, b, c, d, e, f := diff(time.Now(),time.Unix(bootTime, 0))
  tmp := ""

  if a > 0 && a < 2 {tmp += strconv.Itoa(a) + " year "} else if a > 1 {tmp += strconv.Itoa(a) + " years "}
  if b > 0 && b < 2 {tmp += strconv.Itoa(b) + " month "} else if b > 1 {tmp += strconv.Itoa(b) + " months "}
  if c > 0 && c < 2 {tmp += strconv.Itoa(c) + " day "} else if c > 1 {tmp += strconv.Itoa(c) + " days "}
  if d > 0 && d < 2 {tmp += strconv.Itoa(d) + " hour "} else if d > 1 {tmp += strconv.Itoa(d) + " hours "}
  if e > 0 && e < 2 {tmp += strconv.Itoa(e) + " minute "} else if e > 1 {tmp += strconv.Itoa(e) + " minutes "}
  if f > 0 && f < 2 {tmp += strconv.Itoa(f) + " second"} else if f > 1 {tmp += strconv.Itoa(f) + " seconds"}

  return tmp
}

func diff(a, b time.Time) (year, month, day, hour, min, sec int) {
    if a.Location() != b.Location() {
        b = b.In(a.Location())
    }
    if a.After(b) {
        a, b = b, a
    }
    y1, M1, d1 := a.Date()
    y2, M2, d2 := b.Date()

    h1, m1, s1 := a.Clock()
    h2, m2, s2 := b.Clock()

    year = int(y2 - y1)
    month = int(M2 - M1)
    day = int(d2 - d1)
    hour = int(h2 - h1)
    min = int(m2 - m1)
    sec = int(s2 - s1)

    // Normalize negative values
    if sec < 0 {
        sec += 60
        min--
    }
    if min < 0 {
        min += 60
        hour--
    }
    if hour < 0 {
        hour += 24
        day--
    }
    if day < 0 {
        // days in month:
        t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
        day += 32 - t.Day()
        month--
    }
    if month < 0 {
        month += 12
        year--
    }

    return
}

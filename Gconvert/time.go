package Gconvert

import "time"

func Time2Str(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
func Unix2Time(t int64) time.Time {
	return time.Unix(t, 0)
}
func Str2Time(t string) time.Time {
	ret, _ := time.Parse("2006-01-02 15:04:05", t)
	return ret
}

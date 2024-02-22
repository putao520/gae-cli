package types

import (
	"time"
)

type DateTime time.Time

var TimeFormats = []string{"2006-01-02 15:04:05", "20060102150405"}

func (t *DateTime) UnmarshalJSON(data []byte) (err error) {
	// fmt.Println(string(data))
	// 空值不进行解析
	if len(data) == 2 {
		*t = DateTime(time.Time{})
		return
	}

	var now time.Time
	for _, format := range TimeFormats {
		// 指定解析的格式
		if now, err = time.ParseInLocation(format, string(data), time.Local); err == nil {
			*t = DateTime(now)
			return
		}
		// 指定解析的格式
		if now, err = time.ParseInLocation(`"`+format+`"`, string(data), time.Local); err == nil {
			*t = DateTime(now)
			return
		}
	}

	return
}
func (t DateTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TimeFormats[0])+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, TimeFormats[0])
	b = append(b, '"')
	return b, nil
}

func (t DateTime) String() string {
	return time.Time(t).Format(TimeFormats[0])
}

func (t DateTime) Time() time.Time {
	return time.Time(t)
}

func (t DateTime) IsZero() bool {
	return time.Time(t).IsZero()
}

// 秒级时间戳转time
func UnixSecondToTime(second int64) time.Time {
	return time.Unix(second, 0)
}

// 毫秒级时间戳转time
func UnixMilliToTime(milli int64) time.Time {
	return time.Unix(milli/1000, (milli%1000)*(1000*1000))
}

// 纳秒级时间戳转time
func UnixNanoToTime(nano int64) time.Time {
	return time.Unix(nano/(1000*1000*1000), nano%(1000*1000*1000))
}

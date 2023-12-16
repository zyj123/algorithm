package timeUtil

import (
	"fmt"
	"strings"
	"time"
)

// HumanReadTimeSpan 可读性较好的时间长度，example:3天2小时5分钟45秒
func HumanReadTimeSpan(t time.Duration) string {
	var (
		day  int
		hour int
		min  int
		sec  int
		ms   int
	)
	for t > time.Hour*24 {
		day, t = int(t/time.Hour*24), t%time.Hour*24
	}
	for t >= time.Hour {
		hour, t = int(t/time.Hour), t%time.Hour
	}
	for t >= time.Minute {
		min, t = int(t/time.Minute), t%time.Minute
	}
	for t >= time.Second {
		sec, t = int(t/time.Second), t%time.Second
	}
	for t >= time.Millisecond {
		ms, t = int(t/time.Millisecond), t%time.Millisecond
	}
	sb := strings.Builder{}
	switch true {
	case day > 0:
		sb.WriteString(fmt.Sprintf("%d天", day))
		fallthrough
	case day > 0 || hour > 0:
		sb.WriteString(fmt.Sprintf("%d小时", hour))
		fallthrough
	case hour > 0 || min > 0:
		sb.WriteString(fmt.Sprintf("%d分钟", min))
		fallthrough
	case min > 0 || sec > 0:
		sb.WriteString(fmt.Sprintf("%d秒", sec))
		fallthrough
	case sec > 0 || ms > 0:
		sb.WriteString(fmt.Sprintf("%d毫秒", ms))
	default:
	}

	if sb.Len() > 0 {
		return sb.String()
	}
	return "0秒"
}

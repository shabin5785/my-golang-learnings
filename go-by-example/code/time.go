package gobyex

import (
	"fmt"
	"time"
)

func TimeFn() {
	p := fmt.Println
	now := time.Now()
	p(now)

	then := time.Date(2001, 6, 22, 14, 22, 12, 342322341, time.UTC)
	p(then)

	p(now.Month())
	p(now.Location())

	p(then.Before(now))

	diff := now.Sub(then)
	p(diff)
	p(diff.Minutes())

	//epoch
	epoch := now.Unix()
	p(epoch)

	parseTime()
}

func parseTime() {
	p := fmt.Println

	t := time.Now()
	p(t.Format(time.RFC3339))

	t1, e := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")

	p(t1)
	p(t.Format("3:04PM"))
	//custom parse format needs to be below
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))

	p(e)
}

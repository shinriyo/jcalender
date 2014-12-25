package test

import (
	jc "../../jcalender"
	"time"
)

func Main() {
	params := jc.Params{Year: 2000, Month:12, Day:23, Date:time.Time{}}
	jc.GetHolidayName(params)
}




package jcalender

import (
	"math"
	"time"
)

const (
	MONDAY = 0
	TUESDAY
	WEDNESDAY
)

func Main() {
}

// 整数で年を与えると、その年の春分の日が3月の何日であるかを返す
func VernalEquinox(y int) int {
	d := 0

	if y <= 1947 {
		d = 0
	} else if y <= 1979 {
		d = int(math.Floor(20.8357+0.242194*float64(y - 1980)-math.Floor(float64(y - 1980)/4.0)))
	} else if y <= 2099 {
		d = int(math.Floor(20.8431+0.242194*float64(y - 1980)-math.Floor(float64(y - 1980)/4.0)))
	} else if y <= 2150 {
		d = int(math.Floor(21.8510+0.242194*float64(y - 1980)-math.Floor(float64(y - 1980)/4.0)))
	}

	return d
}

// 整数で年を与えると、その年の秋分の日が9月の何日であるかを返す
func AutumnEquinox(y int) int {
	d := 0

	if y <= 1947 {
		d = 0
	} else if y <= 1979 {
		d = int(math.Floor(23.2588+0.242194*float64(y - 1980)-math.Floor(float64(y - 1980)/4.0)))
	} else if y <= 2099 {
		d = int(math.Floor(23.2488+0.242194*float64(y - 1980)-math.Floor(float64(y - 1980)/4.0)))
	} else if y <= 2150 {
		d = int(math.Floor(24.2488+0.242194*float64(y - 1980)-math.Floor(float64(y - 1980)/4.0)))
	} else {
		d = 0
	}

	return d
}

/*
    GetHolidayName() の呼び出し方法は2通りあります。

    1つ目の方法は、3つの引数 year, month, day に整数を渡す方法です。
    もうひとつの方法は前述のキーワード引数 date に datetime.date のオブジェクトを渡す方法です。
    この場合は year, month, day を渡す必要はなく、また、渡したとしても無視されます。

    GetHolidayName() は、その日が祝日であれば 
        Python 2.x 系以前の場合には Unicode 文字列で
        Python 3.x 系以降の場合には文字列で
    祝日名を返します。
    指定した日が祝日でなければ、 Python のバージョンによらず nil を返します。
*/
// year = nil, month = nil, day = nil, p.Date = nil
type Params struct {
	Year  int
	Month time.Month
	Day   int
	Date  time.Time
}

func GetHolidayName(p Params) string {
	name := ""
	//if p.Date == nil {
	now := time.Now()
	if p.Date.IsZero() {
		// func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
		p.Date = time.Date(p.Year, p.Month, p.Day, 0, 0, 0, 0, now.Location())
	}
	//if p.Date < time.Date(1948, 7, 20, 0, 0, 0, 0, now.Location()) {
	if p.Date.Sub(time.Date(1948, 7, 20, 0, 0, 0, 0, now.Location())) < 0 {
		return ""
	} else {
		name = ""
	}

	// 1月
	if p.Date.Month() == time.January {
		if p.Date.Day() == 1 {
			name = "元日"
		} else {
			if p.Date.Year() >= 2000 {
				if int((p.Date.Day() - 1) / 7) == 1 && p.Date.Weekday() == MONDAY {
					name = "成人の日"
				}
			} else {
				if p.Date.Day() == 15 {
					name = "成人の日"
				}
			}
		}
	} else if p.Date.Month() == time.February {
		// 2月
		if p.Date.Day() == 11 && p.Date.Year() >= 1967 {
			name = "建国記念の日"
		} else if p.Date.Year() == 1989 && p.Date.Month() == 2 && p.Date.Day() == 24 {
			name = "昭和天皇の大喪の礼"
		}
	}  else if p.Date.Month() == time.March {
		// 3月
		if p.Date.Day() == VernalEquinox(p.Date.Year()) {
			name = "春分の日"
		}
	} else if p.Date.Month() == time.April {
		// 4月
		if p.Date.Day() == 29 {
			if p.Date.Year() >= 2007 {
				name = "昭和の日"
			} else if p.Date.Year() >= 1989 {
				name = "みどりの日"
			} else {
				name = "天皇誕生日"
			}
		} else if p.Date.Year() == 1959 && p.Date.Month() == 4 && p.Date.Day() == 10 {
			name = "皇太子明仁親王の結婚の儀"
		}
	} else if p.Date.Month() == time.May {
		// 5月
		if p.Date.Day() == 3 {
			name = "憲法記念日"
		} else if p.Date.Day() == 4 {
			if p.Date.Year() >= 2007 {
				name = "みどりの日"
			} else if p.Date.Year() >= 1986 && p.Date.Weekday() != MONDAY {
				name = "国民の休日"
			}
		} else if p.Date.Day() == 5 {
			name = "こどもの日"

		} else if p.Date.Day() == 6 {
			if p.Date.Year() >= 2007 && (p.Date.Weekday() == TUESDAY || p.Date.Weekday() == WEDNESDAY) {
				name = "振替休日"
			}
		}
	} else if p.Date.Month() == time.June {
		// 6月
		if p.Date.Year() == 1993 && p.Date.Month() == 6 && p.Date.Day() == 9 {
			name = "皇太子徳仁親王の結婚の儀"
		}
	} else if p.Date.Month() == time.July {
		// 7月
		if p.Date.Year() >= 2003 {
			if int((p.Date.Day() - 1) / 7) == 2 && p.Date.Weekday() == MONDAY {
				name = "海の日"
			}
		} else if p.Date.Year() >= 1996 && p.Date.Day() == 20 {
			name = "海の日"
		}
	} else if p.Date.Month() == time.August {
		// 8月
		if p.Date.Day() == 11 && p.Date.Year() >= 2016 {
			name = "山の日"
		}
	} else if p.Date.Month() == 9 {
		// 9月
		autumn_equinox := AutumnEquinox(p.Date.Year())
		if p.Date.Day() == autumn_equinox {
			name = "秋分の日"
		} else if p.Date.Year() >= 2003 {
			if (int((p.Date.Day() - 1) / 7.0) == 2) && p.Date.Weekday() == MONDAY {
				name = "敬老の日"
			} else if p.Date.Weekday() == TUESDAY && p.Date.Day() == autumn_equinox-1 {
				name = "国民の休日"
			} else if p.Date.Year() >= 1966 && p.Date.Day() == 15 {
				name = "敬老の日"
			}
		}
	} else if p.Date.Month() == 10 {
		// 10月
		if p.Date.Year() >= 2000 {
			if (int((p.Date.Day() - 1) / 7.0) == 1) && p.Date.Weekday() == MONDAY {
				name = "体育の日"
			}
		} else if p.Date.Year() >= 1966 && p.Date.Day() == 10 {
			name = "体育の日"
		}
	} else if p.Date.Month() == 11 {
		// 11月
		if p.Date.Day() == 3 {
			name = "文化の日"
		} else if p.Date.Day() == 23 {
			name = "勤労感謝の日"
		} else if p.Date.Year() == 1990 && p.Date.Month() == 11 && p.Date.Day() == 12 {
			name = "即位礼正殿の儀"
		}
	}  else if p.Date.Month() == 12 {
		// 12月
		if p.Date.Day() == 23 && p.Date.Year() >= 1989 {
			name = "天皇誕生日"
		}
	}

	// 振替休日
	if name != "" && p.Date.Weekday() == MONDAY {
		// 1日前
		prev := p.Date.AddDate(0, 0, -1)
		params := Params{Year:prev.Year(), Month:prev.Month(), Day:prev.Day(), Date:time.Time{}}
		if GetHolidayName(params) != "" {
			name = "振替休日"
		}
	}

	return name
}

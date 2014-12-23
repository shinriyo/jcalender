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

// Pythonのinがないのでその代わり
func IntInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// 整数で年を与えると、その年の春分の日が3月の何日であるかを返す
func VernalEquinox(y int) float64 {
	var d = 0.0

	if y <= 1947 {
		d = 0
	} else if y <= 1979 {
		d = math.Floor(20.8357+0.242194*float64(y - 1980)-math.Floor(float64(y - 1980)/4.0))
	} else if y <= 2099 {
		d = math.Floor(20.8431+0.242194*float64(y - 1980)-math.Floor(float64(y - 1980)/4.0))
	} else if y <= 2150 {
		d = math.Floor(21.8510+0.242194*float64(y - 1980)-math.Floor(float64(y - 1980)/4.0))
	}

	return d
}

// 整数で年を与えると、その年の秋分の日が9月の何日であるかを返す
func AutumnEquinox(y int) float64 {
	var d = 0.0

	if y <= 1947 {
		d = 0
	} else if y <= 1979 {
		d = math.Floor(23.2588+0.242194*float64(y - 1980)-math.Floor(float64(y - 1980)/4.0))
	} else if y <= 2099 {
		d = math.Floor(23.2488+0.242194*float64(y - 1980)-math.Floor(float64(y - 1980)/4.0))
	} else if y <= 2150 {
		d = math.Floor(24.2488+0.242194*float64(y - 1980)-math.Floor(float64(y - 1980)/4.0))
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
    指定した日が祝日でなければ、 Python のバージョンによらず None を返します。
*/
// year = None, month = None, day = None, p.date = None
type Params struct {
	year  string
	month int
	day   int
	date  Time
}

func GetHolidayName(p Params) string {
	var name = ""
	if p.date == None {
		p.date = Date(p.year, p.month, p.day)
	}
	if p.date < Date(1948, 7, 20) {
		return None
	} else {
		name = None
	}

	// 1月
	if p.date.month == 1 {
		if p.date.day == 1 {
			name = "元日"
		} else {
			if p.datedate.year >= 2000 {
				if int((p.date.day - 1) / 7) == 1 && p.date.Weekday() == MONDAY {
					name = "成人の日"
				}
			} else {
				if p.date.day == 15 {
					name = "成人の日"
				}
			}
		}
	} else if p.date.month == 2 {
		// 2月
		if p.date.day == 11 && p.date.year >= 1967 {
			name = "建国記念の日"
		} else if p.date.year == 1989 && p.date.month == 2 && p.date.day == 24 {
			name = "昭和天皇の大喪の礼"
		}
	}  else if p.date.month == 3 {
		// 3月
		if p.date.day == VernalEquinox(p.date.year) {
			name = "春分の日"
		}
	} else if p.date.month == 4 {
		// 4月
		if p.date.day == 29 {
			if p.date.year >= 2007 {
				name = "昭和の日"
			} else if p.date.year >= 1989 {
				name = "みどりの日"
			} else {
				name = "天皇誕生日"
			}
		} else if p.date.year == 1959 && p.date.month == 4 && p.date.day == 10 {
			name = "皇太子明仁親王の結婚の儀"
		}
	} else if p.date.month == 5 {
		// 5月
		if p.date.day == 3 {
			name = "憲法記念日"
		} else if p.date.day == 4 {
			if p.date.year >= 2007 {
				name = "みどりの日"
			} else if p.date.year >= 1986 && p.date.weekday() != MONDAY {
				name = "国民の休日"
			}
		} else if p.date.day == 5 {
			name = "こどもの日"

		} else if p.date.day == 6 {
			array := [...]int{TUESDAY, WEDNESDAY}
			if p.date.year >= 2007 && IntInSlice(int(p.date.Weekday()), array) {
				name = "振替休日"
			}
			// 6月
		} else if p.date.month == 6 {
			if p.date.year == 1993 && p.date.month == 6 && p.date.day == 9 {
				name = "皇太子徳仁親王の結婚の儀"
			}
			// 7月
		} else if p.date.month == 7 {
			if p.date.year >= 2003 {
				if int((p.date.day - 1) / 7) == 2 && p.date.weekday() == MONDAY {
					name = "海の日"
				}
			} else if p.date.year >= 1996 && p.date.day == 20 {
				name = "海の日"
			}
			// 8月
		} else if p.date.month == 8 {
			if p.date.day == 11 && p.date.year >= 2016 {
				name = "山の日"
				// 9月
			}
		} else if p.date.month == 9 {
			var autumn_equinox = AutumnEquinox(p.date.year)
			if p.date.day == autumn_equinox {
				name = "秋分の日"
			} else if p.date.year >= 2003 {
				if (int((p.date.day - 1) / 7.0) == 2) && p.date.weekday() == MONDAY {
					name = "敬老の日"
				} else if p.date.weekday() == TUESDAY && p.date.day == autumn_equinox-1 {
					name = "国民の休日"
				} else if p.date.year >= 1966 && p.date.day == 15 {
					name = "敬老の日"
				}
			}
		} else if p.date.month == 10 {
			// 10月
			if p.date.year >= 2000 {
				if (int((p.date.day - 1) / 7.0) == 1) && p.date.weekday() == MONDAY {
					name = "体育の日"
				}
			} else if p.date.year >= 1966 && p.date.day == 10 {
				name = "体育の日"
			}
		}
	} else if p.date.month == 11 {
		// 11月
		if p.date.day == 3 {
			name = "文化の日"
		} else if p.date.day == 23 {
			name = "勤労感謝の日"
		} else if p.date.year == 1990 && p.date.month == 11 && p.date.day == 12 {
			name = "即位礼正殿の儀"
		}
	}  else if p.date.month == 12 {
		// 12月
		if p.date.day == 23 && p.date.year >= 1989 {
			name = "天皇誕生日"
		}
	}

	// 振替休日
	if !name && p.date.weekday() == MONDAY {
		// 1日前
		var prev = p.date.AddDate(0, 0, -1)
		if GetHolidayName(Params(prev.year, prev.month, prev.day, None)) != "" {
			name = "振替休日"
		}
	}

	return name
}

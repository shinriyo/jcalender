// 必ず実行するファイルはmainパッケージ
package main

import (
	jc "../../jcalender"
	"time"
	"fmt"
)

// 必ず小文字のmain()メソッド
func main() {
	params := jc.Params{Year: 2000, Month:12, Day:23, Date:time.Time{}}
	name := jc.GetHolidayName(params)
	// 天皇誕生日
	fmt.Printf("Name: %#v\n", name)

	params = jc.Params{Year: 1984, Month:1, Day:1, Date:time.Time{}}
	name = jc.GetHolidayName(params)
	// 元旦
	fmt.Printf("Name: %#v\n", name)

	params = jc.Params{Year: 2014, Month:10, Day:10, Date:time.Time{}}
	name = jc.GetHolidayName(params)
	// 体育の日
	fmt.Printf("Name: %#v\n", name)
}




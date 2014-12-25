jCalender
====

Overview

## Description

日本のカレンダー

http://www.h3.dion.ne.jp/~sakatsu/holiday_logic5.htm#Python
これをベースにした。

## Demo

N/A

## VS. 

N/A

## Requirement

go version go1.3.3 or higher

## Usage

### テスト方法
```
go run test/test.go
```

### インポート方法
```
import (
  jc "github.com/shinriyo/jcalender"
)
```

jcalenderは長いので省略したjcにエイリアス。

### 使い方
```
params = jc.Params{Year: 1984, Month:1, Day:1, Date:time.Time{}}
name = jc.GetHolidayName(params)
// 元旦
fmt.Printf("Name: %#v\n", name)
```

### 注意

2016年以降は対応できないかもしれません。

## Install

```
go get github.com/shinriyo/jcalender
```

## Contribution

sinriyo

## Licence

CopyRight(C) K.Tsunoda(AddinBox) 2001 All Rights Reserved.
(http://www.h3.dion.ne.jp/~sakatsu/index.htm)

## Author

[shinriyo](https://github.com/shinriyo/)

[K.Tsunoda](http://www.h3.dion.ne.jp/~sakatsu/index.htm)


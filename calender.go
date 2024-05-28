package main

import (
	"fmt"
	"time"
)

// カレンダーを表示する関数
func PrintCalendar(year int, month time.Month) {
	// 月初の日時を取得
	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	// 月の最後の日を取得
	lastDay := firstDay.AddDate(0, 1, -1)

	fmt.Printf("\n     %d年 %s\n", year, month)
	fmt.Println("日 月 火 水 木 金 土")

	// 最初の曜日までスペースを埋める
	for i := time.Sunday; i < firstDay.Weekday(); i++ {
		fmt.Print("   ")
	}

	// 各日付を出力
	for day := 1; day <= lastDay.Day(); day++ {
		fmt.Printf("%2d ", day)
		if firstDay.Weekday() == time.Saturday {
			fmt.Println()
		}
		firstDay = firstDay.AddDate(0, 0, 1)
	}
	fmt.Println()
}

func WaitForInput() int {
	for {
		var Input string
		fmt.Print("\n1:Previous")
		fmt.Print("\n2:Next")
		fmt.Print("\n0:終了\n")
		fmt.Scan(&Input)
		switch Input {
		case "1":
			return -1
		case "2":
			return 1
		case "0":
			return 0
		default:
			fmt.Printf("無効な入力です。\n")
	
		}
	}
}

func main() {
	Currentdate := time.Now()
	year := Currentdate.Year()
	month := Currentdate.Month()
	PrintCalendar(year, month)
	for{
		//入力を待つ
		Input:=WaitForInput()
		//終了
		if Input==0 {
			return
		}
		Currentdate=Currentdate.AddDate(0,Input,0)
		year = Currentdate.Year()
		month = Currentdate.Month()
		//カレンダー表示
		PrintCalendar(year, month)
	}
}

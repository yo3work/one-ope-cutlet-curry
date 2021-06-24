package main

import (
	"fmt"
	"os"
	"time"
)

func CookingOneOpe(menu string, cookingTime int) {
	fmt.Printf("%s の調理を開始します。\n", menu)
	start := time.Now()
	time.Sleep(time.Duration(cookingTime) * time.Second)
	end := time.Now()
	fmt.Printf("%s の調理が完了しました。 %f 秒かかりました。\n", menu, (end.Sub(start)).Seconds())
}

func Cooking(menu string, cookingTime int, quit chan bool) {
	fmt.Printf("%s の調理を開始します。\n", menu)
	start := time.Now()
	time.Sleep(time.Duration(cookingTime) * time.Second)
	end := time.Now()
	fmt.Printf("%s の調理が完了しました。 %f 秒かかりました。\n", menu, (end.Sub(start)).Seconds())
	quit <- true
}

func main() {
	StartCooking := time.Now()

	if os.Args[1] == "1" {
		// Case 1 ワンオペ関数を3回呼び出す。期待される結果 -> 調理時間の合計時間かかる
		CookingOneOpe("ごはん", 2)
		CookingOneOpe("カレー", 3)
		CookingOneOpe("トンカツ", 5)
	} else if os.Args[1] == "2" {
		// Case 2 goroutineで呼び出す「だけ」。期待される結果 -> ほぼ0秒で終わる(goroutineの終了を待たずにmain関数が終わっちゃう)
		quit := make(chan bool)
		go Cooking("ごはん", 2, quit)
		go Cooking("カレー", 3, quit)
		go Cooking("トンカツ", 5, quit)
	} else if os.Args[1] == "3" {
		// Case 3 goroutineの終了待ちをする。期待される結果 -> 一番時間のかかる処理の処理時間で終わる
		quitRice := make(chan bool)
		quitCurry := make(chan bool)
		quitCutlet := make(chan bool)
		go Cooking("ごはん", 2, quitRice)
		go Cooking("カレー", 3, quitCurry)
		go Cooking("トンカツ", 5, quitCutlet)
		<-quitRice
		<-quitCurry
		<-quitCutlet
	} else if os.Args[1] == "4" {
		// Case 4 終了待ちの場所を変えてみる。期待される結果 -> この場合はワンオペと同じになる。
		quitRice := make(chan bool)
		quitCurry := make(chan bool)
		quitCutlet := make(chan bool)
		go Cooking("ごはん", 2, quitRice)
		<-quitRice
		go Cooking("カレー", 3, quitCurry)
		<-quitCurry
		go Cooking("トンカツ", 5, quitCutlet)
		<-quitCutlet
	} else {
		fmt.Println("go run main.go の後に数字をいれてね")
		os.Exit(0)
	}

	EndCooking := time.Now()
	fmt.Printf("合計時間は %f 秒でした。\n", (EndCooking.Sub(StartCooking)).Seconds())

}

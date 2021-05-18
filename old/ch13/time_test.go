package ch13

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeDemo(t *testing.T) {

	now := time.Now()
	fmt.Printf("now time %v\n", now)
	year := now.Year()
	fmt.Printf("year %v\n", year)
	month := now.Month()
	fmt.Printf("month %v\n", month)

}

func TestTimestamp(t *testing.T) {
	now := time.Now()
	timestamp := now.Unix()
	timestamp2 := now.UnixNano()
	fmt.Printf("timestamp %v\n", timestamp)
	fmt.Printf("timestamp %v\n", timestamp2)
}

func TestTimestampconvTime(t *testing.T) {

	var timestamp int64
	timestamp = 1620899619
	timeObj := time.Unix(timestamp, 0)
	fmt.Println(timeObj)

	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()

	fmt.Printf("%d-%d-%d\n", month, day, hour)
	mill := time.Millisecond * 1
	fmt.Printf("mill: %d", mill)

}

// func (t Time) Add(d Duration) Time

func TestTimeInterval(t *testing.T) {

	now := time.Now()
	later := now.Add(time.Hour)
	fmt.Println(later)

	fmt.Println(later.Sub(now))

	// time.Sleep(5 * time.Second)
	fmt.Println("end ")

	// ticker := time.Tick(time.Second)
	// for i := range ticker {
	// 	fmt.Println(i)
	// }

	nowe := time.Now()
	fmt.Println(nowe.Format("2007-01-02"))

	fmt.Println(nowe.Format("2006/01/02/ 15:04"))

	fmt.Println(now)
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return

	}
	timeStr := "2020-01-02 20:04:05"
	fmt.Println(loc)
	timeObj, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)

	timeObj1, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj1)

}

package buildin

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestDuration(t *testing.T) {
	Convey("test duration", t, func() {
		Convey("duration string convert", func() {
			twoHours, _ := time.ParseDuration("2h")
			So(twoHours, ShouldEqual, 2*time.Hour)

			So((2 * time.Hour).String(), ShouldEqual, "2h0m0s")
		})

		Convey("duration value", func() {
			twoHours := 2 * time.Hour
			So(twoHours.Minutes(), ShouldEqual, 2*60)
			So(twoHours.Seconds(), ShouldEqual, 2*60*60)
			So(twoHours.Milliseconds(), ShouldEqual, 2*60*60*1000)
			So(twoHours.Microseconds(), ShouldEqual, 2*60*60*1000*1000)
			So(twoHours.Nanoseconds(), ShouldEqual, 2*60*60*1000*1000*1000)
			So(twoHours.Seconds(), ShouldEqual, 2*60*60)
			// 时间截断
			So((2*time.Hour + 30*time.Minute).Round(time.Hour), ShouldEqual, 3*time.Hour)      // 四舍五入
			So((2*time.Hour + 30*time.Minute).Truncate(time.Hour), ShouldEqual, 2*time.Hour)   // 向下取整
			So((-2*time.Hour - 30*time.Minute).Round(time.Hour), ShouldEqual, -3*time.Hour)    // 负数返回正数相反数
			So((-2*time.Hour - 30*time.Minute).Truncate(time.Hour), ShouldEqual, -2*time.Hour) // 负数返回正数相反数
		})
	})
}

func TestTime(t *testing.T) {
	Convey("test time", t, func() {
		Convey("time create", func() {
			t0 := time.Now()
			t1 := time.Date(2020, 1, 2, 3, 4, 5, 6, time.UTC)
			t2, _ := time.Parse(time.RFC3339, "2020-01-02T03:04:05.000000006Z")
			t3 := time.Unix(1577934245, 6)
			So(t0.After(t1), ShouldBeTrue)
			So(t1, ShouldEqual, t2)
			So(t1, ShouldEqual, t3)
			So(t2.String(), ShouldEqual, "2020-01-02 03:04:05.000000006 +0000 UTC")
			So(t2.Format(time.RFC3339), ShouldEqual, "2020-01-02T03:04:05Z")
			So(t3.Unix(), ShouldEqual, 1577934245)
			So(t3.UnixNano(), ShouldEqual, 1577934245000000006)

			// 注意 2006-01-02T15:04:05Z07:00 不是合法的 RFC3339，下面两个例子是合法的 RFC3339
			// 2006-01-02T15:04:05Z
			// 2006-01-02T15:04:05+8:00
			_, err := time.Parse(time.RFC3339, time.RFC3339)
			So(err, ShouldNotBeNil)
		})

		Convey("time status", func() {
			t := time.Date(2020, 1, 2, 3, 4, 5, 6, time.UTC)
			So(t.Year(), ShouldEqual, 2020)
			So(t.Month(), ShouldEqual, time.January)
			So(t.Day(), ShouldEqual, 2)
			So(t.Hour(), ShouldEqual, 3)
			So(t.Minute(), ShouldEqual, 4)
			So(t.Second(), ShouldEqual, 5)
			So(t.Nanosecond(), ShouldEqual, 6)
			So(t.Weekday(), ShouldEqual, time.Thursday) // 星期几
			So(t.YearDay(), ShouldEqual, 2)             // 一年中的第几天， 取值 [1,366]
			year, month, day := t.Date()
			So(year, ShouldEqual, t.Year())
			So(month, ShouldEqual, t.Month())
			So(day, ShouldEqual, t.Day())
			hour, minute, second := t.Clock()
			So(hour, ShouldEqual, t.Hour())
			So(minute, ShouldEqual, t.Minute())
			So(second, ShouldEqual, t.Second())
			year, week := t.ISOWeek() // 返回某年的第几周(从 1 开始计数)
			So(year, ShouldEqual, t.Year())
			So(week, ShouldEqual, 1)
			So(t.IsZero(), ShouldBeFalse)
		})

		Convey("time calculate", func() {
			t1, _ := time.Parse(time.RFC3339, "2020-01-02T03:04:05.000000006Z")
			t2, _ := time.Parse(time.RFC3339, "2020-01-02T05:04:05.000000006Z")
			So(t2.Sub(t1), ShouldEqual, 2*time.Hour)
			So(t1.Add(2*time.Hour), ShouldEqual, t2)
			So(t2.After(t1), ShouldBeTrue)
			So(t1.Before(t2), ShouldBeTrue)
			So(t1.Equal(t1), ShouldBeTrue)
			So(t1.AddDate(0, 1, 3), ShouldEqual, t1.Add(34*24*time.Hour))
			So(t1.Round(time.Hour*24).Format(time.RFC3339), ShouldEqual, "2020-01-02T00:00:00Z")
			So(t1.Truncate(time.Hour*24).Format(time.RFC3339), ShouldEqual, "2020-01-02T00:00:00Z")
		})

		Convey("timezone", func() {
			{
				loc, _ := time.LoadLocation("Asia/Shanghai")
				t1 := time.Date(2020, 1, 2, 3, 4, 5, 6, loc)
				t2, _ := time.Parse(time.RFC3339, "2020-01-02T03:04:05.000000006+08:00")
				So(t1, ShouldEqual, t2)
				So(t1.UTC().Format(time.RFC3339), ShouldEqual, "2020-01-01T19:04:05Z")
				So(t1.Local().Format(time.RFC3339), ShouldEqual, "2020-01-02T03:04:05+08:00")
				So(t1.In(time.UTC), ShouldEqual, t1.UTC())
				name, offset := t1.Zone()
				So(name, ShouldEqual, "CST")
				So(offset, ShouldEqual, 8*3600)
			}
		})

		Convey("time serialize", func() {
			{
				var buf bytes.Buffer
				t1, _ := time.Parse(time.RFC3339, "2020-01-02T03:04:05.000000006+08:00")
				var t2 time.Time
				_ = gob.NewEncoder(&buf).Encode(t1)
				_ = gob.NewDecoder(&buf).Decode(&t2)
				So(t2, ShouldEqual, t1)
			}
			{
				t1, _ := time.Parse(time.RFC3339, "2020-01-02T03:04:05.000000006+08:00")
				var t2 time.Time
				buf, _ := json.Marshal(t1)
				_ = json.Unmarshal(buf, &t2)
				So(string(buf), ShouldEqual, "\"2020-01-02T03:04:05.000000006+08:00\"")
				So(t1, ShouldEqual, t2)
			}
		})
	})
}

func TestTimer(t *testing.T) {
	Convey("test timer", t, func() {
		{
			timer := time.AfterFunc(200*time.Millisecond, func() {
				fmt.Println("hello world")
			})
			So(timer.Stop(), ShouldBeTrue)
		}
		{
			timer := time.AfterFunc(200*time.Millisecond, func() {
				fmt.Println("hello world")
			})
			time.Sleep(250 * time.Millisecond)
			So(timer.Stop(), ShouldBeFalse)
		}
		{
			<-time.After(200 * time.Millisecond)
			fmt.Println("hello world")
		}
		{
			// 一定要放到 for select 外面，否则每次进入 for select 循环都会创建新的对象，进而导致资源泄漏
			every20Ms := time.Tick(20 * time.Millisecond)
			after200Ms := time.After(200 * time.Millisecond)
		out:
			for {
				select {
				case <-after200Ms:
					fmt.Println("quit")
					break out
				case <-every20Ms:
					fmt.Println("tick")
				}
			}
		}
	})
}

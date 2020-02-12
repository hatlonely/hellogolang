package mysql

import (
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	. "github.com/smartystreets/goconvey/convey"
)

type Wisdom struct {
	ID      int       `gorm:"type:bigint(20) auto_increment;primary_key" json:"id"`
	Author  string    `gorm:"type:varchar(64);index:author_idx;default:0" json:"author,omitempty"`
	Content string    `gorm:"type:longtext COLLATE utf8mb4_unicode_520_ci;not null" json:"content,omitempty"`
	CTime   time.Time `gorm:"type:timestamp;column:ctime;not null;default:CURRENT_TIMESTAMP;index:ctime_idx" json:"ctime,omitempty"`
	UTime   time.Time `gorm:"type:timestamp;column:utime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;index:utime_idx" json:"utime,omitempty"`
}

func TestGormCreateTable(t *testing.T) {
	Convey("test gorm", t, func() {
		db, err := gorm.Open("mysql", "hatlonely:keaiduo1@tcp(test-mysql:3306)/hads?charset=utf8mb4&parseTime=True&loc=Local")
		So(err, ShouldBeNil)
		So(db, ShouldNotBeNil)
		defer db.Close()

		db.DropTableIfExists(&Wisdom{})
		if !db.HasTable(&Wisdom{}) {
			So(db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&Wisdom{}).Error, ShouldBeNil)
		} else {
			So(db.AutoMigrate(&Wisdom{}).Error, ShouldBeNil)
		}

		db.DB().SetConnMaxLifetime(60 * time.Second)
		db.DB().SetMaxIdleConns(5)
		db.DB().SetMaxOpenConns(20)
	})
}

func TestGormInsert(t *testing.T) {
	db, _ := gorm.Open("mysql", "hatlonely:keaiduo1@tcp(test-mysql:3306)/hads?charset=utf8mb4&parseTime=True&loc=Local")
	defer db.Close()

	Convey("test insert", t, func() {
		wisdom := &Wisdom{ID: 1, Author: "鲁迅", Content: "真正的勇士，敢于直面惨淡的人生，敢于正视淋漓的鲜血"}
		So(db.Delete(&Wisdom{ID: 1}).Error, ShouldBeNil)

		Convey("test create", func() {
			op1 := db.Create(wisdom)
			So(op1.RowsAffected, ShouldEqual, 1)
			So(op1.Error, ShouldBeNil)

			op2 := db.Create(wisdom)
			So(op2.RowsAffected, ShouldEqual, 0)
			So(op2.Error, ShouldNotBeNil)
		})

		So(db.Delete(&Wisdom{ID: 1}).Error, ShouldBeNil)
	})
}

func TestGormSelect(t *testing.T) {
	db, _ := gorm.Open("mysql", "hatlonely:keaiduo1@tcp(test-mysql:3306)/hads?charset=utf8mb4&parseTime=True&loc=Local")
	defer db.Close()

	Convey("test select", t, func() {
		wisdom1 := &Wisdom{ID: 1, Author: "鲁迅", Content: "真正的勇士，敢于直面惨淡的人生，敢于正视淋漓的鲜血"}
		wisdom2 := &Wisdom{ID: 2, Author: "鲁迅", Content: "不在沉默中爆发，就在沉默中灭亡"}
		So(db.Delete(&Wisdom{ID: 1}).Error, ShouldBeNil)
		So(db.Delete(&Wisdom{ID: 2}).Error, ShouldBeNil)
		So(db.Create(wisdom1).Error, ShouldBeNil)
		So(db.Create(wisdom2).Error, ShouldBeNil)

		Convey("select by id", func() {
			{
				var wisdom Wisdom
				So(db.Find(&wisdom, 1).Error, ShouldBeNil)
				So(wisdom, ShouldResemble, *wisdom1)
			}
			{
				var wisdom Wisdom
				So(db.Find(&wisdom, 2).Error, ShouldBeNil)
				So(wisdom, ShouldResemble, *wisdom2)
			}
		})

		Convey("select by author", func() {
			var wisdoms []Wisdom
			So(db.Where("author=?", "鲁迅").Find(&wisdoms).Error, ShouldBeNil)
			So(len(wisdoms), ShouldEqual, 2)
			So(wisdoms[0], ShouldResemble, *wisdom1)
			So(wisdoms[1], ShouldResemble, *wisdom2)
		})

		Convey("select first", func() {
			{
				var wisdom Wisdom
				So(db.First(&wisdom).Error, ShouldBeNil)
				So(wisdom, ShouldResemble, *wisdom1)
			}
			{
				var wisdom Wisdom
				So(db.Where("author=?", "鲁迅").First(&wisdom).Error, ShouldBeNil)
				So(wisdom, ShouldResemble, *wisdom1)
			}
		})

		Convey("select fields", func() {
			var wisdom Wisdom
			So(db.Select("id, content").First(&wisdom).Error, ShouldBeNil)
			So(wisdom.ID, ShouldResemble, wisdom1.ID)
			So(wisdom.Content, ShouldResemble, wisdom1.Content)
		})
	})
}

func TestUpdate(t *testing.T) {
	db, _ := gorm.Open("mysql", "hatlonely:keaiduo1@tcp(test-mysql:3306)/hads?charset=utf8mb4&parseTime=True&loc=Local")
	defer db.Close()

	Convey("test update", t, func() {
		wisdom1 := &Wisdom{ID: 1, Author: "鲁迅", Content: "真正的勇士，敢于直面惨淡的人生，敢于正视淋漓的鲜血"}
		So(db.Delete(&Wisdom{ID: 1}).Error, ShouldBeNil)
		So(db.Create(wisdom1).Error, ShouldBeNil)

		So(db.Model(&Wisdom{}).Where("id=?", 1).Updates(&Wisdom{Content: "不在沉默中爆发，就在沉默中灭亡"}).Error, ShouldBeNil)
		var wisdom Wisdom
		So(db.Where("id=?", 1).Find(&wisdom).Error, ShouldBeNil)
		So(wisdom.Content, ShouldEqual, "不在沉默中爆发，就在沉默中灭亡")
	})
}

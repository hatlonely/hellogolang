package aliyun

import (
	"fmt"
	"testing"

	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	. "github.com/smartystreets/goconvey/convey"
)

var otsClient *tablestore.TableStoreClient

func init() {
	var err error
	_, accessKeyID, accessKeySecret, err := LoadOSSConfig()
	if err != nil {
		panic(err)
	}
	// https://otsnext.console.aliyun.com/cn-shanghai/hatlonely-ots-sh/list
	otsClient = tablestore.NewClient(
		"https://hatlonely-ots-sh.cn-shanghai.ots.aliyuncs.com",
		"hatlonely-ots-sh",
		accessKeyID,
		accessKeySecret,
	)
}

func TestTable(t *testing.T) {
	Convey("test table", t, func() {
		tables, err := otsClient.ListTable()
		So(err, ShouldBeNil)

		for _, table := range tables.TableNames {
			_, _ = Println(table)
		}
	})

	Convey("describe table", t, func() {
		info, err := otsClient.DescribeTable(&tablestore.DescribeTableRequest{
			TableName: "mysampletable",
		})
		So(err, ShouldBeNil)
		_, _ = Printf("%#v %#v", info.TableOption, info.TableMeta)
	})
}

func TestRow(t *testing.T) {
	Convey("put row", t, func() {
		pk := &tablestore.PrimaryKey{}
		pk.AddPrimaryKeyColumn("uid", "86")
		pk.AddPrimaryKeyColumn("pid", int64(6775))

		rc := &tablestore.PutRowChange{
			TableName:  "mysampletable",
			PrimaryKey: pk,
		}
		rc.AddColumn("name", "hatlonely")
		rc.AddColumn("country", "china")
		rc.SetCondition(tablestore.RowExistenceExpectation_IGNORE)

		res, err := otsClient.PutRow(&tablestore.PutRowRequest{PutRowChange: rc})
		So(err, ShouldBeNil)
		_, _ = Println(res)
	})

	Convey("get row", t, func() {
		pk := &tablestore.PrimaryKey{}
		pk.AddPrimaryKeyColumn("uid", "86")
		pk.AddPrimaryKeyColumn("pid", int64(6775))

		query := &tablestore.SingleRowQueryCriteria{
			TableName:  "mysampletable",
			PrimaryKey: pk,
			MaxVersion: 1,
		}

		res, err := otsClient.GetRow(&tablestore.GetRowRequest{
			SingleRowQueryCriteria: query,
		})

		So(err, ShouldBeNil)

		for _, col := range res.PrimaryKey.PrimaryKeys {
			fmt.Println(col.ColumnName, ":", col.Value)
		}
		for _, col := range res.Columns {
			_, _ = Println(col.ColumnName, ":", col.Value)
		}
	})

	Convey("update row", t, func() {
		pk := &tablestore.PrimaryKey{}
		pk.AddPrimaryKeyColumn("uid", "86")
		pk.AddPrimaryKeyColumn("pid", int64(6775))

		rc := &tablestore.UpdateRowChange{
			TableName:  "mysampletable",
			PrimaryKey: pk,
		}
		rc.PutColumn("name", "playjokes")
		rc.PutColumn("country", "china")
		rc.SetCondition(tablestore.RowExistenceExpectation_EXPECT_EXIST)

		res, err := otsClient.UpdateRow(&tablestore.UpdateRowRequest{UpdateRowChange: rc})
		So(err, ShouldBeNil)
		_, _ = Println(res)
	})

	Convey("del row", t, func() {
		pk := &tablestore.PrimaryKey{}
		pk.AddPrimaryKeyColumn("uid", "86")
		pk.AddPrimaryKeyColumn("pid", int64(6775))

		rc := &tablestore.DeleteRowChange{
			TableName:  "mysampletable",
			PrimaryKey: pk,
		}
		rc.SetCondition(tablestore.RowExistenceExpectation_EXPECT_EXIST)
		rc.SetColumnCondition(tablestore.NewSingleColumnCondition("name", tablestore.CT_EQUAL, "playjokes"))
		res, err := otsClient.DeleteRow(&tablestore.DeleteRowRequest{
			DeleteRowChange: rc,
		})
		So(err, ShouldBeNil)
		_, _ = Println(res)
	})
}

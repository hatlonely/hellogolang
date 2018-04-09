package aws

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	. "github.com/smartystreets/goconvey/convey"
)

func TestListTables(t *testing.T) {
	Convey("ListTables", t, func() {
		session, err := session.NewSession(&aws.Config{
			Region: aws.String(endpoints.ApSoutheast1RegionID),
		})
		So(err, ShouldBeNil)
		db := dynamodb.New(session)

		result, err := db.ListTables(&dynamodb.ListTablesInput{})
		So(err, ShouldBeNil)
		So(result.TableNames, ShouldResemble, []*string{})
	})
}

func TestCreateTable(t *testing.T) {
	Convey("CreateTable", t, func() {
		session, err := session.NewSession(&aws.Config{
			Region: aws.String(endpoints.ApSoutheast1RegionID),
		})
		So(err, ShouldBeNil)
		db := dynamodb.New(session)

		result, err := db.CreateTable(&dynamodb.CreateTableInput{
			AttributeDefinitions: []*dynamodb.AttributeDefinition{{
				AttributeName: aws.String("year"),
				AttributeType: aws.String("N"),
			}, {
				AttributeName: aws.String("title"),
				AttributeType: aws.String("S"),
			}},
			KeySchema: []*dynamodb.KeySchemaElement{{
				AttributeName: aws.String("year"),
				KeyType:       aws.String("HASH"),
			}, {
				AttributeName: aws.String("title"),
				KeyType:       aws.String("RANGE"),
			}},
			ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
				ReadCapacityUnits:  aws.Int64(10),
				WriteCapacityUnits: aws.Int64(10),
			},
			TableName: aws.String("Movies"),
		})
		So(err, ShouldBeNil)
		Println(result)
	})
}

type ItemInfo struct {
	Plot   string  `json:"plot"`
	Rating float64 `json:"rating"`
}

type Item struct {
	Year  int      `json:"year"`
	Title string   `json:"title"`
	Info  ItemInfo `json:"info"`
}

func TestPutItem(t *testing.T) {
	Convey("PutItem", t, func() {
		session, err := session.NewSession(&aws.Config{
			Region: aws.String(endpoints.ApSoutheast1RegionID),
		})
		So(err, ShouldBeNil)
		db := dynamodb.New(session)

		item := Item{
			Year:  2015,
			Title: "The Big New Movie",
			Info: ItemInfo{
				Plot:   "Nothing happens at all.",
				Rating: 0.0,
			},
		}

		av, err := dynamodbattribute.MarshalMap(item)

		result, err := db.PutItem(&dynamodb.PutItemInput{
			Item:      av,
			TableName: aws.String("Movies"),
		})

		So(err, ShouldBeNil)
		Println(result)
	})
}

func TestGetItem(t *testing.T) {
	Convey("GetItem", t, func() {
		session, err := session.NewSession(&aws.Config{
			Region: aws.String(endpoints.ApSoutheast1RegionID),
		})
		So(err, ShouldBeNil)
		db := dynamodb.New(session)

		result, err := db.GetItem(&dynamodb.GetItemInput{
			TableName: aws.String("Movies"),
			Key: map[string]*dynamodb.AttributeValue{
				"year": {
					N: aws.String("2015"),
				},
				"title": {
					S: aws.String("The Big New Movie"),
				},
			},
		})

		Println(result)

		item := &Item{}
		err = dynamodbattribute.UnmarshalMap(result.Item, item)
		So(err, ShouldBeNil)
		So(item.Info.Plot, ShouldEqual, "Nothing happens at all.")
		So(item.Info.Rating, ShouldEqual, 0.0)
	})
}

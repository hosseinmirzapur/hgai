package models

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const (
	FreeTextPrompts  uint = 20
	FreeImagePrompts uint = 5
)

type User struct {
	// Unique User ID
	ID           int64 `json:"id"`
	TextPrompts  uint  `json:"text_prompts"`
	ImagePrompts uint  `json:"image_prompts"`
}

// create a new user instance with default values
// values can be modified afterwards
// each user is given 20 free text prompts
// each user is given 5 free image processing prompts
// if user provides a caption for the image, it's counted as 1 text prompts as well
func NewUser() *User {
	return &User{
		ID:           -1,
		TextPrompts:  FreeImagePrompts,
		ImagePrompts: FreeImagePrompts,
	}
}

// save a non-empty user to dynamodb instance
// this method can be used to update `TextPrompts` and `ImagePrompts` value as well
// `Username` should never be updated with this method or any other method
func (u *User) SaveTo(service *dynamodb.DynamoDB) (*User, error) {
	if u.ID == -1 {
		return nil, fmt.Errorf("user ID not set")
	}

	data, err := dynamodbattribute.MarshalMap(u)
	if err != nil {
		return nil, err
	}

	_, err = service.PutItem(&dynamodb.PutItemInput{
		Item:      data,
		TableName: u.tableName(),
	})
	if err != nil {
		return nil, err
	}

	return u, nil
}

// scans all the data in `users` table in dynamodb instance
// `id` can be obtained from updated message: update.Message.From.ID
func (u *User) FindByIDIn(service *dynamodb.DynamoDB, id int64) (*User, error) {
	data, err := service.GetItem(&dynamodb.GetItemInput{
		TableName: u.tableName(),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String(fmt.Sprint(id)),
			},
		},
	})
	if err != nil {
		return nil, err
	}

	err = dynamodbattribute.UnmarshalMap(data.Item, u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (u *User) tableName() *string {
	return aws.String("users")
}

package pkg

import (
	"cmp"
	"errors"
	"fmt"
	"slices"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/comprehend"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/translate"
	"github.com/hosseinmirzapur/golangchain/pkg/models"
)

type AWS struct {
	sess  *session.Session
	trans *translate.Translate
	compr *comprehend.Comprehend

	inputLang string
}

// establish a new session for aws services
func NewSession() *session.Session {
	sess := session.Must(session.NewSession())

	return sess
}

// return a new comprehend service instance
func NewComprehend(sess *session.Session) *comprehend.Comprehend {
	svc := comprehend.New(sess)

	return svc
}

// return a new translate service instance
func NewTranslate(sess *session.Session) *translate.Translate {
	svc := translate.New(sess)

	return svc
}

func NewDynamoDB(sess *session.Session) *dynamodb.DynamoDB {
	svc := dynamodb.New(sess)

	return svc
}

// detect the language of the input text
func DetectLanguage(cmpr *comprehend.Comprehend, text string) (string, error) {
	input := &comprehend.DetectDominantLanguageInput{
		Text: &text,
	}

	output, err := cmpr.DetectDominantLanguage(input)
	if err != nil {
		return "", err
	}

	dominantLang := slices.MaxFunc(output.Languages, func(a, b *comprehend.DominantLanguage) int {
		return cmp.Compare(*a.Score, *b.Score)
	})

	return *dominantLang.LanguageCode, nil

}

// translate text from any language to any language
// that aws translate service supports
func TranslateTo(trans *translate.Translate, text, from, to string) (string, error) {
	output, err := trans.Text(&translate.TextInput{
		SourceLanguageCode: &from,
		TargetLanguageCode: &to,
		Text:               &text,
	})

	if err != nil {
		return "", err
	}

	return *output.TranslatedText, nil
}

func CreateUsersTable(dynamo *dynamodb.DynamoDB) error {
	params := &dynamodb.CreateTableInput{
		// defining table name
		TableName: aws.String("users"),
		// defining table data schema
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("user_id"),
				AttributeType: aws.String("N"),
			},
		},
		// defining primary key(s)
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("user_id"),
				KeyType:       aws.String("HASH"),
			},
		},
		// define table's throughput ()
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(5),
		},
	}

	_, err := dynamo.CreateTable(params)

	var resourceInUseErr *types.ResourceInUseException
	if errors.As(err, &resourceInUseErr) {
		return nil
	}

	return err
}

// register new bot user and save to dynamodb instance
// if user doesn't exist, gets saved and a message is returned
// if user exists, a message will be returned with no error
func RegisterNewUser(dynamo *dynamodb.DynamoDB, id int64) (string, error) {
	user := models.NewUser()

	existingUser, err := user.FindByIDIn(dynamo, id)
	if err != nil {
		return "", err
	}

	if fmt.Sprint(existingUser.UserID) == "" {
		user.UserID = id
		_, err := user.SaveTo(dynamo)
		if err != nil {
			return "", err
		}

		return "Successful Registeration!", nil
	}

	return "Already Registered!", nil
}

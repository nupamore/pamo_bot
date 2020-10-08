package services

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/translate"
)

// TranslateAWS : translate via aws
func TranslateAWS(source string, target string, text string) (*string, error) {
	response, err := AWStranslate.TranslateText(context.Background(), &translate.TranslateTextInput{
		SourceLanguageCode: aws.String(source),
		TargetLanguageCode: aws.String(target),
		Text:               aws.String(text),
	})

	if err != nil {
		return nil, err
	}

	return response.TranslatedText, nil
}

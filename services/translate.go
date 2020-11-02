package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/monaco-io/request"
	"github.com/nupamore/pamo_bot/configs"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/translate"
)

// TranslateService : translate service
type TranslateService struct{}

// Translate : translate service instance
var Translate = TranslateService{}

// AWS : translate via aws
func (s *TranslateService) AWS(source string, target string, text string) (*translate.TranslateTextOutput, error) {
	response, err := AWStranslate.TranslateText(context.Background(), &translate.TranslateTextInput{
		SourceLanguageCode: aws.String(source),
		TargetLanguageCode: aws.String(target),
		Text:               aws.String(text),
	})

	return response, err
}

type papagoResult struct {
	Source string `json:"srcLangType"`
	Target string `json:"tarLangType"`
	Text   string `json:"translatedText"`
}
type papagoMessage struct {
	Result papagoResult `json:"result"`
}
type papagoResponse struct {
	Message      papagoMessage `json:"message"`
	ErrorCode    string        `json:"errorCode"`
	ErrorMessage string        `json:"errorMessage"`
}

// Papago : translate via papago
func (s *TranslateService) Papago(source string, target string, text string) (*string, error) {
	client := request.Client{
		URL:    configs.Env["NAVER_TRANSLATE"],
		Method: "POST",
		Header: map[string]string{
			"X-Naver-Client-Id":     configs.Env["NAVER_ID"],
			"X-Naver-Client-Secret": configs.Env["NAVER_SECRET"],
		},
		Body: []byte(fmt.Sprintf(`{
            "source": "%s",
            "target": "%s",
            "text": "%s"
        }`, source, target, text)),
	}
	resp, err := client.Do()

	var res papagoResponse
	json.Unmarshal(resp.Data, &res)

	return &res.Message.Result.Text, err
}

type languageInfo struct {
	Code       string  `json:"code"`
	Name       string  `json:"name"`
	Confidence float64 `json:"confidence"`
}

// PapagoDetectResponse : detect lang response
type PapagoDetectResponse struct {
	ErrorType    string         `json:"errorType"`
	Message      string         `json:"message"`
	LanguageInfo []languageInfo `json:"language_info"`
}

// PapagoDetect : detect lang via kakao
func (s *TranslateService) PapagoDetect(text string) (PapagoDetectResponse, error) {
	client := request.Client{
		URL:    configs.Env["KAKAO_DETECT_LANG"],
		Method: "GET",
		Header: map[string]string{"Authorization": configs.Env["KAKAO_KEY"]},
		Params: map[string]string{"query": text},
	}
	resp, err := client.Do()

	var result PapagoDetectResponse
	json.Unmarshal(resp.Data, &result)

	return result, err
}

package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/monaco-io/request"

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

// TranslatePapago : translate via papago
func TranslatePapago(source string, target string, text string) (*string, error) {
	client := request.Client{
		URL:    os.Getenv("NAVER_TRANSLATE"),
		Method: "POST",
		Header: map[string]string{
			"X-Naver-Client-Id":     os.Getenv("NAVER_ID"),
			"X-Naver-Client-Secret": os.Getenv("NAVER_SECRET"),
		},
		Body: []byte(fmt.Sprintf(`{
            "source": "%s",
            "target": "%s",
            "text": "%s"
        }`, source, target, text)),
	}
	resp, err := client.Do()

	if err != nil {
		log.Println(err)
	}

	var res papagoResponse
	json.Unmarshal(resp.Data, &res)

	return &res.Message.Result.Text, err
}

type languageInfo struct {
	Code       string  `json:"code"`
	Name       string  `json:"name"`
	Confidence float64 `json:"confidence"`
}

// LanguageDetectResponse : detect lang response
type LanguageDetectResponse struct {
	ErrorType    string         `json:"errorType"`
	Message      string         `json:"message"`
	LanguageInfo []languageInfo `json:"language_info"`
}

// LanguageDetect : detect lang via kakao
func LanguageDetect(text string) (LanguageDetectResponse, error) {
	client := request.Client{
		URL:    os.Getenv("KAKAO_DETECT_LANG"),
		Method: "GET",
		Header: map[string]string{"Authorization": os.Getenv("KAKAO_KEY")},
		Params: map[string]string{"query": text},
	}
	resp, err := client.Do()

	if err != nil {
		log.Println(err)
	}

	var result LanguageDetectResponse
	json.Unmarshal(resp.Data, &result)

	return result, err
}

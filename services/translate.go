package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

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
	uri := configs.Env["NAVER_API"] + "/v1/papago/n2mt"
	body := bytes.NewBufferString(fmt.Sprintf(`{
        "source": "%s",
        "target": "%s",
        "text": "%s"
    }`, source, target, text))

	req, _ := http.NewRequest("POST", uri, body)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Naver-Client-Id", configs.Env["NAVER_ID"])
	req.Header.Add("X-Naver-Client-Secret", configs.Env["NAVER_SECRET"])

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var res papagoResponse
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &res)

	return &res.Message.Result.Text, err
}

type languageInfo struct {
	Code       string  `json:"code"`
	Name       string  `json:"name"`
	Confidence float64 `json:"confidence"`
}

// KakakoDetectResponse : detect lang response
type KakakoDetectResponse struct {
	ErrorType    string         `json:"errorType"`
	Message      string         `json:"message"`
	LanguageInfo []languageInfo `json:"language_info"`
}

// KakakoDetect : detect lang via kakao
func (s *TranslateService) KakakoDetect(text string) (KakakoDetectResponse, error) {
	uri := configs.Env["KAKAO_API"] + "/v3/translation/language/detect"
	query := "?query=" + url.QueryEscape(text)

	req, _ := http.NewRequest("GET", uri+query, nil)
	req.Header.Add("Authorization", configs.Env["KAKAO_KEY"])

	var res KakakoDetectResponse
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &res)

	return res, err
}

package services

import (
	"crypto/sha1"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"time"

	"github.com/diamondburned/arikawa/discord"
	"github.com/monaco-io/request"

	"github.com/nupamore/pamo_bot/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// Hash : sha1 6 letters to base64
func Hash(input string) string {
	h := sha1.New()
	h.Write([]byte(input))
	bs := h.Sum(nil)
	str := b64.URLEncoding.EncodeToString(bs[:3])

	return str
}

// InitLinks : init links
func InitLinks(ownerID discord.UserID) (models.SimpleDynamicLinkSlice, error) {
	// init already
	links, err := GetLinks(ownerID)
	if len(links) > 0 {
		return links, err
	}

	// create
	rand.Seed(time.Now().UnixNano())
	var newLinks models.SimpleDynamicLinkSlice

	for i := 0; i < 5; i++ {
		rand := strconv.Itoa(rand.Int())
		input := string(ownerID) + rand
		hash := Hash(input)

		link := models.SimpleDynamicLink{
			LinkID:  hash,
			OwnerID: null.UintFrom(uint(ownerID)),
			Status:  null.StringFrom("CREATED"),
			RegDate: null.TimeFrom(time.Now()),
		}
		err = link.Insert(DB, boil.Infer())
		if err != nil {
			isDuplicate, _ := regexp.MatchString("Error 1062", err.Error())
			if isDuplicate {
				i = i - 1
			}
		}
		newLinks = append(newLinks, &link)
	}
	return newLinks, err
}

// GetLinks : get links
func GetLinks(ownerID discord.UserID) (models.SimpleDynamicLinkSlice, error) {
	links, err := models.SimpleDynamicLinks(
		qm.Where("owner_id = ?", ownerID),
	).All(DB)

	return links, err
}

// GetLink : get link
func GetLink(linkID string) (*models.SimpleDynamicLink, error) {
	link, err := models.SimpleDynamicLinks(
		qm.Where("link_id = ?", linkID),
	).One(DB)

	return link, err
}

// UpdateLink : update link
func UpdateLink(linkID string, ownerID discord.UserID, options []byte) error {
	link, err := GetLink(linkID)
	json.Unmarshal(options, &link)
	link.Update(DB, boil.Infer())
	return err
}

// LogLink : log link
func LogLink(linkID string) error {
	now := time.Now()
	lastLog, _ := models.LinkLogs(
		qm.Where("link_id = ?", linkID),
		qm.OrderBy("reg_date DESC"),
	).One(DB)
	if lastLog != nil {
		if 24 < now.Sub(lastLog.ViewDate.Time).Hours() {
			go LinkValidTest(linkID)
		}
	}

	newLog := models.LinkLog{
		LinkID:   null.TimeFrom(now),
		ViewDate: null.TimeFrom(now),
	}

	return newLog.Insert(DB, boil.Infer())
}

// LinkValidTest : link target valid test
func LinkValidTest(linkID string) {
	link, _ := GetLink(linkID)
	ownerID := discord.UserID(*link.OwnerID.Ptr())

	client := request.Client{
		URL:    *link.Target.Ptr(),
		Method: "GET",
	}
	resp, err := client.Do()

	if err != nil || resp.StatusCode() != 200 {
		message := fmt.Sprintf(
			"**Link validation failed**\nlinkID: %s\ntarget: %s",
			linkID,
			*link.Target.Ptr(),
		)
		go SendDM(ownerID, message)

		link.Target = null.StringFrom("404")
		link.Status = null.StringFrom("ERROR")
		link.Update(DB, boil.Infer())
	}
}

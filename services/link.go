package services

import (
	"crypto/sha1"
	b64 "encoding/base64"
	"encoding/json"
	"errors"
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

// LinkService : link service
type LinkService struct{}

// Link : link service instance
var Link = LinkService{}

func newHash(input string) string {
	h := sha1.New()
	h.Write([]byte(input))
	bs := h.Sum(nil)
	str := b64.URLEncoding.EncodeToString(bs[:3])

	return str
}

// Create : init links
func (s *LinkService) Create(ownerID discord.UserID) (models.SimpleDynamicLinkSlice, error) {
	// init already
	links, err := s.List(ownerID)
	if len(links) > 0 {
		return links, err
	}

	// create
	rand.Seed(time.Now().UnixNano())
	var newLinks models.SimpleDynamicLinkSlice

	for i := 0; i < 5; i++ {
		rand := strconv.Itoa(rand.Int())
		input := string(ownerID) + rand
		hash := newHash(input)

		link := models.SimpleDynamicLink{
			LinkID:  hash,
			OwnerID: null.StringFrom(strconv.FormatUint(uint64(ownerID), 10)),
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

// List : get links
func (s *LinkService) List(ownerID discord.UserID) (models.SimpleDynamicLinkSlice, error) {
	links, err := models.SimpleDynamicLinks(
		qm.Where("owner_id = ?", ownerID),
	).All(DB)

	return links, err
}

// Info : get link
func (s *LinkService) Info(linkID string) (*models.SimpleDynamicLink, error) {
	link, err := models.SimpleDynamicLinks(
		qm.Where("link_id = ?", linkID),
	).One(DB)

	return link, err
}

// Update : update link
func (s *LinkService) Update(linkID string, ownerID discord.UserID, options []byte) (models.SimpleDynamicLink, error) {
	link, err := s.Info(linkID)
	json.Unmarshal(options, &link)

	if !s.Test(link) {
		err = errors.New("Link is invalid")
	} else {
		link.Status = null.StringFrom("UPDATED")
		link.Update(DB, boil.Infer())
	}

	return *link, err
}

// Log : log link
func (s *LinkService) Log(linkID string) error {
	now := time.Now()
	tempTime, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	temp := null.TimeFrom(tempTime)

	lastLog, _ := models.LinkLogs(
		qm.Where("link_id = ?", temp),
		qm.OrderBy("view_date DESC"),
	).One(DB)
	if lastLog != nil && 24 < now.Sub(lastLog.ViewDate.Time).Hours() {
		go func() {
			link, _ := s.Info(linkID)
			if !s.Test(link) {
				s.TestFailEvent(link)
			}
		}()
	}

	newLog := models.LinkLog{
		LinkID:   temp,
		ViewDate: null.TimeFrom(now),
	}

	return newLog.Insert(DB, boil.Infer())
}

// Test : link target valid test
func (s *LinkService) Test(link *models.SimpleDynamicLink) bool {
	client := request.Client{
		URL:    *link.Target.Ptr(),
		Method: "GET",
	}
	resp, err := client.Do()

	if err != nil || resp.StatusCode() != 200 {
		return false
	}

	return true
}

// TestFailEvent : link valid fail event
func (s *LinkService) TestFailEvent(link *models.SimpleDynamicLink) {
	// send dm
	ownerID, _ := strconv.ParseUint(*link.OwnerID.Ptr(), 10, 64)
	message := fmt.Sprintf(
		"**Link validation failed**\nlinkID: %s\ntarget: %s",
		link.LinkID,
		*link.Target.Ptr(),
	)
	go SendDM(discord.UserID(ownerID), message)

	// target change 404
	var target string
	if link.Type.Ptr() == nil {
		target = "https://github.com/404"
	} else {
		switch *link.Type.Ptr() {
		case "image":
			target = "https://github.com/404"
		case "video":
			target = "https://github.com/404"
		}
	}
	link.Target = null.StringFrom(target)
	link.Status = null.StringFrom("ERROR")
	link.Update(DB, boil.Infer())
}

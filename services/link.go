package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dchest/uniuri"
	"github.com/diamondburned/arikawa/v2/discord"

	"github.com/nupamore/pamo_bot/configs"
	"github.com/nupamore/pamo_bot/models"
	"github.com/nupamore/pamo_bot/utils"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// LinkService : link service
type LinkService struct{}

// Link : link service instance
var Link = LinkService{}

// Create : init links
func (s *LinkService) Create(ownerID discord.UserID) (models.SimpleDynamicLinkSlice, error) {
	// init already
	links, err := s.List(ownerID)
	if len(links) > 0 {
		return links, err
	}

	// create
	var newLinks models.SimpleDynamicLinkSlice

	for i := 0; i < 5; i++ {
		hash := uniuri.NewLen(4)

		link := models.SimpleDynamicLink{
			LinkID:  hash,
			OwnerID: null.StringFrom(strconv.FormatUint(uint64(ownerID), 10)),
			Status:  null.StringFrom("CREATED"),
			RegDate: null.TimeFrom(time.Now()),
		}
		err = link.Insert(DB, boil.Infer())
		if utils.IsDuplicate(err) {
			i = i - 1
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
	lastLog, _ := models.LinkLogs(
		qm.Where("link_id = ?", linkID),
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
		LinkID:   null.StringFrom(linkID),
		ViewDate: null.TimeFrom(now),
	}

	return newLog.Insert(DB, boil.Infer())
}

// Test : link target valid test
func (s *LinkService) Test(link *models.SimpleDynamicLink) bool {
	resp, err := http.Get(*link.Target.Ptr())
	if err != nil || resp.StatusCode != 200 {
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
	target := configs.Env["LINK_RESERVE_PAGE"]
	reserve := link.ReserveTarget.Ptr()
	typ := link.Type.Ptr()

	if reserve != nil {
		target = *reserve
	} else if typ != nil {
		switch *typ {
		case "image":
			target = configs.Env["LINK_RESERVE_IMAGE"]
		case "video":
			target = configs.Env["LINK_RESERVE_VIDEO"]
		}
	}
	link.Target = null.StringFrom(target)
	link.Status = null.StringFrom("ERROR")
	link.Update(DB, boil.Infer())
}

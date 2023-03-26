package gamelog

import (
	"regexp"
	"strings"

	"github.com/new-pop-corn/internal/model"
)

const (
	playerName = `([A-Z]\.\s[A-ZĀČĒĢĪĶĻŅOŠŪŽa-zāčēģīķļņošūžó'\-A-Za-zčćšđžČĆŠĐŽņģŅĢó]+)`
	make2      = `(makes 2-pt)`
	make3      = `(makes 3-pt)`
	makeft     = `(makes free throw)`
	miss2      = `(misses 2-pt)`
	miss3      = `(misses 3-pt)`
	missft     = `(misses free throw)`
	offReb     = `(Offensive rebound)`
	defReb     = `(Defensive rebound)`
	assist     = `(assist)`
	to         = `(Turnover)`
	steal      = `(steal)`
	block      = `(block)`
	swtichDesc = "enters the game for"
)

func Parse(gameLog string) []*model.PlayerAction {
	var playerActions []*model.PlayerAction
	playerNamePattern := regexp.MustCompile(playerName)
	nameMatch := playerNamePattern.FindAllStringSubmatch(gameLog, -1)
	if strings.Contains(gameLog, swtichDesc) {
		if len(nameMatch) > 1 {
			player1Action := &model.PlayerAction{
				Player: nameMatch[0][0],
				Event:  model.ON,
			}
			player2Action := &model.PlayerAction{
				Player: nameMatch[1][0],
				Event:  model.OFF,
			}
			playerActions = append(playerActions, player1Action, player2Action)
		}
		return playerActions
	}

	patterns := []string{
		make2, make3, makeft, miss2, miss3, missft, offReb, defReb, assist, to, steal, block,
	}

	var player string
	var event model.Event
	if len(nameMatch) > 0 {
		player = nameMatch[0][1]
	}

	for _, p := range patterns {
		regP := regexp.MustCompile(p)
		eventMatch := regP.FindStringSubmatch(gameLog)
		if len(eventMatch) > 0 {
			switch eventMatch[0] {
			case "makes 2-pt":
				event = model.MAKE2
			case "makes 3-pt":
				event = model.MAKE3
			case "misses 2-pt":
				event = model.MISS2
			case "misses 3-pt":
				event = model.MISS3
			case "misses free throw":
				event = model.MISSFT
			case "makes free throw":
				event = model.MAKEFT
			case "Offensive rebound":
				event = model.OFFREBOUND
			case "Defensive rebound":
				event = model.DEFREBOUND
			case "assist":
				event = model.ASSIST
			case "Turnover":
				event = model.TO
			case "steal":
				event = model.STEAL
			case "block":
				event = model.BLOCK
			}
		}
	}
	playerActions = append(playerActions, &model.PlayerAction{
		Player: player,
		Event:  event,
	})
	return playerActions
}

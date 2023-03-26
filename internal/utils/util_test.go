package utils

import (
	"fmt"
	"testing"

	"github.com/new-pop-corn/internal/pkg/model"
	"github.com/new-pop-corn/internal/pkg/utils/gamelog"
)

func TestParseDesc(t *testing.T) {
	testCases := []struct {
		desc          string
		expectedName  string
		expectedEvent model.Event
	}{
		{"J. Tatum makes 2-pt layup at rim", "J. Tatum", model.MAKE2},
		{"M. Smart misses 3-pt jump shot from 24 ft", "M. Smart", model.MISS3},
		{"Offensive rebound by T. Harris", "T. Harris", model.OFFREBOUND},
		{"Defensive rebound by J. Embiid", "J. Embiid", model.DEFREBOUND},
		{"A. Horford assist", "A. Horford", model.ASSIST},
		{"Turnover by L. Ball", "L. Ball", model.TO},
		{"steal by R. Westbrook", "R. Westbrook", model.STEAL},
		{"block by G. Antetokounmpo", "G. Antetokounmpo", model.BLOCK},
		{"Offensive rebound by D. White", "D. White", model.OFFREBOUND},
		{"J. Brown misses 2-pt layup from 5 ft", "J. Brown", model.MISS2},
		{"J. Embiid misses free throw 1 of 2", "J. Embiid", model.MISSFT},
		{"J. Toscano-Anderson misses free throw 1 of 2", "J. Toscano-Anderson", model.MISSFT},
		{"Turnover by B. Bogdanović (bad pass)", "B. Bogdanović", model.TO},
		{"J. Valančiūnas misses 2-pt layup from 4 ft", "J. Valančiūnas", model.MISS2},
		{"J. Hernangómez enters the game for D. Banton", "J. Hernangómez", model.ON},
		{"R. O'Neale misses 2-pt jump shot from 9 ft", "R. O'Neale", model.MISS2},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			playerActions := gamelog.Parse(tc.desc)
			fmt.Println(playerActions[0].Player)
			fmt.Println(tc.expectedName == playerActions[0].Player)
			if len(playerActions) == 1 {
				if playerActions[0].Player != tc.expectedName {
					t.Errorf("Expected player name: %s, got: %s", tc.expectedName, playerActions[0].Player)
				}
				if playerActions[0].Event != tc.expectedEvent {
					t.Errorf("Expected event: %s, got: %s", tc.expectedEvent, playerActions[0].Event)
				}
			}
		})
	}
}

package utils

import (
	"fmt"
	"testing"

	"github.com/new-pop-corn/internal/model"
	"github.com/new-pop-corn/internal/utils/gamelog"
)

func TestParseDesc(t *testing.T) {
	testCases := []struct {
		desc        string
		expectedMap map[string]model.Event
	}{
		{"J. Tatum makes 2-pt layup at rim", map[string]model.Event{"J. Tatum": model.MAKE2}},
		{"M. Smart misses 3-pt jump shot from 24 ft", map[string]model.Event{"M. Smart": model.MISS3}},
		{"Offensive rebound by T. Harris", map[string]model.Event{"T. Harris": model.OFFREBOUND}},
		{"Defensive rebound by J. Embiid", map[string]model.Event{"J. Embiid": model.DEFREBOUND}},
		{"A. Horford assist", map[string]model.Event{"A. Horford": model.ASSIST}},
		{"Turnover by L. Ball", map[string]model.Event{"L. Ball": model.TO}},
		{"J. Toscano-Anderson misses free throw 1 of 2", map[string]model.Event{"J. Toscano-Anderson": model.MISSFT}},
		// {"steal by R. Westbrook", "R. Westbrook", model.STEAL},
		// {"block by G. Antetokounmpo", "G. Antetokounmpo", model.BLOCK},
		// {"Offensive rebound by D. White", "D. White", model.OFFREBOUND},
		// {"J. Brown misses 2-pt layup from 5 ft", "J. Brown", model.MISS2},
		// {"J. Embiid misses free throw 1 of 2", "J. Embiid", model.MISSFT},
		// {"Turnover by B. Bogdanović (bad pass)", "B. Bogdanović", model.TO},
		// {"J. Valančiūnas misses 2-pt layup from 4 ft", "J. Valančiūnas", model.MISS2},
		{"J. Hernangómez enters the game for D. Banton", map[string]model.Event{"J. Hernangómez": model.ON, "D. Banton": model.OFF}},
		{"R. O'Neale misses 2-pt jump shot from 9 ft", map[string]model.Event{"R. O'Neale": model.MISS2}},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			playerActions := gamelog.Parse(tc.desc)
			for _, pa := range playerActions {
				fmt.Printf("playerActions: %s, %s\n", pa.Player, pa.Event)
				if _, ok := tc.expectedMap[pa.Player]; !ok {
					t.Errorf("Expected player name: %s", pa.Player)
				} else {
					if tc.expectedMap[pa.Player] != pa.Event {
						t.Errorf("Expected event: %s, got: %s", tc.expectedMap[pa.Player], pa.Event)
					}
				}
			}
		})
	}
}

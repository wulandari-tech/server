//go:build !prod || full || netent

package trolls

import (
	"github.com/slotopol/server/game"
	"github.com/slotopol/server/util"
)

var Info = game.GameInfo{
	Aliases: []game.GameAlias{
		{Prov: "NetEnt", Name: "Trolls"},
		{Prov: "NetEnt", Name: "Excalibur"},
		{Prov: "NetEnt", Name: "Pandora's Box"},
		{Prov: "NetEnt", Name: "Wild Witches"},
	},
	GP: game.GPsel |
		game.GPretrig |
		game.GPfgmult |
		game.GPscat |
		game.GPwild |
		game.GPwmult,
	SX:  5,
	SY:  3,
	SN:  len(LinePay),
	LN:  len(BetLines),
	BN:  0,
	RTP: game.MakeRtpList(ReelsMap),
}

func init() {
	game.GameList = append(game.GameList, &Info)
	for _, ga := range Info.Aliases {
		var aid = util.ToID(ga.Prov + "/" + ga.Name)
		game.ScanFactory[aid] = CalcStat
		game.GameFactory[aid] = func() any { return NewGame() }
	}
}

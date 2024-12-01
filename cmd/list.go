package cmd

import (
	"fmt"
	"sort"
	"strings"

	cfg "github.com/slotopol/server/config"
	"github.com/slotopol/server/game"
	"github.com/slotopol/server/util"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var listflags *pflag.FlagSet

var (
	fAll, fProp, fRTP bool
)

const listShort = "List of available games released on server"
const listLong = ``
const listExmp = `Get the list of all available games:
  %[1]s list --all
Get the list of available 'NetExt' and 'BetSoft' games only:
  %[1]s list --netent --betsoft
Get the list of available 'Play'n GO' games with RTP list for each:
  %[1]s list --playngo --rtp`

func incinfo(gi *game.GameInfo) bool {
	if fAll {
		return true
	}
	var is bool
	if is, _ = listflags.GetBool("keno"); is && gi.SX == 80 {
		return true
	}
	if is, _ = listflags.GetBool("3reels"); is && gi.SX == 3 {
		return true
	}
	if is, _ = listflags.GetBool("5reels"); is && gi.SX == 5 {
		return true
	}
	if is, _ = listflags.GetBool("3x3"); is && gi.SX == 3 && gi.SY == 3 {
		return true
	}
	if is, _ = listflags.GetBool("5x3"); is && gi.SX == 5 && gi.SY == 3 {
		return true
	}
	if is, _ = listflags.GetBool("5x4"); is && gi.SX == 5 && gi.SY == 4 {
		return true
	}
	if is, _ = listflags.GetBool("fewlines"); is && gi.LN < 20 {
		return true
	}
	if is, _ = listflags.GetBool("multilines"); is && gi.LN >= 20 {
		return true
	}
	if is, _ = listflags.GetBool("megaway"); is && gi.LN > 100 {
		return true
	}
	if is, _ = listflags.GetBool("fg"); is && gi.GP&(game.GPfghas+game.GPretrig) > 0 {
		return true
	}
	if is, _ = listflags.GetBool("bonus"); is && gi.BN > 0 {
		return true
	}
	return false
}

func isProv(prov string) bool {
	var is, _ = listflags.GetBool(util.ToID(prov))
	return is
}

func FormatGameInfo(gi *game.GameInfo, ai int) string {
	var b strings.Builder
	if gi.SN > 0 {
		fmt.Fprintf(&b, "'%s' %s %dx%d videoslot", gi.Aliases[ai].Name, gi.Aliases[ai].Prov, gi.SX, gi.SY)
	} else {
		fmt.Fprintf(&b, "'%s' %s %d spots lottery", gi.Aliases[ai].Name, gi.Aliases[ai].Prov, gi.SX)
	}
	if fProp {
		if gi.SN > 0 {
			fmt.Fprintf(&b, ", %d symbols", gi.SN)
		}
		if gi.LN > 100 {
			fmt.Fprintf(&b, ", %d ways", gi.LN)
		} else if gi.LN > 0 {
			if gi.GP&game.GPsel == 0 {
				fmt.Fprintf(&b, ", constant %d lines", gi.LN)
			} else {
				fmt.Fprintf(&b, ", %d lines", gi.LN)
			}
		}
		if gi.GP&game.GPjack > 0 {
			b.WriteString(", has jackpot")
		}
		if gi.GP&(game.GPfghas+game.GPretrig) > 0 {
			b.WriteString(", ")
			if gi.GP&game.GPretrig > 0 {
				b.WriteString("retriggerable ")
			}
			b.WriteString("free games")
			if gi.GP&game.GPfgmult > 0 {
				b.WriteString(" with multiplier")
			}
			if gi.GP&game.GPfgreel > 0 {
				b.WriteString(" on bonus reels")
			}
		}
		if gi.GP&game.GPscat > 0 {
			b.WriteString(", has scatters")
		}
		if gi.GP&game.GPwild > 0 {
			if gi.GP&game.GPwmult > 0 {
				b.WriteString(", has wilds with multiplier")
			} else {
				b.WriteString(", has wilds")
			}
		}
		if gi.GP&game.GPrwild > 0 {
			if gi.GP&game.GPwmult > 0 {
				b.WriteString(", has reel wilds with multiplier")
			} else {
				b.WriteString(", has reel wilds")
			}
		}
		if gi.GP&game.GPbwild > 0 {
			b.WriteString(", has big wilds")
		}
		if gi.GP&game.GPbsym > 0 {
			b.WriteString(", has big symbols")
		}
	}
	if fRTP && len(gi.RTP) > 0 {
		b.WriteString(", RTP: ")
		for i, rtp := range gi.RTP {
			if i > 0 {
				b.WriteString(", ")
			}
			fmt.Fprintf(&b, "%.2f", rtp)
		}
	}
	return b.String()
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Short:   listShort,
	Long:    listLong,
	Example: fmt.Sprintf(listExmp, cfg.AppName),
	Run: func(cmd *cobra.Command, args []string) {
		var num, alg int
		var prov = map[string]int{}
		for _, gi := range game.GameList {
			var inc = incinfo(gi)
			for _, ga := range gi.Aliases {
				if inc || isProv(ga.Prov) {
					num++
				}
			}
		}

		var gamelist = make([]string, num)
		var i int
		for _, gi := range game.GameList {
			var inc = incinfo(gi)
			var has bool
			for ai, ga := range gi.Aliases {
				if inc || isProv(ga.Prov) {
					has = true
					prov[ga.Prov]++
					gamelist[i] = FormatGameInfo(gi, ai)
					i++
				}
			}
			if has {
				alg++
			}
		}
		var provlist = make([]string, len(prov))
		i = 0
		for p, n := range prov {
			provlist[i] = fmt.Sprintf("%s: %d games", p, n)
			i++
		}

		if is, _ := listflags.GetBool("name"); is {
			fmt.Println()
			sort.Strings(gamelist)
			for _, s := range gamelist {
				fmt.Println(s)
			}
		}
		if is, _ := listflags.GetBool("stat"); is {
			fmt.Println()
			fmt.Printf("total: %d games, %d algorithms, %d providers\n", num, alg, len(prov))
			sort.Strings(provlist)
			for _, s := range provlist {
				fmt.Println(s)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listflags = listCmd.Flags()
	listflags.BoolP("name", "n", true, "list of provided games names")
	listflags.BoolP("stat", "s", true, "summary statistics of provided games")

	listflags.BoolVar(&fAll, "all", false, "include all provided games, overrides any other filters")
	listflags.BoolVar(&fProp, "prop", false, "print properties for each game")
	listflags.BoolVar(&fRTP, "rtp", false, "RTP (Return to Player) percents list of available reels for each game")

	listflags.Bool("agt", false, "include games of 'AGT' provider")
	listflags.Bool("aristocrat", false, "include games of 'Aristocrat' provider")
	listflags.Bool("betsoft", false, "include games of 'BetSoft' provider")
	listflags.Bool("megajack", false, "include games of 'Megajack' provider")
	listflags.Bool("netent", false, "include games of 'NetExt' provider")
	listflags.Bool("novomatic", false, "include games of 'Novomatic' provider")
	listflags.Bool("playngo", false, "include games of 'Play'n GO' provider")
	listflags.Bool("playtech", false, "include games of 'Playtech' provider")
	listflags.Bool("slotopol", false, "include games of this 'Slotopol' provider")

	listflags.Bool("keno", false, "include keno games")
	listflags.Bool("3reels", false, "include games with 3 reels")
	listflags.Bool("5reels", false, "include games with 5 reels")
	listflags.Bool("3x3", false, "include games with 3x3 screen")
	listflags.Bool("5x3", false, "include games with 5x3 screen")
	listflags.Bool("5x4", false, "include games with 5x4 screen")
	listflags.Bool("fewlines", false, "include games with few lines, i.e. with less than 20")
	listflags.Bool("multilines", false, "include games with few lines, i.e. with not less than 20")
	listflags.Bool("megaway", false, "include games with multiways, i.e. with 243, 1024 ways")
	listflags.Bool("fg", false, "include games with any free games")
	listflags.Bool("bonus", false, "include games with bonus games")

	listCmd.MarkFlagsOneRequired("all",
		"agt", "aristocrat", "betsoft", "megajack", "netent", "novomatic", "playngo", "playtech", "slotopol",
		"keno", "3reels", "5reels", "3x3", "5x3", "5x4", "fewlines", "multilines", "megaway", "fg", "bonus")
}

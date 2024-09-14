package goldentour

import (
	"context"
	"fmt"
	"strconv"
	"time"

	slot "github.com/slotopol/server/game/slot"
	"github.com/slotopol/server/util"
)

var Egolfbn float64

func ExpGolf() float64 {
	var sum float64
	for _, v := range Golf {
		sum += float64(v)
	}
	var E = sum / float64(len(Golf))
	fmt.Printf("len = %d, E = %g\n", len(Golf), E)
	return E
}

func CalcStat(ctx context.Context, rn string) float64 {
	fmt.Printf("*bonus games calculations*\n")
	Egolfbn = ExpGolf()
	fmt.Printf("*reels calculations*\n")
	var reels *slot.Reels5x
	if mrtp, _ := strconv.ParseFloat(rn, 64); mrtp != 0 {
		var _, r = FindReels(mrtp)
		reels = r.(*slot.Reels5x)
	} else {
		reels = &Reels145
	}
	var g = NewGame()
	g.SBL = util.MakeBitNum(1, 1)
	var sbl = float64(g.SBL.Num())
	var s slot.Stat

	var total = float64(reels.Reshuffles())
	var dur = func() time.Duration {
		var t0 = time.Now()
		var ctx2, cancel2 = context.WithCancel(ctx)
		defer cancel2()
		go s.Progress(ctx2, time.Tick(2*time.Second), sbl, total)
		slot.BruteForce5x(ctx2, &s, g, reels)
		return time.Since(t0)
	}()

	var reshuf = float64(s.Reshuffles)
	var lrtp, srtp = s.LinePay / reshuf / sbl * 100, s.ScatPay / reshuf / sbl * 100
	var rtpsym = lrtp + srtp
	var qgolfbn = float64(s.BonusCount[golfbon]) / reshuf / sbl
	var rtpgolfbn = Egolfbn * qgolfbn * 100
	var rtp = rtpsym + rtpgolfbn
	fmt.Printf("completed %.5g%%, selected %d lines, time spent %v\n", reshuf/total*100, g.SBL.Num(), dur)
	fmt.Printf("reels lengths [%d, %d, %d, %d, %d], total reshuffles %d\n",
		len(reels.Reel(1)), len(reels.Reel(2)), len(reels.Reel(3)), len(reels.Reel(4)), len(reels.Reel(5)), reels.Reshuffles())
	fmt.Printf("golf bonuses: count %d, rtp = %.6f%%\n", s.BonusCount[golfbon], rtpgolfbn)
	fmt.Printf("golf bonuses frequency: 1/%.5g\n", reshuf/float64(s.BonusCount[golfbon]))
	fmt.Printf("RTP = %.5g(lined) + %.5g(scatter) + %.5g(golf) = %.6f%%\n", lrtp, srtp, rtpgolfbn, rtp)
	return rtp
}

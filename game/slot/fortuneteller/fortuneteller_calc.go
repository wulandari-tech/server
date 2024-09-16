package fortuneteller

import (
	"context"
	"fmt"
	"strconv"
	"time"

	slot "github.com/slotopol/server/game/slot"
)

var Ecards float64

func ExpCards() float64 {
	var sum float64
	for c1 := 1; c1 <= 4; c1++ {
		for c2 := 1; c2 <= 4; c2++ {
			for c3 := 1; c3 <= 4; c3++ {
				sum += CardsWin(c1, c2, c3)
			}
		}
	}
	var E = sum / 4 / 4 / 4
	fmt.Printf("total = %d, E = %g\n", 4*4*4, E)
	return E
}

func CalcStatBon(ctx context.Context, rn string) float64 {
	var reels *slot.Reels5x
	if mrtp, _ := strconv.ParseFloat(rn, 64); mrtp != 0 {
		var _, r = FindReels(mrtp)
		reels = r.(*slot.Reels5x)
	} else {
		reels = &Reels92
	}
	var g = NewGame()
	var sln float64 = 1
	g.Sel.SetNum(int(sln), 1)
	g.FS = 15 // set free spins mode
	var s slot.Stat

	var total = float64(reels.Reshuffles())
	var dur = func() time.Duration {
		var t0 = time.Now()
		var ctx2, cancel2 = context.WithCancel(ctx)
		defer cancel2()
		go s.Progress(ctx2, time.Tick(2*time.Second), sln, total)
		slot.BruteForce5x(ctx2, &s, g, reels)
		return time.Since(t0)
	}()

	var reshuf = float64(s.Reshuffles)
	var lrtp, srtp = s.LinePay / reshuf / sln * 100, s.ScatPay / reshuf / sln * 100
	var rtpsym = lrtp + srtp
	var qcbn = float64(s.BonusCount[cbn]) / reshuf / sln
	var rtpcbn = Ecards * qcbn * 100
	var rtp = rtpsym + rtpcbn
	fmt.Printf("completed %.5g%%, selected %d lines, time spent %v\n", reshuf/total*100, g.Sel.Num(), dur)
	fmt.Printf("reels lengths [%d, %d, %d, %d, %d], total reshuffles %d\n",
		len(reels.Reel(1)), len(reels.Reel(2)), len(reels.Reel(3)), len(reels.Reel(4)), len(reels.Reel(5)), reels.Reshuffles())
	fmt.Printf("symbols: %.5g(lined) + %.5g(scatter) = %.6f%%\n", lrtp, srtp, rtpsym)
	fmt.Printf("cards bonuses: frequency 1/%.5g, rtp = %.6f%%\n", reshuf/float64(s.BonusCount[cbn]), rtpcbn)
	fmt.Printf("RTP = %.5g(sym) + %.5g(cards) = %.6f%%\n", rtpsym, rtpcbn, rtp)
	return rtp
}

func CalcStatReg(ctx context.Context, rn string) float64 {
	fmt.Printf("*bonus games calculations*\n")
	Ecards = ExpCards()
	fmt.Printf("*bonus reels calculations*\n")
	var rtpfs = CalcStatBon(ctx, rn)
	if ctx.Err() != nil {
		return 0
	}
	fmt.Printf("*regular reels calculations*\n")
	var reels *slot.Reels5x
	if mrtp, _ := strconv.ParseFloat(rn, 64); mrtp != 0 {
		var _, r = FindReels(mrtp)
		reels = r.(*slot.Reels5x)
	} else {
		reels = &Reels92
	}
	var g = NewGame()
	var sln float64 = 1
	g.Sel.SetNum(int(sln), 1)
	var s slot.Stat

	var total = float64(reels.Reshuffles())
	var dur = func() time.Duration {
		var t0 = time.Now()
		var ctx2, cancel2 = context.WithCancel(ctx)
		defer cancel2()
		go s.Progress(ctx2, time.Tick(2*time.Second), sln, total)
		slot.BruteForce5x(ctx2, &s, g, reels)
		return time.Since(t0)
	}()

	var reshuf = float64(s.Reshuffles)
	var lrtp, srtp = s.LinePay / reshuf / sln * 100, s.ScatPay / reshuf / sln * 100
	var rtpsym = lrtp + srtp
	var q = float64(s.FreeCount) / reshuf
	var qcbn = float64(s.BonusCount[cbn]) / reshuf / sln
	var rtpcbn = Ecards * qcbn * 100
	var rtp = rtpsym + rtpcbn + q*rtpfs
	fmt.Printf("completed %.5g%%, selected %d lines, time spent %v\n", reshuf/total*100, g.Sel.Num(), dur)
	fmt.Printf("reels lengths [%d, %d, %d, %d, %d], total reshuffles %d\n",
		len(reels.Reel(1)), len(reels.Reel(2)), len(reels.Reel(3)), len(reels.Reel(4)), len(reels.Reel(5)), reels.Reshuffles())
	fmt.Printf("symbols: %.5g(lined) + %.5g(scatter) = %.6f%%\n", lrtp, srtp, rtpsym)
	fmt.Printf("free spins %d, q = %.6f\n", s.FreeCount, q)
	fmt.Printf("free games frequency: 1/%.5g\n", reshuf/float64(s.FreeHits))
	fmt.Printf("cards bonuses: frequency 1/%.5g, rtp = %.6f%%\n", reshuf/float64(s.BonusCount[cbn]), rtpcbn)
	fmt.Printf("RTP = %.5g(sym) + %.5g(cards) + %.5g*%.5g(fg) = %.6f%%\n", rtpsym, rtpcbn, q, rtpfs, rtp)
	return rtp
}

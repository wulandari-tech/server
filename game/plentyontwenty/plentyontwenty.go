package plentyontwenty

import (
	"github.com/slotopol/server/game"
)

// reels lengths [35, 29, 35, 29, 35], total reshuffles 36057875
// RTP = 83.016(lined) + 4.2741(scatter) = 87.290058%
var Reels87 = game.Reels5x{
	{1, 5, 5, 5, 2, 2, 2, 2, 2, 6, 6, 6, 3, 3, 3, 3, 3, 7, 7, 7, 4, 4, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 6, 6, 6, 3, 3, 3, 7, 7, 7, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 2, 2, 6, 6, 6, 3, 3, 3, 3, 3, 7, 7, 7, 4, 4, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 6, 6, 6, 3, 3, 3, 7, 7, 7, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 2, 2, 6, 6, 6, 3, 3, 3, 3, 3, 7, 7, 7, 4, 4, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
}

// reels lengths [35, 32, 29, 32, 35], total reshuffles 36377600
// RTP = 83.781(lined) + 4.2399(scatter) = 88.020691%
var Reels88 = game.Reels5x{
	{1, 5, 5, 5, 2, 2, 2, 2, 2, 6, 6, 6, 3, 3, 3, 3, 3, 7, 7, 7, 4, 4, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 2, 6, 6, 6, 3, 3, 3, 3, 7, 7, 7, 4, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 6, 6, 6, 3, 3, 3, 7, 7, 7, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 2, 6, 6, 6, 3, 3, 3, 3, 7, 7, 7, 4, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 2, 2, 6, 6, 6, 3, 3, 3, 3, 3, 7, 7, 7, 4, 4, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
}

// reels lengths [32, 32, 32, 32, 32], total reshuffles 33554432
// RTP = 84.911(lined) + 4.4458(scatter) = 89.357108%
var Reels89 = game.Reels5x{
	{1, 5, 5, 5, 2, 2, 2, 2, 6, 6, 6, 3, 3, 3, 3, 7, 7, 7, 4, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 2, 6, 6, 6, 3, 3, 3, 3, 7, 7, 7, 4, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 2, 6, 6, 6, 3, 3, 3, 3, 7, 7, 7, 4, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 2, 6, 6, 6, 3, 3, 3, 3, 7, 7, 7, 4, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 2, 6, 6, 6, 3, 3, 3, 3, 7, 7, 7, 4, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
}

// reels lengths [29, 35, 29, 35, 29], total reshuffles 29876525
// RTP = 85.332(lined) + 4.8114(scatter) = 90.142947%
var Reels90 = game.Reels5x{
	{1, 5, 5, 5, 2, 2, 2, 6, 6, 6, 3, 3, 3, 7, 7, 7, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 2, 2, 6, 6, 6, 3, 3, 3, 3, 3, 7, 7, 7, 4, 4, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 6, 6, 6, 3, 3, 3, 7, 7, 7, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 2, 2, 6, 6, 6, 3, 3, 3, 3, 3, 7, 7, 7, 4, 4, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 6, 6, 6, 3, 3, 3, 7, 7, 7, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
}

// reels lengths [31, 31, 31, 31, 31], total reshuffles 28629151
// RTP = 86.309(lined) + 4.9135(scatter) = 91.222125%
var Reels91 = game.Reels5x{
	{1, 5, 5, 5, 2, 2, 2, 6, 6, 6, 3, 3, 3, 3, 7, 7, 7, 4, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 6, 6, 6, 3, 3, 3, 3, 7, 7, 7, 4, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 6, 6, 6, 3, 3, 3, 3, 7, 7, 7, 4, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 6, 6, 6, 3, 3, 3, 3, 7, 7, 7, 4, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 6, 6, 6, 3, 3, 3, 3, 7, 7, 7, 4, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
}

// reels lengths [29, 32, 32, 32, 29], total reshuffles 27557888
// RTP = 86.769(lined) + 5.0413(scatter) = 91.809957%
var Reels92 = game.Reels5x{
	{1, 5, 5, 5, 2, 2, 2, 6, 6, 6, 3, 3, 3, 7, 7, 7, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 2, 6, 6, 6, 3, 3, 3, 3, 7, 7, 7, 4, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 2, 6, 6, 6, 3, 3, 3, 3, 7, 7, 7, 4, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 2, 6, 6, 6, 3, 3, 3, 3, 7, 7, 7, 4, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 6, 6, 6, 3, 3, 3, 7, 7, 7, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
}

// reels lengths [31, 31, 31, 31, 31], total reshuffles 28629151
// RTP = 84.533(lined) + 8.0144(scatter) = 92.546929%
var Reels93 = game.Reels5x{
	{1, 5, 5, 5, 2, 2, 2, 6, 6, 6, 3, 3, 3, 3, 7, 7, 7, 4, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 6, 6, 6, 3, 3, 3, 3, 7, 7, 7, 4, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 8, 6, 6, 6, 3, 3, 3, 3, 7, 7, 7, 4, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 6, 6, 6, 3, 3, 3, 3, 7, 7, 7, 4, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 6, 6, 6, 3, 3, 3, 3, 7, 7, 7, 4, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
}

// reels lengths [29, 32, 29, 32, 29], total reshuffles 24974336
// RTP = 89.099(lined) + 5.365(scatter) = 94.464473%
var Reels94 = game.Reels5x{
	{1, 5, 5, 5, 2, 2, 2, 6, 6, 6, 3, 3, 3, 7, 7, 7, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 2, 6, 6, 6, 3, 3, 3, 3, 7, 7, 7, 4, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 6, 6, 6, 3, 3, 3, 7, 7, 7, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 2, 6, 6, 6, 3, 3, 3, 3, 7, 7, 7, 4, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 6, 6, 6, 3, 3, 3, 7, 7, 7, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
}

// reels lengths [34, 34, 34, 34, 34], total reshuffles 45435424
// RTP = 85.342(lined) + 9.5496(scatter) = 94.891642%
var Reels95 = game.Reels5x{
	{1, 5, 5, 5, 2, 2, 2, 6, 6, 6, 3, 3, 3, 3, 7, 7, 7, 4, 4, 4, 4, 5, 5, 5, 5, 6, 6, 6, 6, 7, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 6, 6, 6, 3, 3, 3, 8, 7, 7, 7, 4, 4, 4, 4, 5, 5, 5, 5, 6, 6, 6, 6, 7, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 6, 6, 6, 3, 3, 3, 3, 7, 7, 7, 4, 4, 4, 4, 5, 5, 5, 5, 6, 6, 6, 6, 7, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 6, 6, 6, 3, 3, 3, 8, 7, 7, 7, 4, 4, 4, 4, 5, 5, 5, 5, 6, 6, 6, 6, 7, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 6, 6, 6, 3, 3, 3, 3, 7, 7, 7, 4, 4, 4, 4, 5, 5, 5, 5, 6, 6, 6, 6, 7, 7, 7, 7, 8},
}

// reels lengths [29, 29, 32, 29, 29], total reshuffles 22632992
// RTP = 90.613(lined) + 5.7071(scatter) = 96.320186%
var Reels96 = game.Reels5x{
	{1, 5, 5, 5, 2, 2, 2, 6, 6, 6, 3, 3, 3, 7, 7, 7, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 6, 6, 6, 3, 3, 3, 7, 7, 7, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 2, 6, 6, 6, 3, 3, 3, 3, 7, 7, 7, 4, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 6, 6, 6, 3, 3, 3, 7, 7, 7, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 6, 6, 6, 3, 3, 3, 7, 7, 7, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
}

// reels lengths [30, 30, 30, 30, 30], total reshuffles 24300000
// RTP = 89.87(lined) + 8.9(scatter) = 98.770288%
var Reels99 = game.Reels5x{
	{1, 5, 5, 5, 2, 2, 2, 2, 6, 6, 6, 3, 3, 3, 7, 7, 7, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 2, 6, 6, 6, 3, 3, 3, 7, 7, 7, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 8, 6, 6, 6, 3, 3, 3, 7, 7, 7, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 2, 6, 6, 6, 3, 3, 3, 7, 7, 7, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 2, 6, 6, 6, 3, 3, 3, 7, 7, 7, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
}

// reels lengths [29, 29, 29, 29, 29], total reshuffles 20511149
// RTP = 94.347(lined) + 6.0684(scatter) = 100.415437%
var Reels100 = game.Reels5x{
	{1, 5, 5, 5, 2, 2, 2, 6, 6, 6, 3, 3, 3, 7, 7, 7, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 6, 6, 6, 3, 3, 3, 7, 7, 7, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 6, 6, 6, 3, 3, 3, 7, 7, 7, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 6, 6, 6, 3, 3, 3, 7, 7, 7, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 6, 6, 6, 3, 3, 3, 7, 7, 7, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
}

// reels lengths [30, 30, 30, 30, 30], total reshuffles 24300000
// RTP = 105.35(lined) + 5.45(scatter) = 110.800453%
var Reels111 = game.Reels5x{
	{1, 5, 5, 5, 2, 2, 2, 2, 6, 6, 6, 3, 3, 3, 7, 7, 7, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 2, 6, 6, 6, 3, 3, 3, 7, 7, 7, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 6, 6, 6, 3, 3, 3, 7, 7, 7, 4, 4, 4, 1, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 2, 6, 6, 6, 3, 3, 3, 7, 7, 7, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 2, 6, 6, 6, 3, 3, 3, 7, 7, 7, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
}

// reels lengths [30, 30, 30, 30, 30], total reshuffles 24300000
// RTP = 100.88(lined) + 22.6(scatter) = 123.483457%
var Reels123 = game.Reels5x{
	{1, 5, 5, 5, 2, 2, 2, 6, 6, 6, 3, 3, 3, 7, 7, 7, 4, 4, 4, 1, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 6, 6, 6, 3, 3, 3, 7, 7, 7, 4, 4, 4, 8, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 6, 6, 6, 3, 3, 3, 7, 7, 7, 4, 4, 4, 8, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 6, 6, 6, 3, 3, 3, 7, 7, 7, 4, 4, 4, 8, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
	{1, 5, 5, 5, 2, 2, 2, 6, 6, 6, 3, 3, 3, 7, 7, 7, 4, 4, 4, 1, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8},
}

// Map with available reels.
var ReelsMap = map[string]*game.Reels5x{
	"87":  &Reels87,
	"88":  &Reels88,
	"89":  &Reels89,
	"90":  &Reels90,
	"91":  &Reels91,
	"92":  &Reels92,
	"93":  &Reels93,
	"94":  &Reels94,
	"95":  &Reels95,
	"96":  &Reels96,
	"99":  &Reels99,
	"100": &Reels100,
	"111": &Reels111,
	"123": &Reels123,
}

// Lined payment.
var LinePay = [8][5]float64{
	{0, 0, 40, 400, 1000}, // seven
	{0, 0, 20, 80, 400},   // bell
	{0, 0, 20, 40, 200},   // melon
	{0, 0, 20, 40, 200},   // plum
	{0, 0, 10, 20, 100},   // orange
	{0, 0, 10, 20, 100},   // lemon
	{0, 0, 10, 20, 100},   // cherry
	{0, 0, 0, 0, 0},       // star
}

// Scatters payment.
var ScatPay = [5]float64{0, 0, 5, 20, 500} // star

const (
	jid = 1 // jackpot ID
)

// Jackpot win combinations.
var Jackpot = [8][5]int{
	{0, 0, 0, 0, 0}, // seven
	{0, 0, 0, 0, 0}, // bell
	{0, 0, 0, 0, 0}, // melon
	{0, 0, 0, 0, 0}, // plum
	{0, 0, 0, 0, 0}, // orange
	{0, 0, 0, 0, 0}, // lemon
	{0, 0, 0, 0, 0}, // cherry
	{0, 0, 0, 0, 0}, // star
}

type Game struct {
	game.Slot5x3 `yaml:",inline"`
}

func NewGame(rd string) *Game {
	return &Game{
		Slot5x3: game.Slot5x3{
			RD:  rd,
			SBL: game.MakeBitNum(20),
			Bet: 1,
		},
	}
}

const wild, scat = 1, 8

var bl = game.BetLinesNvm20

func (g *Game) Scanner(screen game.Screen, wins *game.Wins) {
	g.ScanLined(screen, wins)
	g.ScanScatters(screen, wins)
}

// Lined symbols calculation.
func (g *Game) ScanLined(screen game.Screen, wins *game.Wins) {
	for li := g.SBL.Next(0); li != 0; li = g.SBL.Next(li) {
		var line = bl.Line(li)

		var numw, numl = 0, 5
		var syml game.Sym
		for x := 1; x <= 5; x++ {
			var sx = screen.Pos(x, line)
			if sx == wild {
				if syml == 0 {
					numw = x
				}
			} else if syml == 0 && sx != scat {
				syml = sx
			} else if sx != syml {
				numl = x - 1
				break
			}
		}

		var payw, payl float64
		if numw > 0 {
			payw = LinePay[wild-1][numw-1]
		}
		if numl > 0 && syml > 0 {
			payl = LinePay[syml-1][numl-1]
		}
		if payl > payw {
			*wins = append(*wins, game.WinItem{
				Pay:  g.Bet * payl,
				Mult: 1,
				Sym:  syml,
				Num:  numl,
				Line: li,
				XY:   line.CopyL(numl),
			})
		} else if payw > 0 {
			*wins = append(*wins, game.WinItem{
				Pay:  g.Bet * payw,
				Mult: 1,
				Sym:  wild,
				Num:  numw,
				Line: li,
				XY:   line.CopyL(numw),
				Jack: Jackpot[wild-1][numw-1],
			})
		}
	}
}

// Scatters calculation.
func (g *Game) ScanScatters(screen game.Screen, wins *game.Wins) {
	if count := screen.ScatNum(scat); count >= 2 {
		var pay = ScatPay[count-1]
		*wins = append(*wins, game.WinItem{
			Pay:  g.Bet * float64(g.SBL.Num()) * pay,
			Mult: 1,
			Sym:  scat,
			Num:  count,
			XY:   screen.ScatPos(scat),
		})
	}
}

func (g *Game) Spin(screen game.Screen) {
	screen.Spin(ReelsMap[g.RD])
}

func (g *Game) SetLines(sbl game.Bitset) error {
	var mask game.Bitset = (1<<len(bl) - 1) << 1
	if sbl == 0 {
		return game.ErrNoLineset
	}
	if sbl&^mask != 0 {
		return game.ErrLinesetOut
	}
	if g.FreeSpins() > 0 {
		return game.ErrNoFeature
	}
	g.SBL = sbl
	return nil
}

func (g *Game) SetReels(rd string) error {
	if _, ok := ReelsMap[rd]; !ok {
		return game.ErrNoReels
	}
	g.RD = rd
	return nil
}

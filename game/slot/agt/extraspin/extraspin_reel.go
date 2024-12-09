package extraspin

import (
	"github.com/slotopol/server/game/slot"
)

// reels lengths [52, 87, 87, 87, 52], total reshuffles 1780592112
// symbols: 65.248(lined) + 0(scatter) = 65.247788%
// free spins 184199184, q = 0.10345, sq = 1/(1-q) = 1.115385
// free games frequency: 1/10.008
// RTP = 65.248(sym) + 0.10345*218.33(fg) = 87.833560%
var Reels878 = slot.Reels5x{
	{9, 9, 9, 9, 9, 5, 9, 9, 9, 9, 9, 5, 8, 8, 8, 8, 8, 4, 7, 7, 7, 5, 6, 6, 6, 8, 1, 4, 5, 6, 6, 4, 1, 6, 6, 6, 7, 7, 7, 4, 3, 8, 8, 8, 8, 8, 9, 5, 3, 7, 7, 3},
	{9, 9, 9, 9, 9, 3, 8, 8, 8, 8, 8, 4, 5, 4, 3, 4, 5, 4, 3, 7, 7, 7, 7, 7, 7, 8, 8, 8, 8, 8, 3, 9, 9, 9, 9, 9, 5, 4, 1, 8, 8, 8, 8, 8, 4, 6, 6, 6, 6, 6, 6, 6, 5, 1, 5, 7, 7, 7, 7, 7, 7, 7, 5, 1, 5, 4, 9, 9, 9, 9, 9, 4, 5, 1, 6, 6, 6, 6, 6, 6, 4, 3, 5, 4, 5, 2, 3},
	{9, 9, 9, 9, 9, 4, 1, 4, 3, 4, 9, 9, 9, 9, 9, 5, 3, 8, 8, 8, 8, 8, 4, 8, 8, 8, 8, 8, 5, 1, 3, 5, 3, 5, 6, 6, 6, 6, 6, 6, 6, 5, 4, 5, 4, 9, 9, 9, 9, 9, 4, 5, 1, 3, 7, 7, 7, 7, 7, 7, 7, 2, 6, 6, 6, 6, 6, 6, 5, 3, 5, 8, 8, 8, 8, 8, 5, 4, 7, 7, 7, 7, 7, 7, 4, 1, 4},
	{4, 7, 7, 7, 7, 7, 7, 9, 9, 9, 9, 9, 4, 1, 4, 3, 5, 3, 5, 3, 8, 8, 8, 8, 8, 4, 8, 8, 8, 8, 8, 4, 1, 9, 9, 9, 9, 9, 5, 4, 5, 1, 3, 8, 8, 8, 8, 8, 1, 6, 6, 6, 6, 6, 6, 5, 9, 9, 9, 9, 9, 5, 4, 3, 7, 7, 7, 7, 7, 7, 7, 6, 6, 6, 6, 6, 6, 6, 5, 4, 5, 4, 5, 3, 5, 2, 4},
	{5, 8, 8, 8, 8, 8, 4, 3, 5, 7, 7, 9, 5, 6, 6, 6, 5, 4, 7, 7, 7, 6, 6, 6, 5, 4, 9, 9, 9, 9, 9, 4, 3, 6, 6, 8, 8, 8, 8, 8, 1, 9, 9, 9, 9, 9, 1, 7, 7, 7, 8, 3},
}

// reels lengths [51, 87, 87, 87, 51], total reshuffles 1712766303
// symbols: 65.667(lined) + 0(scatter) = 65.666609%
// free spins 177182721, q = 0.10345, sq = 1/(1-q) = 1.115385
// free games frequency: 1/10.008
// RTP = 65.667(sym) + 0.10345*219.73(fg) = 88.397358%
var Reels883 = slot.Reels5x{
	{6, 6, 3, 6, 6, 6, 4, 5, 4, 8, 8, 8, 6, 6, 6, 8, 8, 8, 7, 7, 7, 1, 4, 5, 7, 7, 7, 3, 1, 4, 8, 5, 3, 9, 9, 8, 8, 8, 9, 9, 9, 5, 9, 9, 9, 7, 7, 7, 9, 9, 9},
	{9, 9, 9, 9, 9, 3, 8, 8, 8, 8, 8, 4, 5, 4, 3, 4, 5, 4, 3, 7, 7, 7, 7, 7, 7, 8, 8, 8, 8, 8, 3, 9, 9, 9, 9, 9, 5, 4, 1, 8, 8, 8, 8, 8, 4, 6, 6, 6, 6, 6, 6, 6, 5, 1, 5, 7, 7, 7, 7, 7, 7, 7, 5, 1, 5, 4, 9, 9, 9, 9, 9, 4, 5, 1, 6, 6, 6, 6, 6, 6, 4, 3, 5, 4, 5, 2, 3},
	{9, 9, 9, 9, 9, 4, 1, 4, 3, 4, 9, 9, 9, 9, 9, 5, 3, 8, 8, 8, 8, 8, 4, 8, 8, 8, 8, 8, 5, 1, 3, 5, 3, 5, 6, 6, 6, 6, 6, 6, 6, 5, 4, 5, 4, 9, 9, 9, 9, 9, 4, 5, 1, 3, 7, 7, 7, 7, 7, 7, 7, 2, 6, 6, 6, 6, 6, 6, 5, 3, 5, 8, 8, 8, 8, 8, 5, 4, 7, 7, 7, 7, 7, 7, 4, 1, 4},
	{4, 7, 7, 7, 7, 7, 7, 9, 9, 9, 9, 9, 4, 1, 4, 3, 5, 3, 5, 3, 8, 8, 8, 8, 8, 4, 8, 8, 8, 8, 8, 4, 1, 9, 9, 9, 9, 9, 5, 4, 5, 1, 3, 8, 8, 8, 8, 8, 1, 6, 6, 6, 6, 6, 6, 5, 9, 9, 9, 9, 9, 5, 4, 3, 7, 7, 7, 7, 7, 7, 7, 6, 6, 6, 6, 6, 6, 6, 5, 4, 5, 4, 5, 3, 5, 2, 4},
	{9, 9, 9, 7, 7, 7, 8, 8, 8, 3, 4, 5, 9, 9, 9, 3, 8, 8, 8, 6, 6, 6, 5, 9, 9, 5, 1, 8, 8, 8, 4, 8, 4, 7, 7, 7, 6, 6, 7, 7, 7, 6, 6, 6, 5, 3, 4, 1, 9, 9, 9},
}

// reels lengths [49, 87, 87, 87, 49], total reshuffles 1581065703
// symbols: 66.397(lined) + 0(scatter) = 66.396664%
// free spins 163558521, q = 0.10345, sq = 1/(1-q) = 1.115385
// free games frequency: 1/10.008
// RTP = 66.397(sym) + 0.10345*222.17(fg) = 89.380124%
var Reels893 = slot.Reels5x{
	{9, 9, 9, 9, 9, 4, 5, 8, 8, 8, 8, 8, 4, 9, 9, 9, 9, 9, 5, 7, 7, 7, 3, 7, 7, 7, 8, 8, 8, 8, 8, 4, 6, 6, 1, 3, 7, 7, 6, 6, 6, 1, 6, 6, 6, 3, 5, 4, 5},
	{9, 9, 9, 9, 9, 3, 8, 8, 8, 8, 8, 4, 5, 4, 3, 4, 5, 4, 3, 7, 7, 7, 7, 7, 7, 8, 8, 8, 8, 8, 3, 9, 9, 9, 9, 9, 5, 4, 1, 8, 8, 8, 8, 8, 4, 6, 6, 6, 6, 6, 6, 6, 5, 1, 5, 7, 7, 7, 7, 7, 7, 7, 5, 1, 5, 4, 9, 9, 9, 9, 9, 4, 5, 1, 6, 6, 6, 6, 6, 6, 4, 3, 5, 4, 5, 2, 3},
	{9, 9, 9, 9, 9, 4, 1, 4, 3, 4, 9, 9, 9, 9, 9, 5, 3, 8, 8, 8, 8, 8, 4, 8, 8, 8, 8, 8, 5, 1, 3, 5, 3, 5, 6, 6, 6, 6, 6, 6, 6, 5, 4, 5, 4, 9, 9, 9, 9, 9, 4, 5, 1, 3, 7, 7, 7, 7, 7, 7, 7, 2, 6, 6, 6, 6, 6, 6, 5, 3, 5, 8, 8, 8, 8, 8, 5, 4, 7, 7, 7, 7, 7, 7, 4, 1, 4},
	{4, 7, 7, 7, 7, 7, 7, 9, 9, 9, 9, 9, 4, 1, 4, 3, 5, 3, 5, 3, 8, 8, 8, 8, 8, 4, 8, 8, 8, 8, 8, 4, 1, 9, 9, 9, 9, 9, 5, 4, 5, 1, 3, 8, 8, 8, 8, 8, 1, 6, 6, 6, 6, 6, 6, 5, 9, 9, 9, 9, 9, 5, 4, 3, 7, 7, 7, 7, 7, 7, 7, 6, 6, 6, 6, 6, 6, 6, 5, 4, 5, 4, 5, 3, 5, 2, 4},
	{9, 9, 9, 9, 9, 5, 9, 9, 9, 9, 9, 5, 4, 3, 1, 6, 6, 7, 7, 3, 4, 8, 8, 8, 8, 8, 5, 4, 6, 6, 6, 1, 3, 8, 8, 8, 8, 8, 4, 5, 7, 7, 7, 6, 6, 6, 7, 7, 7},
}

// reels lengths [47, 87, 87, 87, 47], total reshuffles 1454633127
// symbols: 66.81(lined) + 0(scatter) = 66.810226%
// free spins 150479289, q = 0.10345, sq = 1/(1-q) = 1.115385
// free games frequency: 1/10.008
// RTP = 66.81(sym) + 0.10345*223.56(fg) = 89.936843%
var Reels899 = slot.Reels5x{
	{7, 4, 6, 6, 6, 7, 7, 7, 6, 1, 9, 9, 9, 9, 9, 1, 4, 8, 8, 8, 8, 8, 5, 9, 9, 9, 9, 9, 6, 6, 6, 4, 3, 7, 7, 7, 3, 5, 4, 5, 3, 5, 8, 8, 8, 8, 8},
	{9, 9, 9, 9, 9, 3, 8, 8, 8, 8, 8, 4, 5, 4, 3, 4, 5, 4, 3, 7, 7, 7, 7, 7, 7, 8, 8, 8, 8, 8, 3, 9, 9, 9, 9, 9, 5, 4, 1, 8, 8, 8, 8, 8, 4, 6, 6, 6, 6, 6, 6, 6, 5, 1, 5, 7, 7, 7, 7, 7, 7, 7, 5, 1, 5, 4, 9, 9, 9, 9, 9, 4, 5, 1, 6, 6, 6, 6, 6, 6, 4, 3, 5, 4, 5, 2, 3},
	{9, 9, 9, 9, 9, 4, 1, 4, 3, 4, 9, 9, 9, 9, 9, 5, 3, 8, 8, 8, 8, 8, 4, 8, 8, 8, 8, 8, 5, 1, 3, 5, 3, 5, 6, 6, 6, 6, 6, 6, 6, 5, 4, 5, 4, 9, 9, 9, 9, 9, 4, 5, 1, 3, 7, 7, 7, 7, 7, 7, 7, 2, 6, 6, 6, 6, 6, 6, 5, 3, 5, 8, 8, 8, 8, 8, 5, 4, 7, 7, 7, 7, 7, 7, 4, 1, 4},
	{4, 7, 7, 7, 7, 7, 7, 9, 9, 9, 9, 9, 4, 1, 4, 3, 5, 3, 5, 3, 8, 8, 8, 8, 8, 4, 8, 8, 8, 8, 8, 4, 1, 9, 9, 9, 9, 9, 5, 4, 5, 1, 3, 8, 8, 8, 8, 8, 1, 6, 6, 6, 6, 6, 6, 5, 9, 9, 9, 9, 9, 5, 4, 3, 7, 7, 7, 7, 7, 7, 7, 6, 6, 6, 6, 6, 6, 6, 5, 4, 5, 4, 5, 3, 5, 2, 4},
	{8, 8, 8, 8, 8, 4, 3, 1, 4, 8, 8, 8, 8, 8, 9, 9, 9, 9, 9, 7, 7, 7, 5, 3, 4, 6, 6, 6, 5, 9, 9, 9, 9, 9, 1, 6, 5, 6, 6, 6, 7, 5, 4, 3, 7, 7, 7},
}

// reels lengths [47, 87, 87, 87, 47], total reshuffles 1454633127
// symbols: 67.675(lined) + 0(scatter) = 67.675446%
// free spins 150479289, q = 0.10345, sq = 1/(1-q) = 1.115385
// free games frequency: 1/10.008
// RTP = 67.675(sym) + 0.10345*226.45(fg) = 91.101562%
var Reels911 = slot.Reels5x{
	{9, 9, 9, 7, 7, 5, 1, 5, 8, 8, 8, 7, 7, 7, 5, 7, 7, 7, 4, 3, 1, 9, 9, 9, 6, 6, 6, 5, 6, 6, 6, 3, 4, 8, 8, 8, 9, 9, 9, 6, 6, 4, 8, 8, 8, 4, 3},
	{9, 9, 9, 9, 9, 3, 8, 8, 8, 8, 8, 4, 5, 4, 3, 4, 5, 4, 3, 7, 7, 7, 7, 7, 7, 8, 8, 8, 8, 8, 3, 9, 9, 9, 9, 9, 5, 4, 1, 8, 8, 8, 8, 8, 4, 6, 6, 6, 6, 6, 6, 6, 5, 1, 5, 7, 7, 7, 7, 7, 7, 7, 5, 1, 5, 4, 9, 9, 9, 9, 9, 4, 5, 1, 6, 6, 6, 6, 6, 6, 4, 3, 5, 4, 5, 2, 3},
	{9, 9, 9, 9, 9, 4, 1, 4, 3, 4, 9, 9, 9, 9, 9, 5, 3, 8, 8, 8, 8, 8, 4, 8, 8, 8, 8, 8, 5, 1, 3, 5, 3, 5, 6, 6, 6, 6, 6, 6, 6, 5, 4, 5, 4, 9, 9, 9, 9, 9, 4, 5, 1, 3, 7, 7, 7, 7, 7, 7, 7, 2, 6, 6, 6, 6, 6, 6, 5, 3, 5, 8, 8, 8, 8, 8, 5, 4, 7, 7, 7, 7, 7, 7, 4, 1, 4},
	{4, 7, 7, 7, 7, 7, 7, 9, 9, 9, 9, 9, 4, 1, 4, 3, 5, 3, 5, 3, 8, 8, 8, 8, 8, 4, 8, 8, 8, 8, 8, 4, 1, 9, 9, 9, 9, 9, 5, 4, 5, 1, 3, 8, 8, 8, 8, 8, 1, 6, 6, 6, 6, 6, 6, 5, 9, 9, 9, 9, 9, 5, 4, 3, 7, 7, 7, 7, 7, 7, 7, 6, 6, 6, 6, 6, 6, 6, 5, 4, 5, 4, 5, 3, 5, 2, 4},
	{5, 8, 8, 8, 5, 4, 6, 6, 6, 4, 9, 9, 9, 3, 9, 9, 9, 5, 7, 7, 7, 1, 7, 7, 9, 9, 9, 4, 5, 8, 8, 8, 1, 7, 7, 7, 6, 6, 6, 8, 8, 8, 3, 4, 3, 6, 6},
}

// reels lengths [46, 87, 87, 87, 46], total reshuffles 1393392348
// symbols: 67.902(lined) + 0(scatter) = 67.902417%
// free spins 144144036, q = 0.10345, sq = 1/(1-q) = 1.115385
// free games frequency: 1/10.008
// RTP = 67.902(sym) + 0.10345*227.21(fg) = 91.407100%
var Reels914 = slot.Reels5x{
	{4, 5, 6, 6, 6, 1, 4, 9, 9, 9, 6, 5, 1, 7, 7, 3, 8, 8, 8, 7, 7, 7, 8, 8, 8, 5, 3, 6, 6, 6, 5, 4, 9, 9, 9, 7, 7, 7, 9, 9, 9, 4, 8, 8, 8, 3},
	{9, 9, 9, 9, 9, 3, 8, 8, 8, 8, 8, 4, 5, 4, 3, 4, 5, 4, 3, 7, 7, 7, 7, 7, 7, 8, 8, 8, 8, 8, 3, 9, 9, 9, 9, 9, 5, 4, 1, 8, 8, 8, 8, 8, 4, 6, 6, 6, 6, 6, 6, 6, 5, 1, 5, 7, 7, 7, 7, 7, 7, 7, 5, 1, 5, 4, 9, 9, 9, 9, 9, 4, 5, 1, 6, 6, 6, 6, 6, 6, 4, 3, 5, 4, 5, 2, 3},
	{9, 9, 9, 9, 9, 4, 1, 4, 3, 4, 9, 9, 9, 9, 9, 5, 3, 8, 8, 8, 8, 8, 4, 8, 8, 8, 8, 8, 5, 1, 3, 5, 3, 5, 6, 6, 6, 6, 6, 6, 6, 5, 4, 5, 4, 9, 9, 9, 9, 9, 4, 5, 1, 3, 7, 7, 7, 7, 7, 7, 7, 2, 6, 6, 6, 6, 6, 6, 5, 3, 5, 8, 8, 8, 8, 8, 5, 4, 7, 7, 7, 7, 7, 7, 4, 1, 4},
	{4, 7, 7, 7, 7, 7, 7, 9, 9, 9, 9, 9, 4, 1, 4, 3, 5, 3, 5, 3, 8, 8, 8, 8, 8, 4, 8, 8, 8, 8, 8, 4, 1, 9, 9, 9, 9, 9, 5, 4, 5, 1, 3, 8, 8, 8, 8, 8, 1, 6, 6, 6, 6, 6, 6, 5, 9, 9, 9, 9, 9, 5, 4, 3, 7, 7, 7, 7, 7, 7, 7, 6, 6, 6, 6, 6, 6, 6, 5, 4, 5, 4, 5, 3, 5, 2, 4},
	{4, 5, 4, 6, 3, 4, 5, 4, 5, 8, 8, 8, 9, 9, 9, 1, 3, 8, 8, 8, 7, 7, 7, 8, 8, 8, 5, 6, 6, 6, 3, 6, 6, 6, 9, 9, 9, 7, 7, 7, 1, 7, 7, 9, 9, 9},
}

// reels lengths [44, 87, 87, 87, 44], total reshuffles 1274861808
// symbols: 68.413(lined) + 0(scatter) = 68.413043%
// free spins 131882256, q = 0.10345, sq = 1/(1-q) = 1.115385
// free games frequency: 1/10.008
// RTP = 68.413(sym) + 0.10345*228.92(fg) = 92.094481%
var Reels920 = slot.Reels5x{
	{9, 9, 9, 8, 8, 8, 4, 5, 3, 1, 7, 7, 7, 3, 4, 8, 8, 8, 3, 7, 6, 6, 6, 8, 8, 8, 4, 5, 9, 9, 9, 5, 9, 9, 9, 7, 7, 7, 5, 1, 6, 6, 6, 4},
	{9, 9, 9, 9, 9, 3, 8, 8, 8, 8, 8, 4, 5, 4, 3, 4, 5, 4, 3, 7, 7, 7, 7, 7, 7, 8, 8, 8, 8, 8, 3, 9, 9, 9, 9, 9, 5, 4, 1, 8, 8, 8, 8, 8, 4, 6, 6, 6, 6, 6, 6, 6, 5, 1, 5, 7, 7, 7, 7, 7, 7, 7, 5, 1, 5, 4, 9, 9, 9, 9, 9, 4, 5, 1, 6, 6, 6, 6, 6, 6, 4, 3, 5, 4, 5, 2, 3},
	{9, 9, 9, 9, 9, 4, 1, 4, 3, 4, 9, 9, 9, 9, 9, 5, 3, 8, 8, 8, 8, 8, 4, 8, 8, 8, 8, 8, 5, 1, 3, 5, 3, 5, 6, 6, 6, 6, 6, 6, 6, 5, 4, 5, 4, 9, 9, 9, 9, 9, 4, 5, 1, 3, 7, 7, 7, 7, 7, 7, 7, 2, 6, 6, 6, 6, 6, 6, 5, 3, 5, 8, 8, 8, 8, 8, 5, 4, 7, 7, 7, 7, 7, 7, 4, 1, 4},
	{4, 7, 7, 7, 7, 7, 7, 9, 9, 9, 9, 9, 4, 1, 4, 3, 5, 3, 5, 3, 8, 8, 8, 8, 8, 4, 8, 8, 8, 8, 8, 4, 1, 9, 9, 9, 9, 9, 5, 4, 5, 1, 3, 8, 8, 8, 8, 8, 1, 6, 6, 6, 6, 6, 6, 5, 9, 9, 9, 9, 9, 5, 4, 3, 7, 7, 7, 7, 7, 7, 7, 6, 6, 6, 6, 6, 6, 6, 5, 4, 5, 4, 5, 3, 5, 2, 4},
	{6, 6, 6, 9, 9, 9, 3, 6, 6, 6, 3, 8, 8, 8, 1, 4, 7, 7, 7, 8, 8, 8, 7, 5, 4, 5, 4, 3, 5, 1, 8, 8, 8, 5, 4, 9, 9, 9, 7, 7, 7, 9, 9, 9},
}

// reels lengths [43, 87, 87, 87, 43], total reshuffles 1217572047
// symbols: 68.701(lined) + 0(scatter) = 68.700560%
// free spins 125955729, q = 0.10345, sq = 1/(1-q) = 1.115385
// free games frequency: 1/10.008
// RTP = 68.701(sym) + 0.10345*229.88(fg) = 92.481524%
var Reels924 = slot.Reels5x{
	{4, 6, 6, 6, 5, 4, 3, 4, 9, 9, 9, 6, 6, 6, 8, 8, 8, 9, 9, 9, 5, 3, 8, 8, 8, 3, 8, 8, 8, 7, 7, 7, 1, 7, 7, 7, 5, 1, 4, 5, 9, 9, 9},
	{9, 9, 9, 9, 9, 3, 8, 8, 8, 8, 8, 4, 5, 4, 3, 4, 5, 4, 3, 7, 7, 7, 7, 7, 7, 8, 8, 8, 8, 8, 3, 9, 9, 9, 9, 9, 5, 4, 1, 8, 8, 8, 8, 8, 4, 6, 6, 6, 6, 6, 6, 6, 5, 1, 5, 7, 7, 7, 7, 7, 7, 7, 5, 1, 5, 4, 9, 9, 9, 9, 9, 4, 5, 1, 6, 6, 6, 6, 6, 6, 4, 3, 5, 4, 5, 2, 3},
	{9, 9, 9, 9, 9, 4, 1, 4, 3, 4, 9, 9, 9, 9, 9, 5, 3, 8, 8, 8, 8, 8, 4, 8, 8, 8, 8, 8, 5, 1, 3, 5, 3, 5, 6, 6, 6, 6, 6, 6, 6, 5, 4, 5, 4, 9, 9, 9, 9, 9, 4, 5, 1, 3, 7, 7, 7, 7, 7, 7, 7, 2, 6, 6, 6, 6, 6, 6, 5, 3, 5, 8, 8, 8, 8, 8, 5, 4, 7, 7, 7, 7, 7, 7, 4, 1, 4},
	{4, 7, 7, 7, 7, 7, 7, 9, 9, 9, 9, 9, 4, 1, 4, 3, 5, 3, 5, 3, 8, 8, 8, 8, 8, 4, 8, 8, 8, 8, 8, 4, 1, 9, 9, 9, 9, 9, 5, 4, 5, 1, 3, 8, 8, 8, 8, 8, 1, 6, 6, 6, 6, 6, 6, 5, 9, 9, 9, 9, 9, 5, 4, 3, 7, 7, 7, 7, 7, 7, 7, 6, 6, 6, 6, 6, 6, 6, 5, 4, 5, 4, 5, 3, 5, 2, 4},
	{5, 9, 9, 9, 4, 6, 6, 6, 3, 9, 9, 9, 3, 7, 7, 7, 1, 5, 4, 8, 8, 8, 9, 9, 9, 7, 7, 7, 4, 8, 8, 8, 5, 3, 5, 8, 8, 8, 6, 6, 6, 1, 4},
}

// reels lengths [43, 87, 87, 87, 43], total reshuffles 1217572047
// symbols: 69.162(lined) + 0(scatter) = 69.162021%
// free spins 125955729, q = 0.10345, sq = 1/(1-q) = 1.115385
// free games frequency: 1/10.008
// RTP = 69.162(sym) + 0.10345*231.43(fg) = 93.102720%
var Reels931 = slot.Reels5x{
	{7, 7, 7, 1, 9, 9, 9, 3, 4, 9, 9, 9, 3, 7, 7, 7, 4, 5, 8, 8, 8, 5, 4, 3, 7, 6, 6, 6, 4, 6, 6, 6, 5, 8, 8, 8, 5, 8, 8, 1, 9, 9, 9},
	{9, 9, 9, 9, 9, 3, 8, 8, 8, 8, 8, 4, 5, 4, 3, 4, 5, 4, 3, 7, 7, 7, 7, 7, 7, 8, 8, 8, 8, 8, 3, 9, 9, 9, 9, 9, 5, 4, 1, 8, 8, 8, 8, 8, 4, 6, 6, 6, 6, 6, 6, 6, 5, 1, 5, 7, 7, 7, 7, 7, 7, 7, 5, 1, 5, 4, 9, 9, 9, 9, 9, 4, 5, 1, 6, 6, 6, 6, 6, 6, 4, 3, 5, 4, 5, 2, 3},
	{9, 9, 9, 9, 9, 4, 1, 4, 3, 4, 9, 9, 9, 9, 9, 5, 3, 8, 8, 8, 8, 8, 4, 8, 8, 8, 8, 8, 5, 1, 3, 5, 3, 5, 6, 6, 6, 6, 6, 6, 6, 5, 4, 5, 4, 9, 9, 9, 9, 9, 4, 5, 1, 3, 7, 7, 7, 7, 7, 7, 7, 2, 6, 6, 6, 6, 6, 6, 5, 3, 5, 8, 8, 8, 8, 8, 5, 4, 7, 7, 7, 7, 7, 7, 4, 1, 4},
	{4, 7, 7, 7, 7, 7, 7, 9, 9, 9, 9, 9, 4, 1, 4, 3, 5, 3, 5, 3, 8, 8, 8, 8, 8, 4, 8, 8, 8, 8, 8, 4, 1, 9, 9, 9, 9, 9, 5, 4, 5, 1, 3, 8, 8, 8, 8, 8, 1, 6, 6, 6, 6, 6, 6, 5, 9, 9, 9, 9, 9, 5, 4, 3, 7, 7, 7, 7, 7, 7, 7, 6, 6, 6, 6, 6, 6, 6, 5, 4, 5, 4, 5, 3, 5, 2, 4},
	{5, 8, 8, 3, 8, 8, 8, 4, 5, 9, 9, 9, 5, 1, 7, 7, 7, 5, 8, 8, 8, 4, 6, 6, 6, 4, 6, 6, 6, 1, 7, 3, 7, 7, 7, 9, 9, 9, 4, 3, 9, 9, 9},
}

// reels lengths [43, 87, 87, 87, 43], total reshuffles 1217572047
// symbols: 69.419(lined) + 0(scatter) = 69.418708%
// free spins 125955729, q = 0.10345, sq = 1/(1-q) = 1.115385
// free games frequency: 1/10.008
// RTP = 69.419(sym) + 0.10345*232.29(fg) = 93.448261%
var Reels934 = slot.Reels5x{
	{5, 3, 7, 7, 7, 3, 1, 9, 9, 9, 3, 9, 9, 9, 4, 9, 9, 9, 5, 4, 8, 8, 8, 1, 5, 6, 6, 6, 8, 8, 4, 5, 6, 6, 6, 7, 7, 7, 4, 5, 8, 8, 8},
	{9, 9, 9, 9, 9, 3, 8, 8, 8, 8, 8, 4, 5, 4, 3, 4, 5, 4, 3, 7, 7, 7, 7, 7, 7, 8, 8, 8, 8, 8, 3, 9, 9, 9, 9, 9, 5, 4, 1, 8, 8, 8, 8, 8, 4, 6, 6, 6, 6, 6, 6, 6, 5, 1, 5, 7, 7, 7, 7, 7, 7, 7, 5, 1, 5, 4, 9, 9, 9, 9, 9, 4, 5, 1, 6, 6, 6, 6, 6, 6, 4, 3, 5, 4, 5, 2, 3},
	{9, 9, 9, 9, 9, 4, 1, 4, 3, 4, 9, 9, 9, 9, 9, 5, 3, 8, 8, 8, 8, 8, 4, 8, 8, 8, 8, 8, 5, 1, 3, 5, 3, 5, 6, 6, 6, 6, 6, 6, 6, 5, 4, 5, 4, 9, 9, 9, 9, 9, 4, 5, 1, 3, 7, 7, 7, 7, 7, 7, 7, 2, 6, 6, 6, 6, 6, 6, 5, 3, 5, 8, 8, 8, 8, 8, 5, 4, 7, 7, 7, 7, 7, 7, 4, 1, 4},
	{4, 7, 7, 7, 7, 7, 7, 9, 9, 9, 9, 9, 4, 1, 4, 3, 5, 3, 5, 3, 8, 8, 8, 8, 8, 4, 8, 8, 8, 8, 8, 4, 1, 9, 9, 9, 9, 9, 5, 4, 5, 1, 3, 8, 8, 8, 8, 8, 1, 6, 6, 6, 6, 6, 6, 5, 9, 9, 9, 9, 9, 5, 4, 3, 7, 7, 7, 7, 7, 7, 7, 6, 6, 6, 6, 6, 6, 6, 5, 4, 5, 4, 5, 3, 5, 2, 4},
	{8, 8, 8, 5, 8, 8, 4, 5, 4, 3, 4, 6, 6, 6, 8, 8, 8, 5, 7, 7, 7, 9, 9, 9, 3, 7, 7, 7, 3, 6, 6, 6, 4, 9, 9, 9, 5, 1, 5, 9, 9, 9, 1},
}

// reels lengths [43, 87, 87, 87, 43], total reshuffles 1217572047
// symbols: 69.88(lined) + 0(scatter) = 69.880169%
// free spins 125955729, q = 0.10345, sq = 1/(1-q) = 1.115385
// free games frequency: 1/10.008
// RTP = 69.88(sym) + 0.10345*233.83(fg) = 94.069458%
var Reels940 = slot.Reels5x{
	{5, 4, 9, 9, 5, 1, 6, 6, 6, 7, 7, 7, 9, 9, 9, 5, 9, 9, 9, 8, 8, 3, 8, 8, 8, 4, 8, 8, 8, 4, 3, 4, 3, 5, 6, 6, 6, 7, 1, 7, 7, 7, 5},
	{9, 9, 9, 9, 9, 3, 8, 8, 8, 8, 8, 4, 5, 4, 3, 4, 5, 4, 3, 7, 7, 7, 7, 7, 7, 8, 8, 8, 8, 8, 3, 9, 9, 9, 9, 9, 5, 4, 1, 8, 8, 8, 8, 8, 4, 6, 6, 6, 6, 6, 6, 6, 5, 1, 5, 7, 7, 7, 7, 7, 7, 7, 5, 1, 5, 4, 9, 9, 9, 9, 9, 4, 5, 1, 6, 6, 6, 6, 6, 6, 4, 3, 5, 4, 5, 2, 3},
	{9, 9, 9, 9, 9, 4, 1, 4, 3, 4, 9, 9, 9, 9, 9, 5, 3, 8, 8, 8, 8, 8, 4, 8, 8, 8, 8, 8, 5, 1, 3, 5, 3, 5, 6, 6, 6, 6, 6, 6, 6, 5, 4, 5, 4, 9, 9, 9, 9, 9, 4, 5, 1, 3, 7, 7, 7, 7, 7, 7, 7, 2, 6, 6, 6, 6, 6, 6, 5, 3, 5, 8, 8, 8, 8, 8, 5, 4, 7, 7, 7, 7, 7, 7, 4, 1, 4},
	{4, 7, 7, 7, 7, 7, 7, 9, 9, 9, 9, 9, 4, 1, 4, 3, 5, 3, 5, 3, 8, 8, 8, 8, 8, 4, 8, 8, 8, 8, 8, 4, 1, 9, 9, 9, 9, 9, 5, 4, 5, 1, 3, 8, 8, 8, 8, 8, 1, 6, 6, 6, 6, 6, 6, 5, 9, 9, 9, 9, 9, 5, 4, 3, 7, 7, 7, 7, 7, 7, 7, 6, 6, 6, 6, 6, 6, 6, 5, 4, 5, 4, 5, 3, 5, 2, 4},
	{5, 4, 6, 6, 6, 1, 7, 8, 8, 4, 1, 5, 8, 8, 8, 5, 9, 9, 9, 7, 7, 7, 9, 9, 5, 4, 3, 6, 6, 6, 7, 7, 7, 9, 9, 9, 4, 3, 5, 3, 8, 8, 8},
}

// reels lengths [42, 87, 87, 87, 42], total reshuffles 1161599292
// symbols: 70.199(lined) + 0(scatter) = 70.199402%
// free spins 120165444, q = 0.10345, sq = 1/(1-q) = 1.115385
// free games frequency: 1/10.008
// RTP = 70.199(sym) + 0.10345*234.9(fg) = 94.499195%
var Reels944 = slot.Reels5x{
	{8, 8, 8, 9, 9, 9, 7, 7, 7, 9, 9, 9, 4, 7, 7, 7, 4, 5, 8, 8, 8, 5, 4, 1, 5, 9, 9, 5, 6, 6, 6, 3, 6, 6, 6, 8, 8, 3, 4, 1, 5, 3},
	{9, 9, 9, 9, 9, 3, 8, 8, 8, 8, 8, 4, 5, 4, 3, 4, 5, 4, 3, 7, 7, 7, 7, 7, 7, 8, 8, 8, 8, 8, 3, 9, 9, 9, 9, 9, 5, 4, 1, 8, 8, 8, 8, 8, 4, 6, 6, 6, 6, 6, 6, 6, 5, 1, 5, 7, 7, 7, 7, 7, 7, 7, 5, 1, 5, 4, 9, 9, 9, 9, 9, 4, 5, 1, 6, 6, 6, 6, 6, 6, 4, 3, 5, 4, 5, 2, 3},
	{9, 9, 9, 9, 9, 4, 1, 4, 3, 4, 9, 9, 9, 9, 9, 5, 3, 8, 8, 8, 8, 8, 4, 8, 8, 8, 8, 8, 5, 1, 3, 5, 3, 5, 6, 6, 6, 6, 6, 6, 6, 5, 4, 5, 4, 9, 9, 9, 9, 9, 4, 5, 1, 3, 7, 7, 7, 7, 7, 7, 7, 2, 6, 6, 6, 6, 6, 6, 5, 3, 5, 8, 8, 8, 8, 8, 5, 4, 7, 7, 7, 7, 7, 7, 4, 1, 4},
	{4, 7, 7, 7, 7, 7, 7, 9, 9, 9, 9, 9, 4, 1, 4, 3, 5, 3, 5, 3, 8, 8, 8, 8, 8, 4, 8, 8, 8, 8, 8, 4, 1, 9, 9, 9, 9, 9, 5, 4, 5, 1, 3, 8, 8, 8, 8, 8, 1, 6, 6, 6, 6, 6, 6, 5, 9, 9, 9, 9, 9, 5, 4, 3, 7, 7, 7, 7, 7, 7, 7, 6, 6, 6, 6, 6, 6, 6, 5, 4, 5, 4, 5, 3, 5, 2, 4},
	{6, 6, 6, 5, 4, 9, 9, 9, 5, 7, 7, 7, 4, 3, 9, 9, 8, 8, 8, 3, 4, 1, 7, 7, 7, 5, 1, 4, 5, 3, 6, 6, 6, 9, 9, 9, 8, 8, 5, 8, 8, 8},
}

// reels lengths [43, 87, 87, 87, 43], total reshuffles 1217572047
// symbols: 70.387(lined) + 0(scatter) = 70.386696%
// free spins 125955729, q = 0.10345, sq = 1/(1-q) = 1.115385
// free games frequency: 1/10.008
// RTP = 70.387(sym) + 0.10345*235.52(fg) = 94.751321%
var Reels947 = slot.Reels5x{
	{7, 7, 7, 3, 5, 3, 6, 6, 6, 4, 9, 9, 8, 8, 8, 5, 4, 9, 9, 9, 6, 6, 6, 9, 9, 9, 8, 8, 8, 3, 4, 5, 8, 5, 7, 5, 6, 4, 1, 7, 7, 7, 1},
	{9, 9, 9, 9, 9, 3, 8, 8, 8, 8, 8, 4, 5, 4, 3, 4, 5, 4, 3, 7, 7, 7, 7, 7, 7, 8, 8, 8, 8, 8, 3, 9, 9, 9, 9, 9, 5, 4, 1, 8, 8, 8, 8, 8, 4, 6, 6, 6, 6, 6, 6, 6, 5, 1, 5, 7, 7, 7, 7, 7, 7, 7, 5, 1, 5, 4, 9, 9, 9, 9, 9, 4, 5, 1, 6, 6, 6, 6, 6, 6, 4, 3, 5, 4, 5, 2, 3},
	{9, 9, 9, 9, 9, 4, 1, 4, 3, 4, 9, 9, 9, 9, 9, 5, 3, 8, 8, 8, 8, 8, 4, 8, 8, 8, 8, 8, 5, 1, 3, 5, 3, 5, 6, 6, 6, 6, 6, 6, 6, 5, 4, 5, 4, 9, 9, 9, 9, 9, 4, 5, 1, 3, 7, 7, 7, 7, 7, 7, 7, 2, 6, 6, 6, 6, 6, 6, 5, 3, 5, 8, 8, 8, 8, 8, 5, 4, 7, 7, 7, 7, 7, 7, 4, 1, 4},
	{4, 7, 7, 7, 7, 7, 7, 9, 9, 9, 9, 9, 4, 1, 4, 3, 5, 3, 5, 3, 8, 8, 8, 8, 8, 4, 8, 8, 8, 8, 8, 4, 1, 9, 9, 9, 9, 9, 5, 4, 5, 1, 3, 8, 8, 8, 8, 8, 1, 6, 6, 6, 6, 6, 6, 5, 9, 9, 9, 9, 9, 5, 4, 3, 7, 7, 7, 7, 7, 7, 7, 6, 6, 6, 6, 6, 6, 6, 5, 4, 5, 4, 5, 3, 5, 2, 4},
	{3, 7, 3, 4, 1, 5, 3, 5, 8, 6, 5, 7, 7, 7, 4, 6, 6, 6, 4, 9, 9, 9, 8, 8, 8, 9, 9, 9, 1, 5, 7, 7, 7, 6, 6, 6, 5, 4, 9, 9, 8, 8, 8},
}

// reels lengths [43, 87, 87, 87, 43], total reshuffles 1217572047
// symbols: 70.643(lined) + 0(scatter) = 70.643383%
// free spins 125955729, q = 0.10345, sq = 1/(1-q) = 1.115385
// free games frequency: 1/10.008
// RTP = 70.643(sym) + 0.10345*236.38(fg) = 95.096862%
var Reels950 = slot.Reels5x{
	{9, 9, 5, 7, 4, 5, 6, 6, 6, 7, 7, 7, 6, 6, 6, 1, 8, 8, 8, 3, 4, 3, 9, 9, 9, 8, 1, 4, 7, 7, 7, 4, 5, 9, 9, 9, 3, 5, 8, 8, 8, 5, 4},
	{9, 9, 9, 9, 9, 3, 8, 8, 8, 8, 8, 4, 5, 4, 3, 4, 5, 4, 3, 7, 7, 7, 7, 7, 7, 8, 8, 8, 8, 8, 3, 9, 9, 9, 9, 9, 5, 4, 1, 8, 8, 8, 8, 8, 4, 6, 6, 6, 6, 6, 6, 6, 5, 1, 5, 7, 7, 7, 7, 7, 7, 7, 5, 1, 5, 4, 9, 9, 9, 9, 9, 4, 5, 1, 6, 6, 6, 6, 6, 6, 4, 3, 5, 4, 5, 2, 3},
	{9, 9, 9, 9, 9, 4, 1, 4, 3, 4, 9, 9, 9, 9, 9, 5, 3, 8, 8, 8, 8, 8, 4, 8, 8, 8, 8, 8, 5, 1, 3, 5, 3, 5, 6, 6, 6, 6, 6, 6, 6, 5, 4, 5, 4, 9, 9, 9, 9, 9, 4, 5, 1, 3, 7, 7, 7, 7, 7, 7, 7, 2, 6, 6, 6, 6, 6, 6, 5, 3, 5, 8, 8, 8, 8, 8, 5, 4, 7, 7, 7, 7, 7, 7, 4, 1, 4},
	{4, 7, 7, 7, 7, 7, 7, 9, 9, 9, 9, 9, 4, 1, 4, 3, 5, 3, 5, 3, 8, 8, 8, 8, 8, 4, 8, 8, 8, 8, 8, 4, 1, 9, 9, 9, 9, 9, 5, 4, 5, 1, 3, 8, 8, 8, 8, 8, 1, 6, 6, 6, 6, 6, 6, 5, 9, 9, 9, 9, 9, 5, 4, 3, 7, 7, 7, 7, 7, 7, 7, 6, 6, 6, 6, 6, 6, 6, 5, 4, 5, 4, 5, 3, 5, 2, 4},
	{5, 9, 9, 9, 5, 8, 8, 8, 6, 6, 6, 7, 7, 7, 4, 6, 6, 6, 4, 7, 7, 7, 3, 4, 3, 9, 9, 1, 5, 8, 5, 4, 5, 8, 8, 8, 4, 9, 9, 9, 1, 7, 3},
}

// reels lengths [41, 87, 87, 87, 41], total reshuffles 1106943543
// symbols: 71.331(lined) + 0(scatter) = 71.331040%
// free spins 114511401, q = 0.10345, sq = 1/(1-q) = 1.115385
// free games frequency: 1/10.008
// RTP = 71.331(sym) + 0.10345*238.68(fg) = 96.022554%
var Reels960 = slot.Reels5x{
	{5, 4, 9, 9, 9, 1, 3, 8, 8, 8, 4, 5, 6, 5, 9, 9, 9, 1, 7, 7, 7, 4, 6, 6, 6, 5, 4, 8, 8, 8, 9, 8, 6, 6, 6, 7, 3, 7, 7, 7, 3},
	{9, 9, 9, 9, 9, 3, 8, 8, 8, 8, 8, 4, 5, 4, 3, 4, 5, 4, 3, 7, 7, 7, 7, 7, 7, 8, 8, 8, 8, 8, 3, 9, 9, 9, 9, 9, 5, 4, 1, 8, 8, 8, 8, 8, 4, 6, 6, 6, 6, 6, 6, 6, 5, 1, 5, 7, 7, 7, 7, 7, 7, 7, 5, 1, 5, 4, 9, 9, 9, 9, 9, 4, 5, 1, 6, 6, 6, 6, 6, 6, 4, 3, 5, 4, 5, 2, 3},
	{9, 9, 9, 9, 9, 4, 1, 4, 3, 4, 9, 9, 9, 9, 9, 5, 3, 8, 8, 8, 8, 8, 4, 8, 8, 8, 8, 8, 5, 1, 3, 5, 3, 5, 6, 6, 6, 6, 6, 6, 6, 5, 4, 5, 4, 9, 9, 9, 9, 9, 4, 5, 1, 3, 7, 7, 7, 7, 7, 7, 7, 2, 6, 6, 6, 6, 6, 6, 5, 3, 5, 8, 8, 8, 8, 8, 5, 4, 7, 7, 7, 7, 7, 7, 4, 1, 4},
	{4, 7, 7, 7, 7, 7, 7, 9, 9, 9, 9, 9, 4, 1, 4, 3, 5, 3, 5, 3, 8, 8, 8, 8, 8, 4, 8, 8, 8, 8, 8, 4, 1, 9, 9, 9, 9, 9, 5, 4, 5, 1, 3, 8, 8, 8, 8, 8, 1, 6, 6, 6, 6, 6, 6, 5, 9, 9, 9, 9, 9, 5, 4, 3, 7, 7, 7, 7, 7, 7, 7, 6, 6, 6, 6, 6, 6, 6, 5, 4, 5, 4, 5, 3, 5, 2, 4},
	{9, 9, 9, 5, 7, 5, 9, 5, 3, 7, 7, 7, 5, 1, 3, 9, 9, 9, 4, 8, 8, 8, 6, 6, 6, 4, 6, 8, 8, 8, 4, 1, 4, 6, 6, 6, 3, 7, 7, 7, 8},
}

// reels lengths [40, 87, 87, 87, 40], total reshuffles 1053604800
// symbols: 71.684(lined) + 0(scatter) = 71.684168%
// free spins 108993600, q = 0.10345, sq = 1/(1-q) = 1.115385
// free games frequency: 1/10.008
// RTP = 71.684(sym) + 0.10345*239.87(fg) = 96.497919%
var Reels964 = slot.Reels5x{
	{9, 5, 8, 4, 6, 6, 6, 1, 3, 7, 5, 8, 8, 8, 5, 8, 8, 8, 6, 6, 6, 1, 9, 9, 9, 5, 4, 7, 7, 7, 3, 7, 7, 7, 4, 9, 9, 9, 3, 4},
	{9, 9, 9, 9, 9, 3, 8, 8, 8, 8, 8, 4, 5, 4, 3, 4, 5, 4, 3, 7, 7, 7, 7, 7, 7, 8, 8, 8, 8, 8, 3, 9, 9, 9, 9, 9, 5, 4, 1, 8, 8, 8, 8, 8, 4, 6, 6, 6, 6, 6, 6, 6, 5, 1, 5, 7, 7, 7, 7, 7, 7, 7, 5, 1, 5, 4, 9, 9, 9, 9, 9, 4, 5, 1, 6, 6, 6, 6, 6, 6, 4, 3, 5, 4, 5, 2, 3},
	{9, 9, 9, 9, 9, 4, 1, 4, 3, 4, 9, 9, 9, 9, 9, 5, 3, 8, 8, 8, 8, 8, 4, 8, 8, 8, 8, 8, 5, 1, 3, 5, 3, 5, 6, 6, 6, 6, 6, 6, 6, 5, 4, 5, 4, 9, 9, 9, 9, 9, 4, 5, 1, 3, 7, 7, 7, 7, 7, 7, 7, 2, 6, 6, 6, 6, 6, 6, 5, 3, 5, 8, 8, 8, 8, 8, 5, 4, 7, 7, 7, 7, 7, 7, 4, 1, 4},
	{4, 7, 7, 7, 7, 7, 7, 9, 9, 9, 9, 9, 4, 1, 4, 3, 5, 3, 5, 3, 8, 8, 8, 8, 8, 4, 8, 8, 8, 8, 8, 4, 1, 9, 9, 9, 9, 9, 5, 4, 5, 1, 3, 8, 8, 8, 8, 8, 1, 6, 6, 6, 6, 6, 6, 5, 9, 9, 9, 9, 9, 5, 4, 3, 7, 7, 7, 7, 7, 7, 7, 6, 6, 6, 6, 6, 6, 6, 5, 4, 5, 4, 5, 3, 5, 2, 4},
	{6, 6, 6, 3, 8, 3, 5, 7, 7, 7, 4, 9, 9, 9, 5, 9, 5, 8, 8, 8, 3, 4, 7, 7, 7, 4, 9, 9, 9, 1, 7, 6, 6, 6, 4, 1, 8, 8, 8, 5},
}

// reels lengths [39, 87, 87, 87, 39], total reshuffles 1001583063
// symbols: 72.115(lined) + 0(scatter) = 72.115219%
// free spins 103612041, q = 0.10345, sq = 1/(1-q) = 1.115385
// free games frequency: 1/10.008
// RTP = 72.115(sym) + 0.10345*241.31(fg) = 97.078180%
var Reels970 = slot.Reels5x{
	{8, 4, 6, 6, 6, 5, 7, 7, 7, 9, 9, 9, 5, 8, 8, 8, 7, 7, 7, 5, 9, 1, 9, 9, 9, 8, 8, 8, 5, 3, 4, 3, 4, 3, 4, 1, 6, 6, 6},
	{4, 5, 4, 3, 5, 3, 5, 3, 1, 3, 4, 5, 4, 3, 5, 9, 9, 9, 9, 9, 5, 8, 8, 8, 8, 8, 5, 9, 9, 9, 9, 9, 1, 7, 7, 7, 7, 7, 7, 7, 1, 4, 5, 6, 6, 6, 6, 6, 6, 6, 8, 8, 8, 8, 8, 4, 8, 8, 8, 8, 8, 2, 9, 9, 9, 9, 9, 7, 7, 7, 7, 7, 7, 7, 4, 1, 4, 6, 6, 6, 6, 6, 6, 6, 3, 4, 5},
	{5, 4, 5, 3, 5, 7, 7, 7, 7, 7, 7, 7, 3, 8, 8, 8, 8, 8, 1, 8, 8, 8, 8, 8, 5, 9, 9, 9, 9, 9, 4, 5, 4, 9, 9, 9, 9, 9, 4, 1, 4, 3, 4, 3, 2, 7, 7, 7, 7, 7, 7, 7, 5, 8, 8, 8, 8, 8, 3, 9, 9, 9, 9, 9, 5, 6, 6, 6, 6, 6, 6, 6, 1, 3, 4, 5, 4, 1, 6, 6, 6, 6, 6, 6, 6, 4, 5},
	{6, 6, 6, 6, 6, 6, 6, 2, 5, 4, 6, 6, 6, 6, 6, 6, 6, 4, 5, 4, 3, 5, 9, 9, 9, 9, 9, 1, 9, 9, 9, 9, 9, 4, 1, 7, 7, 7, 7, 7, 7, 7, 5, 9, 9, 9, 9, 9, 4, 1, 8, 8, 8, 8, 8, 4, 3, 5, 4, 8, 8, 8, 8, 8, 4, 8, 8, 8, 8, 8, 3, 5, 3, 5, 4, 3, 5, 3, 7, 7, 7, 7, 7, 7, 7, 1, 5},
	{6, 6, 6, 9, 3, 4, 1, 3, 8, 4, 7, 7, 7, 5, 4, 8, 8, 8, 7, 7, 7, 8, 8, 8, 5, 6, 6, 6, 5, 9, 9, 9, 5, 9, 9, 9, 4, 1, 3},
}

// reels lengths [38, 87, 87, 87, 38], total reshuffles 950878332
// symbols: 73.047(lined) + 0(scatter) = 73.047001%
// free spins 98366724, q = 0.10345, sq = 1/(1-q) = 1.115385
// free games frequency: 1/10.008
// RTP = 73.047(sym) + 0.10345*244.43(fg) = 98.332501%
var Reels983 = slot.Reels5x{
	{6, 6, 6, 5, 7, 7, 7, 1, 4, 3, 4, 9, 6, 6, 6, 5, 3, 5, 9, 9, 9, 8, 8, 8, 9, 9, 9, 1, 3, 8, 8, 8, 5, 4, 7, 7, 7, 4},
	{4, 5, 7, 7, 7, 7, 7, 7, 7, 1, 7, 7, 7, 7, 7, 7, 1, 4, 5, 4, 3, 8, 8, 8, 8, 8, 1, 8, 8, 8, 8, 8, 3, 4, 8, 8, 8, 8, 8, 5, 6, 6, 6, 6, 6, 6, 9, 9, 9, 9, 9, 3, 1, 5, 4, 5, 9, 9, 9, 9, 9, 5, 9, 9, 9, 9, 9, 5, 4, 6, 6, 6, 6, 6, 6, 6, 4, 5, 4, 5, 3, 4, 3, 5, 2, 3, 4},
	{8, 8, 8, 8, 8, 4, 5, 4, 6, 6, 6, 6, 6, 6, 5, 3, 5, 2, 3, 9, 9, 9, 9, 9, 3, 6, 6, 6, 6, 6, 6, 6, 5, 4, 1, 9, 9, 9, 9, 9, 1, 8, 8, 8, 8, 8, 5, 4, 7, 7, 7, 7, 7, 7, 1, 3, 5, 4, 5, 4, 5, 8, 8, 8, 8, 8, 4, 3, 4, 5, 4, 3, 7, 7, 7, 7, 7, 7, 7, 4, 1, 9, 9, 9, 9, 9, 5},
	{4, 3, 9, 9, 9, 9, 9, 3, 4, 3, 2, 9, 9, 9, 9, 9, 1, 5, 4, 5, 4, 5, 8, 8, 8, 8, 8, 5, 4, 1, 4, 5, 8, 8, 8, 8, 8, 3, 5, 6, 6, 6, 6, 6, 6, 6, 4, 6, 6, 6, 6, 6, 6, 8, 8, 8, 8, 8, 5, 3, 5, 7, 7, 7, 7, 7, 7, 7, 5, 4, 5, 3, 4, 1, 9, 9, 9, 9, 9, 4, 7, 7, 7, 7, 7, 7, 1},
	{4, 3, 8, 8, 8, 9, 9, 9, 5, 9, 6, 6, 6, 1, 3, 6, 6, 6, 9, 9, 9, 4, 5, 3, 5, 1, 8, 8, 8, 7, 7, 7, 4, 7, 7, 7, 5, 4},
}

// reels lengths [39, 87, 87, 87, 39], total reshuffles 1001583063
// symbols: 73.714(lined) + 0(scatter) = 73.714288%
// free spins 103612041, q = 0.10345, sq = 1/(1-q) = 1.115385
// free games frequency: 1/10.008
// RTP = 73.714(sym) + 0.10345*246.66(fg) = 99.230772%
var Reels992 = slot.Reels5x{
	{8, 4, 6, 6, 6, 5, 7, 7, 7, 9, 9, 9, 5, 8, 8, 8, 7, 7, 7, 5, 9, 1, 9, 9, 9, 8, 8, 8, 5, 3, 4, 3, 4, 3, 4, 1, 6, 6, 6},
	{4, 3, 8, 8, 8, 8, 8, 4, 5, 9, 9, 9, 9, 9, 5, 1, 4, 8, 8, 8, 8, 8, 3, 1, 5, 8, 8, 8, 8, 8, 5, 7, 7, 7, 7, 7, 7, 3, 4, 5, 4, 5, 1, 5, 3, 4, 6, 6, 6, 6, 6, 6, 4, 5, 4, 7, 7, 7, 7, 7, 7, 4, 3, 5, 9, 9, 9, 9, 9, 5, 4, 2, 3, 6, 6, 6, 6, 6, 6, 3, 1, 3, 9, 9, 9, 9, 9},
	{4, 6, 6, 6, 6, 6, 6, 2, 8, 8, 8, 8, 8, 3, 5, 4, 3, 4, 3, 4, 1, 9, 9, 9, 9, 9, 4, 3, 5, 4, 3, 1, 3, 5, 8, 8, 8, 8, 8, 4, 8, 8, 8, 8, 8, 5, 9, 9, 9, 9, 9, 7, 7, 7, 7, 7, 7, 5, 1, 5, 4, 5, 3, 5, 9, 9, 9, 9, 9, 4, 5, 4, 6, 6, 6, 6, 6, 6, 7, 7, 7, 7, 7, 7, 5, 1, 3},
	{8, 8, 8, 8, 8, 1, 3, 5, 4, 5, 7, 7, 7, 7, 7, 7, 3, 1, 3, 5, 4, 5, 4, 3, 4, 7, 7, 7, 7, 7, 7, 4, 5, 8, 8, 8, 8, 8, 5, 2, 3, 5, 4, 5, 6, 6, 6, 6, 6, 6, 1, 4, 6, 6, 6, 6, 6, 6, 4, 8, 8, 8, 8, 8, 4, 5, 9, 9, 9, 9, 9, 3, 5, 3, 9, 9, 9, 9, 9, 4, 3, 9, 9, 9, 9, 9, 1},
	{6, 6, 6, 9, 3, 4, 1, 3, 8, 4, 7, 7, 7, 5, 4, 8, 8, 8, 7, 7, 7, 8, 8, 8, 5, 6, 6, 6, 5, 9, 9, 9, 5, 9, 9, 9, 4, 1, 3},
}

// reels lengths [37, 87, 87, 87, 37], total reshuffles 901490607
// symbols: 74.101(lined) + 0(scatter) = 74.101316%
// free spins 93257649, q = 0.10345, sq = 1/(1-q) = 1.115385
// free games frequency: 1/10.008
// RTP = 74.101(sym) + 0.10345*247.95(fg) = 99.751771%
var Reels997 = slot.Reels5x{
	{9, 9, 9, 3, 8, 8, 8, 4, 3, 6, 6, 6, 1, 5, 4, 5, 9, 9, 9, 4, 6, 6, 6, 8, 8, 8, 7, 7, 7, 5, 1, 5, 3, 4, 7, 7, 7},
	{4, 5, 7, 7, 7, 7, 7, 7, 7, 1, 7, 7, 7, 7, 7, 7, 1, 4, 5, 4, 3, 8, 8, 8, 8, 8, 1, 8, 8, 8, 8, 8, 3, 4, 8, 8, 8, 8, 8, 5, 6, 6, 6, 6, 6, 6, 9, 9, 9, 9, 9, 3, 1, 5, 4, 5, 9, 9, 9, 9, 9, 5, 9, 9, 9, 9, 9, 5, 4, 6, 6, 6, 6, 6, 6, 6, 4, 5, 4, 5, 3, 4, 3, 5, 2, 3, 4},
	{8, 8, 8, 8, 8, 4, 5, 4, 6, 6, 6, 6, 6, 6, 5, 3, 5, 2, 3, 9, 9, 9, 9, 9, 3, 6, 6, 6, 6, 6, 6, 6, 5, 4, 1, 9, 9, 9, 9, 9, 1, 8, 8, 8, 8, 8, 5, 4, 7, 7, 7, 7, 7, 7, 1, 3, 5, 4, 5, 4, 5, 8, 8, 8, 8, 8, 4, 3, 4, 5, 4, 3, 7, 7, 7, 7, 7, 7, 7, 4, 1, 9, 9, 9, 9, 9, 5},
	{4, 3, 9, 9, 9, 9, 9, 3, 4, 3, 2, 9, 9, 9, 9, 9, 1, 5, 4, 5, 4, 5, 8, 8, 8, 8, 8, 5, 4, 1, 4, 5, 8, 8, 8, 8, 8, 3, 5, 6, 6, 6, 6, 6, 6, 6, 4, 6, 6, 6, 6, 6, 6, 8, 8, 8, 8, 8, 5, 3, 5, 7, 7, 7, 7, 7, 7, 7, 5, 4, 5, 3, 4, 1, 9, 9, 9, 9, 9, 4, 7, 7, 7, 7, 7, 7, 1},
	{9, 9, 9, 3, 4, 5, 6, 6, 6, 5, 8, 8, 8, 7, 7, 7, 8, 8, 8, 1, 4, 9, 9, 9, 6, 6, 6, 4, 3, 5, 4, 1, 5, 3, 7, 7, 7},
}

// reels lengths [35, 87, 87, 87, 35], total reshuffles 806666175
// symbols: 75.113(lined) + 0(scatter) = 75.112697%
// free spins 83448225, q = 0.10345, sq = 1/(1-q) = 1.115385
// free games frequency: 1/10.008
// RTP = 75.113(sym) + 0.10345*251.34(fg) = 101.113246%
var Reels101 = slot.Reels5x{
	{9, 9, 9, 1, 5, 6, 6, 6, 3, 7, 7, 7, 8, 8, 8, 5, 7, 7, 4, 8, 8, 8, 6, 6, 4, 9, 9, 9, 4, 3, 1, 5, 3, 5, 4},
	{4, 5, 7, 7, 7, 7, 7, 7, 7, 1, 7, 7, 7, 7, 7, 7, 1, 4, 5, 4, 3, 8, 8, 8, 8, 8, 1, 8, 8, 8, 8, 8, 3, 4, 8, 8, 8, 8, 8, 5, 6, 6, 6, 6, 6, 6, 9, 9, 9, 9, 9, 3, 1, 5, 4, 5, 9, 9, 9, 9, 9, 5, 9, 9, 9, 9, 9, 5, 4, 6, 6, 6, 6, 6, 6, 6, 4, 5, 4, 5, 3, 4, 3, 5, 2, 3, 4},
	{8, 8, 8, 8, 8, 4, 5, 4, 6, 6, 6, 6, 6, 6, 5, 3, 5, 2, 3, 9, 9, 9, 9, 9, 3, 6, 6, 6, 6, 6, 6, 6, 5, 4, 1, 9, 9, 9, 9, 9, 1, 8, 8, 8, 8, 8, 5, 4, 7, 7, 7, 7, 7, 7, 1, 3, 5, 4, 5, 4, 5, 8, 8, 8, 8, 8, 4, 3, 4, 5, 4, 3, 7, 7, 7, 7, 7, 7, 7, 4, 1, 9, 9, 9, 9, 9, 5},
	{4, 3, 9, 9, 9, 9, 9, 3, 4, 3, 2, 9, 9, 9, 9, 9, 1, 5, 4, 5, 4, 5, 8, 8, 8, 8, 8, 5, 4, 1, 4, 5, 8, 8, 8, 8, 8, 3, 5, 6, 6, 6, 6, 6, 6, 6, 4, 6, 6, 6, 6, 6, 6, 8, 8, 8, 8, 8, 5, 3, 5, 7, 7, 7, 7, 7, 7, 7, 5, 4, 5, 3, 4, 1, 9, 9, 9, 9, 9, 4, 7, 7, 7, 7, 7, 7, 1},
	{4, 9, 9, 9, 5, 7, 7, 5, 4, 6, 6, 4, 3, 4, 9, 9, 9, 3, 5, 1, 5, 8, 8, 8, 3, 8, 8, 8, 1, 7, 7, 7, 6, 6, 6},
}

// reels lengths [37, 87, 87, 87, 37], total reshuffles 901490607
// symbols: 75.858(lined) + 0(scatter) = 75.857731%
// free spins 93257649, q = 0.10345, sq = 1/(1-q) = 1.115385
// free games frequency: 1/10.008
// RTP = 75.858(sym) + 0.10345*253.83(fg) = 102.116176%
var Reels102 = slot.Reels5x{
	{9, 9, 9, 3, 8, 8, 8, 4, 3, 6, 6, 6, 1, 5, 4, 5, 9, 9, 9, 4, 6, 6, 6, 8, 8, 8, 7, 7, 7, 5, 1, 5, 3, 4, 7, 7, 7},
	{4, 7, 7, 7, 7, 7, 7, 8, 8, 8, 8, 8, 3, 4, 5, 4, 1, 5, 3, 5, 4, 5, 2, 3, 5, 9, 9, 9, 9, 9, 1, 9, 9, 9, 9, 9, 3, 5, 9, 9, 9, 9, 9, 5, 7, 7, 7, 7, 7, 7, 3, 8, 8, 8, 8, 8, 3, 4, 8, 8, 8, 8, 8, 5, 6, 6, 6, 6, 6, 6, 5, 4, 3, 4, 3, 1, 6, 6, 6, 6, 6, 6, 4, 1, 4, 5, 4},
	{7, 7, 7, 7, 7, 7, 3, 1, 3, 5, 4, 5, 4, 2, 4, 3, 4, 8, 8, 8, 8, 8, 5, 1, 8, 8, 8, 8, 8, 1, 5, 4, 5, 7, 7, 7, 7, 7, 7, 3, 4, 5, 4, 6, 6, 6, 6, 6, 6, 3, 5, 9, 9, 9, 9, 9, 3, 9, 9, 9, 9, 9, 4, 5, 4, 3, 8, 8, 8, 8, 8, 5, 9, 9, 9, 9, 9, 5, 4, 6, 6, 6, 6, 6, 6, 3, 1},
	{5, 7, 7, 7, 7, 7, 7, 4, 5, 8, 8, 8, 8, 8, 5, 3, 4, 2, 6, 6, 6, 6, 6, 6, 3, 4, 3, 4, 3, 4, 1, 4, 5, 4, 5, 8, 8, 8, 8, 8, 4, 8, 8, 8, 8, 8, 9, 9, 9, 9, 9, 5, 1, 3, 5, 3, 1, 3, 4, 9, 9, 9, 9, 9, 5, 3, 7, 7, 7, 7, 7, 7, 5, 1, 6, 6, 6, 6, 6, 6, 4, 5, 9, 9, 9, 9, 9},
	{9, 9, 9, 3, 4, 5, 6, 6, 6, 5, 8, 8, 8, 7, 7, 7, 8, 8, 8, 1, 4, 9, 9, 9, 6, 6, 6, 4, 3, 5, 4, 1, 5, 3, 7, 7, 7},
}

// reels lengths [33, 87, 87, 87, 33], total reshuffles 717109767
// symbols: 77.67(lined) + 0(scatter) = 77.669881%
// free spins 74183769, q = 0.10345, sq = 1/(1-q) = 1.115385
// free games frequency: 1/10.008
// RTP = 77.67(sym) + 0.10345*259.9(fg) = 104.555609%
var Reels104 = slot.Reels5x{
	{5, 6, 6, 6, 6, 6, 5, 3, 8, 8, 8, 8, 8, 3, 4, 5, 4, 5, 4, 1, 4, 9, 9, 9, 9, 9, 3, 7, 7, 7, 7, 7, 1},
	{4, 5, 7, 7, 7, 7, 7, 7, 7, 1, 7, 7, 7, 7, 7, 7, 1, 4, 5, 4, 3, 8, 8, 8, 8, 8, 1, 8, 8, 8, 8, 8, 3, 4, 8, 8, 8, 8, 8, 5, 6, 6, 6, 6, 6, 6, 9, 9, 9, 9, 9, 3, 1, 5, 4, 5, 9, 9, 9, 9, 9, 5, 9, 9, 9, 9, 9, 5, 4, 6, 6, 6, 6, 6, 6, 6, 4, 5, 4, 5, 3, 4, 3, 5, 2, 3, 4},
	{8, 8, 8, 8, 8, 4, 5, 4, 6, 6, 6, 6, 6, 6, 5, 3, 5, 2, 3, 9, 9, 9, 9, 9, 3, 6, 6, 6, 6, 6, 6, 6, 5, 4, 1, 9, 9, 9, 9, 9, 1, 8, 8, 8, 8, 8, 5, 4, 7, 7, 7, 7, 7, 7, 1, 3, 5, 4, 5, 4, 5, 8, 8, 8, 8, 8, 4, 3, 4, 5, 4, 3, 7, 7, 7, 7, 7, 7, 7, 4, 1, 9, 9, 9, 9, 9, 5},
	{4, 3, 9, 9, 9, 9, 9, 3, 4, 3, 2, 9, 9, 9, 9, 9, 1, 5, 4, 5, 4, 5, 8, 8, 8, 8, 8, 5, 4, 1, 4, 5, 8, 8, 8, 8, 8, 3, 5, 6, 6, 6, 6, 6, 6, 6, 4, 6, 6, 6, 6, 6, 6, 8, 8, 8, 8, 8, 5, 3, 5, 7, 7, 7, 7, 7, 7, 7, 5, 4, 5, 3, 4, 1, 9, 9, 9, 9, 9, 4, 7, 7, 7, 7, 7, 7, 1},
	{8, 8, 8, 8, 8, 1, 4, 3, 4, 5, 1, 9, 9, 9, 9, 9, 5, 7, 7, 7, 7, 7, 3, 5, 4, 3, 6, 6, 6, 6, 6, 4, 5},
}

// reels lengths [37, 87, 87, 87, 37], total reshuffles 901490607
// symbols: 80.316(lined) + 0(scatter) = 80.315770%
// free spins 93257649, q = 0.10345, sq = 1/(1-q) = 1.115385
// free games frequency: 1/10.008
// RTP = 80.316(sym) + 0.10345*268.75(fg) = 108.117382%
var Reels108 = slot.Reels5x{
	{9, 9, 9, 3, 8, 8, 8, 4, 3, 6, 6, 6, 1, 5, 4, 5, 9, 9, 9, 4, 6, 6, 6, 8, 8, 8, 7, 7, 7, 5, 1, 5, 3, 4, 7, 7, 7},
	{5, 6, 6, 6, 6, 6, 6, 4, 1, 3, 7, 7, 7, 7, 7, 7, 3, 9, 9, 9, 9, 9, 9, 9, 3, 5, 3, 2, 3, 4, 3, 4, 1, 3, 5, 4, 5, 8, 8, 8, 8, 8, 8, 8, 5, 4, 3, 8, 8, 8, 8, 8, 8, 8, 5, 4, 5, 4, 5, 1, 7, 7, 7, 7, 7, 7, 5, 4, 1, 9, 9, 9, 9, 9, 9, 9, 5, 3, 4, 6, 6, 6, 6, 6, 6, 4, 3},
	{4, 9, 9, 9, 9, 9, 9, 9, 5, 1, 4, 3, 4, 1, 4, 3, 4, 6, 6, 6, 6, 6, 6, 4, 3, 5, 8, 8, 8, 8, 8, 8, 8, 3, 5, 9, 9, 9, 9, 9, 9, 9, 5, 4, 3, 2, 7, 7, 7, 7, 7, 7, 4, 3, 5, 1, 5, 4, 5, 3, 4, 8, 8, 8, 8, 8, 8, 8, 3, 5, 3, 6, 6, 6, 6, 6, 6, 1, 3, 5, 7, 7, 7, 7, 7, 7, 5},
	{5, 3, 6, 6, 6, 6, 6, 6, 5, 4, 5, 1, 3, 7, 7, 7, 7, 7, 7, 3, 9, 9, 9, 9, 9, 9, 9, 3, 1, 5, 3, 4, 2, 4, 5, 3, 4, 6, 6, 6, 6, 6, 6, 3, 5, 4, 3, 4, 5, 1, 4, 8, 8, 8, 8, 8, 8, 8, 5, 3, 9, 9, 9, 9, 9, 9, 9, 8, 8, 8, 8, 8, 8, 8, 5, 1, 3, 4, 7, 7, 7, 7, 7, 7, 4, 5, 4},
	{9, 9, 9, 3, 4, 5, 6, 6, 6, 5, 8, 8, 8, 7, 7, 7, 8, 8, 8, 1, 4, 9, 9, 9, 6, 6, 6, 4, 3, 5, 4, 1, 5, 3, 7, 7, 7},
}

// reels lengths [35, 87, 87, 87, 35], total reshuffles 806666175
// symbols: 82.265(lined) + 0(scatter) = 82.265098%
// free spins 83448225, q = 0.10345, sq = 1/(1-q) = 1.115385
// free games frequency: 1/10.008
// RTP = 82.265(sym) + 0.10345*275.27(fg) = 110.741478%
var Reels110 = slot.Reels5x{
	{9, 9, 9, 1, 5, 6, 6, 6, 3, 7, 7, 7, 8, 8, 8, 5, 7, 7, 4, 8, 8, 8, 6, 6, 4, 9, 9, 9, 4, 3, 1, 5, 3, 5, 4},
	{5, 6, 6, 6, 6, 6, 6, 4, 1, 3, 7, 7, 7, 7, 7, 7, 3, 9, 9, 9, 9, 9, 9, 9, 3, 5, 3, 2, 3, 4, 3, 4, 1, 3, 5, 4, 5, 8, 8, 8, 8, 8, 8, 8, 5, 4, 3, 8, 8, 8, 8, 8, 8, 8, 5, 4, 5, 4, 5, 1, 7, 7, 7, 7, 7, 7, 5, 4, 1, 9, 9, 9, 9, 9, 9, 9, 5, 3, 4, 6, 6, 6, 6, 6, 6, 4, 3},
	{4, 9, 9, 9, 9, 9, 9, 9, 5, 1, 4, 3, 4, 1, 4, 3, 4, 6, 6, 6, 6, 6, 6, 4, 3, 5, 8, 8, 8, 8, 8, 8, 8, 3, 5, 9, 9, 9, 9, 9, 9, 9, 5, 4, 3, 2, 7, 7, 7, 7, 7, 7, 4, 3, 5, 1, 5, 4, 5, 3, 4, 8, 8, 8, 8, 8, 8, 8, 3, 5, 3, 6, 6, 6, 6, 6, 6, 1, 3, 5, 7, 7, 7, 7, 7, 7, 5},
	{5, 3, 6, 6, 6, 6, 6, 6, 5, 4, 5, 1, 3, 7, 7, 7, 7, 7, 7, 3, 9, 9, 9, 9, 9, 9, 9, 3, 1, 5, 3, 4, 2, 4, 5, 3, 4, 6, 6, 6, 6, 6, 6, 3, 5, 4, 3, 4, 5, 1, 4, 8, 8, 8, 8, 8, 8, 8, 5, 3, 9, 9, 9, 9, 9, 9, 9, 8, 8, 8, 8, 8, 8, 8, 5, 1, 3, 4, 7, 7, 7, 7, 7, 7, 4, 5, 4},
	{4, 9, 9, 9, 5, 7, 7, 5, 4, 6, 6, 4, 3, 4, 9, 9, 9, 3, 5, 1, 5, 8, 8, 8, 3, 8, 8, 8, 1, 7, 7, 7, 6, 6, 6},
}

// Map with available reels.
var ReelsMap = map[float64]*slot.Reels5x{
	87.833560:  &Reels878,
	88.397358:  &Reels883,
	89.380124:  &Reels893,
	89.936843:  &Reels899,
	91.101562:  &Reels911,
	91.407100:  &Reels914,
	92.094481:  &Reels920,
	92.481524:  &Reels924,
	93.102720:  &Reels931,
	93.448261:  &Reels934,
	94.069458:  &Reels940,
	94.499195:  &Reels944,
	94.751321:  &Reels947,
	95.096862:  &Reels950,
	96.022554:  &Reels960,
	96.497919:  &Reels964,
	97.078180:  &Reels970,
	98.332501:  &Reels983,
	99.230772:  &Reels992,
	99.751771:  &Reels997,
	101.113246: &Reels101,
	102.116176: &Reels102,
	104.555609: &Reels104,
	108.117382: &Reels108,
	110.741478: &Reels110,
}

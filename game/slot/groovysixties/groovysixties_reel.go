package groovysixties

import (
	slot "github.com/slotopol/server/game/slot"
)

// reels lengths [52, 55, 56, 55, 52], total reshuffles 458057600
// symbols: 72.595(lined) + 2.7202(scatter) = 75.314745%
// free spins 31150080, q = 0.068005, sq = 1/(1-q) = 1.072967
// free games frequency: 1/73.524
// RTP = 75.315(sym) + 0.068005*161.62(fg) = 86.305698%
var Reels863 = slot.Reels5x{
	{1, 1, 1, 1, 3, 11, 10, 9, 7, 11, 12, 10, 9, 8, 12, 7, 3, 4, 11, 7, 10, 8, 9, 6, 3, 5, 7, 6, 11, 9, 12, 10, 5, 9, 12, 10, 5, 11, 12, 9, 10, 8, 6, 12, 11, 8, 4, 5, 7, 8, 6, 4},
	{1, 1, 1, 1, 12, 7, 9, 3, 5, 10, 8, 11, 9, 10, 4, 6, 12, 11, 10, 9, 5, 12, 11, 8, 3, 9, 7, 6, 2, 12, 3, 11, 8, 2, 4, 6, 10, 8, 5, 9, 7, 4, 12, 11, 7, 8, 10, 6, 12, 11, 10, 7, 9, 5, 2},
	{1, 1, 1, 1, 3, 6, 9, 10, 8, 11, 2, 6, 9, 12, 10, 5, 2, 8, 9, 3, 11, 10, 5, 4, 7, 9, 12, 11, 8, 6, 3, 7, 12, 11, 9, 8, 4, 11, 10, 6, 4, 12, 7, 8, 10, 12, 7, 5, 2, 11, 10, 7, 2, 12, 9, 5},
	{1, 1, 1, 1, 5, 10, 8, 11, 7, 10, 5, 11, 4, 7, 12, 8, 11, 10, 6, 9, 11, 12, 8, 4, 9, 3, 10, 2, 8, 12, 9, 7, 6, 12, 3, 9, 11, 10, 2, 5, 12, 6, 3, 11, 9, 7, 6, 10, 2, 9, 5, 4, 12, 8, 7},
	{1, 1, 1, 1, 10, 4, 5, 9, 12, 7, 5, 6, 12, 7, 8, 10, 4, 9, 8, 6, 12, 11, 7, 4, 9, 12, 3, 10, 8, 11, 9, 6, 5, 3, 10, 11, 12, 9, 6, 7, 3, 11, 10, 8, 12, 11, 7, 5, 10, 11, 9, 8},
}

// reels lengths [52, 55, 56, 55, 52], total reshuffles 458057600
// symbols: 73.486(lined) + 2.7202(scatter) = 76.206004%
// free spins 31150080, q = 0.068005, sq = 1/(1-q) = 1.072967
// free games frequency: 1/73.524
// RTP = 76.206(sym) + 0.068005*163.53(fg) = 87.327022%
var Reels873 = slot.Reels5x{
	{1, 1, 1, 1, 3, 11, 10, 9, 7, 11, 12, 10, 9, 8, 12, 7, 3, 4, 11, 7, 10, 8, 9, 6, 3, 5, 7, 6, 11, 9, 12, 10, 5, 9, 12, 10, 5, 11, 12, 9, 10, 8, 6, 12, 11, 8, 4, 5, 7, 8, 6, 4},
	{1, 1, 1, 1, 12, 7, 9, 3, 5, 10, 8, 11, 9, 10, 4, 6, 12, 11, 10, 9, 5, 12, 11, 8, 3, 9, 7, 6, 2, 12, 3, 11, 8, 2, 4, 6, 10, 8, 5, 9, 7, 4, 12, 11, 7, 8, 10, 6, 12, 11, 10, 7, 9, 5, 2},
	{1, 1, 1, 1, 11, 2, 8, 12, 4, 6, 3, 5, 9, 10, 7, 2, 4, 10, 12, 11, 9, 10, 5, 11, 2, 8, 5, 9, 11, 6, 12, 10, 3, 9, 11, 2, 3, 7, 9, 12, 6, 7, 4, 8, 3, 9, 7, 12, 8, 6, 10, 11, 5, 12, 4, 10},
	{1, 1, 1, 1, 5, 10, 8, 11, 7, 10, 5, 11, 4, 7, 12, 8, 11, 10, 6, 9, 11, 12, 8, 4, 9, 3, 10, 2, 8, 12, 9, 7, 6, 12, 3, 9, 11, 10, 2, 5, 12, 6, 3, 11, 9, 7, 6, 10, 2, 9, 5, 4, 12, 8, 7},
	{1, 1, 1, 1, 10, 4, 5, 9, 12, 7, 5, 6, 12, 7, 8, 10, 4, 9, 8, 6, 12, 11, 7, 4, 9, 12, 3, 10, 8, 11, 9, 6, 5, 3, 10, 11, 12, 9, 6, 7, 3, 11, 10, 8, 12, 11, 7, 5, 10, 11, 9, 8},
}

// reels lengths [52, 55, 56, 55, 52], total reshuffles 458057600
// symbols: 75.283(lined) + 2.7202(scatter) = 78.002808%
// free spins 31150080, q = 0.068005, sq = 1/(1-q) = 1.072967
// free games frequency: 1/73.524
// RTP = 78.003(sym) + 0.068005*167.39(fg) = 89.386041%
var Reels893 = slot.Reels5x{
	{1, 1, 1, 1, 3, 11, 10, 9, 7, 11, 12, 10, 9, 8, 12, 7, 3, 4, 11, 7, 10, 8, 9, 6, 3, 5, 7, 6, 11, 9, 12, 10, 5, 9, 12, 10, 5, 11, 12, 9, 10, 8, 6, 12, 11, 8, 4, 5, 7, 8, 6, 4},
	{1, 1, 1, 1, 10, 8, 6, 5, 11, 7, 12, 8, 9, 2, 10, 7, 11, 8, 3, 10, 9, 12, 11, 7, 3, 5, 11, 9, 12, 6, 2, 4, 9, 12, 3, 4, 10, 12, 6, 4, 8, 2, 5, 11, 10, 12, 9, 5, 3, 6, 7, 10, 11, 9, 4},
	{1, 1, 1, 1, 11, 2, 8, 12, 4, 6, 3, 5, 9, 10, 7, 2, 4, 10, 12, 11, 9, 10, 5, 11, 2, 8, 5, 9, 11, 6, 12, 10, 3, 9, 11, 2, 3, 7, 9, 12, 6, 7, 4, 8, 3, 9, 7, 12, 8, 6, 10, 11, 5, 12, 4, 10},
	{1, 1, 1, 1, 5, 7, 12, 9, 4, 5, 10, 9, 12, 2, 10, 4, 11, 3, 9, 12, 8, 10, 11, 9, 2, 8, 4, 9, 6, 7, 3, 11, 12, 6, 7, 10, 3, 5, 11, 8, 3, 10, 6, 11, 5, 4, 10, 8, 12, 7, 9, 2, 11, 12, 6},
	{1, 1, 1, 1, 10, 4, 5, 9, 12, 7, 5, 6, 12, 7, 8, 10, 4, 9, 8, 6, 12, 11, 7, 4, 9, 12, 3, 10, 8, 11, 9, 6, 5, 3, 10, 11, 12, 9, 6, 7, 3, 11, 10, 8, 12, 11, 7, 5, 10, 11, 9, 8},
}

// reels lengths [52, 55, 56, 55, 52], total reshuffles 458057600
// symbols: 75.671(lined) + 2.7202(scatter) = 78.390952%
// free spins 31150080, q = 0.068005, sq = 1/(1-q) = 1.072967
// free games frequency: 1/73.524
// RTP = 78.391(sym) + 0.068005*168.22(fg) = 89.830827%
var Reels898 = slot.Reels5x{
	{1, 1, 1, 1, 5, 8, 11, 4, 6, 12, 11, 9, 10, 3, 12, 7, 10, 9, 4, 6, 10, 12, 3, 7, 9, 12, 6, 10, 8, 4, 11, 10, 9, 7, 5, 3, 10, 11, 12, 5, 9, 8, 6, 5, 3, 11, 7, 8, 12, 4, 11, 9},
	{1, 1, 1, 1, 8, 11, 2, 10, 7, 9, 8, 11, 12, 4, 5, 8, 10, 9, 7, 12, 8, 3, 9, 12, 11, 6, 5, 4, 7, 10, 11, 5, 2, 3, 11, 12, 8, 6, 11, 9, 4, 10, 6, 9, 3, 7, 12, 10, 9, 2, 4, 10, 6, 5, 12},
	{1, 1, 1, 1, 3, 7, 9, 8, 4, 6, 11, 7, 9, 2, 3, 4, 6, 9, 12, 5, 10, 11, 9, 8, 10, 11, 12, 2, 4, 10, 5, 8, 12, 6, 2, 8, 12, 9, 5, 10, 11, 2, 7, 12, 10, 11, 3, 9, 10, 7, 5, 6, 8, 12, 11, 4},
	{1, 1, 1, 1, 11, 7, 6, 10, 9, 11, 12, 8, 10, 2, 12, 9, 6, 8, 2, 9, 10, 7, 12, 11, 9, 5, 6, 11, 3, 10, 5, 7, 9, 4, 5, 12, 9, 7, 5, 4, 11, 6, 3, 12, 10, 8, 4, 2, 3, 11, 8, 10, 4, 12, 8},
	{1, 1, 1, 1, 9, 8, 6, 12, 10, 5, 8, 9, 11, 5, 7, 9, 10, 3, 6, 12, 10, 5, 6, 7, 3, 12, 11, 9, 4, 3, 11, 10, 9, 12, 6, 11, 10, 12, 8, 11, 7, 10, 4, 5, 3, 9, 4, 8, 7, 11, 4, 12},
}

// reels lengths [52, 55, 56, 55, 52], total reshuffles 458057600
// symbols: 76.569(lined) + 2.7202(scatter) = 79.289354%
// free spins 31150080, q = 0.068005, sq = 1/(1-q) = 1.072967
// free games frequency: 1/73.524
// RTP = 79.289(sym) + 0.068005*170.15(fg) = 90.860336%
var Reels908 = slot.Reels5x{
	{1, 1, 1, 1, 11, 3, 6, 10, 9, 12, 7, 11, 4, 9, 12, 10, 3, 5, 11, 9, 12, 8, 7, 6, 5, 9, 4, 10, 12, 7, 11, 10, 8, 9, 11, 4, 6, 8, 10, 9, 7, 5, 6, 3, 8, 10, 4, 12, 11, 5, 8, 12},
	{1, 1, 1, 1, 10, 8, 6, 5, 11, 7, 12, 8, 9, 2, 10, 7, 11, 8, 3, 10, 9, 12, 11, 7, 3, 5, 11, 9, 12, 6, 2, 4, 9, 12, 3, 4, 10, 12, 6, 4, 8, 2, 5, 11, 10, 12, 9, 5, 3, 6, 7, 10, 11, 9, 4},
	{1, 1, 1, 1, 11, 2, 8, 12, 4, 6, 3, 5, 9, 10, 7, 2, 4, 10, 12, 11, 9, 10, 5, 11, 2, 8, 5, 9, 11, 6, 12, 10, 3, 9, 11, 2, 3, 7, 9, 12, 6, 7, 4, 8, 3, 9, 7, 12, 8, 6, 10, 11, 5, 12, 4, 10},
	{1, 1, 1, 1, 5, 7, 12, 9, 4, 5, 10, 9, 12, 2, 10, 4, 11, 3, 9, 12, 8, 10, 11, 9, 2, 8, 4, 9, 6, 7, 3, 11, 12, 6, 7, 10, 3, 5, 11, 8, 3, 10, 6, 11, 5, 4, 10, 8, 12, 7, 9, 2, 11, 12, 6},
	{1, 1, 1, 1, 10, 12, 8, 3, 10, 11, 7, 6, 9, 11, 4, 7, 9, 8, 10, 7, 9, 12, 4, 3, 11, 12, 9, 6, 10, 11, 5, 12, 4, 6, 5, 12, 9, 11, 4, 3, 8, 10, 12, 11, 8, 7, 5, 9, 10, 8, 5, 6},
}

// reels lengths [52, 55, 56, 55, 52], total reshuffles 458057600
// symbols: 77.856(lined) + 2.7202(scatter) = 80.575900%
// free spins 31150080, q = 0.068005, sq = 1/(1-q) = 1.072967
// free games frequency: 1/73.524
// RTP = 80.576(sym) + 0.068005*172.91(fg) = 92.334632%
var Reels923 = slot.Reels5x{
	{1, 1, 1, 1, 5, 8, 11, 4, 6, 12, 11, 9, 10, 3, 12, 7, 10, 9, 4, 6, 10, 12, 3, 7, 9, 12, 6, 10, 8, 4, 11, 10, 9, 7, 5, 3, 10, 11, 12, 5, 9, 8, 6, 5, 3, 11, 7, 8, 12, 4, 11, 9},
	{1, 1, 1, 1, 10, 8, 6, 5, 11, 7, 12, 8, 9, 2, 10, 7, 11, 8, 3, 10, 9, 12, 11, 7, 3, 5, 11, 9, 12, 6, 2, 4, 9, 12, 3, 4, 10, 12, 6, 4, 8, 2, 5, 11, 10, 12, 9, 5, 3, 6, 7, 10, 11, 9, 4},
	{1, 1, 1, 1, 11, 2, 8, 12, 4, 6, 3, 5, 9, 10, 7, 2, 4, 10, 12, 11, 9, 10, 5, 11, 2, 8, 5, 9, 11, 6, 12, 10, 3, 9, 11, 2, 3, 7, 9, 12, 6, 7, 4, 8, 3, 9, 7, 12, 8, 6, 10, 11, 5, 12, 4, 10},
	{1, 1, 1, 1, 5, 7, 12, 9, 4, 5, 10, 9, 12, 2, 10, 4, 11, 3, 9, 12, 8, 10, 11, 9, 2, 8, 4, 9, 6, 7, 3, 11, 12, 6, 7, 10, 3, 5, 11, 8, 3, 10, 6, 11, 5, 4, 10, 8, 12, 7, 9, 2, 11, 12, 6},
	{1, 1, 1, 1, 9, 8, 6, 12, 10, 5, 8, 9, 11, 5, 7, 9, 10, 3, 6, 12, 10, 5, 6, 7, 3, 12, 11, 9, 4, 3, 11, 10, 9, 12, 6, 11, 10, 12, 8, 11, 7, 10, 4, 5, 3, 9, 4, 8, 7, 11, 4, 12},
}

// reels lengths [52, 55, 55, 55, 52], total reshuffles 449878000
// symbols: 78.602(lined) + 2.7696(scatter) = 81.371846%
// free spins 31150080, q = 0.069241, sq = 1/(1-q) = 1.074392
// free games frequency: 1/72.211
// RTP = 81.372(sym) + 0.069241*174.85(fg) = 93.478703%
var Reels934 = slot.Reels5x{
	{1, 1, 1, 1, 5, 8, 11, 4, 6, 12, 11, 9, 10, 3, 12, 7, 10, 9, 4, 6, 10, 12, 3, 7, 9, 12, 6, 10, 8, 4, 11, 10, 9, 7, 5, 3, 10, 11, 12, 5, 9, 8, 6, 5, 3, 11, 7, 8, 12, 4, 11, 9},
	{1, 1, 1, 1, 10, 8, 6, 5, 11, 7, 12, 8, 9, 2, 10, 7, 11, 8, 3, 10, 9, 12, 11, 7, 3, 5, 11, 9, 12, 6, 2, 4, 9, 12, 3, 4, 10, 12, 6, 4, 8, 2, 5, 11, 10, 12, 9, 5, 3, 6, 7, 10, 11, 9, 4},
	{1, 1, 1, 1, 3, 8, 5, 6, 12, 9, 11, 10, 7, 6, 12, 2, 10, 11, 7, 9, 3, 10, 11, 9, 7, 8, 4, 11, 12, 5, 6, 2, 4, 8, 11, 10, 2, 5, 3, 9, 10, 12, 2, 11, 6, 12, 4, 7, 8, 5, 4, 3, 10, 12, 9},
	{1, 1, 1, 1, 5, 7, 12, 9, 4, 5, 10, 9, 12, 2, 10, 4, 11, 3, 9, 12, 8, 10, 11, 9, 2, 8, 4, 9, 6, 7, 3, 11, 12, 6, 7, 10, 3, 5, 11, 8, 3, 10, 6, 11, 5, 4, 10, 8, 12, 7, 9, 2, 11, 12, 6},
	{1, 1, 1, 1, 9, 8, 6, 12, 10, 5, 8, 9, 11, 5, 7, 9, 10, 3, 6, 12, 10, 5, 6, 7, 3, 12, 11, 9, 4, 3, 11, 10, 9, 12, 6, 11, 10, 12, 8, 11, 7, 10, 4, 5, 3, 9, 4, 8, 7, 11, 4, 12},
}

// reels lengths [52, 54, 56, 54, 52], total reshuffles 441552384
// symbols: 78.998(lined) + 2.8219(scatter) = 81.819898%
// free spins 31150080, q = 0.070547, sq = 1/(1-q) = 1.075901
// free games frequency: 1/70.875
// RTP = 81.82(sym) + 0.070547*176.06(fg) = 94.240376%
var Reels942 = slot.Reels5x{
	{1, 1, 1, 1, 5, 8, 11, 4, 6, 12, 11, 9, 10, 3, 12, 7, 10, 9, 4, 6, 10, 12, 3, 7, 9, 12, 6, 10, 8, 4, 11, 10, 9, 7, 5, 3, 10, 11, 12, 5, 9, 8, 6, 5, 3, 11, 7, 8, 12, 4, 11, 9},
	{1, 1, 1, 1, 6, 7, 2, 3, 12, 11, 2, 3, 5, 12, 6, 2, 9, 8, 3, 12, 9, 7, 10, 5, 8, 3, 10, 11, 4, 6, 9, 11, 10, 6, 12, 4, 10, 5, 12, 7, 10, 11, 8, 4, 9, 10, 11, 12, 7, 5, 9, 11, 4, 8},
	{1, 1, 1, 1, 11, 2, 8, 12, 4, 6, 3, 5, 9, 10, 7, 2, 4, 10, 12, 11, 9, 10, 5, 11, 2, 8, 5, 9, 11, 6, 12, 10, 3, 9, 11, 2, 3, 7, 9, 12, 6, 7, 4, 8, 3, 9, 7, 12, 8, 6, 10, 11, 5, 12, 4, 10},
	{1, 1, 1, 1, 12, 9, 2, 8, 4, 6, 12, 8, 7, 6, 11, 3, 9, 7, 10, 2, 12, 4, 9, 10, 12, 11, 6, 10, 8, 7, 9, 3, 10, 5, 8, 11, 4, 7, 10, 3, 5, 4, 11, 12, 10, 5, 2, 12, 6, 11, 9, 3, 5, 11},
	{1, 1, 1, 1, 9, 8, 6, 12, 10, 5, 8, 9, 11, 5, 7, 9, 10, 3, 6, 12, 10, 5, 6, 7, 3, 12, 11, 9, 4, 3, 11, 10, 9, 12, 6, 11, 10, 12, 8, 11, 7, 10, 4, 5, 3, 9, 4, 8, 7, 11, 4, 12},
}

// reels lengths [52, 54, 55, 54, 52], total reshuffles 433667520
// symbols: 79.848(lined) + 2.8732(scatter) = 82.721047%
// free spins 31150080, q = 0.071829, sq = 1/(1-q) = 1.077388
// free games frequency: 1/69.609
// RTP = 82.721(sym) + 0.071829*178.25(fg) = 95.524304%
var Reels955 = slot.Reels5x{
	{1, 1, 1, 1, 5, 8, 11, 4, 6, 12, 11, 9, 10, 3, 12, 7, 10, 9, 4, 6, 10, 12, 3, 7, 9, 12, 6, 10, 8, 4, 11, 10, 9, 7, 5, 3, 10, 11, 12, 5, 9, 8, 6, 5, 3, 11, 7, 8, 12, 4, 11, 9},
	{1, 1, 1, 1, 6, 7, 2, 3, 12, 11, 2, 3, 5, 12, 6, 2, 9, 8, 3, 12, 9, 7, 10, 5, 8, 3, 10, 11, 4, 6, 9, 11, 10, 6, 12, 4, 10, 5, 12, 7, 10, 11, 8, 4, 9, 10, 11, 12, 7, 5, 9, 11, 4, 8},
	{1, 1, 1, 1, 3, 8, 5, 6, 12, 9, 11, 10, 7, 6, 12, 2, 10, 11, 7, 9, 3, 10, 11, 9, 7, 8, 4, 11, 12, 5, 6, 2, 4, 8, 11, 10, 2, 5, 3, 9, 10, 12, 2, 11, 6, 12, 4, 7, 8, 5, 4, 3, 10, 12, 9},
	{1, 1, 1, 1, 12, 9, 2, 8, 4, 6, 12, 8, 7, 6, 11, 3, 9, 7, 10, 2, 12, 4, 9, 10, 12, 11, 6, 10, 8, 7, 9, 3, 10, 5, 8, 11, 4, 7, 10, 3, 5, 4, 11, 12, 10, 5, 2, 12, 6, 11, 9, 3, 5, 11},
	{1, 1, 1, 1, 9, 8, 6, 12, 10, 5, 8, 9, 11, 5, 7, 9, 10, 3, 6, 12, 10, 5, 6, 7, 3, 12, 11, 9, 4, 3, 11, 10, 9, 12, 6, 11, 10, 12, 8, 11, 7, 10, 4, 5, 3, 9, 4, 8, 7, 11, 4, 12},
}

// reels lengths [52, 53, 56, 53, 52], total reshuffles 425350016
// symbols: 80.184(lined) + 2.9294(scatter) = 83.113525%
// free spins 31150080, q = 0.073234, sq = 1/(1-q) = 1.079021
// free games frequency: 1/68.274
// RTP = 83.114(sym) + 0.073234*179.36(fg) = 96.248956%
var Reels962 = slot.Reels5x{
	{1, 1, 1, 1, 5, 8, 11, 4, 6, 12, 11, 9, 10, 3, 12, 7, 10, 9, 4, 6, 10, 12, 3, 7, 9, 12, 6, 10, 8, 4, 11, 10, 9, 7, 5, 3, 10, 11, 12, 5, 9, 8, 6, 5, 3, 11, 7, 8, 12, 4, 11, 9},
	{1, 1, 1, 1, 10, 6, 4, 12, 9, 2, 5, 8, 12, 7, 10, 9, 11, 5, 3, 10, 12, 11, 9, 3, 10, 6, 8, 4, 2, 12, 8, 7, 11, 5, 10, 2, 6, 9, 3, 11, 4, 12, 7, 9, 11, 12, 7, 8, 11, 4, 6, 5, 3},
	{1, 1, 1, 1, 11, 2, 8, 12, 4, 6, 3, 5, 9, 10, 7, 2, 4, 10, 12, 11, 9, 10, 5, 11, 2, 8, 5, 9, 11, 6, 12, 10, 3, 9, 11, 2, 3, 7, 9, 12, 6, 7, 4, 8, 3, 9, 7, 12, 8, 6, 10, 11, 5, 12, 4, 10},
	{1, 1, 1, 1, 8, 11, 12, 2, 8, 10, 11, 7, 3, 6, 5, 9, 12, 7, 4, 10, 11, 6, 12, 5, 11, 4, 9, 3, 7, 10, 12, 8, 2, 10, 9, 8, 12, 5, 6, 4, 10, 3, 11, 12, 9, 5, 11, 2, 6, 7, 3, 4, 9},
	{1, 1, 1, 1, 9, 8, 6, 12, 10, 5, 8, 9, 11, 5, 7, 9, 10, 3, 6, 12, 10, 5, 6, 7, 3, 12, 11, 9, 4, 3, 11, 10, 9, 12, 6, 11, 10, 12, 8, 11, 7, 10, 4, 5, 3, 9, 4, 8, 7, 11, 4, 12},
}

// reels lengths [52, 53, 55, 53, 52], total reshuffles 417754480
// symbols: 81.039(lined) + 2.9826(scatter) = 84.021455%
// free spins 31150080, q = 0.074566, sq = 1/(1-q) = 1.080574
// free games frequency: 1/67.055
// RTP = 84.021(sym) + 0.074566*181.58(fg) = 97.561265%
var Reels975 = slot.Reels5x{
	{1, 1, 1, 1, 5, 8, 11, 4, 6, 12, 11, 9, 10, 3, 12, 7, 10, 9, 4, 6, 10, 12, 3, 7, 9, 12, 6, 10, 8, 4, 11, 10, 9, 7, 5, 3, 10, 11, 12, 5, 9, 8, 6, 5, 3, 11, 7, 8, 12, 4, 11, 9},
	{1, 1, 1, 1, 10, 6, 4, 12, 9, 2, 5, 8, 12, 7, 10, 9, 11, 5, 3, 10, 12, 11, 9, 3, 10, 6, 8, 4, 2, 12, 8, 7, 11, 5, 10, 2, 6, 9, 3, 11, 4, 12, 7, 9, 11, 12, 7, 8, 11, 4, 6, 5, 3},
	{1, 1, 1, 1, 3, 8, 5, 6, 12, 9, 11, 10, 7, 6, 12, 2, 10, 11, 7, 9, 3, 10, 11, 9, 7, 8, 4, 11, 12, 5, 6, 2, 4, 8, 11, 10, 2, 5, 3, 9, 10, 12, 2, 11, 6, 12, 4, 7, 8, 5, 4, 3, 10, 12, 9},
	{1, 1, 1, 1, 8, 11, 12, 2, 8, 10, 11, 7, 3, 6, 5, 9, 12, 7, 4, 10, 11, 6, 12, 5, 11, 4, 9, 3, 7, 10, 12, 8, 2, 10, 9, 8, 12, 5, 6, 4, 10, 3, 11, 12, 9, 5, 11, 2, 6, 7, 3, 4, 9},
	{1, 1, 1, 1, 9, 8, 6, 12, 10, 5, 8, 9, 11, 5, 7, 9, 10, 3, 6, 12, 10, 5, 6, 7, 3, 12, 11, 9, 4, 3, 11, 10, 9, 12, 6, 11, 10, 12, 8, 11, 7, 10, 4, 5, 3, 9, 4, 8, 7, 11, 4, 12},
}

// reels lengths [52, 53, 54, 53, 52], total reshuffles 410158944
// symbols: 81.925(lined) + 3.0379(scatter) = 84.963012%
// free spins 31150080, q = 0.075946, sq = 1/(1-q) = 1.082188
// free games frequency: 1/65.836
// RTP = 84.963(sym) + 0.075946*183.89(fg) = 98.928936%
var Reels989 = slot.Reels5x{
	{1, 1, 1, 1, 5, 8, 11, 4, 6, 12, 11, 9, 10, 3, 12, 7, 10, 9, 4, 6, 10, 12, 3, 7, 9, 12, 6, 10, 8, 4, 11, 10, 9, 7, 5, 3, 10, 11, 12, 5, 9, 8, 6, 5, 3, 11, 7, 8, 12, 4, 11, 9},
	{1, 1, 1, 1, 10, 6, 4, 12, 9, 2, 5, 8, 12, 7, 10, 9, 11, 5, 3, 10, 12, 11, 9, 3, 10, 6, 8, 4, 2, 12, 8, 7, 11, 5, 10, 2, 6, 9, 3, 11, 4, 12, 7, 9, 11, 12, 7, 8, 11, 4, 6, 5, 3},
	{1, 1, 1, 1, 12, 10, 7, 11, 12, 3, 9, 7, 5, 3, 9, 8, 12, 11, 6, 7, 4, 8, 11, 12, 10, 4, 9, 2, 10, 7, 12, 4, 6, 2, 11, 5, 8, 10, 12, 5, 6, 11, 3, 9, 8, 5, 2, 3, 11, 9, 10, 6, 4, 2},
	{1, 1, 1, 1, 8, 11, 12, 2, 8, 10, 11, 7, 3, 6, 5, 9, 12, 7, 4, 10, 11, 6, 12, 5, 11, 4, 9, 3, 7, 10, 12, 8, 2, 10, 9, 8, 12, 5, 6, 4, 10, 3, 11, 12, 9, 5, 11, 2, 6, 7, 3, 4, 9},
	{1, 1, 1, 1, 9, 8, 6, 12, 10, 5, 8, 9, 11, 5, 7, 9, 10, 3, 6, 12, 10, 5, 6, 7, 3, 12, 11, 9, 4, 3, 11, 10, 9, 12, 6, 11, 10, 12, 8, 11, 7, 10, 4, 5, 3, 9, 4, 8, 7, 11, 4, 12},
}

// reels lengths [51, 53, 54, 53, 52], total reshuffles 402271272
// symbols: 82.968(lined) + 3.0379(scatter) = 86.005589%
// free spins 30551040, q = 0.075946, sq = 1/(1-q) = 1.082188
// free games frequency: 1/65.836
// RTP = 86.006(sym) + 0.075946*186.15(fg) = 100.142889%
var Reels100 = slot.Reels5x{
	{1, 1, 1, 1, 12, 6, 10, 7, 3, 4, 9, 12, 3, 10, 11, 12, 5, 3, 6, 10, 8, 4, 11, 7, 5, 10, 4, 12, 8, 9, 11, 12, 6, 3, 11, 9, 7, 5, 6, 10, 12, 5, 8, 11, 9, 4, 10, 11, 7, 9, 8},
	{1, 1, 1, 1, 10, 6, 4, 12, 9, 2, 5, 8, 12, 7, 10, 9, 11, 5, 3, 10, 12, 11, 9, 3, 10, 6, 8, 4, 2, 12, 8, 7, 11, 5, 10, 2, 6, 9, 3, 11, 4, 12, 7, 9, 11, 12, 7, 8, 11, 4, 6, 5, 3},
	{1, 1, 1, 1, 12, 10, 7, 11, 12, 3, 9, 7, 5, 3, 9, 8, 12, 11, 6, 7, 4, 8, 11, 12, 10, 4, 9, 2, 10, 7, 12, 4, 6, 2, 11, 5, 8, 10, 12, 5, 6, 11, 3, 9, 8, 5, 2, 3, 11, 9, 10, 6, 4, 2},
	{1, 1, 1, 1, 8, 11, 12, 2, 8, 10, 11, 7, 3, 6, 5, 9, 12, 7, 4, 10, 11, 6, 12, 5, 11, 4, 9, 3, 7, 10, 12, 8, 2, 10, 9, 8, 12, 5, 6, 4, 10, 3, 11, 12, 9, 5, 11, 2, 6, 7, 3, 4, 9},
	{1, 1, 1, 1, 9, 8, 6, 12, 10, 5, 8, 9, 11, 5, 7, 9, 10, 3, 6, 12, 10, 5, 6, 7, 3, 12, 11, 9, 4, 3, 11, 10, 9, 12, 6, 11, 10, 12, 8, 11, 7, 10, 4, 5, 3, 9, 4, 8, 7, 11, 4, 12},
}

// reels lengths [50, 53, 54, 53, 50], total reshuffles 379215000
// symbols: 84.419(lined) + 3.0379(scatter) = 87.456878%
// free spins 28800000, q = 0.075946, sq = 1/(1-q) = 1.082188
// free games frequency: 1/65.836
// RTP = 87.457(sym) + 0.075946*189.29(fg) = 101.832736%
var Reels101 = slot.Reels5x{
	{1, 1, 1, 1, 11, 5, 6, 12, 9, 7, 10, 11, 4, 9, 12, 5, 3, 11, 12, 9, 8, 4, 7, 10, 11, 6, 12, 10, 11, 9, 3, 4, 10, 8, 6, 12, 7, 5, 8, 4, 9, 6, 11, 5, 3, 12, 8, 7, 3, 10},
	{1, 1, 1, 1, 10, 6, 4, 12, 9, 2, 5, 8, 12, 7, 10, 9, 11, 5, 3, 10, 12, 11, 9, 3, 10, 6, 8, 4, 2, 12, 8, 7, 11, 5, 10, 2, 6, 9, 3, 11, 4, 12, 7, 9, 11, 12, 7, 8, 11, 4, 6, 5, 3},
	{1, 1, 1, 1, 12, 10, 7, 11, 12, 3, 9, 7, 5, 3, 9, 8, 12, 11, 6, 7, 4, 8, 11, 12, 10, 4, 9, 2, 10, 7, 12, 4, 6, 2, 11, 5, 8, 10, 12, 5, 6, 11, 3, 9, 8, 5, 2, 3, 11, 9, 10, 6, 4, 2},
	{1, 1, 1, 1, 8, 11, 12, 2, 8, 10, 11, 7, 3, 6, 5, 9, 12, 7, 4, 10, 11, 6, 12, 5, 11, 4, 9, 3, 7, 10, 12, 8, 2, 10, 9, 8, 12, 5, 6, 4, 10, 3, 11, 12, 9, 5, 11, 2, 6, 7, 3, 4, 9},
	{1, 1, 1, 1, 11, 6, 9, 12, 3, 7, 9, 4, 8, 10, 6, 3, 5, 12, 7, 11, 3, 10, 12, 4, 7, 11, 12, 9, 10, 11, 5, 9, 8, 4, 5, 10, 11, 6, 8, 5, 12, 11, 4, 8, 6, 3, 12, 7, 9, 10},
}

// reels lengths [48, 51, 51, 51, 48], total reshuffles 305627904
// symbols: 93.469(lined) + 2.6053(scatter) = 96.074809%
// free spins 19906560, q = 0.065133, sq = 1/(1-q) = 1.069671
// free games frequency: 1/76.766
// RTP = 96.075(sym) + 0.065133*205.54(fg) = 109.462111%
var Reels109 = slot.Reels5x{
	{1, 1, 1, 1, 5, 9, 8, 11, 10, 6, 12, 3, 11, 5, 9, 4, 6, 7, 9, 8, 12, 10, 5, 11, 7, 9, 12, 4, 7, 6, 8, 12, 5, 3, 10, 11, 12, 9, 10, 4, 8, 7, 3, 10, 11, 6, 3, 4},
	{1, 1, 1, 1, 3, 12, 5, 4, 7, 8, 11, 12, 7, 9, 5, 11, 3, 12, 6, 2, 4, 10, 12, 3, 9, 4, 12, 10, 2, 8, 7, 10, 11, 9, 5, 8, 4, 6, 10, 7, 9, 3, 8, 6, 5, 9, 11, 6, 10, 2, 11},
	{1, 1, 1, 1, 7, 8, 6, 3, 5, 2, 11, 6, 10, 3, 11, 9, 8, 4, 5, 12, 7, 10, 8, 12, 2, 6, 11, 9, 10, 7, 8, 9, 11, 3, 4, 6, 9, 12, 4, 5, 7, 9, 11, 4, 10, 12, 5, 3, 2, 12, 10},
	{1, 1, 1, 1, 7, 12, 4, 11, 10, 3, 2, 11, 7, 12, 5, 3, 11, 9, 4, 12, 8, 6, 9, 5, 11, 7, 3, 12, 8, 6, 4, 5, 9, 6, 7, 10, 4, 8, 9, 10, 6, 8, 2, 10, 12, 11, 2, 9, 3, 10, 5},
	{1, 1, 1, 1, 6, 12, 4, 3, 8, 12, 10, 9, 11, 8, 10, 9, 6, 4, 12, 3, 10, 5, 7, 9, 12, 11, 7, 4, 10, 11, 9, 6, 10, 7, 12, 5, 3, 11, 4, 5, 9, 8, 3, 11, 6, 8, 7, 5},
}

// Map with available reels.
var ReelsMap = map[float64]*slot.Reels5x{
	86.305698:  &Reels863,
	87.327022:  &Reels873,
	89.386041:  &Reels893,
	89.830827:  &Reels898,
	90.860336:  &Reels908,
	92.334632:  &Reels923,
	93.478703:  &Reels934,
	94.240376:  &Reels942,
	95.524304:  &Reels955,
	96.248956:  &Reels962,
	97.561265:  &Reels975,
	98.928936:  &Reels989,
	100.142889: &Reels100,
	101.832736: &Reels101,
	109.462111: &Reels109,
}

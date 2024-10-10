local path = arg[0]:match("(.*[/\\])")
dofile(path.."reelgen.lua")

local symset = {
	1, --  1 knife
	1, --  2 sneakers
	3, --  3 tent
	3, --  4 drum
	4, --  5 camel
	4, --  6 king
	4, --  7 queen
	4, --  8 jack
	4, --  9 ten
	4, -- 10 nine
	1, -- 11 wild
	1, -- 12 scatter
}

local neighbours = {
	--1, 2, 3, 4, 5, 6, 7, 8, 9,10,11,12,
	{ 2, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, }, --  1
	{ 1, 2, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, }, --  2
	{ 1, 1, 2, 1, 0, 0, 0, 0, 0, 0, 0, 0, }, --  3
	{ 1, 1, 1, 2, 0, 0, 0, 0, 0, 0, 0, 0, }, --  4
	{ 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, }, --  5
	{ 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, }, --  6
	{ 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, }, --  7
	{ 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, }, --  8
	{ 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, }, --  9
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, }, -- 10
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 2, }, -- 11
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 2, }, -- 12
}

math.randomseed(os.time())
local reel = MakeReel(symset)
print("reel length: " .. #reel)
ShuffleReel(reel)
local iter = CorrectReel(reel, neighbours)
PrintReel(reel, iter)

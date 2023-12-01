-- local inspect = require("inspect")

require("busted.runner")()

function Mysplit(input_str, sep)
	if sep == nil then
		sep = "%s"
	end
	local t = {}
	for str in string.gmatch(input_str, "([^" .. sep .. "]+)") do
		table.insert(t, str)
	end
	return t
end

local part_1_total_score = 0
local example_total_score = 0
local part_2_total_score = 0

local scores_table_part_1 = {
	A = { X = 4, Y = 8, Z = 3 },
	B = { X = 1, Y = 5, Z = 9 },
	C = { X = 7, Y = 2, Z = 6 },
}

local scores_table_part_2 = {
	A = { X = 3, Y = 4, Z = 8 },
	B = { X = 1, Y = 5, Z = 9 },
	C = { X = 2, Y = 6, Z = 7 },
}

local function calculateRound(line, score_table)
	local opp, mine = table.unpack(Mysplit(line))
	return score_table[opp][mine]
end

local function day2(filename, score_table, score)
	for line in io.lines(filename) do
		score = score + calculateRound(line, score_table)
	end
	return score
end

local testRun = day2("./dev_input", scores_table_part_1, example_total_score)

describe("Day 2", function()
	it("Example assertion", function()
		assert.are.equal(15, testRun)
	end)
	it("Part 1", function()
		assert.are.equal(13526, day2("./real_input", scores_table_part_1, part_1_total_score))
	end)
	it("Part 2", function()
		assert.are.equal(14204, day2("./real_input", scores_table_part_2, part_2_total_score))
	end)
end)

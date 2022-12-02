local singleElf = { 0 }

for line in io.lines("./real_input") do
	if line ~= "" then
		singleElf[#singleElf] = singleElf[#singleElf] + tonumber(line)
	else
		table.insert(singleElf, 0)
	end
end

table.sort(singleElf, function(a, b)
	return a > b
end)

print("Top elf: " .. singleElf[1])
print("Top 3 elves: " .. singleElf[1] + singleElf[2] + singleElf[3])

local function part1(filename)
	local testElf = { 0 }
	for line in io.lines(filename) do
		if line ~= "" then
			testElf[#testElf] = testElf[#testElf] + tonumber(line)
		else
			table.insert(testElf, 0)
		end
	end

	table.sort(testElf, function(a, b)
		return a > b
	end)
	return testElf[1]
end

describe("Day 1 part 1", function()
	it("part 1", function()
		assert.are.equal(part1("./real_input"), 70698)
	end)
end)

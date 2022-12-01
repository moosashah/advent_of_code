local singleElf = { 0 }

for line in io.lines("./day1_input") do
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

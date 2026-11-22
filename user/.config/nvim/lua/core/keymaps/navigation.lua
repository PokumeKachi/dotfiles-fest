local map = require("utils.keymap").map

local function is_nonblank(line)
	return vim.fn.getline(line):match("%S") ~= nil
end

local function find_line(start_line, direction, predicate)
	local limit = direction == 1 and vim.fn.line("$") or 1

	for line = start_line + direction, limit, direction do
		if predicate(line) then
			return line
		end
	end
end

local function same_indent_motion(direction)
	return function()
		local current_line = vim.fn.line(".")
		local base_line = current_line

		-- Find a reference line if the current line is blank.
		if not is_nonblank(base_line) then
			base_line = find_line(base_line, direction, is_nonblank)

			if not base_line then
				return
			end
		end

		local base_indent = vim.fn.indent(base_line)

		-- Find the next nonblank line with matching indentation.
		local target_line = find_line(base_line, direction, function(line)
			return is_nonblank(line) and vim.fn.indent(line) == base_indent
		end)

		if target_line then
			vim.cmd("normal! " .. target_line .. "G")
		end
	end
end

map({ "n", "c", "x", "t" }, "<leader>]", same_indent_motion(1), {
	desc = "Next Same Indent",
})

map({ "n", "c", "x", "t" }, "<leader>[", same_indent_motion(-1), {
	desc = "Previous Same Indent",
})

map("n", "<Tab>", ":bnext<CR>")
map("n", "<S-tab>", ":bprev<CR>")
map("n", "<leader><Tab>", "<C-^>")

map("n", "<C-d>", "<C-d>zz")
map("n", "<C-u>", "<C-u>zz")

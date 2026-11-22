-- ~/.config/nvim/lua/plugins/codesnap.lua
local map = require("utils.keymap").lazy_map

return {
	"mistricky/codesnap.nvim",
	-- tag = "v2.1.1", -- pin to latest release
	build = "make",
	cmd = {
		"CodeSnap",
		"CodeSnapSave",
		"CodeSnapASCII",
		"CodeSnapHighlight",
		"CodeSnapSaveHighlight",
	},
	keys = {
		-- Quick access mappings (optional)
		map({ "x" }, "<leader>cs", "<cmd>CodeSnap<cr>", { desc = "📸 CodeSnap: Copy snapshot" }),
		map(
			{ "x" },
			"<leader>cS",
			"<cmd>CodeSnapSave<cr>",
			{ desc = "💾 CodeSnap: Save snapshot" }
		),
		map(
			{ "x" },
			"<leader>ca",
			"<cmd>CodeSnapASCII<cr>",
			{ desc = "🔤 CodeSnap: ASCII snapshot" }
		),
		map(
			{ "x" },
			"<leader>ch",
			"<cmd>CodeSnapHighlight<cr>",
			{ desc = "✨ CodeSnap: Highlight snapshot" }
		),
		map(
			{ "x" },
			"<leader>cH",
			"<cmd>CodeSnapSaveHighlight<cr>",
			{ desc = "✨ CodeSnap: Save highlight snapshot" }
		),
	},
	opts = {
		-- ─── General ────────────────────────────────────────────────────────
		show_line_number = true,
		highlight_color = "#ffffff20",
		show_workspace = true,

		-- ─── Snapshot config ───────────────────────────────────────────────
		snapshot_config = {
			-- 🖌️ Theme: uses the beautiful "One Hunter Vercel" VSCode theme
			--    via the Asset URL parser — no local files needed!
			-- theme = "vercel@https://raw.githubusercontent.com/Railly/one-hunter-vscode/refs/heads/main/themes/OneHunter-Vercel-color-theme.json",
			-- themes_folders = { "~/.config/codesnap/themes" },
			-- fonts_folders = { "~/.config/codesnap/fonts" },

			line_number_color = "#7C8A9C",

			-- 🪟 macOS-style window bar with soft shadow + border
			window = {
				mac_window_bar = true,
				shadow = {
					radius = 24,
					color = "#00000060",
				},
				margin = {
					x = 72,
					y = 72,
				},
				border = {
					width = 1,
					color = "#ffffff25",
				},
				title_config = {
					color = "#E4E9F0",
					font_family = "Pacifico",
				},
			},

			-- 🖥️ Terminal-style command output
			command_output_config = {
				prompt = "❯",
				font_family = "CaskaydiaCove Nerd Font",
				prompt_color = "#F78FB3",
				command_color = "#98C379",
				string_arg_color = "#E5C07B",
			},

			-- 💻 Code block styling
			code_config = {
				font_family = "CaskaydiaCove Nerd Font",
				breadcrumbs = {
					enable = true,
					separator = " › ",
					color = "#80848b",
					font_family = "CaskaydiaCove Nerd Font",
				},
			},

			-- 💧 Watermark
			watermark = {
				content = "✦ CodeSnap.nvim ✦",
				font_family = "Pacifico",
				color = "#ffffffcc",
			},

			-- 🌈 Beautiful 4-stop aurora gradient background
			background = {
				start = {
					x = 0,
					y = 0,
				},
				["end"] = {
					x = "max",
					y = "max",
				},
				stops = {
					{
						position = 0,
						color = "#EBECB2",
					},
					{
						position = 0.28,
						color = "#F3B0F7",
					},
					{
						position = 0.73,
						color = "#92B5F0",
					},
					{
						position = 0.94,
						color = "#AEF0F8",
					},
				},
			},
		},
	},
}

package main

import "github.com/emad-elsaid/fest"

func main() {
	fest.Main()
}

func init() {
	base()
	coreSystem()
	packageManagement()
	systemSecurity()
	filesystem()
	hardware()
	shells()
	editors()
	toolchains()
	kde()
	windowManagers()
	browsers()
	networking()
	audio()
	fonts()
	office()
	creative()
	virtualization()
	syncTools()
	media()
	monitoring()
	printing()
}

func base() {
}

func coreSystem() {
	fest.Package(
		"base",
		"base-devel",
		"linux-cachyos",
		"linux-cachyos-lts",
		"linux-cachyos-headers",
		"linux-cachyos-lts-headers",
		"linux-firmware",
		"linux-firmware-amdgpu",
		"linux-firmware-intel",
		"linux-firmware-nvidia",
		"linux-firmware-radeon",
		"amd-ucode",
		"intel-ucode",
		"limine",
		"limine-mkinitcpio-hook",
		"limine-snapper-sync",
		"mkinitcpio",
		"plymouth",
		"zram-generator",
	)
}

func packageManagement() {
	fest.Package(
		"pacman",
		"pacman-contrib",
		"paru",
		"reflector",
		"archlinux-keyring",
		"cachyos-keyring",
		"cachyos-mirrorlist",
		"cachyos-rate-mirrors",
		"rebuild-detector",
		"cachyos-hooks",
		"cachyos-settings",
		"cachyos-hello",
		"cachyos-kernel-manager",
		"cachyos-packageinstaller",
	)
}

func systemSecurity() {
	fest.Package(
		"systemd",
		"sudo",
		"polkit",
		"ufw",
		"ufw-extras",
		"gnupg",
		"pinentry",
		"pcsclite",
		"tpm2-tss",
	)
}

func filesystem() {
	fest.Package(
		"btrfs-progs",
		"snapper",
		"snap-pac",
		"exfatprogs",
		"dosfstools",
		"f2fs-tools",
		"xfsprogs",
		"nilfs-utils",
		"mdadm",
		"lvm2",
		"udisks2",
	)
}

func hardware() {
	fest.Package(
		"smartmontools",
		"hdparm",
		"dmidecode",
		"nvme-cli",
		"upower",
		"power-profiles-daemon",
		"thinkfan",
		"ananicy-cpp",
		"scx-scheds",
	)
}

func shells() {
	fest.Package(
		"fish",
		"fish-autopair",
		"fish-pure-prompt",
		"fisher",
		"kitty",
		"alacritty",
		"foot",
	)
}

func editors() {
	fest.Package(
		"neovim",
		"micro",
		"code",
		"qtcreator",
		"git",
		"github-cli",
		"lazygit",
		"just",
		"direnv",
		"sd",
		"stow",
	)
}

func toolchains() {
	fest.Package(
		"go",
		"rustup",
		"nodejs",
		"npm",
		"yarn",
		"deno",
		"dart",
		"fvm",
		"python",
		"python-pip",
		"php",
		"uv",
		"cmake",
		"jdk-openjdk",
	)
}

func kde() {
	fest.Package(
		"ark",
		"dolphin",
		"spectacle",
		"filelight",
	)
}

func windowManagers() {
	fest.Package(
		"niri",
		"dwl",
		"way-edges-git",
		"noctalia",
		"shelly",
		"fuzzel",
		"bemenu",
		"bemenu-wayland",
		"seatd",
		"swaybg",
	)
}

func browsers() {
	fest.Package(
		"firefox",
		"google-chrome",
		"chromium",
		"brave-origin-bin",
		"qutebrowser",
		"tor-browser-bin",
	)
}

func networking() {
	fest.Package(
		"networkmanager",
		"networkmanager-openvpn",
		"iwd",
		"wpa_supplicant",
		"openssh",
		"tailscale",
		"tor",
		"openvpn",
		"avahi",
		"samba",
		"nfs-utils",
		"cifs-utils",
		"bluez",
		"bluez-utils",
		"bluez-hid2hci",
	)
}

func audio() {
	fest.Package(
		"pipewire",
		"pipewire-alsa",
		"pipewire-audio",
		"pipewire-pulse",
		"wireplumber",
		"alsa-utils",
		"alsa-firmware",
		"pavucontrol",
		"easyeffects",
	)
}

func fonts() {
	fest.Package(
		"ttf-fantasque-nerd",
		"ttf-jetbrains-mono-nerd",
		"ttf-meslo-nerd",
		"ttf-fira-sans",
		"ttf-opensans",
		"noto-fonts-cjk",
		"noto-fonts-emoji",
		"powerline-fonts",
		"awesome-terminal-fonts",
		"cantarell-fonts",
	)
}

func office() {
	fest.Package(
		"libreoffice-fresh",
		"pandoc-cli",
		"zathura",
		"zathura-pdf-mupdf",
		"foliate",
		"zotero",
		"readest",
		"keepassxc",
		"bitwarden-bin",
		"ente-auth",
	)
}

func creative() {
	fest.Package(
		"blender",
		"krita",
		"inkscape",
		"freecad",
		"kicad",
		"kicad-library",
		"aseprite",
	)
}

func virtualization() {
	fest.Package(
		"qemu-full",
		"libvirt",
		"virt-manager",
		"virt-install",
		"edk2-ovmf",
		"spice-gtk",
	)
}

func syncTools() {
	fest.Package(
		"syncthing",
		"syncthingtray",
		"vesktop",
	)
}

func media() {
	fest.Package(
		"vlc",
		"qbittorrent",
		"zoom",
		"itch-bin",
		"prismlauncher-offline-git",
		"gpu-screen-recorder",
		"sunshine",
	)
}

func monitoring() {
	fest.Package(
		"btop",
		"glances",
		"fastfetch",
		"inxi",
		"hwinfo",
		"yt-dlp-git",
		"android-tools",
	)
}

func printing() {
	fest.Package(
		"cups",
		"cups-filters",
		"system-config-printer",
	)
}

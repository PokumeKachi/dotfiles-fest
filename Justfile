_default:
    @just --choose

save:
    go run . save

diff:
    go run . diff > .output 2>&1 && "${EDITOR:-vi}" .output

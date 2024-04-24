module github.com/igadmg/vector-demo

go 1.22.2

replace github.com/gen2brain/raylib-go/raylib => ./pkg/raylib-go/raylib

replace github.com/gen2brain/raylib-go/raygui => ./pkg/raylib-go/raygui

replace github.com/EliCDavis/vector => ./pkg/raylib-go/vector

require (
	github.com/EliCDavis/vector v1.6.0
	github.com/gen2brain/raylib-go/raygui v0.0.0-20240421191056-278df68f40bb
	github.com/gen2brain/raylib-go/raylib v0.0.0-20240421191056-278df68f40bb
)

require (
	github.com/ebitengine/purego v0.7.1 // indirect
	golang.org/x/exp v0.0.0-20240416160154-fe59bbe5cc7f // indirect
	golang.org/x/sys v0.19.0 // indirect
)

package main

import (
	"fmt"
	"image/color"
	"math"

	"github.com/EliCDavis/vector/mathex"
	"github.com/EliCDavis/vector/vector2"
	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	modes       []Mode
	currentMode *Mode

	vecA vector2.Float32 = vector2.New[float32](100, 0)
	vecB vector2.Float32 = vector2.New[float32](0, 200)
)

type Mode struct {
	Name    string
	DrawFn  rg.BoundsFn
	InputFn rg.BoundsFn
}

func TwoVectorInput(bounds rl.Rectangle) {
	if rl.IsMouseButtonDown(0) {
		xy := rl.GetMousePosition()
		if bounds.Contains(xy) {
			vecA = xy.Sub(bounds.Center())
		}
	}
	if rl.IsMouseButtonDown(1) {
		xy := rl.GetMousePosition()
		if bounds.Contains(xy) {
			vecB = xy.Sub(bounds.Center())
		}
	}
}

func main() {
	rl.InitWindow(800, 600, "Vector demo")

	modes = make([]Mode, 0)
	modes = append(modes, Mode{
		Name: "Project",
		DrawFn: func(bounds rl.Rectangle) {
			draw_vector(bounds, vecA, rl.Blue)
			draw_axis(bounds, vecB, 10, rl.Blue)

			draw_vector(bounds, vecA.Project(vecB), rl.Red)
		},
		InputFn: TwoVectorInput,
	})
	modes = append(modes, Mode{
		Name: "Reject",
		DrawFn: func(bounds rl.Rectangle) {
			draw_vector(bounds, vecA, rl.Blue)
			draw_axis(bounds, vecB, 10, rl.Blue)

			draw_vector(bounds, vecA.Reject(vecB), rl.Red)
		},
		InputFn: TwoVectorInput,
	})
	modes = append(modes, Mode{
		Name: "Reflect",
		DrawFn: func(bounds rl.Rectangle) {
			draw_vector(bounds, vecA, rl.Blue)
			draw_axis(bounds, vecB, 10, rl.Blue)

			draw_vector(bounds, vecA.Reflect(vecB.Normalized()), rl.Red)

		},
		InputFn: TwoVectorInput,
	})
	modes = append(modes, Mode{
		Name: "Angle",
		DrawFn: func(bounds rl.Rectangle) {
			draw_vector(bounds, vecA, rl.Green)
			draw_vector(bounds, vecB, rl.Red)

			a := vecA.Angle(vecB)
			rl.DrawTextLayout(rl.GetFontDefault(), fmt.Sprintf("%02f Degrees", a*180/math.Pi), 20, 2, rl.Blue,
				func(wh rl.Vector2) rl.Rectangle {
					cl := rg.CanvasLayout(bounds.ShrinkXYWH(8, 8, 8, 8))
					return cl.Layout(rl.AnchorCenter, rl.AnchorTopLeft, wh)
				})
		},
		InputFn: TwoVectorInput,
	})

	for !rl.WindowShouldClose() {
		draw_frame()
	}
}

func draw_frame() {
	rl.BeginDrawing()
	defer rl.EndDrawing()

	rl.ClearBackground(rl.GetColor((uint)(rg.GetStyle(rg.DEFAULT, rg.BACKGROUND_COLOR))))

	hl := rg.HorizontalLayout(rl.NewRectangle(0, 0, 800, 600), 0)
	rg.GroupBoxEx(hl.Pie(0.2).ShrinkXYWH(4, 8, 4, 4), "Sidebar", func(bounds rl.Rectangle) {
		vl := rg.VerticalLayout(bounds, 6)
		for _, mode := range modes {
			if rg.Button(vl.Layout(rl.NewVector2(0, 32), rg.JustifyFill), mode.Name) {
				currentMode = &mode
			}
		}
	})
	mode_bounds := hl.Fill(0, rg.JustifyFill).ShrinkXYWH(4, 4, 4, 4)
	draw_mode_background(mode_bounds)
	if currentMode != nil {
		currentMode.InputFn(mode_bounds)
		currentMode.DrawFn(mode_bounds)
	}
}

func draw_mode_background(bounds rl.Rectangle) {
	gs := bounds.Width() / mathex.Round(bounds.Width()/20)
	rg.Grid(bounds, "", gs, 3, nil)
}

func draw_vector(bounds rl.Rectangle, v vector2.Float32, col color.RGBA) {
	origin := bounds.Center()

	rl.DrawLineV(origin, v.Add(origin), col)
}

func draw_axis(bounds rl.Rectangle, v vector2.Float32, stripe float32, col color.RGBA) {
	origin := bounds.Center()

	sections := int(mathex.Ceil(2 * v.LengthF() / stripe))
	section := v.Normalized().ScaleF(stripe)
	a := v.Negated()
	for i := range sections {
		b := a.Add(section)

		if i%2 == 0 {
			rl.DrawLineV(a.Add(origin), b.Add(origin), col)
		}

		a = b
	}
}

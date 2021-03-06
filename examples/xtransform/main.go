package main

import (
	"math"

	"github.com/gabstv/ecs/v2"
	"github.com/gabstv/primen"
	"github.com/gabstv/primen/components"
	"github.com/gabstv/primen/core"
	"github.com/gabstv/primen/core/debug"
)

func main() {
	debug.Draw = true
	engine := primen.NewEngine(&primen.NewEngineInput{
		Width:     800,
		Height:    600,
		Resizable: true,
		OnReady:   ready,
	})
	engine.Run()
}

func ready(engine primen.Engine) {
	w := engine.NewWorldWithDefaults(0)
	tr := primen.NewRootNode(w)
	coretr := tr.Transform()
	coretr.SetX(400).SetY(300)
	tr2 := primen.NewChildNode(tr)
	coretr2 := tr2.Transform()
	coretr2.SetX(10).SetY(20)
	coretr.SetAngle(-math.Pi / 2)
	components.SetFunctionComponentData(w, tr.Entity(), components.Function{
		Update: func(ctx core.UpdateCtx, e ecs.Entity) {
			dd := components.GetTransformComponentData(w, e)
			dd.SetAngle(dd.Angle() + math.Pi*ctx.DT())
		},
	})
	//
	tr3 := primen.NewChildNode(tr2)
	coretr3 := tr3.Transform()
	coretr3.SetX(-40).SetY(50)
	//
	rnd := primen.NewRootNode(w)
	//
	components.SetFunctionComponentData(w, tr2.Entity(), components.Function{
		Update: func(ctx core.UpdateCtx, e ecs.Entity) {
			dd := components.GetTransformComponentData(w, e)
			dd.SetAngle(dd.Angle() - .5*math.Pi*ctx.DT())
			x, y, _ := components.GetTransformSystem(w).LocalToGlobal(0, 0, tr3.Entity())
			rnd.Transform().SetX(x - 10).SetY(y - 10)
		},
	})
}

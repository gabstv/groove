package primen

import (
	"image"
	"image/color"

	"github.com/gabstv/ecs"
	"github.com/gabstv/primen/core"
	"github.com/hajimehoshi/ebiten"
	"golang.org/x/image/font"
)

type Label struct {
	*WorldItem
	*TransformItem
	label     *core.Label
	drawLayer *core.DrawLayer
}

func NewLabel(w *ecs.World, fontFace font.Face, layer Layer, parent TransformGetter) *Label {
	lbl := &Label{}
	lbl.WorldItem = newWorldItem(w.NewEntity(), w)
	lbl.label = &core.Label{
		ScaleX: 1,
		ScaleY: 1,
		Color:  color.White,
		Face:   fontFace,
		Filter: ebiten.FilterDefault,
	}
	lbl.drawLayer = &core.DrawLayer{
		Layer:  layer,
		ZIndex: core.ZIndexTop,
	}
	lbl.TransformItem = newTransformItem(lbl.entity, lbl.world, parent)
	if err := w.AddComponentToEntity(lbl.entity, w.Component(core.CNDrawable), lbl.label); err != nil {
		panic(err)
	}
	if err := w.AddComponentToEntity(lbl.entity, w.Component(core.CNDrawLayer), lbl.drawLayer); err != nil {
		panic(err)
	}
	return lbl
}

func (e *Engine) NewLabel(fontFace font.Face, layer Layer, parent TransformGetter) *Label {
	return NewLabel(e.Default(), fontFace, layer, parent)
}

func (l *Label) SetText(t string) {
	if l.label.Text == t {
		return
	}
	l.label.Text = t
	l.label.SetDirty()
}

func (l *Label) Text() string {
	return l.label.Text
}

func (l *Label) SetArea(w, h int) {
	l.label.Area = image.Point{
		X: w,
		Y: h,
	}
	l.label.SetDirty()
}

func (l *Label) SetFilter(filter ebiten.Filter) {
	l.label.Filter = filter
	l.label.SetDirty()
}

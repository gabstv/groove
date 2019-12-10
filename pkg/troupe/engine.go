package troupe

import (
	"os"
	"path"
	"sort"
	"sync"
	"time"

	"github.com/gabstv/troupe/pkg/troupe/fs"
	osfs "github.com/gabstv/troupe/pkg/troupe/fs/os"
	"github.com/hajimehoshi/ebiten"
)

// Engine is what controls the ECS of troupe.
type Engine struct {
	lock         sync.Mutex
	lt           time.Time
	worlds       []worldContainer
	defaultWorld *World
	dmap         Dict
	options      EngineOptions
	f            fs.Filesystem
}

// NewEngineInput is the input data of NewEngine
type NewEngineInput struct {
	Width  int
	Height int
	Scale  float64
	Title  string
	FS     fs.Filesystem
}

// EngineOptions is used to setup Ebiten @ Engine.boot
type EngineOptions struct {
	Width  int
	Height int
	Scale  float64
	Title  string
}

// Options will create a EngineOptions struct to be used in
// an *Engine
func (i *NewEngineInput) Options() EngineOptions {
	opt := EngineOptions{
		Width:  i.Width,
		Height: i.Height,
		Scale:  i.Scale,
		Title:  i.Title,
	}
	return opt
}

// NewEngine returns a new Engine
func NewEngine(v *NewEngineInput) *Engine {
	fbase := ""
	if len(os.Args) > 0 {
		fbase = path.Dir(os.Args[0])
	}
	if v == nil {
		v = &NewEngineInput{
			Width:  800,
			Height: 600,
			Scale:  1,
			Title:  "Troupe",
			FS:     osfs.New(fbase),
		}
	} else {
		if v.Scale < 1 {
			v.Scale = 1
		}
		if v.Width <= 0 {
			v.Width = 320
		}
		if v.Height <= 0 {
			v.Height = 240
		}
		if v.FS == nil {
			v.FS = osfs.New(fbase)
		}
	}
	// assign the default systems and controllers

	e := &Engine{
		options: v.Options(),
		f:       v.FS,
	}

	// create the default world
	dw := NewWorld(e)

	e.worlds = []worldContainer{
		worldContainer{
			priority: 0,
			world:    dw,
		},
	}
	e.defaultWorld = dw

	// start default components and systems
	startDefaultComponents(e)
	startDefaultSystems(e)

	return e
}

// AddWorld adds a world to the engine.
// The priority is used to sort world execution, from hight to low.
func (e *Engine) AddWorld(w *World, priority int) {
	e.lock.Lock()
	defer e.lock.Unlock()
	if e.worlds == nil {
		e.worlds = make([]worldContainer, 0, 2)
	}
	e.worlds = append(e.worlds, worldContainer{
		priority: priority,
		world:    w,
	})
	// sort by priority
	sort.Sort(sortedWorldContainer(e.worlds))
}

// RemoveWorld removes a *World
func (e *Engine) RemoveWorld(w *World) bool {
	e.lock.Lock()
	defer e.lock.Unlock()
	wi := -1
	for k, ww := range e.worlds {
		if ww.world == w {
			wi = k
			ww.world = nil
			break
		}
	}
	if wi == -1 {
		return false
	}
	// splice
	e.worlds = append(e.worlds[:wi], e.worlds[wi+1:]...)
	if w == e.defaultWorld {
		e.defaultWorld = nil
	}
	return true
}

// Default world
func (e *Engine) Default() *World {
	return e.defaultWorld
}

// Run boots up the game engine
func (e *Engine) Run() error {
	e.lock.Lock()
	width, height, scale, title := e.options.Width, e.options.Height, e.options.Scale, e.options.Title
	e.lt = time.Now()
	e.lock.Unlock()
	return ebiten.Run(e.loop, width, height, scale, title)
}

func (e *Engine) loop(screen *ebiten.Image) error {
	e.lock.Lock()
	now := time.Now()
	ld := now.Sub(e.lt).Seconds()
	e.lt = now
	e.dmap.Set(TagDelta, ld)
	worlds := e.worlds
	e.lock.Unlock()

	for _, w := range worlds {
		w.world.RunWithoutTag(WorldTagDraw, screen, ld)
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	for _, w := range worlds {
		w.world.RunWithTag(WorldTagDraw, screen, ld)
	}

	return nil
}

// Get an item from the global map
func (e *Engine) Get(key string) interface{} {
	return e.dmap.Get(key)
}

// Set an item to the global map
func (e *Engine) Set(key string, value interface{}) {
	e.dmap.Set(key, value)
}

// FS returns the filesystem
func (e *Engine) FS() fs.Filesystem {
	return e.f
}
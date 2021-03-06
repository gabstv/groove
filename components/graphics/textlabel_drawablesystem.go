// Code generated by ecs https://github.com/gabstv/ecs; DO NOT EDIT.

package graphics

import (
    
    "sort"

    "github.com/gabstv/ecs/v2"
    
    "github.com/gabstv/primen/components"
    
)









const uuidDrawableTextLabelSystem = "70EC2F13-4C71-4A3F-9F6D-FF11F5DE9384"

type viewDrawableTextLabelSystem struct {
    entities []VIDrawableTextLabelSystem
    world ecs.BaseWorld
    
}

type VIDrawableTextLabelSystem struct {
    Entity ecs.Entity
    
    TextLabel *TextLabel 
    
    Transform *components.Transform 
    
}

type sortedVIDrawableTextLabelSystems []VIDrawableTextLabelSystem
func (a sortedVIDrawableTextLabelSystems) Len() int           { return len(a) }
func (a sortedVIDrawableTextLabelSystems) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortedVIDrawableTextLabelSystems) Less(i, j int) bool { return a[i].Entity < a[j].Entity }

func newviewDrawableTextLabelSystem(w ecs.BaseWorld) *viewDrawableTextLabelSystem {
    return &viewDrawableTextLabelSystem{
        entities: make([]VIDrawableTextLabelSystem, 0),
        world: w,
    }
}

func (v *viewDrawableTextLabelSystem) Matches() []VIDrawableTextLabelSystem {
    
    return v.entities
    
}

func (v *viewDrawableTextLabelSystem) indexof(e ecs.Entity) int {
    i := sort.Search(len(v.entities), func(i int) bool { return v.entities[i].Entity >= e })
    if i < len(v.entities) && v.entities[i].Entity == e {
        return i
    }
    return -1
}

// Fetch a specific entity
func (v *viewDrawableTextLabelSystem) Fetch(e ecs.Entity) (data VIDrawableTextLabelSystem, ok bool) {
    
    i := v.indexof(e)
    if i == -1 {
        return VIDrawableTextLabelSystem{}, false
    }
    return v.entities[i], true
}

func (v *viewDrawableTextLabelSystem) Add(e ecs.Entity) bool {
    
    
    // MUST NOT add an Entity twice:
    if i := v.indexof(e); i > -1 {
        return false
    }
    v.entities = append(v.entities, VIDrawableTextLabelSystem{
        Entity: e,
        TextLabel: GetTextLabelComponent(v.world).Data(e),
Transform: components.GetTransformComponentData(v.world, e),

    })
    if len(v.entities) > 1 {
        if v.entities[len(v.entities)-1].Entity < v.entities[len(v.entities)-2].Entity {
            sort.Sort(sortedVIDrawableTextLabelSystems(v.entities))
        }
    }
    return true
}

func (v *viewDrawableTextLabelSystem) Remove(e ecs.Entity) bool {
    
    
    if i := v.indexof(e); i != -1 {

        v.entities = append(v.entities[:i], v.entities[i+1:]...)
        return true
    }
    return false
}

func (v *viewDrawableTextLabelSystem) clearpointers() {
    
    
    for i := range v.entities {
        e := v.entities[i].Entity
        
        v.entities[i].TextLabel = nil
        
        v.entities[i].Transform = nil
        
        _ = e
    }
}

func (v *viewDrawableTextLabelSystem) rescan() {
    
    
    for i := range v.entities {
        e := v.entities[i].Entity
        
        v.entities[i].TextLabel = GetTextLabelComponent(v.world).Data(e)
        
        v.entities[i].Transform = components.GetTransformComponentData(v.world, e)
        
        _ = e
        
    }
}

// DrawableTextLabelSystem implements ecs.BaseSystem
type DrawableTextLabelSystem struct {
    initialized bool
    world       ecs.BaseWorld
    view        *viewDrawableTextLabelSystem
    enabled     bool
    
}

// GetDrawableTextLabelSystem returns the instance of the system in a World
func GetDrawableTextLabelSystem(w ecs.BaseWorld) *DrawableTextLabelSystem {
    return w.S(uuidDrawableTextLabelSystem).(*DrawableTextLabelSystem)
}

// Enable system
func (s *DrawableTextLabelSystem) Enable() {
    s.enabled = true
}

// Disable system
func (s *DrawableTextLabelSystem) Disable() {
    s.enabled = false
}

// Enabled checks if enabled
func (s *DrawableTextLabelSystem) Enabled() bool {
    return s.enabled
}

// UUID implements ecs.BaseSystem
func (DrawableTextLabelSystem) UUID() string {
    return "70EC2F13-4C71-4A3F-9F6D-FF11F5DE9384"
}

func (DrawableTextLabelSystem) Name() string {
    return "DrawableTextLabelSystem"
}

// ensure matchfn
var _ ecs.MatchFn = matchDrawableTextLabelSystem

// ensure resizematchfn
var _ ecs.MatchFn = resizematchDrawableTextLabelSystem

func (s *DrawableTextLabelSystem) match(eflag ecs.Flag) bool {
    return matchDrawableTextLabelSystem(eflag, s.world)
}

func (s *DrawableTextLabelSystem) resizematch(eflag ecs.Flag) bool {
    return resizematchDrawableTextLabelSystem(eflag, s.world)
}

func (s *DrawableTextLabelSystem) ComponentAdded(e ecs.Entity, eflag ecs.Flag) {
    if s.match(eflag) {
        if s.view.Add(e) {
            // TODO: dispatch event that this entity was added to this system
            
        }
    } else {
        if s.view.Remove(e) {
            // TODO: dispatch event that this entity was removed from this system
            
        }
    }
}

func (s *DrawableTextLabelSystem) ComponentRemoved(e ecs.Entity, eflag ecs.Flag) {
    if s.match(eflag) {
        if s.view.Add(e) {
            // TODO: dispatch event that this entity was added to this system
            
        }
    } else {
        if s.view.Remove(e) {
            // TODO: dispatch event that this entity was removed from this system
            
        }
    }
}

func (s *DrawableTextLabelSystem) ComponentResized(cflag ecs.Flag) {
    if s.resizematch(cflag) {
        s.view.rescan()
        
    }
}

func (s *DrawableTextLabelSystem) ComponentWillResize(cflag ecs.Flag) {
    if s.resizematch(cflag) {
        
        s.view.clearpointers()
    }
}

func (s *DrawableTextLabelSystem) V() *viewDrawableTextLabelSystem {
    return s.view
}

func (*DrawableTextLabelSystem) Priority() int64 {
    return 10
}

func (s *DrawableTextLabelSystem) Setup(w ecs.BaseWorld) {
    if s.initialized {
        panic("DrawableTextLabelSystem called Setup() more than once")
    }
    s.view = newviewDrawableTextLabelSystem(w)
    s.world = w
    s.enabled = true
    s.initialized = true
    
}


func init() {
    ecs.RegisterSystem(func() ecs.BaseSystem {
        return &DrawableTextLabelSystem{}
    })
}

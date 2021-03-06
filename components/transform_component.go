// Code generated by ecs https://github.com/gabstv/ecs; DO NOT EDIT.

package components

import (
    "sort"
    

    "github.com/gabstv/ecs/v2"
)








const uuidTransformComponent = "45E8849D-7EA9-4CDC-8AB1-86DB8705C253"
const capTransformComponent = 256

type drawerTransformComponent struct {
    Entity ecs.Entity
    Data   Transform
}

// WatchTransform is a helper struct to access a valid pointer of Transform
type WatchTransform interface {
    Entity() ecs.Entity
    Data() *Transform
}

type slcdrawerTransformComponent []drawerTransformComponent
func (a slcdrawerTransformComponent) Len() int           { return len(a) }
func (a slcdrawerTransformComponent) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a slcdrawerTransformComponent) Less(i, j int) bool { return a[i].Entity < a[j].Entity }


type mWatchTransform struct {
    c *TransformComponent
    entity ecs.Entity
}

func (w *mWatchTransform) Entity() ecs.Entity {
    return w.entity
}

func (w *mWatchTransform) Data() *Transform {
    
    
    id := w.c.indexof(w.entity)
    if id == -1 {
        return nil
    }
    return &w.c.data[id].Data
}

// TransformComponent implements ecs.BaseComponent
type TransformComponent struct {
    initialized bool
    flag        ecs.Flag
    world       ecs.BaseWorld
    wkey        [4]byte
    data        []drawerTransformComponent
    
}

// GetTransformComponent returns the instance of the component in a World
func GetTransformComponent(w ecs.BaseWorld) *TransformComponent {
    return w.C(uuidTransformComponent).(*TransformComponent)
}

// SetTransformComponentData updates/adds a Transform to Entity e
func SetTransformComponentData(w ecs.BaseWorld, e ecs.Entity, data Transform) {
    GetTransformComponent(w).Upsert(e, data)
}

// GetTransformComponentData gets the *Transform of Entity e
func GetTransformComponentData(w ecs.BaseWorld, e ecs.Entity) *Transform {
    return GetTransformComponent(w).Data(e)
}

// WatchTransformComponentData gets a pointer getter of an entity's Transform.
//
// The pointer must not be stored because it may become invalid overtime.
func WatchTransformComponentData(w ecs.BaseWorld, e ecs.Entity) WatchTransform {
    return &mWatchTransform{
        c: GetTransformComponent(w),
        entity: e,
    }
}

// UUID implements ecs.BaseComponent
func (TransformComponent) UUID() string {
    return "45E8849D-7EA9-4CDC-8AB1-86DB8705C253"
}

// Name implements ecs.BaseComponent
func (TransformComponent) Name() string {
    return "TransformComponent"
}

func (c *TransformComponent) indexof(e ecs.Entity) int {
    i := sort.Search(len(c.data), func(i int) bool { return c.data[i].Entity >= e })
    if i < len(c.data) && c.data[i].Entity == e {
        return i
    }
    return -1
}

// Upsert creates or updates a component data of an entity.
// Not recommended to be used directly. Use SetTransformComponentData to change component
// data outside of a system loop.
func (c *TransformComponent) Upsert(e ecs.Entity, data interface{}) {
    v, ok := data.(Transform)
    if !ok {
        panic("data must be Transform")
    }
    
    id := c.indexof(e)
    
    if id > -1 {
        
        dwr := &c.data[id]
        dwr.Data = v
        
        return
    }
    
    rsz := false
    if cap(c.data) == len(c.data) {
        rsz = true
        c.world.CWillResize(c, c.wkey)
        c.willresize()
    }
    newindex := len(c.data)
    c.data = append(c.data, drawerTransformComponent{
        Entity: e,
        Data:   v,
    })
    if len(c.data) > 1 {
        if c.data[newindex].Entity < c.data[newindex-1].Entity {
            c.world.CWillResize(c, c.wkey)
            c.willresize()
            sort.Sort(slcdrawerTransformComponent(c.data))
            rsz = true
        }
    }
    
    if rsz {
        c.resized()
        c.world.CResized(c, c.wkey)
        c.world.Dispatch(ecs.Event{
            Type: ecs.EvtComponentsResized,
            ComponentName: "TransformComponent",
            ComponentID: "45E8849D-7EA9-4CDC-8AB1-86DB8705C253",
        })
    }
    c.setupTransform(e)
    c.world.CAdded(e, c, c.wkey)
    c.world.Dispatch(ecs.Event{
        Type: ecs.EvtComponentAdded,
        ComponentName: "TransformComponent",
        ComponentID: "45E8849D-7EA9-4CDC-8AB1-86DB8705C253",
        Entity: e,
    })
}

// Remove a Transform data from entity e
//
// Warning: DO NOT call remove inside the system entities loop
func (c *TransformComponent) Remove(e ecs.Entity) {
    
    
    i := c.indexof(e)
    if i == -1 {
        return
    }
    
    //c.data = append(c.data[:i], c.data[i+1:]...)
    c.data = c.data[:i+copy(c.data[i:], c.data[i+1:])]
    c.world.CRemoved(e, c, c.wkey)
    c.removed(e)
    c.world.Dispatch(ecs.Event{
        Type: ecs.EvtComponentRemoved,
        ComponentName: "TransformComponent",
        ComponentID: "45E8849D-7EA9-4CDC-8AB1-86DB8705C253",
        Entity: e,
    })
}

func (c *TransformComponent) Data(e ecs.Entity) *Transform {
    
    
    index := c.indexof(e)
    if index > -1 {
        return &c.data[index].Data
    }
    return nil
}

// Flag returns the 
func (c *TransformComponent) Flag() ecs.Flag {
    return c.flag
}

// Setup is called by ecs.BaseWorld
//
// Do not call this directly
func (c *TransformComponent) Setup(w ecs.BaseWorld, f ecs.Flag, key [4]byte) {
    if c.initialized {
        panic("TransformComponent called Setup() more than once")
    }
    c.flag = f
    c.world = w
    c.wkey = key
    c.data = make([]drawerTransformComponent, 0, 256)
    c.initialized = true
    
}


func init() {
    ecs.RegisterComponent(func() ecs.BaseComponent {
        return &TransformComponent{}
    })
}

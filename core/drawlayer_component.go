// Code generated by ecs https://github.com/gabstv/ecs; DO NOT EDIT.

package core

import (
    "sort"
    

    "github.com/gabstv/ecs/v2"
)








const uuidDrawLayerComponent = "2D35C735-7275-4195-A61F-F559F8346D46"
const capDrawLayerComponent = 256

type drawerDrawLayerComponent struct {
    Entity ecs.Entity
    Data   DrawLayer
}

// WatchDrawLayer is a helper struct to access a valid pointer of DrawLayer
type WatchDrawLayer interface {
    Entity() ecs.Entity
    Data() *DrawLayer
}

type slcdrawerDrawLayerComponent []drawerDrawLayerComponent
func (a slcdrawerDrawLayerComponent) Len() int           { return len(a) }
func (a slcdrawerDrawLayerComponent) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a slcdrawerDrawLayerComponent) Less(i, j int) bool { return a[i].Entity < a[j].Entity }


type mWatchDrawLayer struct {
    c *DrawLayerComponent
    entity ecs.Entity
}

func (w *mWatchDrawLayer) Entity() ecs.Entity {
    return w.entity
}

func (w *mWatchDrawLayer) Data() *DrawLayer {
    
    
    id := w.c.indexof(w.entity)
    if id == -1 {
        return nil
    }
    return &w.c.data[id].Data
}

// DrawLayerComponent implements ecs.BaseComponent
type DrawLayerComponent struct {
    initialized bool
    flag        ecs.Flag
    world       ecs.BaseWorld
    wkey        [4]byte
    data        []drawerDrawLayerComponent
    
}

// GetDrawLayerComponent returns the instance of the component in a World
func GetDrawLayerComponent(w ecs.BaseWorld) *DrawLayerComponent {
    return w.C(uuidDrawLayerComponent).(*DrawLayerComponent)
}

// SetDrawLayerComponentData updates/adds a DrawLayer to Entity e
func SetDrawLayerComponentData(w ecs.BaseWorld, e ecs.Entity, data DrawLayer) {
    GetDrawLayerComponent(w).Upsert(e, data)
}

// GetDrawLayerComponentData gets the *DrawLayer of Entity e
func GetDrawLayerComponentData(w ecs.BaseWorld, e ecs.Entity) *DrawLayer {
    return GetDrawLayerComponent(w).Data(e)
}

// WatchDrawLayerComponentData gets a pointer getter of an entity's DrawLayer.
//
// The pointer must not be stored because it may become invalid overtime.
func WatchDrawLayerComponentData(w ecs.BaseWorld, e ecs.Entity) WatchDrawLayer {
    return &mWatchDrawLayer{
        c: GetDrawLayerComponent(w),
        entity: e,
    }
}

// UUID implements ecs.BaseComponent
func (DrawLayerComponent) UUID() string {
    return "2D35C735-7275-4195-A61F-F559F8346D46"
}

// Name implements ecs.BaseComponent
func (DrawLayerComponent) Name() string {
    return "DrawLayerComponent"
}

func (c *DrawLayerComponent) indexof(e ecs.Entity) int {
    i := sort.Search(len(c.data), func(i int) bool { return c.data[i].Entity >= e })
    if i < len(c.data) && c.data[i].Entity == e {
        return i
    }
    return -1
}

// Upsert creates or updates a component data of an entity.
// Not recommended to be used directly. Use SetDrawLayerComponentData to change component
// data outside of a system loop.
func (c *DrawLayerComponent) Upsert(e ecs.Entity, data interface{}) {
    v, ok := data.(DrawLayer)
    if !ok {
        panic("data must be DrawLayer")
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
        
    }
    newindex := len(c.data)
    c.data = append(c.data, drawerDrawLayerComponent{
        Entity: e,
        Data:   v,
    })
    if len(c.data) > 1 {
        if c.data[newindex].Entity < c.data[newindex-1].Entity {
            c.world.CWillResize(c, c.wkey)
            
            sort.Sort(slcdrawerDrawLayerComponent(c.data))
            rsz = true
        }
    }
    
    if rsz {
        
        c.world.CResized(c, c.wkey)
        c.world.Dispatch(ecs.Event{
            Type: ecs.EvtComponentsResized,
            ComponentName: "DrawLayerComponent",
            ComponentID: "2D35C735-7275-4195-A61F-F559F8346D46",
        })
    }
    
    c.world.CAdded(e, c, c.wkey)
    c.world.Dispatch(ecs.Event{
        Type: ecs.EvtComponentAdded,
        ComponentName: "DrawLayerComponent",
        ComponentID: "2D35C735-7275-4195-A61F-F559F8346D46",
        Entity: e,
    })
}

// Remove a DrawLayer data from entity e
//
// Warning: DO NOT call remove inside the system entities loop
func (c *DrawLayerComponent) Remove(e ecs.Entity) {
    
    
    i := c.indexof(e)
    if i == -1 {
        return
    }
    
    //c.data = append(c.data[:i], c.data[i+1:]...)
    c.data = c.data[:i+copy(c.data[i:], c.data[i+1:])]
    c.world.CRemoved(e, c, c.wkey)
    
    c.world.Dispatch(ecs.Event{
        Type: ecs.EvtComponentRemoved,
        ComponentName: "DrawLayerComponent",
        ComponentID: "2D35C735-7275-4195-A61F-F559F8346D46",
        Entity: e,
    })
}

func (c *DrawLayerComponent) Data(e ecs.Entity) *DrawLayer {
    
    
    index := c.indexof(e)
    if index > -1 {
        return &c.data[index].Data
    }
    return nil
}

// Flag returns the 
func (c *DrawLayerComponent) Flag() ecs.Flag {
    return c.flag
}

// Setup is called by ecs.BaseWorld
//
// Do not call this directly
func (c *DrawLayerComponent) Setup(w ecs.BaseWorld, f ecs.Flag, key [4]byte) {
    if c.initialized {
        panic("DrawLayerComponent called Setup() more than once")
    }
    c.flag = f
    c.world = w
    c.wkey = key
    c.data = make([]drawerDrawLayerComponent, 0, 256)
    c.initialized = true
    
}


func init() {
    ecs.RegisterComponent(func() ecs.BaseComponent {
        return &DrawLayerComponent{}
    })
}

// Code generated by ecs https://github.com/gabstv/ecs; DO NOT EDIT.

package core

import (
    "sort"
    

    "github.com/gabstv/ecs/v2"
)








const uuidSpriteComponent = "80C95DEC-DBBF-4529-BD27-739A69055BA0"
const capSpriteComponent = 256

type drawerSpriteComponent struct {
    Entity ecs.Entity
    Data   Sprite
}

// WatchSprite is a helper struct to access a valid pointer of Sprite
type WatchSprite interface {
    Entity() ecs.Entity
    Data() *Sprite
}

type slcdrawerSpriteComponent []drawerSpriteComponent
func (a slcdrawerSpriteComponent) Len() int           { return len(a) }
func (a slcdrawerSpriteComponent) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a slcdrawerSpriteComponent) Less(i, j int) bool { return a[i].Entity < a[j].Entity }


type mWatchSprite struct {
    c *SpriteComponent
    entity ecs.Entity
}

func (w *mWatchSprite) Entity() ecs.Entity {
    return w.entity
}

func (w *mWatchSprite) Data() *Sprite {
    
    
    id := w.c.indexof(w.entity)
    if id == -1 {
        return nil
    }
    return &w.c.data[id].Data
}

// SpriteComponent implements ecs.BaseComponent
type SpriteComponent struct {
    initialized bool
    flag        ecs.Flag
    world       ecs.BaseWorld
    wkey        [4]byte
    data        []drawerSpriteComponent
    
}

// GetSpriteComponent returns the instance of the component in a World
func GetSpriteComponent(w ecs.BaseWorld) *SpriteComponent {
    return w.C(uuidSpriteComponent).(*SpriteComponent)
}

// SetSpriteComponentData updates/adds a Sprite to Entity e
func SetSpriteComponentData(w ecs.BaseWorld, e ecs.Entity, data Sprite) {
    GetSpriteComponent(w).Upsert(e, data)
}

// GetSpriteComponentData gets the *Sprite of Entity e
func GetSpriteComponentData(w ecs.BaseWorld, e ecs.Entity) *Sprite {
    return GetSpriteComponent(w).Data(e)
}

// WatchSpriteComponentData gets a pointer getter of an entity's Sprite.
//
// The pointer must not be stored because it may become invalid overtime.
func WatchSpriteComponentData(w ecs.BaseWorld, e ecs.Entity) WatchSprite {
    return &mWatchSprite{
        c: GetSpriteComponent(w),
        entity: e,
    }
}

// UUID implements ecs.BaseComponent
func (SpriteComponent) UUID() string {
    return "80C95DEC-DBBF-4529-BD27-739A69055BA0"
}

// Name implements ecs.BaseComponent
func (SpriteComponent) Name() string {
    return "SpriteComponent"
}

func (c *SpriteComponent) indexof(e ecs.Entity) int {
    i := sort.Search(len(c.data), func(i int) bool { return c.data[i].Entity >= e })
    if i < len(c.data) && c.data[i].Entity == e {
        return i
    }
    return -1
}

// Upsert creates or updates a component data of an entity.
// Not recommended to be used directly. Use SetSpriteComponentData to change component
// data outside of a system loop.
func (c *SpriteComponent) Upsert(e ecs.Entity, data interface{}) {
    v, ok := data.(Sprite)
    if !ok {
        panic("data must be Sprite")
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
    c.data = append(c.data, drawerSpriteComponent{
        Entity: e,
        Data:   v,
    })
    if len(c.data) > 1 {
        if c.data[newindex].Entity < c.data[newindex-1].Entity {
            c.world.CWillResize(c, c.wkey)
            
            sort.Sort(slcdrawerSpriteComponent(c.data))
            rsz = true
        }
    }
    
    if rsz {
        
        c.world.CResized(c, c.wkey)
        c.world.Dispatch(ecs.Event{
            Type: ecs.EvtComponentsResized,
            ComponentName: "SpriteComponent",
            ComponentID: "80C95DEC-DBBF-4529-BD27-739A69055BA0",
        })
    }
    
    c.world.CAdded(e, c, c.wkey)
    c.world.Dispatch(ecs.Event{
        Type: ecs.EvtComponentAdded,
        ComponentName: "SpriteComponent",
        ComponentID: "80C95DEC-DBBF-4529-BD27-739A69055BA0",
        Entity: e,
    })
}

// Remove a Sprite data from entity e
//
// Warning: DO NOT call remove inside the system entities loop
func (c *SpriteComponent) Remove(e ecs.Entity) {
    
    
    i := c.indexof(e)
    if i == -1 {
        return
    }
    
    //c.data = append(c.data[:i], c.data[i+1:]...)
    c.data = c.data[:i+copy(c.data[i:], c.data[i+1:])]
    c.world.CRemoved(e, c, c.wkey)
    
    c.world.Dispatch(ecs.Event{
        Type: ecs.EvtComponentRemoved,
        ComponentName: "SpriteComponent",
        ComponentID: "80C95DEC-DBBF-4529-BD27-739A69055BA0",
        Entity: e,
    })
}

func (c *SpriteComponent) Data(e ecs.Entity) *Sprite {
    
    
    index := c.indexof(e)
    if index > -1 {
        return &c.data[index].Data
    }
    return nil
}

// Flag returns the 
func (c *SpriteComponent) Flag() ecs.Flag {
    return c.flag
}

// Setup is called by ecs.BaseWorld
//
// Do not call this directly
func (c *SpriteComponent) Setup(w ecs.BaseWorld, f ecs.Flag, key [4]byte) {
    if c.initialized {
        panic("SpriteComponent called Setup() more than once")
    }
    c.flag = f
    c.world = w
    c.wkey = key
    c.data = make([]drawerSpriteComponent, 0, 256)
    c.initialized = true
    c.onCompSetup()
}


func init() {
    ecs.RegisterComponent(func() ecs.BaseComponent {
        return &SpriteComponent{}
    })
}

// Code generated by ecs https://github.com/gabstv/ecs; DO NOT EDIT.

package layerexample

import (
    "sort"
    

    "github.com/gabstv/ecs/v2"
)








const uuidOrbitalMovementComponent = "DAD60C25-6B0D-4D3D-BF8E-5EB424FD8F1B"
const capOrbitalMovementComponent = 256

type drawerOrbitalMovementComponent struct {
    Entity ecs.Entity
    Data   OrbitalMovement
}

// WatchOrbitalMovement is a helper struct to access a valid pointer of OrbitalMovement
type WatchOrbitalMovement interface {
    Entity() ecs.Entity
    Data() *OrbitalMovement
}

type slcdrawerOrbitalMovementComponent []drawerOrbitalMovementComponent
func (a slcdrawerOrbitalMovementComponent) Len() int           { return len(a) }
func (a slcdrawerOrbitalMovementComponent) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a slcdrawerOrbitalMovementComponent) Less(i, j int) bool { return a[i].Entity < a[j].Entity }


type mWatchOrbitalMovement struct {
    c *OrbitalMovementComponent
    entity ecs.Entity
}

func (w *mWatchOrbitalMovement) Entity() ecs.Entity {
    return w.entity
}

func (w *mWatchOrbitalMovement) Data() *OrbitalMovement {
    
    
    id := w.c.indexof(w.entity)
    if id == -1 {
        return nil
    }
    return &w.c.data[id].Data
}

// OrbitalMovementComponent implements ecs.BaseComponent
type OrbitalMovementComponent struct {
    initialized bool
    flag        ecs.Flag
    world       ecs.BaseWorld
    wkey        [4]byte
    data        []drawerOrbitalMovementComponent
    
}

// GetOrbitalMovementComponent returns the instance of the component in a World
func GetOrbitalMovementComponent(w ecs.BaseWorld) *OrbitalMovementComponent {
    return w.C(uuidOrbitalMovementComponent).(*OrbitalMovementComponent)
}

// SetOrbitalMovementComponentData updates/adds a OrbitalMovement to Entity e
func SetOrbitalMovementComponentData(w ecs.BaseWorld, e ecs.Entity, data OrbitalMovement) {
    GetOrbitalMovementComponent(w).Upsert(e, data)
}

// GetOrbitalMovementComponentData gets the *OrbitalMovement of Entity e
func GetOrbitalMovementComponentData(w ecs.BaseWorld, e ecs.Entity) *OrbitalMovement {
    return GetOrbitalMovementComponent(w).Data(e)
}

// WatchOrbitalMovementComponentData gets a pointer getter of an entity's OrbitalMovement.
//
// The pointer must not be stored because it may become invalid overtime.
func WatchOrbitalMovementComponentData(w ecs.BaseWorld, e ecs.Entity) WatchOrbitalMovement {
    return &mWatchOrbitalMovement{
        c: GetOrbitalMovementComponent(w),
        entity: e,
    }
}

// UUID implements ecs.BaseComponent
func (OrbitalMovementComponent) UUID() string {
    return "DAD60C25-6B0D-4D3D-BF8E-5EB424FD8F1B"
}

// Name implements ecs.BaseComponent
func (OrbitalMovementComponent) Name() string {
    return "OrbitalMovementComponent"
}

func (c *OrbitalMovementComponent) indexof(e ecs.Entity) int {
    i := sort.Search(len(c.data), func(i int) bool { return c.data[i].Entity >= e })
    if i < len(c.data) && c.data[i].Entity == e {
        return i
    }
    return -1
}

// Upsert creates or updates a component data of an entity.
// Not recommended to be used directly. Use SetOrbitalMovementComponentData to change component
// data outside of a system loop.
func (c *OrbitalMovementComponent) Upsert(e ecs.Entity, data interface{}) {
    v, ok := data.(OrbitalMovement)
    if !ok {
        panic("data must be OrbitalMovement")
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
    c.data = append(c.data, drawerOrbitalMovementComponent{
        Entity: e,
        Data:   v,
    })
    if len(c.data) > 1 {
        if c.data[newindex].Entity < c.data[newindex-1].Entity {
            c.world.CWillResize(c, c.wkey)
            
            sort.Sort(slcdrawerOrbitalMovementComponent(c.data))
            rsz = true
        }
    }
    
    if rsz {
        
        c.world.CResized(c, c.wkey)
        c.world.Dispatch(ecs.Event{
            Type: ecs.EvtComponentsResized,
            ComponentName: "OrbitalMovementComponent",
            ComponentID: "DAD60C25-6B0D-4D3D-BF8E-5EB424FD8F1B",
        })
    }
    
    c.world.CAdded(e, c, c.wkey)
    c.world.Dispatch(ecs.Event{
        Type: ecs.EvtComponentAdded,
        ComponentName: "OrbitalMovementComponent",
        ComponentID: "DAD60C25-6B0D-4D3D-BF8E-5EB424FD8F1B",
        Entity: e,
    })
}

// Remove a OrbitalMovement data from entity e
//
// Warning: DO NOT call remove inside the system entities loop
func (c *OrbitalMovementComponent) Remove(e ecs.Entity) {
    
    
    i := c.indexof(e)
    if i == -1 {
        return
    }
    
    //c.data = append(c.data[:i], c.data[i+1:]...)
    c.data = c.data[:i+copy(c.data[i:], c.data[i+1:])]
    c.world.CRemoved(e, c, c.wkey)
    
    c.world.Dispatch(ecs.Event{
        Type: ecs.EvtComponentRemoved,
        ComponentName: "OrbitalMovementComponent",
        ComponentID: "DAD60C25-6B0D-4D3D-BF8E-5EB424FD8F1B",
        Entity: e,
    })
}

func (c *OrbitalMovementComponent) Data(e ecs.Entity) *OrbitalMovement {
    
    
    index := c.indexof(e)
    if index > -1 {
        return &c.data[index].Data
    }
    return nil
}

// Flag returns the 
func (c *OrbitalMovementComponent) Flag() ecs.Flag {
    return c.flag
}

// Setup is called by ecs.BaseWorld
//
// Do not call this directly
func (c *OrbitalMovementComponent) Setup(w ecs.BaseWorld, f ecs.Flag, key [4]byte) {
    if c.initialized {
        panic("OrbitalMovementComponent called Setup() more than once")
    }
    c.flag = f
    c.world = w
    c.wkey = key
    c.data = make([]drawerOrbitalMovementComponent, 0, 256)
    c.initialized = true
}


func init() {
    ecs.RegisterComponent(func() ecs.BaseComponent {
        return &OrbitalMovementComponent{}
    })
}
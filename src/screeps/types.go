package screeps

type LookAtResult []struct {
	Type   string
	Object RoomObject
}

type LookAtAreaResult []struct {
	X      int
	Y      int
	Type   string
	Object RoomObject
}

type LookForAtAreaResult []struct {
	X      int
	Y      int
	Object RoomObject
}

type FindPathResult []struct {
	x         int
	y         int
	dx        int
	dy        int
	direction DirectionConstant
}

type FindPathOpts struct {
	IgnoreCreeps                 *bool
	IgnoreDestructibleStructures *bool
	IgnoreRoads                  *bool
	//CostCallback - I don't know how to port that
	MaxOps          *uint
	HeuristicWeight *float32
	Serialize       *bool
	MaxRooms        *uint
	Range           *uint
	PlainCost       *uint
	SwampCost       *uint
}

type MoveToOpts struct {
	FindPathOpts

	ReusePath          *uint
	SerializeMemory    *bool
	NoPathFinding      *bool
	VisualizePathStyle *PolyStyle
}

type FindClosestByPathAlgorithm string

const (
	ALGORITHM_ASTAR    FindClosestByPathAlgorithm = "astar"
	ALGORITHM_DIJKSTRA FindClosestByPathAlgorithm = "dijkstra"
)

type FindClosestByPathOpts struct {
	FindPathOpts
	//Filter - I don't know how to port that
	Algorithm *FindClosestByPathAlgorithm
}

type Effect struct {
	Effect         EffectId
	Level          *int
	TicksRemaining int
}

type BodyPart struct {
	Boost *ResourceConstant
	Type  BodyPartConstant
	Hits  int
}

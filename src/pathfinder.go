package src

import . "screepsgo/screeps-go"

func pathfinderExample() {
	flag1, ok1 := Game.Flags()["Flag1"]
	flag2, ok2 := Game.Flags()["Flag2"]
	if ok1 && ok2 {
		goals := []PathFinderGoal{
			{
				Pos:   flag2.Pos(),
				Range: 0,
			},
		}
		var roomCb RoomCallback = func(roomName string) *CostMatrix {
			result := NewCostMatrix()

			flags := Game.Flags()
			for _, flag := range flags {
				if flag.Color() == COLOR_RED {
					result.Set(flag.Pos().X, flag.Pos().Y, 255)
				}
			}

			return &result
		}
		opts := PathFinderOpts{
			RoomCallback: &roomCb,
		}
		path := PathFinder.Search(flag1.Pos(), goals, &opts)
		if path.Incomplete {
			Console.Log("Flag1-Flag2 Path incomplete")
		} else {
			stepsInCurrentRoom := []RoomPosition{}
			for i, step := range path.Path {
				if i > 0 && step.RoomName != path.Path[i-1].RoomName {
					roomName := stepsInCurrentRoom[0].RoomName
					visual := NewRoomVisual(&roomName)
					visual.Poly(stepsInCurrentRoom, nil)
					stepsInCurrentRoom = []RoomPosition{}
				}

				stepsInCurrentRoom = append(stepsInCurrentRoom, step)
			}
			if len(stepsInCurrentRoom) > 0 {
				roomName := stepsInCurrentRoom[0].RoomName
				visual := NewRoomVisual(&roomName)
				visual.Poly(stepsInCurrentRoom, nil)
			}
		}
	}
}

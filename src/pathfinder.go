package main

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
			stepsByRoom := map[string][]RoomPosition{}
			for _, step := range path.Path {
				_, ok := stepsByRoom[step.RoomName]
				if !ok {
					stepsByRoom[step.RoomName] = []RoomPosition{}
				}
				stepsByRoom[step.RoomName] = append(stepsByRoom[step.RoomName], step)
			}

			for roomName, pathInRoom := range stepsByRoom {
				visual := NewRoomVisual(&roomName)
				visual.Poly(pathInRoom, nil)
			}
		}
	}
}

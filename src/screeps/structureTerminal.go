package screeps

import "syscall/js"

type StructureTerminal struct {
	OwnedStructure
}

func (terminal StructureTerminal) Cooldown() int {
	return terminal.ref.Get("cooldown").Int()
}

func (terminal StructureTerminal) Store() Store {
	var store = terminal.ref.Get("store")
	return Store{
		ref: store,
	}
}

func (terminal StructureTerminal) Send(resourceType ResourceConstant, amount int, destination string, description *string) ErrorCode {
	var jsDescription js.Value
	if description == nil {
		jsDescription = js.Null()
	} else {
		jsDescription = js.ValueOf(*description)
	}
	var result = terminal.ref.Call("send", string(resourceType), amount, destination, jsDescription).Int()
	return ErrorCode(result)
}

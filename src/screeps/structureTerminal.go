package screeps

import "syscall/js"

type StructureTerminal struct {
	OwnedStructure
}

func (terminal StructureTerminal) Cooldown() int {
	return terminal.ref.Get("cooldown").Int()
}

func (terminal StructureTerminal) Store() Store {
	jsStore := terminal.ref.Get("store")
	return Store{
		ref: jsStore,
	}
}

func (terminal StructureTerminal) Send(resourceType ResourceConstant, amount int, destination string, description *string) ErrorCode {
	var jsDescription js.Value
	if description == nil {
		jsDescription = js.Null()
	} else {
		jsDescription = js.ValueOf(*description)
	}
	result := terminal.ref.Call("send", string(resourceType), amount, destination, jsDescription).Int()
	return ErrorCode(result)
}

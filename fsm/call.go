// +build callback

package main

import (
	"fmt"
	"github.com/looplab/fsm"
)

type Door struct {
	To  string
	FSM *fsm.FSM
}

func NewDoor(to string) *Door {
	d := &Door{
		To: to,
	}

	d.FSM = fsm.NewFSM(
		"closed",
		fsm.Events{
			{Name: "open", Src: []string{"closed"}, Dst: "open"},
			{Name: "close", Src: []string{"open"}, Dst: "closed"},
		},
		fsm.Callbacks{
			"enter_state": func(e *fsm.Event) { d.enterState(e) },
		},
	)

	return d
}

func (d *Door) enterState(e *fsm.Event) {
	fmt.Printf("The door to %s is %s\n", d.To, e.Dst)
}

func main() {
	fsm := fsm.NewFSM(
		"green",
		fsm.Events{
			//状态事件的名称   该事件的起始状态Src         该事件的结束状态Dst
			//即：状态事件warn（警告事件）表示事物的状态从状态green到状态yellow
			{Name: "warn", Src: []string{"green"}, Dst: "yellow"},
			{Name: "panic", Src: []string{"yellow"}, Dst: "red"},
			{Name: "calm", Src: []string{"red"}, Dst: "yellow"},
		},
		fsm.Callbacks{
			//表示任一事件
			"before_event": func(e *fsm.Event) {
				fmt.Println("before_event")
			},
			"leave_state": func(e *fsm.Event) {
				fmt.Println("leave_state")
			},
			// //根据自定义状态或事件所定义的状态事件函数
			// "before_yellow": func(e *fsm.Event) {
			// 	fmt.Println("before_yellow")
			// },
			"before_warn": func(e *fsm.Event) {
				fmt.Println("before_warn")
			},
		},
	)

	//打印当前状态，输出是默认状态green
	fmt.Println(fsm.Current())
	//触发warn状态事件，状态将会从green转变到yellow
	//本地触发"before_warn"、"before_event"、"leave_state"函数
	fsm.Event("warn")
	//打印当前状态，输出状态是yellow
	fmt.Println(fsm.Current())

}

// NewFSM constructs a FSM from events and callbacks.
//
// The events and transitions are specified as a slice of Event structs
// specified as Events. Each Event is mapped to one or more internal
// transitions from Event.Src to Event.Dst.
//
// Callbacks are added as a map specified as Callbacks where the key is parsed
// as the callback event as follows, and called in the same order:
//
// 1. before_<EVENT> - called before event named <EVENT>
//
// 2. before_event - called before all events
//
// 3. leave_<OLD_STATE> - called before leaving <OLD_STATE>
//
// 4. leave_state - called before leaving all states
//
// 5. enter_<NEW_STATE> - called after entering <NEW_STATE>
//
// 6. enter_state - called after entering all states
//
// 7. after_<EVENT> - called after event named <EVENT>
//
// 8. after_event - called after all events
//
// There are also two short form versions for the most commonly used callbacks.
// They are simply the name of the event or state:
//
// 1. <NEW_STATE> - called after entering <NEW_STATE>
//
// 2. <EVENT> - called after event named <EVENT>
//
// If both a shorthand version and a full version is specified it is undefined
// which version of the callback will end up in the internal map. This is due
// to the psuedo random nature of Go maps. No checking for multiple keys is
// currently performed.

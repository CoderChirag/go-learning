package enums

import "fmt"
type ServerState int
const (
	IDLE ServerState = iota
	CONNECTED
	ERROR
	RETRYING
)
var StateEnum = map[ServerState]string {
	IDLE: "idle",
	CONNECTED: "connected",
	ERROR: "error",
	RETRYING: "retrying",
}
// Implementing Stringer Interface
func (ss ServerState) String() string {
	return StateEnum[ss]
}
func Transition(s ServerState) ServerState {
	switch s {
	case IDLE:
		return CONNECTED
	case CONNECTED, RETRYING:
		return IDLE
	case ERROR:
		return ERROR
	default:
		panic(fmt.Errorf("unknown state %s", s))
	}
}
func EnumExample() {
	ns := Transition(IDLE)
	fmt.Println(ns)
	ns2 := Transition(ns)
	fmt.Println(ns2)
}
package states

var StateMachineInstance *StateMachine

type StateMachine struct {
	States map[StateUser]State
}

func NewStateMachine() *StateMachine {
	return &StateMachine{
		States: make(map[StateUser]State),
	}
}

func InitStateMachine() {
	StateMachineInstance = NewStateMachine()
}

func (sm *StateMachine) AddState(user StateUser, state State) {
	sm.States[user] = state
}

func (sm *StateMachine) RemoveState(user StateUser) {
	delete(sm.States, user)
}

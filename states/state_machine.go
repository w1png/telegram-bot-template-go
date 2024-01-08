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

func (sm *StateMachine) GetState(user StateUser) (State, bool) {
	state, ok := sm.States[user]
	return state, ok
}

func (sm *StateMachine) SetState(user StateUser, state State) {
	sm.States[user] = state
}

func (sm *StateMachine) DeleteState(user StateUser) {
	delete(sm.States, user)
}

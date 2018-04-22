// Code generated by go generate; DO NOT EDIT
// This file was generated by robots at 2018-04-22 21:34:55.734459881 +0000 UTC
package main

type GhostState int

type GhostStimulus int

const (
	WanderTheMaze     GhostState = 0
	ChasePacman       GhostState = 1
	RunAwayFromPacman GhostState = 2
	ReturnToBase      GhostState = 3
)

const (
	SightPacman         GhostStimulus = 0
	PacmanEatsPowerPill GhostStimulus = 1
	LoseSightOfPacman   GhostStimulus = 2
	GetEatenByPacman    GhostStimulus = 3
	PowerPillWearsOff   GhostStimulus = 4
	ReachTheBase        GhostStimulus = 5
)

type GhostStateMachine struct {
	CurrentState GhostState
	Data         interface{}
}

type fsmError struct {
	s string
}

func (e *fsmError) Error() string {
	return e.s
}

func NewGhostStateMachine(data interface{}) *GhostStateMachine {
	return &GhostStateMachine{WanderTheMaze, data}
}

func GetGhostStimulusString(stimulus GhostStimulus) string {
	switch stimulus {
	case SightPacman:
		return "SightPacman"
	case PacmanEatsPowerPill:
		return "PacmanEatsPowerPill"
	case LoseSightOfPacman:
		return "LoseSightOfPacman"
	case GetEatenByPacman:
		return "GetEatenByPacman"
	case PowerPillWearsOff:
		return "PowerPillWearsOff"
	case ReachTheBase:
		return "ReachTheBase"
	default:
		return ""
	}
}

func GetGhostStateString(state GhostState) string {
	switch state {
	case WanderTheMaze:
		return "WanderTheMaze"
	case ChasePacman:
		return "ChasePacman"
	case RunAwayFromPacman:
		return "RunAwayFromPacman"
	case ReturnToBase:
		return "ReturnToBase"
	default:
		return ""
	}
}

func (fsm *GhostStateMachine) MoveNextState(stimulus GhostStimulus) (GhostState, error) {
	nextState, err := fsm.GetNextState(stimulus)
	if err != nil {
		return fsm.CurrentState, err
	}
	cbs := GetOnExitCallback(fsm.CurrentState)
	for _, cb := range cbs {
		cb(fsm.Data)
	}
	if fsm.CurrentState == nextState {
		cbs = GetOnLoopCallback(fsm.CurrentState)
		for _, cb := range cbs {
			cb(fsm.Data)
		}
	} else {
		cbs = GetOnEntryCallback(nextState)
		for _, cb := range cbs {
			cb(fsm.Data)
		}
	}
	fsm.CurrentState = nextState
	return nextState, nil
}

func (fsm *GhostStateMachine) GetNextState(stimulus GhostStimulus) (GhostState, error) {
	switch fsm.CurrentState {
	case WanderTheMaze:
		switch stimulus {
		case SightPacman:
			if pacmanClose(fsm.Data) == true {
				return ChasePacman, nil
			}
			return fsm.CurrentState, &fsmError{"Conditional callback returned false"}
		case PacmanEatsPowerPill:
			return RunAwayFromPacman, nil
		default:
			return fsm.CurrentState, &fsmError{"Invalid stimulus"}
		}
	case ChasePacman:
		switch stimulus {
		case LoseSightOfPacman:
			return WanderTheMaze, nil
		case PacmanEatsPowerPill:
			return RunAwayFromPacman, nil
		default:
			return fsm.CurrentState, &fsmError{"Invalid stimulus"}
		}
	case RunAwayFromPacman:
		switch stimulus {
		case GetEatenByPacman:
			return ReturnToBase, nil
		case PowerPillWearsOff:
			return WanderTheMaze, nil
		default:
			return fsm.CurrentState, &fsmError{"Invalid stimulus"}
		}
	case ReturnToBase:
		switch stimulus {
		case ReachTheBase:
			return WanderTheMaze, nil
		default:
			return fsm.CurrentState, &fsmError{"Invalid stimulus"}
		}
	default:
		return fsm.CurrentState, &fsmError{"Invalid stimulus"}
	}
}

func GetOnEntryCallback(state GhostState) []func(interface{}) {
	switch state {
	case WanderTheMaze:
		return []func(interface{}){walkAround}
	case ChasePacman:
		return []func(interface{}){chase}
	case RunAwayFromPacman:
		return []func(interface{}){run}
	case ReturnToBase:
		return []func(interface{}){reset}
	default:
		return []func(interface{}){}
	}
}

func GetOnExitCallback(state GhostState) []func(interface{}) {
	switch state {
	default:
		return []func(interface{}){}
	}
}

func GetOnLoopCallback(state GhostState) []func(interface{}) {
	switch state {
	default:
		return []func(interface{}){}
	}
}

package character

type CharacterState uint8

const (
	Idle CharacterState = iota
	Attacking
	Running
	JumpPrep
	JumpAscent
	JumpDescent
	JumpLanding
	JumpReload
)

var _CharacterStateNames = [...]string{
	"Idle",
	"Attacking",
	"Running",
	"JumpPrep",
	"JumpAscent",
	"JumpDescent",
	"JumpLanding",
	"JumpReload",
}

func (s CharacterState) String() string {
	return _CharacterStateNames[s]
}

func (s CharacterState) Tiles() (first, last, repeat int) {
	repeat = -1
	switch s {
	case Idle:
		first = 0
		last = 5
	case Running:
		first = 15
		last = 23
	case JumpPrep:
		first = 24
		last = 25
		repeat = 0
	case JumpAscent:
		first = 26
		last = 29
	case JumpDescent:
		first = 33
		last = 36
	case JumpLanding:
		first = 37
		last = 39
		repeat = 0
	case JumpReload:
		first = 30
		last = 32
		repeat = 0
	}
	return
}

func (s CharacterState) Next() CharacterState {
	next := s
	switch s {
	case JumpPrep:
		next = JumpAscent
	case JumpLanding:
		next = Idle
	case JumpReload:
		next = JumpAscent
	}
	return next
}

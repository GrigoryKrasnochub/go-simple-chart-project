package calc

const (
	Minimize Type = "min"
	Maximize Type = "max"
)

type Type string

func (t Type) String() string {
	return string(t)
}

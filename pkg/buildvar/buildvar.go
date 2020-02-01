package buildvar

type Op int

const (
    Set Op = iota
    SetDefault
    Append
)

type T struct {
	name, value string
	operation Op;
}

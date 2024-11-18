package dnode

type DValue interface {
	val()
}

type DIntVal struct {
	Val int64
}

func (DIntVal) val() {}

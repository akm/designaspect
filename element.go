package designaspect

type Element interface {
	InService()
	InMethod()
	InPayload()
	InHTTP()
}

type ElementBase struct {
	inMethod  func()
	inPayload func()
	inHTTP    func()
}

func NewElementBase(inMethod, inPayload, inHTTP func()) *ElementBase {
	return &ElementBase{
		inMethod:  inMethod,
		inPayload: inPayload,
		inHTTP:    inHTTP,
	}
}

// Service内で呼ばれた場合は InMethod を呼び出す
func (el *ElementBase) InService() {
	el.InMethod()
}

func (el *ElementBase) InMethod() {
	if el.inMethod != nil {
		el.inMethod()
	}
}

func (el *ElementBase) InPayload() {
	if el.inPayload != nil {
		el.inPayload()
	}
}

func (el *ElementBase) InHTTP() {
	if el.inHTTP != nil {
		el.inHTTP()
	}
}

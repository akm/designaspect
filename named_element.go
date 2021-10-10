package designaspect

type NamedElement struct {
	Name     string
	required bool // in payload
	*ElementBase
}

func NewNamedElement(name string, inMethod, inPayload, inHTTP func(string)) *NamedElement {
	r := &NamedElement{Name: name}
	r.ElementBase = NewElementBase(r.bind(inMethod), r.bind(inPayload), r.bind(inHTTP))
	return r
}

func (el *NamedElement) bind(src func(string)) func() {
	if src == nil {
		return nil
	}
	return func() {
		src(el.Name)
	}
}

func (el *NamedElement) GetName() string {
	return el.Name
}

func (el *NamedElement) IsRequired() bool {
	return el.required
}

func (el *NamedElement) Clone() *NamedElement {
	r := *el
	return &r
}

func (el *NamedElement) Required() *NamedElement {
	r := el.Clone()
	r.required = true
	return r
}

func (el *NamedElement) Optional() *NamedElement {
	r := el.Clone()
	r.required = false
	return r
}

func (el *NamedElement) WithInMethod(f func(name string)) *NamedElement {
	el.ElementBase.inMethod = el.bind(f)
	return el
}

func (el *NamedElement) WithInPayload(f func(name string)) *NamedElement {
	el.ElementBase.inPayload = el.bind(f)
	return el
}

func (el *NamedElement) WithInHTTP(f func(name string)) *NamedElement {
	el.ElementBase.inHTTP = el.bind(f)
	return el
}

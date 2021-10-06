package designaspect

type Elements []Element

func AppendElements(elementsSlice ...Elements) Elements {
	r := Elements{}
	for _, elements := range elementsSlice {
		if elements != nil {
			r = append(r, elements...)
		}
	}
	return r
}

func (s Elements) InService() {
	for _, i := range s {
		i.InService()
	}
}

func (s Elements) InMethod() {
	for _, i := range s {
		i.InMethod()
	}
}

func (s Elements) InPayload() {
	for _, i := range s {
		i.InPayload()
	}
}

func (s Elements) InHTTP() {
	for _, i := range s {
		i.InHTTP()
	}
}

func (s Elements) RequiredNames() []string {
	r := []string{}
	for _, i := range s {
		if v, ok := i.(*NamedElement); ok {
			if v.IsRequired() {
				r = append(r, v.Name)
			}
		}
	}
	return r
}

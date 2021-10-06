package designaspect

type MethodMap map[string]*Method

func (mm MethodMap) Define(names ...string) {
	for _, name := range names {
		if m, ok := mm[name]; ok && m != nil {
			m.Define()
		}
	}
}

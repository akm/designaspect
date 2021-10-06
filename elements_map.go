package designaspect

type ElementsMap map[string]Elements

func (em ElementsMap) Merge(other ElementsMap) ElementsMap {
	for key, elements := range other {
		em[key] = AppendElements(em[key], elements)
	}
	return em
}

func ElementsMapWithMultiKeys(element Element, keys ...string) ElementsMap {
	r := ElementsMap{}
	for _, key := range keys {
		r[key] = Elements{element}
	}
	return r
}

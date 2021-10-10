package designaspect

type AttributeElement interface {
	GetName() string
	IsRequired() bool
}

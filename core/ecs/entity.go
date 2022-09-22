package ecs

type Entity interface {
	Id() int
	Components() *ComponentCollection
}

type SimpleEntity struct {
	id         int
	components *ComponentCollection
}

func NewSimpleEntity(id int) *SimpleEntity {
	return &SimpleEntity{id: id, components: NewEmptyComponentCollection()}
}

func (s *SimpleEntity) Id() int {
	return s.id
}

func (s *SimpleEntity) Components() *ComponentCollection {
	return s.components
}

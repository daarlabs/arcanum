package auth

import "github.com/daarlabs/arcanum/crest"

var (
	Entity = crest.Entity[UserEntity]()
)

type UserEntity struct {
	crest.EntityBuilder
}

func (e UserEntity) Table() string {
	return "users"
}

func (e UserEntity) Alias() string {
	return "u"
}

func (e UserEntity) Fields() []crest.Field {
	return []crest.Field{
		e.Id(),
	}
}

func (e UserEntity) Id() crest.Field {
	return e.Field(crest.Id).Serial().PrimaryKey()
}

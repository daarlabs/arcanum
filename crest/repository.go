package crest

import (
	"github.com/daarlabs/arcanum/quirk"
)

type RepositoryManager[E entity] interface {
	Find(builders ...QueryBuilder) FindRepository
	Save(builders ...QueryBuilder) SaveRepository
	Remove(builders ...QueryBuilder) RemoveRepository
}

type repository[E entity] struct {
	db     *quirk.DB
	entity *E
}

type result interface{}

func Repository[E entity](db *quirk.DB) RepositoryManager[E] {
	return &repository[E]{
		db:     db,
		entity: Entity[E](),
	}
}

func (r *repository[E]) Find(builders ...QueryBuilder) FindRepository {
	tree := createTree(builders...)
	return &findRepository[E]{
		repository:    r,
		filters:       tree.filters,
		relationships: tree.relationships,
		selectors:     tree.selectors,
		shapes:        tree.shapes,
	}
}

func (r *repository[E]) Save(builders ...QueryBuilder) SaveRepository {
	tree := createTree(builders...)
	return &saveRepository[E]{
		repository:    r,
		filters:       tree.filters,
		relationships: tree.relationships,
		selectors:     tree.selectors,
		temporaries:   tree.temporaries,
		values:        tree.values,
	}
}

func (r *repository[E]) Remove(builders ...QueryBuilder) RemoveRepository {
	tree := createTree(builders...)
	return &removeRepository[E]{
		repository: r,
		filters:    tree.filters,
		selectors:  tree.selectors,
	}
}

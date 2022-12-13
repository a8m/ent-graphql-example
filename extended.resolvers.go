package todo

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"todo/ent"
)

// CreateChildren is the resolver for the createChildren field.
func (r *createTodoInputResolver) CreateChildren(ctx context.Context, obj *ent.CreateTodoInput, data []*ent.CreateTodoInput) error {
	c := ent.FromContext(ctx)
	builders := make([]*ent.TodoCreate, len(data))
	for i := range data {
		builders[i] = c.Todo.Create().SetInput(*data[i])
	}
	todos, err := c.Todo.CreateBulk(builders...).Save(ctx)
	if err != nil {
		return err
	}
	ids := make([]int, len(todos))
	for i := range todos {
		ids[i] = todos[i].ID
	}
	obj.ChildIDs = append(obj.ChildIDs, ids...)
	return nil
}

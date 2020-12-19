package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	//"log"
	"math/rand"

	"github.com/furiko/gqlgen-todo/graph/generated"
	"github.com/furiko/gqlgen-todo/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Text:   input.Text,
		ID:     fmt.Sprintf("%d", rand.Intn(1000000)),
		Done: false,
		UserID: input.UserID,
	}
	log.Printf("todo id: %v\n", todo.ID)
	err := r.DB.Insert(todo)
	if err != nil {
		log.Fatalln(err)
	}
	//r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, input model.DeleteTodo) (*model.Todo, error) {
	var index int
	for i, v := range r.todos {
		if v.ID == input.ID {
			index = i
		}
	}
	todo := r.todos[index]
	r.todos = append(r.todos[0:index], r.todos[index+1:]...)
	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	var todos []*model.Todo
	_, err := r.DB.Select(&todos, "SELECT * FROM todos")
	if err != nil {
		log.Fatalln(err)
	}
	for _, v := range todos {
		log.Println(v)
	}
	r.todos = todos
	return r.todos, nil
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }

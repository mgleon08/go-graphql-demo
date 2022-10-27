package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"example/graph/generated"
	"example/graph/model"
	"example/internal/links"
	"example/internal/users"
	"fmt"
	"strconv"
)

// CreateLink is the resolver for the createLink field.
func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	// panic(fmt.Errorf("not implemented: CreateLink - createLink"))

	var link links.Link
	link.Title = input.Title
	link.Address = input.Address
	link.Userid = input.Userid
	linkID := link.Save()
	return &model.Link{ID: strconv.FormatInt(linkID, 10), Title: link.Title, Address: link.Address}, nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	// panic(fmt.Errorf("not implemented: CreateUser - createUser"))

	var user users.User
	user.Username = input.Username
	user.Password = input.Password
	userID := user.Save()
	return &model.User{ID: strconv.FormatInt(userID, 10), Name: user.Username}, nil
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented: Login - login"))
}

// RefreshToken is the resolver for the refreshToken field.
func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented: RefreshToken - refreshToken"))
}

// Links is the resolver for the links field.
func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	var resultLinks []*model.Link
	var dbLinks []links.Link
	dbLinks = links.GetAll()
	for _, link := range dbLinks {
		user := &model.User{ID: link.User.ID, Name: link.User.Username}
		resultLinks = append(resultLinks, &model.Link{ID: link.ID, Title: link.Title, Address: link.Address, User: user})
	}
	return resultLinks, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, count *int) ([]*model.User, error) {
	var resultUsers []*model.User
	var dbUsers []users.User
	dbUsers = users.GetAll()
	for _, user := range dbUsers {
		resultUsers = append(resultUsers, &model.User{ID: user.ID, Name: user.Username})
	}

	resultUsers = resultUsers[0:*count]
	return resultUsers, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

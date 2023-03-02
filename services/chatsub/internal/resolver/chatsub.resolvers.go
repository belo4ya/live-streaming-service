package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.25

import (
	"context"
	"fmt"

	"github.com/belo4ya/live-streaming-service/api/chatsub/v1"
)

// SendMessage is the resolver for the sendMessage field.
func (r *mutationResolver) SendMessage(ctx context.Context, input *v1.NewMessage) (*v1.Message, error) {
	panic(fmt.Errorf("not implemented: SendMessage - sendMessage"))
}

// History is the resolver for the history field.
func (r *queryResolver) History(ctx context.Context, offset *int) ([]*v1.Message, error) {
	panic(fmt.Errorf("not implemented: History - history"))
}

// NewMessage is the resolver for the newMessage field.
func (r *subscriptionResolver) NewMessage(ctx context.Context, channelID string) (<-chan *v1.Message, error) {
	panic(fmt.Errorf("not implemented: NewMessage - newMessage"))
}

// Mutation returns v1.MutationResolver implementation.
func (r *Resolver) Mutation() v1.MutationResolver { return &mutationResolver{r} }

// Query returns v1.QueryResolver implementation.
func (r *Resolver) Query() v1.QueryResolver { return &queryResolver{r} }

// Subscription returns v1.SubscriptionResolver implementation.
func (r *Resolver) Subscription() v1.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }

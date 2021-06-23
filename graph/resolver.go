//go:generate go run github.com/99designs/gqlgen

package graph

import "ent-graphql/ent"

type Resolver struct {
	entClient *ent.Client
}

func NewResolver(entClient *ent.Client) *Resolver {
	return &Resolver{entClient: entClient}
}



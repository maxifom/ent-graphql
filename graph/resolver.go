//go:generate go run ../scripts/gqlgen.go

package graph

import "ent-graphql/ent"

type Resolver struct {
	entClient *ent.Client
}

func NewResolver(entClient *ent.Client) *Resolver {
	return &Resolver{entClient: entClient}
}



package out

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/siongleng89/gqlgen version v0.17.30-dev

import (
	"context"
	"fmt"

	customresolver "github.com/siongleng89/gqlgen/plugin/resolvergen/testdata/singlefile/out"
)

func (r *queryCustomResolverType) Resolver(ctx context.Context) (*customresolver.Resolver, error) {
	panic(fmt.Errorf("not implemented: Resolver - resolver"))
}

// This comment was added manually after code generation. It will remain after re-generation.
func (r *resolverCustomResolverType) Name(ctx context.Context, obj *customresolver.Resolver) (string, error) {
	panic(fmt.Errorf("not implemented: Name - name"))
}

func (r *CustomResolverType) Query() customresolver.QueryResolver { return &queryCustomResolverType{r} }

func (r *CustomResolverType) Resolver() customresolver.ResolverResolver {
	return &resolverCustomResolverType{r}
}

type queryCustomResolverType struct{ *CustomResolverType }
type resolverCustomResolverType struct{ *CustomResolverType }

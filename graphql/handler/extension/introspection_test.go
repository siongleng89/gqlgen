package extension

import (
	"context"
	"testing"

	"github.com/siongleng89/gqlgen/graphql"
	"github.com/stretchr/testify/require"
)

func TestIntrospection(t *testing.T) {
	rc := &graphql.OperationContext{
		DisableIntrospection: true,
	}
	require.Nil(t, Introspection{}.MutateOperationContext(context.Background(), rc))
	require.Equal(t, false, rc.DisableIntrospection)
}

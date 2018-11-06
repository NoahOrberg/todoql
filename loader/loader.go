package loader

import (
	"context"
	"fmt"

	"github.com/NoahOrberg/todoql/repository"
	"github.com/graph-gophers/dataloader"
)

type Loaders interface {
	Attach(context.Context) context.Context
}

func New(repo repository.Repo) Loaders {
	return &loaders{
		batchFuncs: map[string]dataloader.BatchFunc{
			todoLoaderKey: newTodoLoader(repo),
		},
	}
}

type loaders struct {
	batchFuncs map[string]dataloader.BatchFunc
}

func (c *loaders) Attach(ctx context.Context) context.Context {
	for key, batchFn := range c.batchFuncs {
		ctx = context.WithValue(ctx, key, dataloader.NewBatchedLoader(batchFn))
	}
	return ctx
}

func getLoader(ctx context.Context, key string) (*dataloader.Loader, error) {
	ldr, ok := ctx.Value(key).(*dataloader.Loader)
	if !ok {
		return nil, fmt.Errorf("unable to find %s loader from the request context", key)
	}
	return ldr, nil
}

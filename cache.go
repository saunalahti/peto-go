package main

import (
	"context"

	"github.com/dgraph-io/ristretto"
	"github.com/eko/gocache/lib/v4/cache"
	"github.com/eko/gocache/lib/v4/store"
	ristretto_store "github.com/eko/gocache/store/ristretto/v4"
)

var Cache *cache.Cache[string]

func ConnectCache() {
	ristrettoCache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1000,
		MaxCost:     100,
		BufferItems: 64,
	})
	if err != nil {
		panic(err)
	}
	ristrettoStore := ristretto_store.NewRistretto(ristrettoCache)

	Cache = cache.New[string](ristrettoStore)
	err = Cache.Set(context.Background(), "my-key", "my-value", store.WithCost(2))
	if err != nil {
		panic(err)
	}
}

package repo

import "github.com/google/wire"

// RepoSet model 注入
var RepoSet = wire.NewSet(
	TransSet,
	DemoSet,
	UserSet,
)

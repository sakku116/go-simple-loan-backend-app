package model

import "github.com/op/go-logging"

var logger = logging.MustGetLogger("model")

type MongoIndex struct {
	Key       string
	Unique    bool
	Direction int
}

type ModelProps struct {
	MinioBucketName      string
	QueriableFields      []string
	SortableFields       []string
	DefaultSortableField string
}

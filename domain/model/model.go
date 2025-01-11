package model

import "github.com/op/go-logging"

var logger = logging.MustGetLogger("model")

type MongoIndex struct {
	Key       string
	Unique    bool
	Direction int
}

type ModelProps struct {
	BucketName           string
	QueriableFields      []string
	SortableFields       []string
	DefaultSortableField string
}

type IModel interface {
	GetProps() ModelProps
}

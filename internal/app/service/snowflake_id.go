package service

import (
	"github.com/three-little-dragons/my-whiteboard-server/internal/pkg/snowflake"
)

var node *snowflake.Node

func init() {
	var err error
	node, err = snowflake.NewNode(1)
	if err != nil {
		println(err)
	}
}

func GenerateId() int64 {
	return node.Generate().Int64()
}

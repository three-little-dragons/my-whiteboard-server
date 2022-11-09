package http

import (
	"github.com/gin-gonic/gin"
	"github.com/three-little-dragons/my-whiteboard-server/internal/pkg/com"
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

func SnowflakeId(c *gin.Context) {
	com.Success(c, node.Generate().Int64())
}

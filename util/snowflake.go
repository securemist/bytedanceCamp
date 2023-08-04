/********************************************************************************
* @author: Yakult
* @date: 2023/8/4 17:36
* @description:
********************************************************************************/

package util

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"go.uber.org/zap"
	"time"
)

var node *snowflake.Node

func init() {
	startTime := "2023-08-04"
	var machineID int64 = 1
	st, err := time.Parse("2006-01-02", startTime)
	if err != nil {
		fmt.Println(err)
		return
	}

	snowflake.Epoch = st.UnixNano() / 1e6
	node, err = snowflake.NewNode(machineID)
	if err != nil {
		zap.S().Errorf("初始化雪花算法失败：%s", err)
	}
}

// GenID 生成 64 位的 雪花 ID
func GenID() int64 {
	return node.Generate().Int64()
}

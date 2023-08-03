/********************************************************************************
* @author: Yakult
* @date: 2023/8/2 15:33
* @description:
********************************************************************************/

package initialize

import (
	"bytedanceCamp/config"
	"bytedanceCamp/util/log"
)

func Initialize() {
	log.InitLogger(config.GetConfig().Log.Path, config.GetConfig().Log.Level)
}

package main

import (
	"go.uber.org/zap"
	log "yhgolang/log/util"
)

func main() {
	//log.Error("11",zap.String("1","2"))
	log.Errorf("11%s","11")
	//{"level":"error","ts":"2019-11-23T11:47:47.379+0800","caller":"log/main.go:9","msg":"1111"}
	log.Info("11",zap.String("1","2"))
	//log.Warn("11",zap.String("1","2"))
	//log.Panic("11",zap.String("1","2"))
}

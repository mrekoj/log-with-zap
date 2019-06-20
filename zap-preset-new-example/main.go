package main

import (
	"fmt"
	"go.uber.org/zap"
	"time"
)

func main() {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()

	sugar.Debug("Sugar log DEBUG")
	sugar.Debugf("Sugar log %s", "DEBUGF")
	sugar.Debugw("Sugar log DEBUGW",
		"debug", "log",
		"time", time.Now())

	fmt.Println("================================================================================================")

	sugar.Info("Sugar log INFO")
	sugar.Infof("Sugar log %s", "INFOF")
	sugar.Infow("Sugar log INFOW",
		"info", "log",
		"time", time.Now())

	fmt.Println("================================================================================================")

	sugar.Warn("Sugar log WARN")
	sugar.Warnf("Sugar log %s", "WARNF")
	sugar.Warnw("Sugar log WARNW",
		"warn", "log",
		"time", time.Now())

	fmt.Println("================================================================================================")
	sugar.Error("Sugar log ERROR")
	sugar.Errorf("Sugar log %s", "ERRORF")
	sugar.Errorw("Sugar log ERRORW",
		"error", "log",
		"time", time.Now())

	fmt.Println("================================================================================================")
	sugar.DPanic("Sugar log DPANIC")
	sugar.DPanicf("Sugar log %s", "DPANICF")
	sugar.DPanicw("Sugar log DPANICW",
		"dpanic", "log",
		"time", time.Now())

	fmt.Println("================================================================================================")
	sugar.Panic("Sugar log PANIC")
	sugar.Panicf("Sugar log %s", "PANICF")
	sugar.Panicw("Sugar log PANICW",
		"panic", "log",
		"time", time.Now())

	fmt.Println("================================================================================================")
	sugar.Fatal("Sugar log FATAL")
	sugar.Fatalf("Sugar log %s", "FATALF")
	sugar.Fatalw("Sugar log FATALW",
		"panic", "log",
		"time", time.Now())
}

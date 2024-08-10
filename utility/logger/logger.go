package logger

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/grpool"
	"path"
	"runtime"
)

func getCallerInfo() string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return "unknown:0"
	}
	return fmt.Sprintf("%s:%d", path.Base(file), line)
}

func Debug(ctx context.Context, v ...interface{}) {
	g.Log().Debug(ctx, v...)
}

func Info(ctx context.Context, v ...interface{}) {
	g.Log().Info(ctx, v...)
}

func Error(ctx context.Context, v ...interface{}) {
	g.Log().Error(ctx, v...)
}

func Debugf(ctx context.Context, format string, v ...interface{}) {
	message := append([]interface{}{getCallerInfo()}, v...)
	_ = grpool.AddWithRecover(gctx.NeverDone(ctx), func(ctx context.Context) { g.Log().Debugf(ctx, format, message...) }, nil)
}

func Infof(ctx context.Context, format string, v ...interface{}) {
	message := append([]interface{}{getCallerInfo()}, v...)
	_ = grpool.AddWithRecover(gctx.NeverDone(ctx), func(ctx context.Context) { g.Log().Infof(ctx, format, message...) }, nil)
}

func Errorf(ctx context.Context, format string, v ...interface{}) {
	message := append([]interface{}{getCallerInfo()}, v...)
	g.Log().Errorf(ctx, format, message...)
}

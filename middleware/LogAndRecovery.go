package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"lab_sys/global"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()             // 请求开始时间
		path := c.Request.URL.Path      // 请求路径
		query := c.Request.URL.RawQuery // 请求参数
		c.Next()                        // 处理请求

		cost := time.Since(start)     // 请求处理耗时
		if c.Writer.Status() != 200 { // 检查响应状态码
			// 记录日志
			zap.L().Info(path,
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", c.ClientIP()),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
				zap.Duration("cost", cost),
			)
		}
	}
}

func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}

					}
				}
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					global.Lg.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// 由于连接已断开，无法向客户端写入状态。
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}
				// 根据stack参数决定是否记录堆栈信息。
				if stack {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				// 中止请求处理并返回500内部服务器错误。
				c.AbortWithStatus(http.StatusInternalServerError)
			}

		}()
		c.Next()
	}

}

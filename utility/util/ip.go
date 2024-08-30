package util

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gipv4"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/iimeta/fastapi/internal/config"
	"github.com/iimeta/fastapi/utility/logger"
	"time"
)

var localIp = "127.0.0.1"

func init() {

	ctx := gctx.New()

	if len(config.Cfg.Local.PublicIp) > 0 {

		for _, url := range config.Cfg.Local.PublicIp {

			client := g.Client().Timeout(config.Cfg.Http.Timeout * time.Second)

			response, _ := client.Get(ctx, url)

			if response != nil {

				result := gstr.Trim(response.ReadAllString())
				if result != "" && gipv4.Validate(result) {
					localIp = result
					logger.Infof(ctx, "PUBLIC_IP: %s", localIp)
					_ = response.Close()
					break
				}

				_ = response.Close()
			}
		}
	} else {
		// get intranet ip
		defer func() {
			if r := recover(); r != nil {
				// handle the panic error
				logger.Errorf(ctx, "Error while getting intranet ip: %v", r)
			}
		}()

		// call the method that may panic
		localIp = gipv4.MustGetIntranetIp()
	}

	logger.Infof(ctx, "LOCAL_IP: %s", localIp)

}

func GetLocalIp() string {
	return localIp
}

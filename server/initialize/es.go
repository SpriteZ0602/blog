package initialize

import (
	"os"
	"server/global"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
	"go.uber.org/zap"
)

func ConnectES() *elasticsearch.TypedClient {
	esConfigGlobal := global.Config.ES

	esConfig := elasticsearch.Config{
		Addresses: []string{esConfigGlobal.URL},
		Username:  esConfigGlobal.Username,
		Password:  esConfigGlobal.Password,
	}

	if esConfigGlobal.IsConsolePrint {
		esConfig.Logger = &elastictransport.ColorLogger{
			Output:             os.Stdout,
			EnableRequestBody:  true,
			EnableResponseBody: true,
		}
	}

	client, err := elasticsearch.NewTypedClient(esConfig)

	if err != nil {
		global.Log.Error("连接ES异常", zap.Error(err))
		os.Exit(0)
	}

	return client
}

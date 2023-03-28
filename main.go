package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gotd/td/telegram"
	"github.com/tobycroft/Calc"
	"main.go/config/app_conf"
	"main.go/route"
	"os"
)

func init() {
	if app_conf.TestMode == false {
		s, err := os.Stat("./log/")

		if err != nil {
			os.Mkdir("./log", 0755)
		} else if s.IsDir() {
			os.Mkdir("./log", 0755)
		}
	}
}

func main() {
	// https://core.telegram.org/api/obtaining_api_id
	client := telegram.NewClient(25409358, "e53ea5c1e8a0321d19e21421d92e0b90", telegram.Options{})
	if err := client.Run(context.Background(), func(ctx context.Context) error {
		// It is only valid to use client while this function is not returned
		// and ctx is not cancelled.
		api := client.API()
		fmt.Println(api)
		// Now you can invoke MTProto RPC requests by calling the API.
		// ...

		// Return to close client connection and free up resources.
		return nil
	}); err != nil {
		panic(err)
	}
	// Client is closed.
}

func main2() {

	Calc.RefreshBaseNum()
	mainroute := gin.Default()
	//gin.SetMode(gin.ReleaseMode)
	//gin.DefaultWriter = ioutil.Discard
	mainroute.SetTrustedProxies([]string{"0.0.0.0/0"})
	mainroute.SecureJsonPrefix(app_conf.SecureJsonPrefix)
	route.OnRoute(mainroute)
	mainroute.Run(":80")
	mainroute.Run(":81")

}

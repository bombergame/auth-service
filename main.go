package main

import (
	"github.com/bombergame/auth-service/repositories/mysql"
	"github.com/bombergame/common/logs"
	"os"
	"os/signal"
)

func main() {
	logger := logs.NewLogger()

	conn := mysql.NewConnection()
	if err := conn.Open(); err != nil {
		logger.Fatal(err)
		return
	}
	defer func() {
		if err := conn.Close(); err != nil {
			logger.Error(err)
		}
	}()

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)

	//go func() {
	//	if err := restSrv.Run(); err != nil {
	//		logger.Fatal(err)
	//	}
	//}()
	//
	//go func() {
	//	if err := grpcSrv.Run(); err != nil {
	//		logger.Fatal(err)
	//	}
	//}()

	<-ch

	//if err := restSrv.Shutdown(); err != nil {
	//	logger.Fatal(err)
	//}
}

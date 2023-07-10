package main

import (
	"os"
	"os/signal"
	"smf/service"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var flags = []cli.Flag{
	cli.StringFlag{
		Name:  "config, c",
		Usage: "Load configuration from `FILE`",
	},
	cli.StringFlag{
		Name:  "log, l",
		Usage: "Output logs to `FILE`",
	},
}

func main() {
	logrus.Println("Start Session Management Function")
	app := cli.NewApp()
	app.Name = "smf"
	app.Usage = "5G SMF"
	app.Action = action
	app.Flags = flags

	if err := app.Run(os.Args); err != nil {
		//log
		logrus.Fatal("Fail to start SMF", err)
	} else {
		quit := make(chan struct{})
		sigch := make(chan os.Signal, 1)
		signal.Notify(sigch, os.Interrupt, syscall.SIGTERM)
		go func() {
			<-sigch
			// if nf != nil {
			// 	service.Stop(nf)
			// }
			service.Terminate()
			logrus.Info("Received a kill signal")
			quit <- struct{}{}
		}()
		<-quit
		logrus.Info("Shutdowned SMF")
	}
}

func action(c *cli.Context) (err error) {

	if err = service.Start(); err != nil {
		return
	}

	return nil
}

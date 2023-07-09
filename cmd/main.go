package main

import (
	"os"
	"os/signal"
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
			logrus.Info("Received a kill signal")
			quit <- struct{}{}
		}()
		<-quit
		logrus.Info("Shutdowned SMF")
	}
}

func action(c *cli.Context) (err error) {

	// var cfg context.UpmfConfig
	// if cfg, err = config.LoadConfig("config/smf.json"); err != nil {
	// 	logrus.Errorf(err.Error())
	// 	return
	// }
	// if nf, err = service.New(&cfg); err != nil {
	// 	logrus.Errorf("Fail to create UPMF", err)
	// 	return
	// }
	// service.Start(nf)

	return nil
}

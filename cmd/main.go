package main

import (
	"os"
	"os/signal"
	"syscall"

	"upmf/internal/context"
	"upmf/pkg/config"
	"upmf/pkg/service"

	log "github.com/sirupsen/logrus"
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

// var upmanager *upman.UpManager
var nf *context.UPMF

func main() {
	log.Println("Start User Plane Management Function")
	app := cli.NewApp()
	app.Name = "upmf"
	app.Usage = "5G UPMF"
	app.Action = action
	app.Flags = flags

	if err := app.Run(os.Args); err != nil {
		//log
		log.Fatal("Fail to start application", err)
	} else {
		quit := make(chan struct{})
		sigch := make(chan os.Signal, 1)
		signal.Notify(sigch, os.Interrupt, syscall.SIGTERM)
		go func() {
			<-sigch
			if nf != nil {
				service.Stop(nf)
			}
			log.Info("Received a kill signal")
			quit <- struct{}{}
		}()
		<-quit
		log.Info("Good bye the world")
	}
}

func action(c *cli.Context) (err error) {
	log.SetLevel(log.InfoLevel)

	var cfg context.UpmfConfig
	if cfg, err = config.LoadConfig("config/upmf.json"); err != nil {
		log.Errorf(err.Error())
		return
	}
	if nf, err = service.New(&cfg); err != nil {
		log.Errorf("Fail to create UPMF", err)
		return
	}
	service.Start(nf)

	return nil
}

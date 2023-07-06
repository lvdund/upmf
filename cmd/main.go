package main

import (
	"os"
	"os/signal"
	"syscall"
	"upf/config"
	"upf/models"
	"upf/service"

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

var nf *service.UPF

func main() {
	logrus.Infoln("User Plane Function")

	app := cli.NewApp()
	app.Name = "smf"
	app.Usage = "Etri 5G SMF"
	app.Action = action
	app.Flags = flags

	if err := app.Run(os.Args); err != nil {
		//log
		logrus.Fatal("Fail to start application", err)
	} else {
		quit := make(chan struct{})
		sigch := make(chan os.Signal, 1)
		signal.Notify(sigch, os.Interrupt, syscall.SIGTERM)
		go func() {
			<-sigch
			// if nf != nil {
			// 	nf.Terminate()
			// }
			nf.Terminate()
			quit <- struct{}{}
		}()
		<-quit
		logrus.Info("BYE!")
	}
}

func action(ctx *cli.Context) (err error) {

	var cfg models.UpfConfig
	filename := ctx.String("config")
	if cfg, err = config.LoadConfig(filename); err != nil {
		logrus.Errorf("Fail to parse UPF configuration:", err)
		return
	}
	nf, err = service.New(&cfg)
	if err != nil {
		logrus.Errorf("Fail to setup UPF:", err)
		return
	}
	nf.Print()
	if err = nf.Start(); err != nil {
		logrus.Errorf("Error running UPF:", err)
		return
	}

	return
}

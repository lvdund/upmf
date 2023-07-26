package main

import (
	"os"
	"os/signal"
	"syscall"

	"upmf/internal/context"
	"upmf/pkg/config"
	"upmf/pkg/service"

	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var log *logrus.Entry

func init() {
	log = logrus.WithFields(logrus.Fields{"mod": "upmf"})
}

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

	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigch
		log.Printf("Shutdown UPMF")
		service.Stop(nf)
	}()

	app.Run(os.Args)
}

func action(c *cli.Context) (err error) {

	var cfg *context.UpmfConfig
	if cfg, err = config.LoadConfig("config/upmf.json"); err != nil {
		log.Errorln("Cannot load Config:", err.Error())
		return
	}
	// Print(cfg)
	if nf, err = service.New(cfg); err != nil {
		log.Errorln("Fail to Setup UPMF:", err)
		return
	}
	service.Start(nf)

	return nil
}

func Print(cfg *context.UpmfConfig) {
	spew.Config.Indent = "\t"
	str := spew.Sdump(cfg)
	log.Println("==================================================")
	log.Printf("%s", str)
	log.Println("==================================================")
}

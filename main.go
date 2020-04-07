package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/masonj188/dotfiler/config"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s [flags] [command]\n\n", strings.TrimPrefix(os.Args[0], "./"))
		flag.PrintDefaults()
	}
	defHostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Error setting hostname automatically, must use -h flag")
	}
	var backupPath = flag.String("b", "./backup/config.yml.bak", "Specify the path for backup or restore to use")
	var configPath = flag.String("c", "./config.yml", "Specify path and name of config file to use")
	var hostname = flag.String("h", defHostname, "Set hostname to use for this run")
	flag.Parse()

	command := flag.Arg(0)
	if command == "" {
		flag.Usage()
	}
	myConfig := &config.Config{}
	err = myConfig.Parse(*configPath)
	if err != nil {
		log.Fatalln("Error parsing config file: ", err)
	}

	switch command {
	case "update":
		myConfig.Update(*hostname)
	case "apply":
		err = myConfig.Apply(*hostname)
		if err != nil {
			log.Fatalln(err)
		}
	case "backup":
		myConfig.Backup(*hostname, *backupPath)
	case "restore":
		myConfig.Restore(*hostname, *backupPath)
	default:
		flag.Usage()
	}
	//fmt.Println(myConfig)

}

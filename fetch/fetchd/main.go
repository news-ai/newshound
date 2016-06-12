package main

import (
	"flag"
	"log"
	"time"

	"github.com/jprobinson/go-utils/utils"
	"gopkg.in/mgo.v2"

	"github.com/news-ai/newshound"
	"github.com/news-ai/newshound/fetch"
)

const logPath = "/var/log/newshound/fetchd.log"

var (
	logArg  = flag.String("log", logPath, "log path")
	reparse = flag.Bool("r", false, "reparse all alerts and events")
)

func main() {

	flag.Parse()

	if *logArg != "stderr" {
		logSetup := utils.NewDefaultLogSetup(*logArg)
		logSetup.SetupLogging()
		go utils.ListenForLogSignal(logSetup)
	} else {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	}

	config := newshound.NewConfig()

	sess, err := config.MgoSession()
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	if *reparse {
		if err := fetch.ReParse(config, sess); err != nil {
			log.Fatal(err)
		}
		return
	}

	go fetchMail(config, sess)

	mapReduce(sess)
}

func mapReduce(sess *mgo.Session) {
	for {
		err := fetch.MapReduce(sess)
		if err != nil {
			log.Print("problems performing mapreduce: ", err)

			time.Sleep(5 * time.Minute)
			continue
		}

		time.Sleep(1 * time.Hour)
	}
}

func fetchMail(config *newshound.Config, sess *mgo.Session) {
	for {
		fetch.FetchMail(config, sess)
		time.Sleep(30 * time.Second)
	}
}

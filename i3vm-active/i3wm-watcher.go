package main

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"go.i3wm.org/i3"
	"os"
	"time"
)

type Event struct {
	StartTimestamp time.Time
	EndTimestamp time.Time
	FocusDurationSeconds float64
	i3.WindowEvent
}

func main () {
	fset := pflag.NewFlagSet("i3wm-watcher", pflag.ExitOnError)
	outputFileName := fset.String("output", "events.log", "file to output the events")
	if err := fset.Parse(os.Args[1:]); err != nil {
		logrus.WithError(err).Panic()
	}

	w, err := os.OpenFile(*outputFileName, os.O_APPEND |os.O_CREATE | os.O_WRONLY, 0600)
	if err != nil {
		logrus.WithError(err).Panic("cannot open file")
	}
	defer func() {
		if err := w.Close(); err != nil {
			logrus.WithError(err).Errorln("cannot properly close file")
		}
	}()

	eventEncoder := json.NewEncoder(w)

	var last *Event
	recv := i3.Subscribe(
		i3.WindowEventType,
	)
	defer recv.Close()

	for {
		select {
		default:
			if !recv.Next() {
				return
			}

			ev := recv.Event() // .(*i3.WindowEvent)
			switch ev.(type) {
			case *i3.WindowEvent:
				ev := ev.(*i3.WindowEvent)
				b, _ := json.MarshalIndent(ev, "", "\t")
				if !ev.Container.Focused { // new aren't focused by default
					logrus.Warnln("Container not focused\n", ev.Container.Focused, string(b))
					continue
				}

				if last != nil {
					last.EndTimestamp = time.Now()
					last.FocusDurationSeconds = last.EndTimestamp.Sub(last.StartTimestamp).Seconds()

					logrus.WithFields(logrus.Fields{
						"change": ev.Change,
						"name": ev.Container.Name,
					}).Infof("[%s] %s for %v", last.Container.WindowProperties.Class, last.Container.WindowProperties.Title, last.FocusDurationSeconds)
					if err  := eventEncoder.Encode(last); err != nil {
						logrus.WithError(err).Panicln("cannot encode json")
					}
				}
				last = &Event{
					StartTimestamp: time.Now(),
					WindowEvent: *ev,
				}
			}
		}
	}
}

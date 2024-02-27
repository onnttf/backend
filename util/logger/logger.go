package logger

import (
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

var once sync.Once

var zeroLogger zerolog.Logger

func Get() zerolog.Logger {
	once.Do(func() {
		zerolog.TimeFieldFormat = time.RFC3339

		consoleWriter := zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}

		fileWriter := &lumberjack.Logger{
			Filename:   "logs/app.log",
			MaxSize:    5,
			MaxBackups: 10,
			MaxAge:     14,
			Compress:   true,
		}

		output := zerolog.MultiLevelWriter(consoleWriter, fileWriter)

		zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
			short := file
			for i := len(file) - 1; i > 0; i-- {
				if file[i] == '/' {
					short = file[i+1:]
					break
				}
			}
			file = short
			return file + ":" + strconv.Itoa(line)
		}

		zeroLogger = zerolog.New(output).
			Level(zerolog.InfoLevel).
			With().
			Caller().
			Timestamp().
			Logger()

		zerolog.DefaultContextLogger = &zeroLogger
	})
	return zeroLogger
}

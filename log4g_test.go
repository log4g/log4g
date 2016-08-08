package log4g

import (
	"fmt"
	"testing"

	log "github.com/cihub/seelog"
)

func TestFormat1(t *testing.T) {
	logger, err := log.LoggerFromConfigAsString(`
	    <seelog minlevel="info">
        <outputs formatid="main">
            <console />
        </outputs>
        <formats>
            <format id="main" format="%Date(2006-01-02 15:04:05.999)|%Lev|%RelFile:%Line > %Msg%n"/>
        </formats>
    </seelog>
	`)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer logger.Flush()

	logger.Debug("NOT Printed")
	logger.Info("Printed")
}

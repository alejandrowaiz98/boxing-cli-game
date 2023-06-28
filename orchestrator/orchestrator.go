package orchestrator

import (
	"bufio"
	"os"

	"github.com/alejandrowaiz98/boxing-cli-game/config"
)

type Orchestrator struct {
}

func New() Orchestrator {

	return Orchestrator{}

}

var logger = config.GetLogger()
var scanner = bufio.NewScanner(os.Stdin)

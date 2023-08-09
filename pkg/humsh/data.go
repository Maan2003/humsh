package humsh

import (
	"errors"
	"os/exec"

	"github.com/emirpasic/gods/sets/treeset"
)

type ArgOrder int
const (
	PROGRAM ArgOrder = 100
	SUBCOMMAND ArgOrder = 200
	FLAG ArgOrder = 300
	POSITIONAL ArgOrder = 400
)

type Arg struct {
	order ArgOrder
	value string
}

type CommandLine struct {
	args orderedset.OrderedSet[Arg]
}

func (c CommandLine) ToCmd() (*exec.Cmd, error) {
	if c.args.Size() < 1 {
		return nil, errors.New("program name must be present")
	}
	args := []string{}
	for _, arg := range c.args.Values() {
		args = append(args, arg.value)
	}
	return exec.Command(args[0], args[1:]...), nil
}

func (c CommandLine*) AddArg(ArgOrder order, string value) {
	c.args.Add()
}

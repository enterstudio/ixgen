package bgpqworkers

import (
	"bytes"
	"encoding/json"
	"github.com/ipcjk/ixgen/ixtypes"
	"os/exec"
)

type BGPQ3Config struct {
	Executable string
	Arguments  []string
	Style      string
}

type BGPQ3Worker struct {
	BGPQ3Config
}

func NewBGPQ3Worker(Config BGPQ3Config) BGPQ3Worker {
	return BGPQ3Worker{BGPQ3Config: Config}
}

func (b *BGPQ3Worker) GenPrefixList(prefixListName, asMacro string, ipProtocol int) (ixtypes.PrefixFilters, error) {
	var w = new(bytes.Buffer)
	var ipParameter string
	var prefixFilters ixtypes.PrefixFilters

	if ipProtocol == 4 {
		ipParameter = "-4"
	} else {
		ipParameter = "-6"
	}

	cmd := exec.Command(b.Executable, ipParameter, "-j", "-l", prefixListName, asMacro)
	cmd.Stdout = w
	cmd.Stderr = w

	err := cmd.Run()
	if err != nil {
		return ixtypes.PrefixFilters{}, err
	}

	err = json.Unmarshal(w.Bytes(), &prefixFilters)
	if err != nil {
		return ixtypes.PrefixFilters{}, err
	}

	return prefixFilters, nil

}
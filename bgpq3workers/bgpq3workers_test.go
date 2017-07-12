package bgpqworkers_test

import (
	. "github.com/ipcjk/ixgen/bgpq3workers"
	"testing"
)

func TestRunBGP3Worker(t *testing.T) {
	testAsMacro := "AS196922"

	Config := BGPQ3Config{
		Executable: "/Users/joerg/Documents/Programmierung/bgpq3-0.1.21/bgpq3",
		Style:      "brocade/mlx",
		Arguments:  []string{"-4"},
	}

	bgpWorker := NewBGPQ3Worker(Config)

	prefixFilters, err := bgpWorker.GenPrefixList("as196922p4", testAsMacro, 4)
	if err != nil {
		t.Errorf("Cant run bgpq3: %s", err)
	}

	if prefixFilters.PrefixName != "as196922p4" {
		t.Error("Cant find my home ipv4 PrefixName")
	}

	if len(prefixFilters.PrefixFilters) <= 4 ||
		len(prefixFilters.PrefixFilters) >= 30 {
		t.Error("Found too less or too many ipv4 prefixes, cant be!")
	}

	prefixFilters, err = bgpWorker.GenPrefixList("as196922p6", testAsMacro, 6)
	if err != nil {
		t.Errorf("Cant run bgpq3: %s", err)
	}

	if prefixFilters.PrefixName != "as196922p6" {
		t.Error("Cant find my home ipv6 PrefixName")
	}

	if len(prefixFilters.PrefixFilters) == 0 ||
		len(prefixFilters.PrefixFilters) >= 30 {
		t.Error("Found too less or too many ipv6 prefixes, cant be!")
	}

}
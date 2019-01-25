package pkg1

import (
	candp "github.com/WhisperingChaos/configAndPrint"
)

func ConfigPrint() {

	c := candp.Config{C: "Config"}
	c.PrintIt()
}

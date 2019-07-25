package fileops

import (
	"io"
	"os"

	"github.com/aporeto-inc/benchmark-suite/catapult/common"
)

// copyPlanFile is a throw-away function to copy example plan files for first few runs of simulator.
func copyPlanFile(toDir string) {
	from, err := os.Open(paths["simFiles"] + "/example-plan.yaml")
	common.CheckError("copyPlanFile: Reading source plan file", err)
	defer from.Close()

	to, err := os.OpenFile(toDir+"/plan.yaml", os.O_RDWR|os.O_TRUNC, 0666)
	common.CheckError("copyPlanFile: Handle to destination plan file", err)
	defer to.Close()

	_, err = io.Copy(to, from)
	common.CheckError("copyPlanFile: io.Copy phase", err)
}

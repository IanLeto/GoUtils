package utils

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

type Testcopy_testSuite struct {
	suite.Suite
}

func (s *Testcopy_testSuite) TestFunc() {
	a := map[string]interface{}{
		"k": "v",
	}
	b := map[string]interface{}{}
	MapCopy(a, b)
	s.Equal(a["k"], b["k"])

	al := []map[string]interface{}{{"k1": "v1"}, {"k2": "v2"}}
	var bl []map[string]interface{}
	for _, value := range al {
		bl = append(bl, MapCopyValue(value))
	}
	s.Equal(bl[1]["k2"], al[1]["k2"])
	s.Equal(bl[0]["k1"], al[0]["k1"])
}

func (s *Testcopy_testSuite) SetupTest() {
	fmt.Println("system panic: runtime error: index out of range\ntransfer/utils.RecoverError\n\t/workDir/Go/goPath/src/transfer/utils/errors.go:37\nruntime.gopanic\n\t/usr/local/go/src/runtime/panic.go:522\nruntime.panicindex\n\t/usr/local/go/src/runtime/panic.go:44\ntransfer/esb.GetAllTaskInfo2\n\t/workDir/Go/goPath/src/transfer/esb/utils.go:355\ntransfer/esb.(*CCApiClient).VisitAllHost\n\t/workDir/Go/goPath/src/transfer/esb/cc.go:230\ntransfer/scheduler.(*CCHostUpdater).UpdateTo\n\t/workDir/Go/goPath/src/transfer/scheduler/cache.go:120\ntransfer/scheduler.NewCCHostUpdateTask.func4\n\t/workDir/Go/goPath/src/transfer/scheduler/cache.go:230\ntransfer/define.(*PeriodTask).Start\n\t/workDir/Go/goPath/src/transfer/define/task.go:148\ntransfer/define.(*TaskManager).Start.func1\n\t/workDir/Go/goPath/src/transfer/define/task.go:48\ntransfer/define.(*TaskManager).ForEach\n\t/workDir/Go/goPath/src/transfer/define/task.go:23\ntransfer/define.(*TaskManager).Start\n\t/workDir/Go/goPath/src/transfer/define/task.go:47\ntransfer/scheduler.(*Scheduler).Start\n\t/workDir/Go/goPath/src/transfer/scheduler/scheduler.go:222\ntransfer/cmd.glob..func14\n\t/workDir/Go/goPath/src/transfer/cmd/transfer.go:125\ngithub.com/spf13/cobra.(*Command).execute\n\t/workDir/Go/goPath/pkg/mod/github.com/spf13/cobra@v0.0.3/command.go:766\ngithub.com/spf13/cobra.(*Command).ExecuteC\n\t/workDir/Go/goPath/pkg/mod/github.com/spf13/cobra@v0.0.3/command.go:852\ngithub.com/spf13/cobra.(*Command).Execute\n\t/workDir/Go/goPath/pkg/mod/github.com/spf13/cobra@v0.0.3/command.go:800\ntransfer/cmd.Execute\n\t/workDir/Go/goPath/src/transfer/cmd/root.go:54\nmain.main\n\t/workDir/Go/goPath/src/transfer/main.go:127\nruntime.main\n\t/usr/local/go/src/runtime/proc.go:200\nruntime.goexit\n\t/usr/local/go/src/runtime/asm_amd64.s:1337")

}

func TestMapCopySuite(t *testing.T) {
	suite.Run(t, new(Testcopy_testSuite))
}

package upgrade

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/wangliujing/goctl/rpc/execx"
)

// upgrade gets the latest goctl by
// go install github.com/wangliujing/goctl@latest
func upgrade(_ *cobra.Command, _ []string) error {
	cmd := `GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go install github.com/wangliujing/goctl@latest`
	if runtime.GOOS == "windows" {
		cmd = `set GOPROXY=https://goproxy.cn,direct && go install github.com/wangliujing/goctl@latest`
	}
	info, err := execx.Run(cmd, "")
	if err != nil {
		return err
	}

	fmt.Print(info)
	return nil
}

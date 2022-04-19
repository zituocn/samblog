package conn

import (
	"github.com/zituocn/gow"
	"github.com/zituocn/gow/lib/config"
	"github.com/zituocn/gow/lib/logy"
	"os"
)

const (
	serverName = "sam-blog"
)

// InitLog init logy
func InitLog() {
	runMode := config.DefaultString("run_mode", "dev")
	if runMode == gow.ProdMode {
		logy.SetOutput(
			logy.MultiWriter(
				logy.WithColor(logy.NewWriter(os.Stdout)),
				logy.NewFileWriter(logy.FileWriterOptions{
					Dir:           "./logs",
					Prefix:        "web",
					StorageMaxDay: 7,
				}),
			),
			serverName,
		)
	} else {
		logy.SetOutput(
			logy.WithColor(logy.NewWriter(os.Stdout)),
			serverName,
		)
	}

	logy.Infof("-------------------------------------------")
	logy.Infof("Start %s ......", serverName)
	logy.Infof("-------------------------------------------")
}

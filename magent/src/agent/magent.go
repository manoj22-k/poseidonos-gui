package agent

import (
	"context"
	"log"
	"magent/src/inputs"
	"magent/src/models"
	"magent/src/outputs"
	"os"
	"os/signal"
	"syscall"
)

func Start() {
	dataChan := make(chan models.ClientPoint, 100)
	defer close(dataChan)
	ctx, cancel := context.WithCancel(context.Background())

	//The following code is for capturing OS signals and to close the application
	signals := make(chan os.Signal)
	signal.Notify(signals, os.Interrupt, syscall.SIGHUP,
		syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-signals
		log.Printf("Exit Signal received")
		cancel()
	}()
	go func() {
		startInputs(ctx, dataChan)
	}()
	outputs.WriteToDB(ctx, "http", dataChan)
}

func startInputs(ctx context.Context, dataChan chan models.ClientPoint) {
	defer log.Println("All inputs Closed")
	go inputs.WatchFiles(ctx, "/tmp", `(?i)^air_[\S]+\.json$`, "air", "air", "default_rp", dataChan)
	go inputs.TailFile(ctx, false, "/var/log/ibofos/report.log", "text", "rebuilding_status", "autogen", dataChan)
	go inputs.CollectCPUData(ctx, dataChan)
	go inputs.CollectMemoryData(ctx, dataChan)
	go inputs.CollectDiskData(ctx, dataChan)
	go inputs.CollectEthernetData(ctx, dataChan)
	go inputs.CollectNetworkData(ctx, dataChan)
	<-ctx.Done()
}

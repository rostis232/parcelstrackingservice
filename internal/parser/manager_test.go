package parser_test

import (
	"fmt"
	"github.com/rostis232/parcelstrackingservice/internal/parser"
	"github.com/rostis232/parcelstrackingservice/models"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

var tasks = []models.Task{
	{TrackNumber: "CI202518045DE", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001531337983CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001483004131CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001521012562CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001521662170CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "LZ134747205CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001505903961CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001464968841CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001451182489CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001467049264CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001525671239CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001523188937CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "RG027843019CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001510422358CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3060152493898CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001464212104CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001500079104CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "LP650662462FR", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001505419783CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001503747765CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "CC464900714DE", OutChannel: make(chan *models.Data)},
	{TrackNumber: "CQ343941106DE", OutChannel: make(chan *models.Data)},
	{TrackNumber: "RB318450992SG", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001513480065CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001481136711CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001480641254CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001524214536CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001509206259CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001506045899CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "CNUSUP00009319533", OutChannel: make(chan *models.Data)},
	{TrackNumber: "5P50D00165849", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001475547882CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "LP650665384FR", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001462228114CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001503358282CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001516249645CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001418781762CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3060152386514CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001527099413CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "CNUSUP00009986372", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001509804493CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001534491644CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001513017605CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "00340434498939362090", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001531445824CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001508240636CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001484959874CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3060152507625CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001507401809CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001515508750CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001445689758CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "RG028087126CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001522542068CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001518393162CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001522720176CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001517956061CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001488473024CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001504320011CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001488486635CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001489396285CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001494145274CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "CQ343778484DE", OutChannel: make(chan *models.Data)},
	{TrackNumber: "LZ131399965CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001511453818CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001488621791CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001497178712CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001483646973CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001493971768CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "CQ343902249DE", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001522801615CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001470119510CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001487645309CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "ML182338679MH", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001490556531CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "ML182340350MH", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001511216734CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001507900012CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001506488269CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001501404073CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001522853584CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001484500960CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001506905230CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001498964000CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001488493831CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "90378946838", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001489723203CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001525269238CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001501689784CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001505222673CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "RG027407155CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "LB398503847SG", OutChannel: make(chan *models.Data)},
	{TrackNumber: "LB400328653SG", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001504537563CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001519040591CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "CI202282772DE", OutChannel: make(chan *models.Data)},
	{TrackNumber: "CI202350282DE", OutChannel: make(chan *models.Data)},
	{TrackNumber: "LP650200519FR", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001464487939CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "RG027385582CN", OutChannel: make(chan *models.Data)},
	{TrackNumber: "4PX3001455510264CN", OutChannel: make(chan *models.Data)},
}

func TestParsingManager_AddTask1(t *testing.T) {
	parsingManager := parser.New(10*time.Second, 10)

	var (
		mu      sync.Mutex
		success int
		noData  int
		wg      sync.WaitGroup
	)

	wg.Add(len(tasks))

	for _, task := range tasks {
		go func() {
			data := <-task.OutChannel
			close(task.OutChannel)
			if data == nil {
				fmt.Println("-", task.TrackNumber, "nothing")
				mu.Lock()
				noData++
				mu.Unlock()
			} else {
				fmt.Println("+", task.TrackNumber, data.OriginCountry, "->", data.DestinationCountry, "checkpoints:", len(data.Checkpoints))
				mu.Lock()
				success++
				mu.Unlock()
			}
			assert.NotNil(t, data)
			wg.Done()
		}()
	}

	for _, task := range tasks {
		parsingManager.AddTask(task)
	}
	wg.Wait()
	fmt.Println(len(tasks), success, noData)
	assert.Equal(t, len(tasks), success)
	assert.Equal(t, noData, 0)
}

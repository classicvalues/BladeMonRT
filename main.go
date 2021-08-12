package main

import (
	"io/ioutil"
	"log"
	"runtime"
	"os"
	"syscall"
	"os/signal"
	"github.com/microsoft/BladeMonRT/logging"
	"github.com/microsoft/BladeMonRT/configs"
	"encoding/json"
	"github.com/microsoft/BladeMonRT/workflows"
)

type Main struct {
	logger          *log.Logger
	workflowScheduler	*WorkflowScheduler
}

func main() {
	// Set GOMAXPROCS such that all operations execute on a single thread.
	runtime.GOMAXPROCS(1)

	workflowsJson, err := ioutil.ReadFile(configs.WORKFLOW_FILE)
	if err != nil {
		log.Fatal(err)
	}
	var workflowFactory WorkflowFactory = newWorkflowFactory(workflowsJson, NodeFactory{})

	schedulesJson, err := ioutil.ReadFile(configs.SCHEDULE_FILE)
	if err != nil {
		log.Fatal(err)
	}

	var mainObj *Main = NewMain()
	mainObj.setupWorkflows(schedulesJson,workflowFactory)
	mainObj.logger.Println("Initialized main.")

	// Setup main such that main does not exit unless there is a keyboard interrupt.
	quitChannel := make(chan os.Signal, 1)
    signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM) 
	<-quitChannel
}

func NewMain() *Main{
	var logger *log.Logger = logging.LoggerFactory{}.ConstructLogger("Main")
	return &Main{workflowScheduler: newWorkflowScheduler(), logger: logger}
}

func (main *Main) setupWorkflows(schedulesJson []byte, workflowFactory WorkflowFactory) {
	// Parse the schedules JSON and add the schedules to the workflow scheduler.
	var schedules map[string][]ScheduleDescription
	json.Unmarshal([]byte(schedulesJson), &schedules)
	for _, schedule := range schedules["schedules"] {
		switch schedule.ScheduleType {
			case "on_win_event":
				var workflow workflows.InterfaceWorkflow = workflowFactory.constructWorkflow(schedule.WorkflowName)	
				var eventQueries []WinEventSubscribeQuery = parseEventSubscribeQueries(schedule.WinEventSubscribeQueries)			
				main.workflowScheduler.addWinEventBasedSchedule(workflow, eventQueries) 
			default:
				main.workflowScheduler.logger.Println("Given schedule type not supported.")
		}
	}
}
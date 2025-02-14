package utils

import (
	"fmt"
	"github.com/antchfx/xmlquery"
	"github.com/microsoft/BladeMonRT/logging"
	"log"
	"strconv"
	"strings"
	"time"
)

type UtilsInterface interface {
	ParseEventXML(eventXML string) EtwEvent
}

/** Class that contains utilities used in BladeMonRT classes. */
type Utils struct {
	logger *log.Logger
}

func NewUtils() *Utils {
	var logger *log.Logger = logging.LoggerFactory{}.ConstructLogger("Utils")
	return &Utils{logger: logger}
}

/** Class that represents an ETW event. */
type EtwEvent struct {
	Provider      string
	EventID       int
	TimeCreated   time.Time
	EventRecordID int
}

/** Parses out the event `Provider`, `EventID`, TimeCreated(`SystemTime`), `EventRecordID` (which is different from event ID) from the event XML. */
func (utils *Utils) ParseEventXML(eventXML string) EtwEvent {
	parsedXML, err := xmlquery.Parse(strings.NewReader(eventXML))
	if err != nil {
		utils.logger.Println("Error parsing XML.")
		return EtwEvent{}
	}

	root := xmlquery.FindOne(parsedXML, fmt.Sprintf("//Event"))
	var provider string = root.SelectElement("//Provider/@Name").InnerText()

	eventID, err := strconv.Atoi(root.SelectElement("//EventID").InnerText())
	if err != nil {
		utils.logger.Println("Wrong format of event ID.")
	}

	timeCreated, err := time.Parse("2006-01-02T15:04:05.0000000Z", root.SelectElement("//TimeCreated/@SystemTime").InnerText())
	if err != nil {
		utils.logger.Println("Wrong format of time.")
	}

	eventRecordID, err := strconv.Atoi(root.SelectElement("//EventRecordID").InnerText())
	if err != nil {
		utils.logger.Println("Wrong format of event record ID.")
	}

	return EtwEvent{Provider: provider, EventID: eventID, TimeCreated: timeCreated, EventRecordID: eventRecordID}
}

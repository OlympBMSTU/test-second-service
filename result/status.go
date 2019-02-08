package result

import (
	"fmt"
	"strconv"
)

type IBaseStatus interface {
	IsError() bool
	GetCode() int
	GetDescription() string
}

type IExtendedStatus interface {
	GetLine() int
	GetFile() string
	GetError() error
	Format(fmtName string) string // output data as html, json or raw
}

// Status is a struct to store all occured evets (like errors or not errors - good exit)
type Status interface {
	// IsError() bool
	// GetCode() int
	// GetDescription() string
	IExtendedStatus
	IBaseStatus
	// GetLine() int
	// GetFile() string
	// GetError() error
	// Format(fmtName string) string // output data as html, json or raw
}

type BaseStatus struct {
	line     int
	fileName string
	err      error
}

// func (status *BaseStatus) GetCode() int {
// 	return -1
// }

// func (status *BaseStatus) GetDescription() string {
// 	return ""
// }

func (status BaseStatus) GetError() error {
	return status.err
}

func (status BaseStatus) GetLine() int {
	return status.line
}

func (status BaseStatus) GetFile() string {
	return status.fileName
}

// func (status *AbstractStatus) IsError() bool {
// 	return true
// }

func (status BaseStatus) Format(format string) string {
	var strRepresentation string
	switch format {
	case "html":
	case "json":
	default:
		strRepresentation = "Line: " + strconv.Itoa(status.line) + "\n" +
			"File: " + status.fileName + "\n" +
			"Error: " + fmt.Sprintf("%v", status.err)
	}
	return strRepresentation
}

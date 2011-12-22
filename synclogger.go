// Copyright 2011 Cloud Instruments Co. Ltd. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sealog

import (
	. "github.com/cihub/sealog/common"
	cfg "github.com/cihub/sealog/config"
)

// SyncLogger performs logging in the same goroutine where 'Trace/Debug/...'
// func was called
type SyncLogger struct {
	commonLogger 
}

// NewSyncLogger creates a new synchronous logger
func NewSyncLogger(config *cfg.LogConfig) (*SyncLogger){
	syncLogger := new(SyncLogger)
	
	syncLogger.commonLogger = *newCommonLogger(config, syncLogger)
	
	return syncLogger
}

func (cLogger *SyncLogger) log(
    level LogLevel, 
	format string, 
	params []interface{}) {
	
	context, err := SpecificContext(3)
	if err != nil {
		reportInternalError(err)
		return
	}
		
	cLogger.processLogMsg(level, format, params, context)
}

func (syncLogger *SyncLogger) Close() {
	syncLogger.config.RootDispatcher.Close()
}

func (syncLogger *SyncLogger) Flush() {
	syncLogger.config.RootDispatcher.Flush()
}
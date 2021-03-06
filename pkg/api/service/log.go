package service

import (
	"github.com/spf13/viper"
	"time"
	"zeus/pkg/api/dao"
	"zeus/pkg/api/dto"
	"zeus/pkg/api/model"
)

type LoginLog = model.LoginLog
type OperationLog = model.OperationLog

var loginLogDao = dao.LoginLogDao{}
var operationLogDao = dao.OperationLogDao{}

// LogService
type LogService struct {
}

// LoginLogInfoOfId log login detail
func (LogService) LoginLogInfoOfId(dto dto.GeneralGetDto) LoginLog {
	return loginLogDao.Detail(dto.Id)
}

// List - users list with pagination
func (LogService) LoginLogLists(dto dto.LoginLogListDto) ([]dao.LoginLogList, int64) {
	return loginLogDao.Lists(dto)
}

func (LogService) OperationLogInfoOfId(dto dto.GeneralGetDto) OperationLog {
	return operationLogDao.Detail(dto.Id)
}

// List - users list with pagination
func (LogService) OperationLogLists(dto dto.OperationLogListDto) ([]dao.OperationLogList, int64) {
	return operationLogDao.Lists(dto)
}

//Insert Operation Log
func (LogService) InsertOperationLog(orLogDto dto.OperationLogDto) error {
	return operationLogDao.Create(orLogDto)
}

//Insert Operation Log
//func (LogService) DeleteLatestPwdUpdate(gDto dto.OperationLogDto) error {
//	//return operationLogDao.Create(orLogDto)
//}

// CheckIdleTooLong check duration between now and  last action time
// true - means too long time user not doing anything,we should kick user out of admin pages
func (LogService) CheckAccountIdleTooLong(uDto dto.GeneralGetDto) bool {
	if viper.GetInt("security.level") == 0 {
		return false
	}
	// pick the latest access record of account
	// then judge if it pass over 1 hour
	oLog := operationLogDao.GetLatestLogOfAccount(uDto.Id)
	if oLog.Id < 1 || time.Now().Sub(oLog.CreateTime).Seconds() > viper.GetFloat64("login.idleDuration") {
		return true
	}
	return false
}

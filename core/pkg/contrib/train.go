package contrib

import (
	"github.com/unarxiv/cvpm/pkg/entity"
	"time"
)

func insertTrainRecordToDB(rayId string, vendor string, packageName string, solverName string, dataId string) bool {
	trainRecord := entity.TrainRecord{
		RayID:     rayId,
		Vendor:    vendor,
		DataId:    dataId,
		Package:   packageName,
		Solver:    solverName,
		Status:    "in progress",
		CreatedAt: time.Now().Format(time.RFC3339),
	}
	sess := GetDBInstance()
	fileCollection := sess.Collection("train")
	err := fileCollection.InsertReturning(&trainRecord)
	if err != nil {
		CloseDB(sess)
		return false
	}
	return true
}

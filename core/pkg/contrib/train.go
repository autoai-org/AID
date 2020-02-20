package contrib

import (
	"github.com/unarxiv/cvpm/pkg/entity"
	"log"
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

func fetchAllTrainRecord() []entity.TrainRecord {
	sess := GetDBInstance()
	trainCollection := sess.Collection("train")
	res := trainCollection.Find()
	var trainRecords []entity.TrainRecord
	err := res.All(&trainRecords)
	if err != nil {
		log.Fatalf("res.All(): %q\n", err)
	}
	CloseDB(sess)
	return trainRecords
}


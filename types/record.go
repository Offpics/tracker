package types

import (
	"fmt"
	"strconv"
	"time"
	"tracker/utils"
)

type Record struct {
	activityName string
	time         int
	date         time.Time
}

func (record Record) ToStringArray() []string {
	return []string{record.activityName, fmt.Sprintf("%v", record.time), record.date.String()}
}

func RecordFromStringArray(record []string) Record {
	activityName := record[0]
	date, err := time.Parse(time.RFC3339, record[2])
	utils.Check(err)
	time, err := strconv.Atoi(record[1])
	utils.Check(err)

	return Record{activityName, time, date}
}

func MakeRecord(elapsedTime int, activityName string) Record {
	return Record{activityName, elapsedTime, time.Now()}
}

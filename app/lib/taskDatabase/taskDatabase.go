package taskDatabase

import (
	"encoding/json"
	"log"

	"github.com/pchan37/tasky/app/lib/dbManager"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

var database *mgo.Database

func Insert(task Task) (success bool) {
	updateSelector := bson.M{"index": bson.M{"$gte": 0}}
	updateUpdator := bson.M{"$inc": bson.M{"index": 1}}
	_, err := database.C("tasks").UpdateAll(updateSelector, updateUpdator)
	if err != nil {
		return
	}
	if err := database.C("tasks").Insert(task); err == nil {
		success = true
	}
	return
}

func Get() {

}

func GetAll() (result []byte) {
	var queryResult []Task
	if err := database.C("tasks").Find(bson.M{}).All(&queryResult); err == nil {
		for i, j := 0, len(queryResult)-1; i < j; i, j = i+1, j-1 {
			queryResult[i], queryResult[j] = queryResult[j], queryResult[i]
		}
		if result, err := json.Marshal(queryResult); err == nil {
			return result
		}
	}
	return nil
}

func Update(task Task) (success bool) {
	selector := bson.M{"index": task.Index}
	updator := bson.M{"$set": bson.M{"title": task.Title, "time": task.Time, "body": task.Body}}
	var err error
	if err = database.C("tasks").Update(selector, updator); err == nil {
		success = true
	}
	log.Println(err)
	return
}

func UpdatePosition(taskPosition TaskPosition) (success bool) {
	task := Task{}

	findSelector := bson.M{"index": taskPosition.StartIndex}
	if err := database.C("tasks").Find(findSelector).One(&task); err == nil {
		updateSelector := bson.M{"index": bson.M{"$gt": taskPosition.StartIndex, "$lte": taskPosition.EndIndex}}
		updateUpdator := bson.M{"$inc": bson.M{"index": -1}}
		_, err = database.C("tasks").UpdateAll(updateSelector, updateUpdator)
		updateTaskPositionSelector := bson.M{"index": task.Index, "title": task.Title, "time": task.Time, "body": task.Body}
		updateTaskPositionUpdator := bson.M{"$set": bson.M{"index": taskPosition.EndIndex}}
		err = database.C("tasks").Update(updateTaskPositionSelector, updateTaskPositionUpdator)
		if err == nil {
			success = true
		}
	}
	return
}

func Remove(taskPosition TaskPosition) (success bool) {
	selector := bson.M{"index": taskPosition.StartIndex}
	if err := database.C("tasks").Remove(selector); err == nil {
		updateIndexSelector := bson.M{"index": bson.M{"$gt": taskPosition.StartIndex}}
		updateIndexUpdator := bson.M{"$inc": bson.M{"index": -1}}
		_, err = database.C("tasks").UpdateAll(updateIndexSelector, updateIndexUpdator)
		if err == nil {
			success = true
		}
	}
	return
}

func InitializeDatabase() (manager *dbManager.DBManager) {
	manager = dbManager.New("tasky", "127.0.0.1:27017")
	database = manager.Database
	return
}

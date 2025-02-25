package utils

import (
	"encoding/json"
	"fmt"
	"github.com/go-openapi/strfmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

var PrettyPrint = func(i interface{}) {
	j, _ := json.MarshalIndent(i, "", "\t")
	fmt.Println(string(j))
}

var ConvertBsonMToMinimalJSONResponse = func(record mongo.SingleResult) (map[string]interface{}, error) {
	var result bson.M
	err := record.Decode(&result)
	if err != nil {
		return nil, err
	}
	// Convert `_id` to string if present
	if id, ok := result["_id"].(primitive.ObjectID); ok {
		result["id"] = id.Hex()
		delete(result, "_id") // Remove original _id key
	}

	// Convert `primitive.DateTime` fields to `strfmt.DateTime`
	for key, value := range result {
		if dt, ok := value.(primitive.DateTime); ok {
			result[key] = strfmt.DateTime(dt.Time()) // Convert to expected format
		}
	}

	// Convert `bson.M` to a minimal JSON response
	response := make(map[string]interface{})
	for key, value := range result {
		response[key] = value
	}

	return response, nil
}

var GerFieldsProjection = func(fieldsParam *string) bson.M {
	if fieldsParam == nil || *fieldsParam == "" {
		return nil // Return an empty map instead of nil
	}

	fields := strings.Split(*fieldsParam, ",")
	projection := bson.M{"_id": 1} // Always include ID

	for _, field := range fields {
		projection[field] = 1
	}

	return projection
}

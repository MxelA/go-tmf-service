package utils

import (
	"encoding/json"
	"fmt"
	"github.com/go-openapi/strfmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
	"strings"
	"time"
)

func PrettyPrint(i interface{}) {
	j, _ := json.MarshalIndent(i, "", "\t")
	fmt.Println(string(j))
}

func ConvertBsonMToMinimalJSONResponse(record mongo.SingleResult) (map[string]interface{}, error) {
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

func GerFieldsProjection(fieldsParam *string) bson.M {
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

func ValidatePaginationParams(offset *int64, limit *int64) (*int64, *int64) {
	// Handle nil values for limit and offset
	var offsetDefault int64 = 0 // Default offset
	var limitDefault int64 = 25 // Default limit
	var limitMax int64 = 500    // Max limit

	if offset == nil {
		offset = &offsetDefault
	}
	if limit == nil {
		limit = &limitDefault
	}
	if limit != nil && *limit > limitMax {
		limit = &limitMax
	}

	return offset, limit
}

func BuildMongoFilter(queryParams map[string][]string) bson.M {
	delete(queryParams, "fields")
	delete(queryParams, "limit")
	delete(queryParams, "offset")

	filter := bson.M{}

	for key, values := range queryParams {
		if len(values) == 0 {
			continue
		}

		// Extract operator (e.g., "completionDate.gt" -> field: "completionDate", operator: "$gt")
		parts := strings.SplitN(key, ".", 2)
		field := parts[0]
		operator := "$eq" // Default to equality

		if len(parts) > 1 {
			switch parts[1] {
			case "gt":
				operator = "$gt"
			case "gte":
				operator = "$gte"
			case "lt":
				operator = "$lt"
			case "lte":
				operator = "$lte"
			case "ne":
				operator = "$ne"
			case "in":
				operator = "$in"
			case "nin":
				operator = "$nin"
			}
		}

		// Convert values
		value := values[0] // Always take the first value

		// Try converting to int, float, or date automatically
		if intValue, err := strconv.Atoi(value); err == nil {
			filter[field] = bson.M{operator: intValue}
		} else if floatValue, err := strconv.ParseFloat(value, 64); err == nil {
			filter[field] = bson.M{operator: floatValue}
		} else if dateValue, err := time.Parse(time.RFC3339, value); err == nil {
			filter[field] = bson.M{operator: dateValue}
		} else {
			// Default to string
			filter[field] = bson.M{operator: value}
		}
	}

	return filter
}

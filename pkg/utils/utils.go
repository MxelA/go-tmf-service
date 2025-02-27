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
	// Exclude pagination & projection params
	delete(queryParams, "fields")
	delete(queryParams, "limit")
	delete(queryParams, "offset")

	filter := bson.M{}

	for key, values := range queryParams {
		if len(values) == 0 {
			continue
		}

		// Extract field name and operator (e.g., "serviceOrderItem.action.eq")
		parts := strings.Split(key, ".") // Split into field and operator
		filterOperator := parts[len(parts)-1]
		filterField := key
		operator := "$eq" // Default to equality

		if len(parts) > 1 {
			opMap := map[string]string{
				"eq":  "$eq",
				"gt":  "$gt",
				"gte": "$gte",
				"lt":  "$lt",
				"lte": "$lte",
				"ne":  "$ne",
				"in":  "$in",
				"nin": "$nin",
			}
			if op, exists := opMap[filterOperator]; exists {
				operator = op
				filterField = strings.Join(parts[:len(parts)-1], ".")
			}
		}

		// Convert value to correct type
		value := values[0]
		var parsedValue interface{} = value

		if intValue, err := strconv.Atoi(value); err == nil {
			parsedValue = intValue
		} else if floatValue, err := strconv.ParseFloat(value, 64); err == nil {
			parsedValue = floatValue
		} else if dateValue, err := time.Parse(time.RFC3339, value); err == nil {
			parsedValue = dateValue
		}

		// If the field already exists, merge the conditions
		if existing, exists := filter[filterField]; exists {
			if existingMap, ok := existing.(bson.M); ok {
				existingMap[operator] = parsedValue
				filter[filterField] = existingMap
			}
		} else {
			filter[filterField] = bson.M{operator: parsedValue}
		}
	}

	return filter
}

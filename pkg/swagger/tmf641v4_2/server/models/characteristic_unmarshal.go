package models

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

// UnmarshalBSON implements the bson.Unmarshaler interface for Characteristic.
func (c *Characteristic) UnmarshalBSON(data []byte) error {
	// Define a temporary struct to hold all fields from the original struct
	aux := &struct {
		ID                         string                        `bson:"id,omitempty"`
		Name                       *string                       `bson:"name,omitempty"`
		Value                      bson.RawValue                 `bson:"value"`
		ValueType                  string                        `bson:"valueType,omitempty"`
		CharacteristicRelationship []*CharacteristicRelationship `json:"characteristicRelationship" bson:"characteristicRelationship"`
	}{}

	// Unmarshal the raw BSON data into the temporary struct
	if err := bson.Unmarshal(data, &aux); err != nil {
		return fmt.Errorf("failed to unmarshal BSON data: %v", err)
	}

	// Debug: Print the temporary struct to verify all fields are populated
	//fmt.Printf("Temporary struct (aux): %+v\n", aux)

	// Copy all fields from aux to c
	c.ID = aux.ID
	c.Name = aux.Name
	c.ValueType = aux.ValueType
	c.CharacteristicRelationship = aux.CharacteristicRelationship

	// Deserialize the Value field based on its type
	switch aux.Value.Type {
	case bsontype.EmbeddedDocument:
		// If the value is an embedded document (object), deserialize it into a map
		var obj map[string]interface{}
		if err := aux.Value.Unmarshal(&obj); err != nil {
			return fmt.Errorf("failed to unmarshal value as object: %v", err)
		}
		c.Value = obj
	case bsontype.Array:
		// If the value is an array, deserialize it into a slice
		var arr []interface{}
		if err := aux.Value.Unmarshal(&arr); err != nil {
			return fmt.Errorf("failed to unmarshal value as array: %v", err)
		}
		c.Value = arr
	case bsontype.String:
		// If the value is a string, deserialize it into a string
		var s string
		if err := aux.Value.Unmarshal(&s); err != nil {
			return fmt.Errorf("failed to unmarshal value as string: %v", err)
		}
		c.Value = s
	case bsontype.Int32, bsontype.Int64:
		// If the value is an integer, deserialize it into an int
		var i int
		if err := aux.Value.Unmarshal(&i); err != nil {
			return fmt.Errorf("failed to unmarshal value as int: %v", err)
		}
		c.Value = i
	case bsontype.Double:
		// If the value is a float, deserialize it into a float64
		var f float64
		if err := aux.Value.Unmarshal(&f); err != nil {
			return fmt.Errorf("failed to unmarshal value as float: %v", err)
		}
		c.Value = f
	default:
		// If the type is unknown, store the raw BSON value
		c.Value = aux.Value
	}

	// Debug: Print the final struct to verify all fields are populated
	//fmt.Printf("Final struct (c): %+v\n", c)

	return nil
}

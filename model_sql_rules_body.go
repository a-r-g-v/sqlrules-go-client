/*
SQLRules API - Evaluation of test coverage for SQL database queries

A set of services to evaluate the coverage of SQL database queries. Coverage criteria are implemented in a set of rules, that when evaluated with respect to a given database determine the coverage of the database with respect to the query.  Two kind of coverage rules are generated, Full Predicate Coverage rules for SQL (SQLFpc) and Mutants for SQL (SQLMutation)

API version: 3.3.19
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sqlrules

import (
	"encoding/json"
)

// checks if the SqlRulesBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SqlRulesBody{}

// SqlRulesBody An object that includes the sql, schema and an optional string with parameters
type SqlRulesBody struct {
	// The sql to generate the coverage rules
	Sql *string `json:"sql,omitempty"`
	Schema *DbSchema `json:"schema,omitempty"`
	// A set of additional options to modify the behaviour of the rule generation (strings separated by space). Allowed values are in documented https://in2test.lsi.uniovi.es/sqlrules/api-doc.html. Some of them are dependent of the kind of rules to be generated, <br/>Example. `lang=en noboundary` specifies that the generated sqlfpc coverage rules must contain an english description of the rule and rules for checking boundary values must not been generated.
	Options *string `json:"options,omitempty"`
}

// NewSqlRulesBody instantiates a new SqlRulesBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSqlRulesBody() *SqlRulesBody {
	this := SqlRulesBody{}
	return &this
}

// NewSqlRulesBodyWithDefaults instantiates a new SqlRulesBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSqlRulesBodyWithDefaults() *SqlRulesBody {
	this := SqlRulesBody{}
	return &this
}

// GetSql returns the Sql field value if set, zero value otherwise.
func (o *SqlRulesBody) GetSql() string {
	if o == nil || IsNil(o.Sql) {
		var ret string
		return ret
	}
	return *o.Sql
}

// GetSqlOk returns a tuple with the Sql field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SqlRulesBody) GetSqlOk() (*string, bool) {
	if o == nil || IsNil(o.Sql) {
		return nil, false
	}
	return o.Sql, true
}

// HasSql returns a boolean if a field has been set.
func (o *SqlRulesBody) HasSql() bool {
	if o != nil && !IsNil(o.Sql) {
		return true
	}

	return false
}

// SetSql gets a reference to the given string and assigns it to the Sql field.
func (o *SqlRulesBody) SetSql(v string) {
	o.Sql = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *SqlRulesBody) GetSchema() DbSchema {
	if o == nil || IsNil(o.Schema) {
		var ret DbSchema
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SqlRulesBody) GetSchemaOk() (*DbSchema, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *SqlRulesBody) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given DbSchema and assigns it to the Schema field.
func (o *SqlRulesBody) SetSchema(v DbSchema) {
	o.Schema = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *SqlRulesBody) GetOptions() string {
	if o == nil || IsNil(o.Options) {
		var ret string
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SqlRulesBody) GetOptionsOk() (*string, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *SqlRulesBody) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given string and assigns it to the Options field.
func (o *SqlRulesBody) SetOptions(v string) {
	o.Options = &v
}

func (o SqlRulesBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SqlRulesBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Sql) {
		toSerialize["sql"] = o.Sql
	}
	if !IsNil(o.Schema) {
		toSerialize["schema"] = o.Schema
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	return toSerialize, nil
}

type NullableSqlRulesBody struct {
	value *SqlRulesBody
	isSet bool
}

func (v NullableSqlRulesBody) Get() *SqlRulesBody {
	return v.value
}

func (v *NullableSqlRulesBody) Set(val *SqlRulesBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSqlRulesBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSqlRulesBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSqlRulesBody(val *SqlRulesBody) *NullableSqlRulesBody {
	return &NullableSqlRulesBody{value: val, isSet: true}
}

func (v NullableSqlRulesBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSqlRulesBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



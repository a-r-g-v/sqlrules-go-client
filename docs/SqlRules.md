# SqlRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RulesClass** | Pointer to **string** | The class of the rules generated (&#x60;sqlfpc&#x60; or &#x60;sqlmutation&#x60;) | [optional] 
**Version** | Pointer to **string** | The version number of the service that generates this rule | [optional] 
**Environment** | Pointer to **string** | The environment of the service that generates this rule | [optional] 
**Summary** | Pointer to **map[string]string** | A map of additional properties to store information about the execution of the rules against the database | [optional] 
**Sql** | Pointer to **string** | The sql of the query that generates the rules | [optional] 
**Parsedsql** | Pointer to **string** | The sql after being parsed (only if specified by the &#x60;options&#x60; used when calling the service) | [optional] 
**Error** | Pointer to **string** | If empty, the service successfully obtained the rules, if not, indicates the error occurred, eg. the query is not syntactically correct. This field can be used to store runtime errors when executing the sql query | [optional] 
**Rules** | Pointer to [**[]SqlRule**](SqlRule.md) | The set of rules generated | [optional] 

## Methods

### NewSqlRules

`func NewSqlRules() *SqlRules`

NewSqlRules instantiates a new SqlRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlRulesWithDefaults

`func NewSqlRulesWithDefaults() *SqlRules`

NewSqlRulesWithDefaults instantiates a new SqlRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRulesClass

`func (o *SqlRules) GetRulesClass() string`

GetRulesClass returns the RulesClass field if non-nil, zero value otherwise.

### GetRulesClassOk

`func (o *SqlRules) GetRulesClassOk() (*string, bool)`

GetRulesClassOk returns a tuple with the RulesClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesClass

`func (o *SqlRules) SetRulesClass(v string)`

SetRulesClass sets RulesClass field to given value.

### HasRulesClass

`func (o *SqlRules) HasRulesClass() bool`

HasRulesClass returns a boolean if a field has been set.

### GetVersion

`func (o *SqlRules) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SqlRules) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SqlRules) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SqlRules) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetEnvironment

`func (o *SqlRules) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *SqlRules) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *SqlRules) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *SqlRules) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetSummary

`func (o *SqlRules) GetSummary() map[string]string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *SqlRules) GetSummaryOk() (*map[string]string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *SqlRules) SetSummary(v map[string]string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *SqlRules) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetSql

`func (o *SqlRules) GetSql() string`

GetSql returns the Sql field if non-nil, zero value otherwise.

### GetSqlOk

`func (o *SqlRules) GetSqlOk() (*string, bool)`

GetSqlOk returns a tuple with the Sql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSql

`func (o *SqlRules) SetSql(v string)`

SetSql sets Sql field to given value.

### HasSql

`func (o *SqlRules) HasSql() bool`

HasSql returns a boolean if a field has been set.

### GetParsedsql

`func (o *SqlRules) GetParsedsql() string`

GetParsedsql returns the Parsedsql field if non-nil, zero value otherwise.

### GetParsedsqlOk

`func (o *SqlRules) GetParsedsqlOk() (*string, bool)`

GetParsedsqlOk returns a tuple with the Parsedsql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParsedsql

`func (o *SqlRules) SetParsedsql(v string)`

SetParsedsql sets Parsedsql field to given value.

### HasParsedsql

`func (o *SqlRules) HasParsedsql() bool`

HasParsedsql returns a boolean if a field has been set.

### GetError

`func (o *SqlRules) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SqlRules) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SqlRules) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *SqlRules) HasError() bool`

HasError returns a boolean if a field has been set.

### GetRules

`func (o *SqlRules) GetRules() []SqlRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *SqlRules) GetRulesOk() (*[]SqlRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *SqlRules) SetRules(v []SqlRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *SqlRules) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



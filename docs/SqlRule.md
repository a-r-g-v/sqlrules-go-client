# SqlRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Summary** | Pointer to **map[string]string** | A map of additional properties to store information about the execution of the rule against the database | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Maintype** | Pointer to **string** |  | [optional] 
**Subtype** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Equivalent** | Pointer to **string** |  | [optional] 
**Sql** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** | This field can be used to store runtime errors when executing the sql query that represents this rule | [optional] 

## Methods

### NewSqlRule

`func NewSqlRule() *SqlRule`

NewSqlRule instantiates a new SqlRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlRuleWithDefaults

`func NewSqlRuleWithDefaults() *SqlRule`

NewSqlRuleWithDefaults instantiates a new SqlRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummary

`func (o *SqlRule) GetSummary() map[string]string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *SqlRule) GetSummaryOk() (*map[string]string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *SqlRule) SetSummary(v map[string]string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *SqlRule) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetId

`func (o *SqlRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SqlRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SqlRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SqlRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCategory

`func (o *SqlRule) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *SqlRule) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *SqlRule) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *SqlRule) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetMaintype

`func (o *SqlRule) GetMaintype() string`

GetMaintype returns the Maintype field if non-nil, zero value otherwise.

### GetMaintypeOk

`func (o *SqlRule) GetMaintypeOk() (*string, bool)`

GetMaintypeOk returns a tuple with the Maintype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintype

`func (o *SqlRule) SetMaintype(v string)`

SetMaintype sets Maintype field to given value.

### HasMaintype

`func (o *SqlRule) HasMaintype() bool`

HasMaintype returns a boolean if a field has been set.

### GetSubtype

`func (o *SqlRule) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *SqlRule) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *SqlRule) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *SqlRule) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetLocation

`func (o *SqlRule) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SqlRule) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SqlRule) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SqlRule) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetEquivalent

`func (o *SqlRule) GetEquivalent() string`

GetEquivalent returns the Equivalent field if non-nil, zero value otherwise.

### GetEquivalentOk

`func (o *SqlRule) GetEquivalentOk() (*string, bool)`

GetEquivalentOk returns a tuple with the Equivalent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquivalent

`func (o *SqlRule) SetEquivalent(v string)`

SetEquivalent sets Equivalent field to given value.

### HasEquivalent

`func (o *SqlRule) HasEquivalent() bool`

HasEquivalent returns a boolean if a field has been set.

### GetSql

`func (o *SqlRule) GetSql() string`

GetSql returns the Sql field if non-nil, zero value otherwise.

### GetSqlOk

`func (o *SqlRule) GetSqlOk() (*string, bool)`

GetSqlOk returns a tuple with the Sql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSql

`func (o *SqlRule) SetSql(v string)`

SetSql sets Sql field to given value.

### HasSql

`func (o *SqlRule) HasSql() bool`

HasSql returns a boolean if a field has been set.

### GetDescription

`func (o *SqlRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SqlRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SqlRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SqlRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetError

`func (o *SqlRule) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SqlRule) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SqlRule) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *SqlRule) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



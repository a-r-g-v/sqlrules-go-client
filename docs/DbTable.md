# DbTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of this table | [optional] 
**Tabletype** | Pointer to **string** | The type of this table as returned by jdbc, eg. &#x60;table&#x60;,&#x60;view&#x60;, &#x60;type&#x60;, NOTE, &#x60;type&#x60; tables hold User Defined Types (UDT), also called object, record or row (as named in SQL99) | [optional] 
**Extended** | Pointer to **map[string]string** | A map of additional table properties to store information required by other applications | [optional] 
**Columns** | Pointer to [**[]DbColumn**](DbColumn.md) | The set of columns in this table | [optional] 
**Checks** | Pointer to [**[]DbCheck**](DbCheck.md) | The set of check constraints associated to the columns of this table | [optional] 
**Ddls** | Pointer to [**[]Ddl**](Ddl.md) | Sql statements (dml) to create and/or drop this table Used by applications that require manage the table included in this schema | [optional] 

## Methods

### NewDbTable

`func NewDbTable() *DbTable`

NewDbTable instantiates a new DbTable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbTableWithDefaults

`func NewDbTableWithDefaults() *DbTable`

NewDbTableWithDefaults instantiates a new DbTable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DbTable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DbTable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DbTable) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DbTable) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTabletype

`func (o *DbTable) GetTabletype() string`

GetTabletype returns the Tabletype field if non-nil, zero value otherwise.

### GetTabletypeOk

`func (o *DbTable) GetTabletypeOk() (*string, bool)`

GetTabletypeOk returns a tuple with the Tabletype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTabletype

`func (o *DbTable) SetTabletype(v string)`

SetTabletype sets Tabletype field to given value.

### HasTabletype

`func (o *DbTable) HasTabletype() bool`

HasTabletype returns a boolean if a field has been set.

### GetExtended

`func (o *DbTable) GetExtended() map[string]string`

GetExtended returns the Extended field if non-nil, zero value otherwise.

### GetExtendedOk

`func (o *DbTable) GetExtendedOk() (*map[string]string, bool)`

GetExtendedOk returns a tuple with the Extended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtended

`func (o *DbTable) SetExtended(v map[string]string)`

SetExtended sets Extended field to given value.

### HasExtended

`func (o *DbTable) HasExtended() bool`

HasExtended returns a boolean if a field has been set.

### GetColumns

`func (o *DbTable) GetColumns() []DbColumn`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *DbTable) GetColumnsOk() (*[]DbColumn, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *DbTable) SetColumns(v []DbColumn)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *DbTable) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetChecks

`func (o *DbTable) GetChecks() []DbCheck`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *DbTable) GetChecksOk() (*[]DbCheck, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *DbTable) SetChecks(v []DbCheck)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *DbTable) HasChecks() bool`

HasChecks returns a boolean if a field has been set.

### GetDdls

`func (o *DbTable) GetDdls() []Ddl`

GetDdls returns the Ddls field if non-nil, zero value otherwise.

### GetDdlsOk

`func (o *DbTable) GetDdlsOk() (*[]Ddl, bool)`

GetDdlsOk returns a tuple with the Ddls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdls

`func (o *DbTable) SetDdls(v []Ddl)`

SetDdls sets Ddls field to given value.

### HasDdls

`func (o *DbTable) HasDdls() bool`

HasDdls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# SqlTableListBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sql** | Pointer to **string** | The sql that contains this list of tables | [optional] 
**Error** | Pointer to **string** | If empty, the service successfully obtained this object, if not, indicates the error occurred | [optional] 
**Tables** | Pointer to **[]string** | The list of tables used by this query | [optional] 

## Methods

### NewSqlTableListBody

`func NewSqlTableListBody() *SqlTableListBody`

NewSqlTableListBody instantiates a new SqlTableListBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlTableListBodyWithDefaults

`func NewSqlTableListBodyWithDefaults() *SqlTableListBody`

NewSqlTableListBodyWithDefaults instantiates a new SqlTableListBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSql

`func (o *SqlTableListBody) GetSql() string`

GetSql returns the Sql field if non-nil, zero value otherwise.

### GetSqlOk

`func (o *SqlTableListBody) GetSqlOk() (*string, bool)`

GetSqlOk returns a tuple with the Sql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSql

`func (o *SqlTableListBody) SetSql(v string)`

SetSql sets Sql field to given value.

### HasSql

`func (o *SqlTableListBody) HasSql() bool`

HasSql returns a boolean if a field has been set.

### GetError

`func (o *SqlTableListBody) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SqlTableListBody) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SqlTableListBody) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *SqlTableListBody) HasError() bool`

HasError returns a boolean if a field has been set.

### GetTables

`func (o *SqlTableListBody) GetTables() []string`

GetTables returns the Tables field if non-nil, zero value otherwise.

### GetTablesOk

`func (o *SqlTableListBody) GetTablesOk() (*[]string, bool)`

GetTablesOk returns a tuple with the Tables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTables

`func (o *SqlTableListBody) SetTables(v []string)`

SetTables sets Tables field to given value.

### HasTables

`func (o *SqlTableListBody) HasTables() bool`

HasTables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



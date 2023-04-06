# SqlParametersBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sql** | Pointer to **string** | The source (non-parametrized) | [optional] 
**Parsedsql** | Pointer to **string** | The transformed parametrized sql | [optional] 
**Error** | Pointer to **string** | If empty, the service successfully obtained this object, if not, indicates the error occurred | [optional] 
**Parameters** | Pointer to [**[]SqlParam**](SqlParam.md) | The list of parameters that have been extracted from the query | [optional] 

## Methods

### NewSqlParametersBody

`func NewSqlParametersBody() *SqlParametersBody`

NewSqlParametersBody instantiates a new SqlParametersBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlParametersBodyWithDefaults

`func NewSqlParametersBodyWithDefaults() *SqlParametersBody`

NewSqlParametersBodyWithDefaults instantiates a new SqlParametersBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSql

`func (o *SqlParametersBody) GetSql() string`

GetSql returns the Sql field if non-nil, zero value otherwise.

### GetSqlOk

`func (o *SqlParametersBody) GetSqlOk() (*string, bool)`

GetSqlOk returns a tuple with the Sql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSql

`func (o *SqlParametersBody) SetSql(v string)`

SetSql sets Sql field to given value.

### HasSql

`func (o *SqlParametersBody) HasSql() bool`

HasSql returns a boolean if a field has been set.

### GetParsedsql

`func (o *SqlParametersBody) GetParsedsql() string`

GetParsedsql returns the Parsedsql field if non-nil, zero value otherwise.

### GetParsedsqlOk

`func (o *SqlParametersBody) GetParsedsqlOk() (*string, bool)`

GetParsedsqlOk returns a tuple with the Parsedsql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParsedsql

`func (o *SqlParametersBody) SetParsedsql(v string)`

SetParsedsql sets Parsedsql field to given value.

### HasParsedsql

`func (o *SqlParametersBody) HasParsedsql() bool`

HasParsedsql returns a boolean if a field has been set.

### GetError

`func (o *SqlParametersBody) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SqlParametersBody) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SqlParametersBody) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *SqlParametersBody) HasError() bool`

HasError returns a boolean if a field has been set.

### GetParameters

`func (o *SqlParametersBody) GetParameters() []SqlParam`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *SqlParametersBody) GetParametersOk() (*[]SqlParam, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *SqlParametersBody) SetParameters(v []SqlParam)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *SqlParametersBody) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



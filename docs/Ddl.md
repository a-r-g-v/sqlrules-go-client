# Ddl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | **string** | Indicates the kind of update to perform (&#x60;create&#x60; or &#x60;drop&#x60;) | 
**Sql** | **string** | The sql to execute the update indicated by &#x60;command&#x60; | 

## Methods

### NewDdl

`func NewDdl(command string, sql string, ) *Ddl`

NewDdl instantiates a new Ddl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDdlWithDefaults

`func NewDdlWithDefaults() *Ddl`

NewDdlWithDefaults instantiates a new Ddl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *Ddl) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *Ddl) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *Ddl) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetSql

`func (o *Ddl) GetSql() string`

GetSql returns the Sql field if non-nil, zero value otherwise.

### GetSqlOk

`func (o *Ddl) GetSqlOk() (*string, bool)`

GetSqlOk returns a tuple with the Sql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSql

`func (o *Ddl) SetSql(v string)`

SetSql sets Sql field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



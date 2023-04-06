# SqlRulesBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sql** | Pointer to **string** | The sql to generate the coverage rules | [optional] 
**Schema** | Pointer to [**DbSchema**](DbSchema.md) |  | [optional] 
**Options** | Pointer to **string** | A set of additional options to modify the behaviour of the rule generation (strings separated by space). Allowed values are in documented https://in2test.lsi.uniovi.es/sqlrules/api-doc.html. Some of them are dependent of the kind of rules to be generated, &lt;br/&gt;Example. &#x60;lang&#x3D;en noboundary&#x60; specifies that the generated sqlfpc coverage rules must contain an english description of the rule and rules for checking boundary values must not been generated. | [optional] 

## Methods

### NewSqlRulesBody

`func NewSqlRulesBody() *SqlRulesBody`

NewSqlRulesBody instantiates a new SqlRulesBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlRulesBodyWithDefaults

`func NewSqlRulesBodyWithDefaults() *SqlRulesBody`

NewSqlRulesBodyWithDefaults instantiates a new SqlRulesBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSql

`func (o *SqlRulesBody) GetSql() string`

GetSql returns the Sql field if non-nil, zero value otherwise.

### GetSqlOk

`func (o *SqlRulesBody) GetSqlOk() (*string, bool)`

GetSqlOk returns a tuple with the Sql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSql

`func (o *SqlRulesBody) SetSql(v string)`

SetSql sets Sql field to given value.

### HasSql

`func (o *SqlRulesBody) HasSql() bool`

HasSql returns a boolean if a field has been set.

### GetSchema

`func (o *SqlRulesBody) GetSchema() DbSchema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *SqlRulesBody) GetSchemaOk() (*DbSchema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *SqlRulesBody) SetSchema(v DbSchema)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *SqlRulesBody) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetOptions

`func (o *SqlRulesBody) GetOptions() string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SqlRulesBody) GetOptionsOk() (*string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SqlRulesBody) SetOptions(v string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *SqlRulesBody) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



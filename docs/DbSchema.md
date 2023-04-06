# DbSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Catalog** | Pointer to **string** | The name of the database catalog as returned by jdbc | [optional] 
**Schema** | Pointer to **string** | The name of the database schema as returned by jdbc | [optional] 
**Dbms** | Pointer to **string** | A string representing the database vendor name returned by jdbc, eg. &#x60;postgres&#x60;, &#x60;oracle&#x60;, &#x60;sqlserver&#x60;. To be used by applications to handle vendor specific database features | [optional] 
**Tables** | Pointer to [**[]DbTable**](DbTable.md) | The set of tables included in this schema | [optional] 

## Methods

### NewDbSchema

`func NewDbSchema() *DbSchema`

NewDbSchema instantiates a new DbSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbSchemaWithDefaults

`func NewDbSchemaWithDefaults() *DbSchema`

NewDbSchemaWithDefaults instantiates a new DbSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalog

`func (o *DbSchema) GetCatalog() string`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *DbSchema) GetCatalogOk() (*string, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *DbSchema) SetCatalog(v string)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *DbSchema) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### GetSchema

`func (o *DbSchema) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *DbSchema) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *DbSchema) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *DbSchema) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetDbms

`func (o *DbSchema) GetDbms() string`

GetDbms returns the Dbms field if non-nil, zero value otherwise.

### GetDbmsOk

`func (o *DbSchema) GetDbmsOk() (*string, bool)`

GetDbmsOk returns a tuple with the Dbms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbms

`func (o *DbSchema) SetDbms(v string)`

SetDbms sets Dbms field to given value.

### HasDbms

`func (o *DbSchema) HasDbms() bool`

HasDbms returns a boolean if a field has been set.

### GetTables

`func (o *DbSchema) GetTables() []DbTable`

GetTables returns the Tables field if non-nil, zero value otherwise.

### GetTablesOk

`func (o *DbSchema) GetTablesOk() (*[]DbTable, bool)`

GetTablesOk returns a tuple with the Tables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTables

`func (o *DbSchema) SetTables(v []DbTable)`

SetTables sets Tables field to given value.

### HasTables

`func (o *DbSchema) HasTables() bool`

HasTables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



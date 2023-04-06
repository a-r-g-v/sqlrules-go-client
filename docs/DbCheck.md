# DbCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Column** | Pointer to **string** | The name of the column that has this constraint (if associated to only one column) | [optional] 
**Name** | Pointer to **string** | The name constraint as determined by jdbc | [optional] 
**Constraint** | **string** | The constraint expression as determined by jdbc | 

## Methods

### NewDbCheck

`func NewDbCheck(constraint string, ) *DbCheck`

NewDbCheck instantiates a new DbCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbCheckWithDefaults

`func NewDbCheckWithDefaults() *DbCheck`

NewDbCheckWithDefaults instantiates a new DbCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumn

`func (o *DbCheck) GetColumn() string`

GetColumn returns the Column field if non-nil, zero value otherwise.

### GetColumnOk

`func (o *DbCheck) GetColumnOk() (*string, bool)`

GetColumnOk returns a tuple with the Column field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumn

`func (o *DbCheck) SetColumn(v string)`

SetColumn sets Column field to given value.

### HasColumn

`func (o *DbCheck) HasColumn() bool`

HasColumn returns a boolean if a field has been set.

### GetName

`func (o *DbCheck) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DbCheck) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DbCheck) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DbCheck) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConstraint

`func (o *DbCheck) GetConstraint() string`

GetConstraint returns the Constraint field if non-nil, zero value otherwise.

### GetConstraintOk

`func (o *DbCheck) GetConstraintOk() (*string, bool)`

GetConstraintOk returns a tuple with the Constraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraint

`func (o *DbCheck) SetConstraint(v string)`

SetConstraint sets Constraint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



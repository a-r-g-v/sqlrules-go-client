# \SqlrulesApi

All URIs are relative to *https://in2test.lsi.uniovi.es/sqlrules/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MutantsPost**](SqlrulesApi.md#MutantsPost) | **Post** /mutants | Gets the set of mutants for a given sql (SQLMutation). Each rule is a mutant of the original SQL query. The mutant is dead (i.e. the rule is covered) if the execution against the database returns the same rows than the original query.
[**RulesPost**](SqlrulesApi.md#RulesPost) | **Post** /rules | Gets the SQLFpc coverage rules for a given sql. Based on Modified Condition Decision Coverage (MCDC) for SQL. Each rule is an SQL statement that is covered if the execution against the database returns at least one row.
[**SqlParametersPost**](SqlrulesApi.md#SqlParametersPost) | **Post** /sql/parameters | Transforms a non parametrized query into a parametrized one by converting each literal into a parameter and returns the mapping parameter-values
[**SqlTablesPost**](SqlrulesApi.md#SqlTablesPost) | **Post** /sql/tables | Gets the set of tables referenced in a sql query
[**TablesPost**](SqlrulesApi.md#TablesPost) | **Post** /tables | DEPRECATED - Gets the set of tables referenced in a sql query



## MutantsPost

> SqlRules MutantsPost(ctx).SqlRulesBody(sqlRulesBody).Execute()

Gets the set of mutants for a given sql (SQLMutation). Each rule is a mutant of the original SQL query. The mutant is dead (i.e. the rule is covered) if the execution against the database returns the same rows than the original query.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/a-r-g-v/sqlrules-go-client"
)

func main() {
    sqlRulesBody := *openapiclient.NewSqlRulesBody() // SqlRulesBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SqlrulesApi.MutantsPost(context.Background()).SqlRulesBody(sqlRulesBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SqlrulesApi.MutantsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MutantsPost`: SqlRules
    fmt.Fprintf(os.Stdout, "Response from `SqlrulesApi.MutantsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMutantsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sqlRulesBody** | [**SqlRulesBody**](SqlRulesBody.md) |  | 

### Return type

[**SqlRules**](SqlRules.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json;charset=utf-8
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesPost

> SqlRules RulesPost(ctx).SqlRulesBody(sqlRulesBody).Execute()

Gets the SQLFpc coverage rules for a given sql. Based on Modified Condition Decision Coverage (MCDC) for SQL. Each rule is an SQL statement that is covered if the execution against the database returns at least one row.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/a-r-g-v/sqlrules-go-client"
)

func main() {
    sqlRulesBody := *openapiclient.NewSqlRulesBody() // SqlRulesBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SqlrulesApi.RulesPost(context.Background()).SqlRulesBody(sqlRulesBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SqlrulesApi.RulesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RulesPost`: SqlRules
    fmt.Fprintf(os.Stdout, "Response from `SqlrulesApi.RulesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRulesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sqlRulesBody** | [**SqlRulesBody**](SqlRulesBody.md) |  | 

### Return type

[**SqlRules**](SqlRules.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json;charset=utf-8
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SqlParametersPost

> SqlParametersBody SqlParametersPost(ctx).Body(body).Execute()

Transforms a non parametrized query into a parametrized one by converting each literal into a parameter and returns the mapping parameter-values

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/a-r-g-v/sqlrules-go-client"
)

func main() {
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SqlrulesApi.SqlParametersPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SqlrulesApi.SqlParametersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SqlParametersPost`: SqlParametersBody
    fmt.Fprintf(os.Stdout, "Response from `SqlrulesApi.SqlParametersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSqlParametersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** |  | 

### Return type

[**SqlParametersBody**](SqlParametersBody.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json;charset=utf-8
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SqlTablesPost

> SqlTableListBody SqlTablesPost(ctx).Body(body).Execute()

Gets the set of tables referenced in a sql query

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/a-r-g-v/sqlrules-go-client"
)

func main() {
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SqlrulesApi.SqlTablesPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SqlrulesApi.SqlTablesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SqlTablesPost`: SqlTableListBody
    fmt.Fprintf(os.Stdout, "Response from `SqlrulesApi.SqlTablesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSqlTablesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** |  | 

### Return type

[**SqlTableListBody**](SqlTableListBody.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json;charset=utf-8
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TablesPost

> SqlTableListBody TablesPost(ctx).Body(body).Execute()

DEPRECATED - Gets the set of tables referenced in a sql query

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/a-r-g-v/sqlrules-go-client"
)

func main() {
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SqlrulesApi.TablesPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SqlrulesApi.TablesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TablesPost`: SqlTableListBody
    fmt.Fprintf(os.Stdout, "Response from `SqlrulesApi.TablesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTablesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** |  | 

### Return type

[**SqlTableListBody**](SqlTableListBody.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json;charset=utf-8
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


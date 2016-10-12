# resolve-ip.DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**healthCheck**](DefaultApi.md#healthCheck) | **GET** /healthcheck | 
[**locationForIP**](DefaultApi.md#locationForIP) | **GET** /ip/{ip} | 


<a name="healthCheck"></a>
# **healthCheck**
> healthCheck()



Checks if the service is healthy

### Example
```javascript
var resolve-ip = require('@clever/resolve-ip');

var apiInstance = new resolve-ip.DefaultApi();
apiInstance.healthCheck().then(function() {
  console.log('API called successfully.');
}, function(error) {
  console.error(error);
});

```

### Parameters
This endpoint does not need any parameter.

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="locationForIP"></a>
# **locationForIP**
> IP locationForIP(ip)



Gets the lat/lon for a given IP.

### Example
```javascript
var resolve-ip = require('@clever/resolve-ip');

var apiInstance = new resolve-ip.DefaultApi();

var ip = "ip_example"; // String | The IP to try to locate

apiInstance.locationForIP(ip).then(function(data) {
  console.log('API called successfully. Returned data: ' + data);
}, function(error) {
  console.error(error);
});

```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ip** | **String**| The IP to try to locate | 

### Return type

[**IP**](IP.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


<a name="module_resolve-ip"></a>

## resolve-ip
resolve-ip client library.


* [resolve-ip](#module_resolve-ip)
    * [ResolveIP](#exp_module_resolve-ip--ResolveIP) ⏏
        * [new ResolveIP(options)](#new_module_resolve-ip--ResolveIP_new)
        * _instance_
            * [.close()](#module_resolve-ip--ResolveIP+close)
            * [.healthCheck([options], [cb])](#module_resolve-ip--ResolveIP+healthCheck) ⇒ <code>Promise</code>
            * [.locationForIP(ip, [options], [cb])](#module_resolve-ip--ResolveIP+locationForIP) ⇒ <code>Promise</code>
        * _static_
            * [.RetryPolicies](#module_resolve-ip--ResolveIP.RetryPolicies)
                * [.Exponential](#module_resolve-ip--ResolveIP.RetryPolicies.Exponential)
                * [.Single](#module_resolve-ip--ResolveIP.RetryPolicies.Single)
                * [.None](#module_resolve-ip--ResolveIP.RetryPolicies.None)
            * [.Errors](#module_resolve-ip--ResolveIP.Errors)
                * [.BadRequest](#module_resolve-ip--ResolveIP.Errors.BadRequest) ⇐ <code>Error</code>
                * [.InternalError](#module_resolve-ip--ResolveIP.Errors.InternalError) ⇐ <code>Error</code>
                * [.NotFound](#module_resolve-ip--ResolveIP.Errors.NotFound) ⇐ <code>Error</code>
            * [.DefaultCircuitOptions](#module_resolve-ip--ResolveIP.DefaultCircuitOptions)

<a name="exp_module_resolve-ip--ResolveIP"></a>

### ResolveIP ⏏
resolve-ip client

**Kind**: Exported class  
<a name="new_module_resolve-ip--ResolveIP_new"></a>

#### new ResolveIP(options)
Create a new client object.


| Param | Type | Default | Description |
| --- | --- | --- | --- |
| options | <code>Object</code> |  | Options for constructing a client object. |
| [options.address] | <code>string</code> |  | URL where the server is located. Must provide this or the discovery argument |
| [options.discovery] | <code>bool</code> |  | Use clever-discovery to locate the server. Must provide this or the address argument |
| [options.timeout] | <code>number</code> |  | The timeout to use for all client requests, in milliseconds. This can be overridden on a per-request basis. Default is 5000ms. |
| [options.keepalive] | <code>bool</code> |  | Set keepalive to true for client requests. This sets the forever: true attribute in request. Defaults to true. |
| [options.retryPolicy] | [<code>RetryPolicies</code>](#module_resolve-ip--ResolveIP.RetryPolicies) | <code>RetryPolicies.Single</code> | The logic to determine which requests to retry, as well as how many times to retry. |
| [options.logger] | <code>module:kayvee.Logger</code> | <code>logger.New(&quot;resolve-ip-wagclient&quot;)</code> | The Kayvee logger to use in the client. |
| [options.circuit] | <code>Object</code> |  | Options for constructing the client's circuit breaker. |
| [options.circuit.forceClosed] | <code>bool</code> |  | When set to true the circuit will always be closed. Default: true. |
| [options.circuit.maxConcurrentRequests] | <code>number</code> |  | the maximum number of concurrent requests the client can make at the same time. Default: 100. |
| [options.circuit.requestVolumeThreshold] | <code>number</code> |  | The minimum number of requests needed before a circuit can be tripped due to health. Default: 20. |
| [options.circuit.sleepWindow] | <code>number</code> |  | how long, in milliseconds, to wait after a circuit opens before testing for recovery. Default: 5000. |
| [options.circuit.errorPercentThreshold] | <code>number</code> |  | the threshold to place on the rolling error rate. Once the error rate exceeds this percentage, the circuit opens. Default: 90. |

<a name="module_resolve-ip--ResolveIP+close"></a>

#### resolveIP.close()
Releases handles used in client

**Kind**: instance method of [<code>ResolveIP</code>](#exp_module_resolve-ip--ResolveIP)  
<a name="module_resolve-ip--ResolveIP+healthCheck"></a>

#### resolveIP.healthCheck([options], [cb]) ⇒ <code>Promise</code>
Checks if the service is healthy

**Kind**: instance method of [<code>ResolveIP</code>](#exp_module_resolve-ip--ResolveIP)  
**Fulfill**: <code>undefined</code>  
**Reject**: [<code>BadRequest</code>](#module_resolve-ip--ResolveIP.Errors.BadRequest)  
**Reject**: [<code>InternalError</code>](#module_resolve-ip--ResolveIP.Errors.InternalError)  
**Reject**: <code>Error</code>  

| Param | Type | Description |
| --- | --- | --- |
| [options] | <code>object</code> |  |
| [options.timeout] | <code>number</code> | A request specific timeout |
| [options.retryPolicy] | [<code>RetryPolicies</code>](#module_resolve-ip--ResolveIP.RetryPolicies) | A request specific retryPolicy |
| [cb] | <code>function</code> |  |

<a name="module_resolve-ip--ResolveIP+locationForIP"></a>

#### resolveIP.locationForIP(ip, [options], [cb]) ⇒ <code>Promise</code>
Gets the lat/lon for a given IP.

**Kind**: instance method of [<code>ResolveIP</code>](#exp_module_resolve-ip--ResolveIP)  
**Fulfill**: <code>Object</code>  
**Reject**: [<code>BadRequest</code>](#module_resolve-ip--ResolveIP.Errors.BadRequest)  
**Reject**: [<code>NotFound</code>](#module_resolve-ip--ResolveIP.Errors.NotFound)  
**Reject**: [<code>InternalError</code>](#module_resolve-ip--ResolveIP.Errors.InternalError)  
**Reject**: <code>Error</code>  

| Param | Type | Description |
| --- | --- | --- |
| ip | <code>string</code> | The IP to try to locate |
| [options] | <code>object</code> |  |
| [options.timeout] | <code>number</code> | A request specific timeout |
| [options.retryPolicy] | [<code>RetryPolicies</code>](#module_resolve-ip--ResolveIP.RetryPolicies) | A request specific retryPolicy |
| [cb] | <code>function</code> |  |

<a name="module_resolve-ip--ResolveIP.RetryPolicies"></a>

#### ResolveIP.RetryPolicies
Retry policies available to use.

**Kind**: static property of [<code>ResolveIP</code>](#exp_module_resolve-ip--ResolveIP)  

* [.RetryPolicies](#module_resolve-ip--ResolveIP.RetryPolicies)
    * [.Exponential](#module_resolve-ip--ResolveIP.RetryPolicies.Exponential)
    * [.Single](#module_resolve-ip--ResolveIP.RetryPolicies.Single)
    * [.None](#module_resolve-ip--ResolveIP.RetryPolicies.None)

<a name="module_resolve-ip--ResolveIP.RetryPolicies.Exponential"></a>

##### RetryPolicies.Exponential
The exponential retry policy will retry five times with an exponential backoff.

**Kind**: static constant of [<code>RetryPolicies</code>](#module_resolve-ip--ResolveIP.RetryPolicies)  
<a name="module_resolve-ip--ResolveIP.RetryPolicies.Single"></a>

##### RetryPolicies.Single
Use this retry policy to retry a request once.

**Kind**: static constant of [<code>RetryPolicies</code>](#module_resolve-ip--ResolveIP.RetryPolicies)  
<a name="module_resolve-ip--ResolveIP.RetryPolicies.None"></a>

##### RetryPolicies.None
Use this retry policy to turn off retries.

**Kind**: static constant of [<code>RetryPolicies</code>](#module_resolve-ip--ResolveIP.RetryPolicies)  
<a name="module_resolve-ip--ResolveIP.Errors"></a>

#### ResolveIP.Errors
Errors returned by methods.

**Kind**: static property of [<code>ResolveIP</code>](#exp_module_resolve-ip--ResolveIP)  

* [.Errors](#module_resolve-ip--ResolveIP.Errors)
    * [.BadRequest](#module_resolve-ip--ResolveIP.Errors.BadRequest) ⇐ <code>Error</code>
    * [.InternalError](#module_resolve-ip--ResolveIP.Errors.InternalError) ⇐ <code>Error</code>
    * [.NotFound](#module_resolve-ip--ResolveIP.Errors.NotFound) ⇐ <code>Error</code>

<a name="module_resolve-ip--ResolveIP.Errors.BadRequest"></a>

##### Errors.BadRequest ⇐ <code>Error</code>
BadRequest

**Kind**: static class of [<code>Errors</code>](#module_resolve-ip--ResolveIP.Errors)  
**Extends**: <code>Error</code>  
**Properties**

| Name | Type |
| --- | --- |
| message | <code>string</code> | 

<a name="module_resolve-ip--ResolveIP.Errors.InternalError"></a>

##### Errors.InternalError ⇐ <code>Error</code>
InternalError

**Kind**: static class of [<code>Errors</code>](#module_resolve-ip--ResolveIP.Errors)  
**Extends**: <code>Error</code>  
**Properties**

| Name | Type |
| --- | --- |
| message | <code>string</code> | 

<a name="module_resolve-ip--ResolveIP.Errors.NotFound"></a>

##### Errors.NotFound ⇐ <code>Error</code>
NotFound

**Kind**: static class of [<code>Errors</code>](#module_resolve-ip--ResolveIP.Errors)  
**Extends**: <code>Error</code>  
**Properties**

| Name | Type |
| --- | --- |
| message | <code>string</code> | 

<a name="module_resolve-ip--ResolveIP.DefaultCircuitOptions"></a>

#### ResolveIP.DefaultCircuitOptions
Default circuit breaker options.

**Kind**: static constant of [<code>ResolveIP</code>](#exp_module_resolve-ip--ResolveIP)  

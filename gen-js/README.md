<a name="module_resolve-ip"></a>

## resolve-ip
resolve-ip client library.


* [resolve-ip](#module_resolve-ip)
    * [ResolveIP](#exp_module_resolve-ip--ResolveIP) ⏏
        * [new ResolveIP(options)](#new_module_resolve-ip--ResolveIP_new)
        * _instance_
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
| [options.discovery] | <code>bool</code> |  | Use @clever/discovery to locate the server. Must provide this or the address argument |
| [options.timeout] | <code>number</code> |  | The timeout to use for all client requests, in milliseconds. This can be overridden on a per-request basis. |
| [options.retryPolicy] | <code>[RetryPolicies](#module_resolve-ip--ResolveIP.RetryPolicies)</code> | <code>RetryPolicies.Single</code> | The logic to determine which requests to retry, as well as how many times to retry. |

<a name="module_resolve-ip--ResolveIP+healthCheck"></a>

#### resolveIP.healthCheck([options], [cb]) ⇒ <code>Promise</code>
Checks if the service is healthy

**Kind**: instance method of <code>[ResolveIP](#exp_module_resolve-ip--ResolveIP)</code>  
**Fulfill**: <code>undefined</code>  
**Reject**: <code>[BadRequest](#module_resolve-ip--ResolveIP.Errors.BadRequest)</code>  
**Reject**: <code>[InternalError](#module_resolve-ip--ResolveIP.Errors.InternalError)</code>  
**Reject**: <code>Error</code>  

| Param | Type | Description |
| --- | --- | --- |
| [options] | <code>object</code> |  |
| [options.timeout] | <code>number</code> | A request specific timeout |
| [options.span] | <code>[Span](https://doc.esdoc.org/github.com/opentracing/opentracing-javascript/class/src/span.js~Span.html)</code> | An OpenTracing span - For example from the parent request |
| [options.retryPolicy] | <code>[RetryPolicies](#module_resolve-ip--ResolveIP.RetryPolicies)</code> | A request specific retryPolicy |
| [cb] | <code>function</code> |  |

<a name="module_resolve-ip--ResolveIP+locationForIP"></a>

#### resolveIP.locationForIP(ip, [options], [cb]) ⇒ <code>Promise</code>
Gets the lat/lon for a given IP.

**Kind**: instance method of <code>[ResolveIP](#exp_module_resolve-ip--ResolveIP)</code>  
**Fulfill**: <code>Object</code>  
**Reject**: <code>[BadRequest](#module_resolve-ip--ResolveIP.Errors.BadRequest)</code>  
**Reject**: <code>[NotFound](#module_resolve-ip--ResolveIP.Errors.NotFound)</code>  
**Reject**: <code>[InternalError](#module_resolve-ip--ResolveIP.Errors.InternalError)</code>  
**Reject**: <code>Error</code>  

| Param | Type | Description |
| --- | --- | --- |
| ip | <code>string</code> | The IP to try to locate |
| [options] | <code>object</code> |  |
| [options.timeout] | <code>number</code> | A request specific timeout |
| [options.span] | <code>[Span](https://doc.esdoc.org/github.com/opentracing/opentracing-javascript/class/src/span.js~Span.html)</code> | An OpenTracing span - For example from the parent request |
| [options.retryPolicy] | <code>[RetryPolicies](#module_resolve-ip--ResolveIP.RetryPolicies)</code> | A request specific retryPolicy |
| [cb] | <code>function</code> |  |

<a name="module_resolve-ip--ResolveIP.RetryPolicies"></a>

#### ResolveIP.RetryPolicies
Retry policies available to use.

**Kind**: static property of <code>[ResolveIP](#exp_module_resolve-ip--ResolveIP)</code>  

* [.RetryPolicies](#module_resolve-ip--ResolveIP.RetryPolicies)
    * [.Exponential](#module_resolve-ip--ResolveIP.RetryPolicies.Exponential)
    * [.Single](#module_resolve-ip--ResolveIP.RetryPolicies.Single)
    * [.None](#module_resolve-ip--ResolveIP.RetryPolicies.None)

<a name="module_resolve-ip--ResolveIP.RetryPolicies.Exponential"></a>

##### RetryPolicies.Exponential
The exponential retry policy will retry five times with an exponential backoff.

**Kind**: static constant of <code>[RetryPolicies](#module_resolve-ip--ResolveIP.RetryPolicies)</code>  
<a name="module_resolve-ip--ResolveIP.RetryPolicies.Single"></a>

##### RetryPolicies.Single
Use this retry policy to retry a request once.

**Kind**: static constant of <code>[RetryPolicies](#module_resolve-ip--ResolveIP.RetryPolicies)</code>  
<a name="module_resolve-ip--ResolveIP.RetryPolicies.None"></a>

##### RetryPolicies.None
Use this retry policy to turn off retries.

**Kind**: static constant of <code>[RetryPolicies](#module_resolve-ip--ResolveIP.RetryPolicies)</code>  
<a name="module_resolve-ip--ResolveIP.Errors"></a>

#### ResolveIP.Errors
Errors returned by methods.

**Kind**: static property of <code>[ResolveIP](#exp_module_resolve-ip--ResolveIP)</code>  

* [.Errors](#module_resolve-ip--ResolveIP.Errors)
    * [.BadRequest](#module_resolve-ip--ResolveIP.Errors.BadRequest) ⇐ <code>Error</code>
    * [.InternalError](#module_resolve-ip--ResolveIP.Errors.InternalError) ⇐ <code>Error</code>
    * [.NotFound](#module_resolve-ip--ResolveIP.Errors.NotFound) ⇐ <code>Error</code>

<a name="module_resolve-ip--ResolveIP.Errors.BadRequest"></a>

##### Errors.BadRequest ⇐ <code>Error</code>
BadRequest

**Kind**: static class of <code>[Errors](#module_resolve-ip--ResolveIP.Errors)</code>  
**Extends:** <code>Error</code>  
**Properties**

| Name | Type |
| --- | --- |
| message | <code>string</code> | 

<a name="module_resolve-ip--ResolveIP.Errors.InternalError"></a>

##### Errors.InternalError ⇐ <code>Error</code>
InternalError

**Kind**: static class of <code>[Errors](#module_resolve-ip--ResolveIP.Errors)</code>  
**Extends:** <code>Error</code>  
**Properties**

| Name | Type |
| --- | --- |
| message | <code>string</code> | 

<a name="module_resolve-ip--ResolveIP.Errors.NotFound"></a>

##### Errors.NotFound ⇐ <code>Error</code>
NotFound

**Kind**: static class of <code>[Errors](#module_resolve-ip--ResolveIP.Errors)</code>  
**Extends:** <code>Error</code>  
**Properties**

| Name | Type |
| --- | --- |
| message | <code>string</code> | 


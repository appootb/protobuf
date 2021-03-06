
# block.proto API Document

*Document generated by protoc-gen-markdown. DO NOT EDIT!*

> APIs


* [Block](#block) - Block service.


	* [Add (/relation/block/{id})](#add) - Block the specified account ID.


	* [Cancel (/relation/block/{id})](#cancel) - Unblock the specified account ID.





<h2 id="block">Block</h2>

>  Block service.



<h3 id="add">Add</h3>

>  Block the specified account ID.



* HTTP Gateway

	* URL: `/relation/block/{id}`
	* Method: `POST`
	* Content-Type: `application/json`

* Request Type: ***UniqueId***

>  Unique ID.

|Field|proto type|JSON type|Comment|Default|Required|
|---|---|---|---|---|---|
|id|int64|string| Unique ID|-|true|




> JSON Demo

```json
{
  "id": "string($int64)"
}
```



* Response Type: ***Empty***

>  A generic empty message that you can re-use to avoid defining duplicated
>  empty messages in your APIs. A typical example is to use it as the request
>  or the response type of an API method. For instance:
> 
>      service Foo {
>        rpc Bar(google.protobuf.Empty) returns (google.protobuf.Empty);
>      }
> 
>  The JSON representation for `Empty` is empty JSON object `{}`.

|Field|proto type|JSON type|Comment|Default|Required|
|---|---|---|---|---|---|



> JSON Demo

```json
{}
```




<h3 id="cancel">Cancel</h3>

>  Unblock the specified account ID.



* HTTP Gateway

	* URL: `/relation/block/{id}`
	* Method: `DELETE`


* Request Type: ***UniqueId***

>  Unique ID.

|Field|proto type|JSON type|Comment|Default|Required|
|---|---|---|---|---|---|
|id|int64|string| Unique ID|-|true|






* Response Type: ***Empty***

>  A generic empty message that you can re-use to avoid defining duplicated
>  empty messages in your APIs. A typical example is to use it as the request
>  or the response type of an API method. For instance:
> 
>      service Foo {
>        rpc Bar(google.protobuf.Empty) returns (google.protobuf.Empty);
>      }
> 
>  The JSON representation for `Empty` is empty JSON object `{}`.

|Field|proto type|JSON type|Comment|Default|Required|
|---|---|---|---|---|---|



> JSON Demo

```json
{}
```






********

## *Embed Messages*















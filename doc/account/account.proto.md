
# account.proto API Document

*Document generated by protoc-gen-markdown. DO NOT EDIT!*

> APIs


* [Account](#account) - Account info service.


	* [GetInfo (/account/info/{id})](#getinfo) - Get account info.


	* [GetInfos (/account/infos)](#getinfos) - Get account infos.


	* [UpdateInfo (/account/info)](#updateinfo) - Update account info.





<h2 id="account">Account</h2>

>  Account info service.



<h3 id="getinfo">GetInfo</h3>

>  Get account info.



* HTTP Gateway

	* URL: `/account/info/{id}`
	* Method: `GET`


* Request Type: ***UniqueId***

>  Unique ID (uint64 type).

|Field|proto type|JSON type|Comment|Default|Required|
|---|---|---|---|---|---|
|id|int64|string| ID value|-|true|






* Response Type: ***AccountInfo***

>  Account information.

|Field|proto type|JSON type|Comment|Default|Required|
|---|---|---|---|---|---|
|unique_id|int64|string| Account ID|-|true|
|nickname|string|string| Nickname|-|true|
|avatar|string|string| Avatar|-|true|
|signature|string|string| Account signature|-|false|
|gender|enum [Gender](#gender)|string/integer| Gender|-|false|
|signs|string|string| Zodiac signs|-|false|
|location|string|string| Location|-|false|
|secret|[Secret](#secret)|object| Account secret info|-|false|
|create_at|[Timestamp](#timestamp)|string ("1972-01-01T10:00:20.021Z")| Account created time|-|false|
|extend|[Any](#any)|object| Account extend info|-|false|



> JSON Demo

```json
{
  "unique_id": "string($int64)",
  "nickname": "string",
  "avatar": "string",
  "signature": "string",
  "gender": "GENDER_UNSPECIFIED (0) | GENDER_MALE (1) | GENDER_FEMALE (2) | GENDER_OTHER (3)",
  "signs": "string",
  "location": "string",
  "secret": {
    "token": "string"
  },
  "create_at": "1972-01-01T10:00:20.021Z",
  "extend": {
    "type_url": "string",
    "value": "YmFzZTY0IHN0cmluZw=="
  }
}
```




<h3 id="getinfos">GetInfos</h3>

>  Get account infos.



* HTTP Gateway

	* URL: `/account/infos`
	* Method: `GET`


* Request Type: ***UniqueIds***

>  Unique IDs (uint64 array).

|Field|proto type|JSON type|Comment|Default|Required|
|---|---|---|---|---|---|
|ids|array [int64]|string| ID list|-|false|






* Response Type: ***AccountInfos***

>  Accounts information.

|Field|proto type|JSON type|Comment|Default|Required|
|---|---|---|---|---|---|
|infos|array [[AccountInfo](#accountinfo)]|array| Accounts|-|false|



> JSON Demo

```json
{
  "infos": [
    {
      "unique_id": "string($int64)",
      "nickname": "string",
      "avatar": "string",
      "signature": "string",
      "gender": "GENDER_UNSPECIFIED (0) | GENDER_MALE (1) | GENDER_FEMALE (2) | GENDER_OTHER (3)",
      "signs": "string",
      "location": "string",
      "secret": {
        "token": "string"
      },
      "create_at": "1972-01-01T10:00:20.021Z",
      "extend": {
        "type_url": "string",
        "value": "YmFzZTY0IHN0cmluZw=="
      }
    }
  ]
}
```




<h3 id="updateinfo">UpdateInfo</h3>

>  Update account info.



* HTTP Gateway

	* URL: `/account/info`
	* Method: `PATCH`


* Request Type: ***VariableInfo***

>  Update account info request.

|Field|proto type|JSON type|Comment|Default|Required|
|---|---|---|---|---|---|
|nickname|string|string| Nickname|-|false|
|avatar|string|string| Avatar|-|false|
|signature|string|string| Signature|-|false|
|gender|enum [Gender](#gender)|string/integer| Gender|-|false|






* Response Type: ***AccountInfo***

>  Account information.

|Field|proto type|JSON type|Comment|Default|Required|
|---|---|---|---|---|---|
|unique_id|int64|string| Account ID|-|true|
|nickname|string|string| Nickname|-|true|
|avatar|string|string| Avatar|-|true|
|signature|string|string| Account signature|-|false|
|gender|enum [Gender](#gender)|string/integer| Gender|-|false|
|signs|string|string| Zodiac signs|-|false|
|location|string|string| Location|-|false|
|secret|[Secret](#secret)|object| Account secret info|-|false|
|create_at|[Timestamp](#timestamp)|string ("1972-01-01T10:00:20.021Z")| Account created time|-|false|
|extend|[Any](#any)|object| Account extend info|-|false|



> JSON Demo

```json
{
  "unique_id": "string($int64)",
  "nickname": "string",
  "avatar": "string",
  "signature": "string",
  "gender": "GENDER_UNSPECIFIED (0) | GENDER_MALE (1) | GENDER_FEMALE (2) | GENDER_OTHER (3)",
  "signs": "string",
  "location": "string",
  "secret": {
    "token": "string"
  },
  "create_at": "1972-01-01T10:00:20.021Z",
  "extend": {
    "type_url": "string",
    "value": "YmFzZTY0IHN0cmluZw=="
  }
}
```






********

## *Embed Messages*





<h3 id="gender">Gender</h3>

>  Account gender enums.

* Enum

|Name (string)|Value (integer)|Comment|
|---|---|---|
|GENDER_UNSPECIFIED|0| Unspecified|
|GENDER_MALE|1| Male|
|GENDER_FEMALE|2| Female|
|GENDER_OTHER|3| Other|








<h3 id="accountinfo">AccountInfo</h3>

>  Account information.

* Fields

|Field|proto type|JSON type|Comment|Default|Required|
|---|---|---|---|---|---|
|unique_id|int64|string| Account ID|-|true|
|nickname|string|string| Nickname|-|true|
|avatar|string|string| Avatar|-|true|
|signature|string|string| Account signature|-|false|
|gender|enum [Gender](#gender)|string/integer| Gender|-|false|
|signs|string|string| Zodiac signs|-|false|
|location|string|string| Location|-|false|
|secret|[Secret](#secret)|object| Account secret info|-|false|
|create_at|[Timestamp](#timestamp)|string ("1972-01-01T10:00:20.021Z")| Account created time|-|false|
|extend|[Any](#any)|object| Account extend info|-|false|


<h3 id="secret">Secret</h3>

>  Account secret.

* Fields

|Field|proto type|JSON type|Comment|Default|Required|
|---|---|---|---|---|---|
|token|string|string| Account token|-|true|


<h3 id="any">Any</h3>

>  `Any` contains an arbitrary serialized protocol buffer message along with a
>  URL that describes the type of the serialized message.
> 
>  Protobuf library provides support to pack/unpack Any values in the form
>  of utility functions or additional generated methods of the Any type.
> 
>  Example 1: Pack and unpack a message in C++.
> 
>      Foo foo = ...;
>      Any any;
>      any.PackFrom(foo);
>      ...
>      if (any.UnpackTo(&foo)) {
>        ...
>      }
> 
>  Example 2: Pack and unpack a message in Java.
> 
>      Foo foo = ...;
>      Any any = Any.pack(foo);
>      ...
>      if (any.is(Foo.class)) {
>        foo = any.unpack(Foo.class);
>      }
> 
>   Example 3: Pack and unpack a message in Python.
> 
>      foo = Foo(...)
>      any = Any()
>      any.Pack(foo)
>      ...
>      if any.Is(Foo.DESCRIPTOR):
>        any.Unpack(foo)
>        ...
> 
>   Example 4: Pack and unpack a message in Go
> 
>       foo := &pb.Foo{...}
>       any, err := ptypes.MarshalAny(foo)
>       ...
>       foo := &pb.Foo{}
>       if err := ptypes.UnmarshalAny(any, foo); err != nil {
>         ...
>       }
> 
>  The pack methods provided by protobuf library will by default use
>  'type.googleapis.com/full.type.name' as the type URL and the unpack
>  methods only use the fully qualified type name after the last '/'
>  in the type URL, for example "foo.bar.com/x/y.z" will yield type
>  name "y.z".
> 
> 
>  JSON
>  ====
>  The JSON representation of an `Any` value uses the regular
>  representation of the deserialized, embedded message, with an
>  additional field `@type` which contains the type URL. Example:
> 
>      package google.profile;
>      message Person {
>        string first_name = 1;
>        string last_name = 2;
>      }
> 
>      {
>        "@type": "type.googleapis.com/google.profile.Person",
>        "firstName": <string>,
>        "lastName": <string>
>      }
> 
>  If the embedded message type is well-known and has a custom JSON
>  representation, that representation will be embedded adding a field
>  `value` which holds the custom JSON in addition to the `@type`
>  field. Example (for message [google.protobuf.Duration][]):
> 
>      {
>        "@type": "type.googleapis.com/google.protobuf.Duration",
>        "value": "1.212s"
>      }

* Fields

|Field|proto type|JSON type|Comment|Default|Required|
|---|---|---|---|---|---|
|type_url|string|string| A URL/resource name that uniquely identifies the type of the serialized<br> protocol buffer message. This string must contain at least<br> one "/" character. The last segment of the URL's path must represent<br> the fully qualified name of the type (as in<br> `path/google.protobuf.Duration`). The name should be in a canonical form<br> (e.g., leading "." is not accepted).<br><br> In practice, teams usually precompile into the binary all types that they<br> expect it to use in the context of Any. However, for URLs which use the<br> scheme `http`, `https`, or no scheme, one can optionally set up a type<br> server that maps type URLs to message definitions as follows:<br><br> * If no scheme is provided, `https` is assumed.<br> * An HTTP GET on the URL must yield a [google.protobuf.Type][]<br>   value in binary format, or produce an error.<br> * Applications are allowed to cache lookup results based on the<br>   URL, or have them precompiled into a binary to avoid any<br>   lookup. Therefore, binary compatibility needs to be preserved<br>   on changes to types. (Use versioned type names to manage<br>   breaking changes.)<br><br> Note: this functionality is not currently available in the official<br> protobuf release, and it is not used for type URLs beginning with<br> type.googleapis.com.<br><br> Schemes other than `http`, `https` (or the empty scheme) might be<br> used with implementation specific semantics.|-|false|
|value|bytes|base64 string| Must be a valid serialized protocol buffer of the above specified type.|-|false|


<h3 id="timestamp">Timestamp</h3>

>  A Timestamp represents a point in time independent of any time zone or local
>  calendar, encoded as a count of seconds and fractions of seconds at
>  nanosecond resolution. The count is relative to an epoch at UTC midnight on
>  January 1, 1970, in the proleptic Gregorian calendar which extends the
>  Gregorian calendar backwards to year one.
> 
>  All minutes are 60 seconds long. Leap seconds are "smeared" so that no leap
>  second table is needed for interpretation, using a [24-hour linear
>  smear](https://developers.google.com/time/smear).
> 
>  The range is from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59.999999999Z. By
>  restricting to that range, we ensure that we can convert to and from [RFC
>  3339](https://www.ietf.org/rfc/rfc3339.txt) date strings.
> 
>  # Examples
> 
>  Example 1: Compute Timestamp from POSIX `time()`.
> 
>      Timestamp timestamp;
>      timestamp.set_seconds(time(NULL));
>      timestamp.set_nanos(0);
> 
>  Example 2: Compute Timestamp from POSIX `gettimeofday()`.
> 
>      struct timeval tv;
>      gettimeofday(&tv, NULL);
> 
>      Timestamp timestamp;
>      timestamp.set_seconds(tv.tv_sec);
>      timestamp.set_nanos(tv.tv_usec * 1000);
> 
>  Example 3: Compute Timestamp from Win32 `GetSystemTimeAsFileTime()`.
> 
>      FILETIME ft;
>      GetSystemTimeAsFileTime(&ft);
>      UINT64 ticks = (((UINT64)ft.dwHighDateTime) << 32) | ft.dwLowDateTime;
> 
>      // A Windows tick is 100 nanoseconds. Windows epoch 1601-01-01T00:00:00Z
>      // is 11644473600 seconds before Unix epoch 1970-01-01T00:00:00Z.
>      Timestamp timestamp;
>      timestamp.set_seconds((INT64) ((ticks / 10000000) - 11644473600LL));
>      timestamp.set_nanos((INT32) ((ticks % 10000000) * 100));
> 
>  Example 4: Compute Timestamp from Java `System.currentTimeMillis()`.
> 
>      long millis = System.currentTimeMillis();
> 
>      Timestamp timestamp = Timestamp.newBuilder().setSeconds(millis / 1000)
>          .setNanos((int) ((millis % 1000) * 1000000)).build();
> 
> 
>  Example 5: Compute Timestamp from current time in Python.
> 
>      timestamp = Timestamp()
>      timestamp.GetCurrentTime()
> 
>  # JSON Mapping
> 
>  In JSON format, the Timestamp type is encoded as a string in the
>  [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format. That is, the
>  format is "{year}-{month}-{day}T{hour}:{min}:{sec}[.{frac_sec}]Z"
>  where {year} is always expressed using four digits while {month}, {day},
>  {hour}, {min}, and {sec} are zero-padded to two digits each. The fractional
>  seconds, which can go up to 9 digits (i.e. up to 1 nanosecond resolution),
>  are optional. The "Z" suffix indicates the timezone ("UTC"); the timezone
>  is required. A proto3 JSON serializer should always use UTC (as indicated by
>  "Z") when printing the Timestamp type and a proto3 JSON parser should be
>  able to accept both UTC and other timezones (as indicated by an offset).
> 
>  For example, "2017-01-15T01:30:15.01Z" encodes 15.01 seconds past
>  01:30 UTC on January 15, 2017.
> 
>  In JavaScript, one can convert a Date object to this format using the
>  standard
>  [toISOString()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString)
>  method. In Python, a standard `datetime.datetime` object can be converted
>  to this format using
>  [`strftime`](https://docs.python.org/2/library/time.html#time.strftime) with
>  the time format spec '%Y-%m-%dT%H:%M:%S.%fZ'. Likewise, in Java, one can use
>  the Joda Time's [`ISODateTimeFormat.dateTime()`](
>  http://www.joda.org/joda-time/apidocs/org/joda/time/format/ISODateTimeFormat.html#dateTime%2D%2D
>  ) to obtain a formatter capable of generating timestamps in this format.

* Fields

|Field|proto type|JSON type|Comment|Default|Required|
|---|---|---|---|---|---|
|seconds|int64|string| Represents seconds of UTC time since Unix epoch<br> 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to<br> 9999-12-31T23:59:59Z inclusive.|-|false|
|nanos|int|number/string| Non-negative fractions of a second at nanosecond resolution. Negative<br> second values with fractions must still have non-negative nanos values<br> that count forward in time. Must be from 0 to 999,999,999<br> inclusive.|-|false|





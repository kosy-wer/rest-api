# RestSharp - Simple .NET REST Client 

![](https://img.shields.io/nuget/dt/RestSharp) [![](https://img.shields.io/nuget/v/RestSharp)](https://www.nuget.org/packages/RestSharp) [![](https://img.shields.io/nuget/vpre/RestSharp)](https://www.nuget.org/packages/RestSharp#versions-body-tab)

RestSharp is a lightweight HTTP API client library. It's a wrapper around `HttpClient`, not a full-fledged client on 
its own.

What RestSharp adds to `HttpClient`:
- Default parameters of any kind, not just headers
- Add a parameter of any kind to requests, like query, URL segment, header, cookie, or body
- Multiple ways to add a request body, including JSON, XML, URL-encoded form data, multipart form data with and 
  without files
- Built-in serialization and deserilization of JSON, XML, and CSV, as well as the ability to add custom serializers
- Rich support for authentication

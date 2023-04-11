---
title: WunderGraph open telemetry Example
pageTitle: WunderGraph - Examples - Open Telemetry
description:
---

[Check the example](https://github.com/wundergraph/wundergraph/tree/main/examples/open-telemetry)

## Configuration

Check the `wundergraph.config.ts` confile:

```typescript
// wundergraph.config.ts
configureWunderGraphApplication({
  apis: [spacex],
  options: {
    openTelemetry: {
      enabled: true, // default: false
      exporterHttpEndpoint: '', // standard otel http exporter endpoint
      exporterJaegerEndpoint: '', // 'http://localhost:14268/api/traces' we recommend to use it for development
      authToken: '', // jwt for authentication, format: 'Bearer ...', use for development only, adds authentication header to the exporter
    },
  },
  // ...
});
```

Empty values could be omitted.

## Getting started

Install the dependencies and run the complete example in one command:

```shell
npm install && npm start
```

## Get Dragons

```shell
curl -H "traceparent: 00-80e1afed08e019fc1110464cfa66635c-7a085853722dc6d2-01" http://localhost:9991/operations/Dragons
```

## Test the tracing

You can access Jaeger from http://localhost:16686/

To obtain meaningful results, it is recommended to change the trace id in the request header every time you send a request.
If you don't provide a traceparent header, a new trace id will be generated automatically.

Header format: `traceparent: {version}-{trace_id}-{span_id}-{trace_flags}`
https://www.w3.org/TR/trace-context/#trace-context-http-headers-format

Configure hooks in the `wundergraph.server.ts` and check how the tracing data changes.

## Environment variables configuration

The Open Telemetry configuration can also be set via environment variables.

```shell
  WG_OTEL_ENABLED=true
  WG_OTEL_EXPORTER_HTTP_ENDPOINT=''
  WG_OTEL_EXPORTER_JAEGER_ENDPOINT='http://localhost:14268/api/traces'
  WG_OTEL_JWT='Bearer ...' # production mode
```

Empty values could be omitted.
Please note that `wundergraph.config.ts` configuration has higher priority than environment variables.
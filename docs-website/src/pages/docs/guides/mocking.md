---
title: Mocking
pageTitle: WunderGraph - Testing
description: This guide shows how to mock your WunderGraph HTTP datasources or APIs
---

## Mocking

WunderGraph testing library comes with builtin support for mocking your HTTP datasources or APIs.
This allows you to test your application without having to rely on external services.

Our testing library is framework agnostic, so you can use it with your preferred test
framework. We use [Vitest](https://vitest.dev/) in this guide. It is a Jest Compatible test runner, but it should be easily adaptable to any testing framework.

### Initial setup

Make sure your `wundergraph.config.ts` is set up to generate the WunderGraph testing library. Take a closer look at the `url`. We use an environment variable to set the URL of the datasource. This allows us to replace the URL with the mock server URL when running the tests.

```typescript
const countries = introspect.graphql({
  apiNamespace: 'countries',
  url: new EnvironmentVariable('COUNTRIES_URL', 'https://countries.trevorblades.com/'),
});

configureWunderGraphApplication({
  apis: [countries],
  server,
  operations,
  codeGenerators: [
    {
      templates: [
        // Generate all TypeScript based templates, which includes the testing library
        ...templates.typescript.all,
      ],
    },
  ],
});
```

### Test setup

Within a test, you can import `createTestAndMockServer()` to create the test and mock server instances.
Note that depending on your project's settings, the path might be slightly different depending on where you generate your templates.
We use `beforeAll` to start the test and mock server instances. The `ts.start()` method returns a cleanup function to shutdown all servers. In vitest you can return an async function to do a cleanup. We also use `mockedAPIs` to replace the environment variables with the mock server URL.

{% callout type="note" %}
Before you run the tests and import the `createTestAndMockServer` function, you must first run `wundergraph generate` to generate the testing library.
Some testing frameworks like Vitest or Jest provides a `globalSetup` hook that allows you to run a script before running the tests. In that way, you can automate it.
{% /callout %}

```ts
import { createTestAndMockServer, TestServers } from '../.wundergraph/generated/testing';

let ts: TestServers;

beforeAll(async () => {
  ts = createTestAndMockServer({
    // The directory where your wundergraph directory is located
    dir: join(__dirname, '..'),
  });

  return ts.start({
    // Environment variables replaced by the test mock server URL
    mockedAPIs: ['COUNTRIES_URL'],
  });
});
```

### Writing tests

We recommended creating as few testing server instances as possible. These minimizes the number of times the server starts and stops, making your tests faster.

{% callout type="warning" %}
Due to the nature of the mock server, it is not possible right now to run multiple tests `it` concurrently.
If you have a demanding test suite, you can create multiple test files and vitest will run them in parallel by default.
{% /callout %}

#### Mocking HTTP datasources

When you setup a mock with `mock` and the request match, the interceptor is removed internally from the list. This means that you can intercept **2** or more calls to the same URL and return different things on each of them. It also means that you must setup one interceptor for each request you are going to have otherwise it will throw an error because that URL was not present in the interceptor list.
The first argument of the `mock` function is a predicate that is used to match the request. The second argument is the function that returns the response. You can use test assertions inside the response function to verify that the request is correct. If an error is thrown the mock is considered as failed. The error is propagated to the `scope.done()` call.

```ts
const scope = ts.mockServer.mock(
  async ({ url, method }) => {
    return url.path === '/' && method === 'POST';
  },
  async ({ json }) => {
    const body = await json();

    expect(body.variables.code).toEqual('ES');
    expect(body.query).toEqual(
      'query($code: String){countries_countries: countries(filter: {code: {eq: $code}}){code name capital}}'
    );

    return {
      status: 200,
      headers: {
        'X-Foo': 'Bar',
      },
      body: {
        data: {
          countries_countries: [
            {
              code: 'ES',
              name: 'Spain',
              capital: 'Madrid',
            },
          ],
        },
      },
    };
  }
);
```

### Make a request and validate the mock

The `createTestAndMockServer` returns a `WunderGraphTestServer` object that wraps the test server and the auto-generated type-safe client.

```ts
const result = await ts.testServer.client().query({
  operationName: 'CountryByCode',
  input: {
    code: 'ES',
  },
});

/**
 * Make sure to call scope.done() to verify that the request have been made
 * If the request was not made, the test will fail
 */
scope.done();

expect(result.error).toBeUndefined();
expect(result.data).toBeDefined();
expect(result.data?.countries_countries?.[0].capital).toBe('Madrid');
```

For a full example please check the example in the [WunderGraph repository](https://github.com/wundergraph/wundergraph/tree/main/packages/testsuite/apps/mock/test/mock-datasource.test.ts)

## Summary

This guide showed how easy it is to mock HTTP datasources. It gives you power to write integration tests fast and reliable.
You can apply the same principles to mock any HTTP API in your TypeScript functions.
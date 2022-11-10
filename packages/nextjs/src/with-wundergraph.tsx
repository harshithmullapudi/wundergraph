import * as React from 'react';
import ssrPrepass from 'react-ssr-prepass';
import { NextComponentType } from 'next';
import { AppContextType, AppPropsType, NextPageContext } from 'next/dist/shared/lib/utils';
import { NextRouter } from 'next/router';
import { User } from '@wundergraph/sdk/client';
import { serialize } from '@wundergraph/sdk/internal';
import { userSWRKey, SWRConfig } from '@wundergraph/swr';
import { BareFetcher, PublicConfiguration, SWRHook } from 'swr/_internal';
import { WithWunderGraphOptions, SSRCache } from './types';

import { useWunderGraphContext, WunderGraphProvider } from './context';
import { Key, Middleware } from 'swr';

const getOperationNameFromKey = (key: Key) => {
	if (key && !Array.isArray(key) && typeof key === 'object' && 'operationName' in key) {
		return key.operationName;
	}
};

interface SSRConfig extends PublicConfiguration {
	ssr?: boolean;
}

const SSRMiddleWare = ((useSWRNext: SWRHook) => {
	return <Data, Error>(key: Key, fetcher: BareFetcher, config: SSRConfig) => {
		const swr = useSWRNext(key, fetcher, config);

		const isSSR = typeof window === 'undefined' && config.ssr !== false;
		const context = useWunderGraphContext();

		const operationName = getOperationNameFromKey(key);

		if (!operationName || !context || !isSSR || !key) {
			return swr;
		}

		const _key = serialize(key);

		const { ssrCache, client, user } = context;

		const shouldAuthenticate = client.isAuthenticatedOperation(operationName) && !user;

		if (ssrCache && !ssrCache[_key] && fetcher && !shouldAuthenticate) {
			ssrCache[_key] = fetcher(key);
		}

		return swr;
	};
}) as unknown as Middleware;

export const withWunderGraph = (options: WithWunderGraphOptions) => {
	return (AppOrPage: NextComponentType<any, any, any>): NextComponentType => {
		const { client, userCacheKey = userSWRKey, context } = options;

		const WithWunderGraph = (props: AppPropsType<NextRouter, any> & { ssrCache: SSRCache; user: User }) => {
			const { ssrCache = {}, user } = props;

			return (
				<WunderGraphProvider context={context} value={{ ssrCache, client, user }}>
					<SWRConfig value={{ fallback: ssrCache, use: [SSRMiddleWare] }}>
						<AppOrPage {...props} />
					</SWRConfig>
				</WunderGraphProvider>
			);
		};

		if (AppOrPage.getInitialProps || options.ssr) {
			WithWunderGraph.getInitialProps = async (appOrPageCtx: AppContextType) => {
				const AppTree = appOrPageCtx.AppTree;
				const isApp = !!appOrPageCtx.Component;
				const ctx = isApp ? appOrPageCtx.ctx : (appOrPageCtx as any as NextPageContext);

				const ssrCache: SSRCache = {};

				let pageProps: Record<string, unknown> = {};
				if (AppOrPage.getInitialProps) {
					const originalProps = await AppOrPage.getInitialProps(appOrPageCtx);

					const originalPageProps = isApp ? originalProps.pageProps ?? {} : originalProps;

					pageProps = {
						...originalPageProps,
						...pageProps,
					};
				}

				const getAppTreeProps = (props: Record<string, unknown>) => (isApp ? { pageProps: props } : props);

				if (typeof window !== 'undefined' || !options.ssr) {
					// we're on the client
					// no need to do all the SSR stuff.
					return getAppTreeProps(pageProps);
				}

				const cookieHeader = ctx.req?.headers.cookie;
				if (typeof cookieHeader === 'string') {
					client.setExtraHeaders({
						Cookie: cookieHeader,
					});
				}

				let ssrUser: User | null = null;

				if (options?.fetchUserSSR !== false) {
					try {
						ssrUser = await client.fetchUser();
					} catch (e) {}
				}

				const start = options?.logPrerenderTime ? process.hrtime() : undefined;

				if (ssrUser) {
					ssrCache[userCacheKey] = ssrUser;
				}

				const prepassProps = {
					pageProps: {
						ssrCache,
						user: ssrUser,
						...pageProps,
					},
				};

				await ssrPrepass(<AppTree {...prepassProps} />);

				const keys = Object.keys(ssrCache)
					.filter((key) => typeof ssrCache[key].then === 'function')
					.map((key) => ({
						key,
						value: ssrCache[key],
					})) as { key: string; value: Promise<any> }[];
				if (keys.length !== 0) {
					const promises = keys.map((key) => key.value);
					const results = await Promise.all(promises);
					for (let i = 0; i < keys.length; i++) {
						const key = keys[i].key;
						ssrCache[key] = results[i];
					}
				}

				if (options?.logPrerenderTime && start) {
					const precision = 3; // 3 decimal places
					const elapsed = process.hrtime(start)[1] / 1000000; // divide by a million to get nano to milli
					console.log(process.hrtime(start)[0] + ' s, ' + elapsed.toFixed(precision) + ' ms - render'); // print message + time
				}

				return getAppTreeProps({ ...pageProps, ssrCache, user: ssrUser });
			};
		}

		WithWunderGraph.displayName = AppOrPage.displayName || AppOrPage.name || 'WithWunderGraph';
		if ((AppOrPage as any).getLayout) {
			// @ts-ignore
			WithWunderGraph.getLayout = AppOrPage.getLayout;
		}

		return WithWunderGraph as any;
	};
};
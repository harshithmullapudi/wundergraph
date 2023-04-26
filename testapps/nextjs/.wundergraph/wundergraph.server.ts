import { configureWunderGraphServer } from '@wundergraph/sdk/server';
import { GraphQLEnumType, GraphQLList, GraphQLNonNull, GraphQLObjectType, GraphQLSchema, GraphQLString } from 'graphql';

const testEnum = new GraphQLEnumType({
	name: 'TestEnum',
	values: {
		EnumValueA: {
			value: 'EnumValueA',
		},
		EnumValueB: {
			value: 'EnumValueA',
		},
	},
});

export default configureWunderGraphServer(() => ({
	hooks: {
		authentication: {
			mutatingPostAuthentication: async (hook) => {
				console.log('mutatingPostAuthentication', JSON.stringify(hook.user));
				return {
					status: 'ok',
					user: {
						...hook.user,
						roles: ['user', 'admin'],
					},
				};
			},
			postLogout: async (hook) => {
				console.log('postLogout', JSON.stringify(hook.user));
			},
		},
		queries: {
			Hello: {
				preResolve: async (hook) => {
					console.log('###preResolve', hook);
				},
				mutatingPreResolve: async (hook) => {
					console.log('###mutatingPreResolve', hook);
					return hook.input;
				},
				postResolve: async (hook) => {
					console.log('###postResolve', hook);
				},
				mutatingPostResolve: async (hook) => {
					console.log('###mutatingPostResolve', hook);

					hook.response?.errors;

					const fakeWeather = await hook.internalClient.queries.FakeWeather();

					const { data, error } = await hook.operations.query({
						operationName: 'Hello',
						input: {
							hello: '1',
						},
					});

					const weather = await hook.operations.query({
						operationName: 'users/get',
						input: {
							id: 1,
						},
					});
					console.log(weather);
					return {
						data: {
							gql_hello: weather.data,
						},
						errors: [weather.error],
					};
				},
			},
			FakeWeather: {
				mockResolve: async (hook) => {
					const randomWeatherDescription = ['sunny', 'cloudy', 'rainy', 'stormy', 'snowy', 'windy'];
					const randomTemperatureCelsius = Math.floor(Math.random() * 40) - 20;
					// '0°, cloudy'
					const description = `${randomTemperatureCelsius}°, ${
						randomWeatherDescription[Math.floor(Math.random() * randomWeatherDescription.length)]
					}`;
					return {
						data: {
							getCityByName: {
								id: '1',
								name: 'Berlin',
								weather: {
									summary: {
										title: 'Weather for Berlin',
										description,
									},
								},
							},
						},
					};
				},
			},
		},
		mutations: {
			SetName: {
				preResolve: async (hook) => {
					console.log('###preResolve', hook);
				},
				mutatingPreResolve: async (hook) => {
					console.log('###mutatingPreResolve', hook);
					return hook.input;
				},
				postResolve: async (hook) => {
					console.log('###postResolve', hook);
				},
				mutatingPostResolve: async (hook) => {
					console.log('###mutatingPostResolve', hook);
					return hook.response;
				},
			},
		},
	},
	graphqlServers: [
		{
			apiNamespace: 'gql',
			serverName: 'gql',
			schema: new GraphQLSchema({
				query: new GraphQLObjectType({
					name: 'Query',
					fields: {
						hello: {
							type: GraphQLString,
							args: {
								name: {
									type: new GraphQLNonNull(GraphQLString),
								},
							},
							resolve(root, args) {
								return args.name + 'world';
							},
						},
						testField: {
							type: new GraphQLObjectType({
								name: 'TestResponse',
								fields: {
									enumList: {
										type: new GraphQLNonNull(new GraphQLList(testEnum)),
									},
									stringList: {
										type: new GraphQLList(GraphQLString),
									},
								},
							}),
							args: {
								inputEnum: {
									type: new GraphQLNonNull(new GraphQLList(testEnum)),
								},
							},
							resolve: (source, args, context, info) => {
								return {
									enumList: args.inputEnum,
									stringList: ['a', 'b', 'c'],
								};
							},
						},
					},
				}),
				mutation: new GraphQLObjectType<any, any>({
					name: 'Mutation',
					fields: {
						setName: {
							type: GraphQLString,
							args: {
								name: {
									type: new GraphQLNonNull(GraphQLString),
								},
							},
							resolve: async (source, args, context, info) => {
								console.log(JSON.stringify({ setName: args }));
								return args.name;
							},
						},
					},
				}),
			}),
		},
	],
}));

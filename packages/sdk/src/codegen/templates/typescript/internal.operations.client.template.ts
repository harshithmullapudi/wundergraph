// TODO: Unify error handling

//language=handlebars
export const template = `
import type { OperationsClient, InternalOperationsDefinition } from "@wundergraph/sdk/server";
import type { ClientOperationErrors } from "@wundergraph/sdk/client";
import type { OperationErrors } from "./ts-operation-errors";
import { {{ modelImports }} } from "./models"

export type Queries = {
{{#if hasInternalQueries}}
    {{#each internalQueries}}
        '{{operationPath}}': {
					input: {{#if hasInternalInput}}Internal{{operationName}}Input{{ else }}undefined{{/if}},
          response: {{#if isTypeScriptOperation}}{ data?: {{operationName}}ResponseData, error?: OperationErrors['{{operationPath}}'] }{{else}}{ data?: {{operationName}}Response['data'], error?: ClientOperationErrors }{{/if}}
				};
    {{/each}}
{{/if}}
}

export type Mutations = {
{{#if hasInternalMutations}}
    {{#each internalMutations}}
			'{{operationPath}}': {
				input: {{#if hasInternalInput}}Internal{{operationName}}Input{{ else }}undefined{{/if}},
      	response: {{#if isTypeScriptOperation}}{ data?: {{operationName}}ResponseData, error?: OperationErrors['{{operationPath}}'] }{{else}}{ data?: {{operationName}}Response['data'], error?: ClientOperationErrors }{{/if}}
			};
    {{/each}}
{{/if}}
}

export type Subscriptions = {
{{#if hasInternalSubscriptions}}
		{{#each internalSubscriptions}}
			'{{operationPath}}': {
				input: {{#if hasInternalInput}}Internal{{operationName}}Input{{ else }}undefined{{/if}},
      	response: {{#if isTypeScriptOperation}}{ data?: {{operationName}}ResponseData, error?: OperationErrors['{{operationPath}}'] }{{else}}{ data?: {{operationName}}Response['data'], error?: ClientOperationErrors }{{/if}}
			};
		{{/each}}
{{/if}}
}

export type InternalOperations = InternalOperationsDefinition<Queries, Mutations, Subscriptions>

export type InternalOperationsClient = OperationsClient<InternalOperations>
`;

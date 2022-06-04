# TheGraph Go code generation

This folder is to generate type-safe Go code to request TheGraph API.

## genqlient

The Go-code generator is [genqlient](https://github.com/Khan/genqlient):

- Compile-time validation of GraphQL queries: no invalid GraphQL query again!
- Type-safe response objects: genqlient generates the right type for each query,
  so you know the response will unmarshal correctly and never need to use `interface{}`.

## Step 1: Download the GraphQL schema

Put it in `pkg/provider/the-graph/<provider-name>/schema-xxxx.graphql`.
Format = [Schema Definition Language (SDL)](https://graphql.org/learn/schema/#type-language).

```graphql
type Query {
  item(id: ID): Item
  items(): [Item]
}

type Item {
  id: ID!
}
```

## Step 2: Write your GraphQL query

Put it in `pkg/provider/generated/<provider-name>/queries.graphql`.
Format = standard [GraphQL syntax](https://graphql.org/learn/queries/)
(supports queries and mutations).

```graphql
query getUser($login: String!) {
  user(login: $login) {
    name
  }
}
```

💡 Use an interactive explorer like [GraphiQL](https://github.com/graphql/graphiql/tree/main/packages/graphiql#readme).

## Step 3: Configuration file

Create a default configuration file in
`pkg/provider/generated/<provider-name>/config.yaml`:

    cd pkg/provider/generated/<provider-name>
    go run github.com/Khan/genqlient --init config.yaml

Filenames and other aspects of `genqlient` are configurable.
You can also configure how `genqlient` converts specific parts
of your query with the `@genqlient` directive.
See <https://github.com/Khan/genqlient/blob/main/docs/genqlient.yaml>
for the full range of options.

## Step 4: Generate the Go code

From your queries generate `pkg/provider/generated/<provider-name>/generated.go`:

[`go generate ./...`](https://go.dev/blog/generate)

## Step 5: Alternative Go code generation

If you prefer, you can specify your queries as string-constants in your Go source,
prefixed with `# @genqlient`. Example:

```go
const _ = `# @genqlient
  query getUser($login: String!) {
    user(login: $login) {
      name
    }
  }`

resp, err := getUser(...)
```

You don't need to do anything with the constant,
just keep it somewhere in the source as documentation
and for the next time you run genqlient.
In this case you'll need to update `config.yaml`
to tell it to look at your Go code.

## Step 6: Use your queries

The generated code will expose a function with the same name as your query name
like the following example:

```go
func getUser(ctx context.Context, client graphql.Client, login string) (*getUserResponse, error)
```

As for the arguments:

- for `ctx`, pass your local context (see [`go doc context`](https://pkg.go.dev/context)) or `context.Background()` if you don't need one
- for `client`, call [`graphql.NewClient`](https://pkg.go.dev/github.com/Khan/genqlient/graphql), e.g. `graphql.NewClient("https://api.provider.com/xx", http.DefaultClient)`
- for `login`, pass username (or whatever the arguments to your query are)

The response object is a `struct` with fields corresponding to each GraphQL field.
Example:

```go
ctx := context.Background()
client := graphql.NewClient("https://api.github.com/graphql", http.DefaultClient)
resp, err := getUser(ctx, client, "benjaminjkraft")
fmt.Println(resp.User.Name, err)
```

## Step 7: Repeat

Over time, as you add or change queries, you'll just need to run
[`go generate ./...`](https://go.dev/blog/generate)
to re-generate `generated.go`.

If you're using an editor or IDE plugin backed by
[gopls](https://github.com/golang/tools/blob/master/gopls/README.md)
(which is most of them), keep `generated.go` open in the background,
and reload it after each run, so your plugin knows about the automated changes.

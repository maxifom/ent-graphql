package validator

import "github.com/vektah/gqlparser/v2/ast"

var Prelude = &ast.Source{
	Name:    "prelude.graphql",
	Input:   "# This file defines all the implicitly declared types that are required by the graphql spec. It is implicitly included by calls to LoadSchema\n\n\"The `Int` scalar type represents non-fractional signed whole numeric values. Int can represent values between -(2^31) and 2^31 - 1.\"\nscalar Int\n\n\"The `Float` scalar type represents signed double-precision fractional values as specified by [IEEE 754](http://en.wikipedia.org/wiki/IEEE_floating_point).\"\nscalar Float\n\n\"The `String`scalar type represents textual data, represented as UTF-8 character sequences. The String type is most often used by GraphQL to represent free-form human-readable text.\"\nscalar String\n\n\"The `Boolean` scalar type represents `true` or `false`.\"\nscalar Boolean\n\n\"\"\"The `ID` scalar type represents a unique identifier, often used to refetch an object or as key for a cache. The ID type appears in a JSON response as a String; however, it is not intended to be human-readable. When expected as an input type, any string (such as \"4\") or integer (such as 4) input value will be accepted as an ID.\"\"\"\nscalar ID\n\n\"The @include directive may be provided for fields, fragment spreads, and inline fragments, and allows for conditional inclusion during execution as described by the if argument.\"\ndirective @include(if: Boolean!) on FIELD | FRAGMENT_SPREAD | INLINE_FRAGMENT\n\n\"The @skip directive may be provided for fields, fragment spreads, and inline fragments, and allows for conditional exclusion during execution as described by the if argument.\"\ndirective @skip(if: Boolean!) on FIELD | FRAGMENT_SPREAD | INLINE_FRAGMENT\n\n\"The @deprecated directive is used within the type system definition language to indicate deprecated portions of a GraphQL service’s schema, such as deprecated fields on a type or deprecated enum values.\"\ndirective @deprecated(reason: String = \"No longer supported\") on FIELD_DEFINITION | ENUM_VALUE\n\ntype __Schema {\n    types: [__Type!]!\n    queryType: __Type!\n    mutationType: __Type\n    subscriptionType: __Type\n    directives: [__Directive!]!\n}\n\ntype __Type {\n    kind: __TypeKind!\n    name: String\n    description: String\n\n    # OBJECT and INTERFACE only\n    fields(includeDeprecated: Boolean = false): [__Field!]\n\n    # OBJECT only\n    interfaces: [__Type!]\n\n    # INTERFACE and UNION only\n    possibleTypes: [__Type!]\n\n    # ENUM only\n    enumValues(includeDeprecated: Boolean = false): [__EnumValue!]\n\n    # INPUT_OBJECT only\n    inputFields: [__InputValue!]\n\n    # NON_NULL and LIST only\n    ofType: __Type\n}\n\ntype __Field {\n    name: String!\n    description: String\n    args: [__InputValue!]!\n    type: __Type!\n    isDeprecated: Boolean!\n    deprecationReason: String\n}\n\ntype __InputValue {\n    name: String!\n    description: String\n    type: __Type!\n    defaultValue: String\n}\n\ntype __EnumValue {\n    name: String!\n    description: String\n    isDeprecated: Boolean!\n    deprecationReason: String\n}\n\nenum __TypeKind {\n    SCALAR\n    OBJECT\n    INTERFACE\n    UNION\n    ENUM\n    INPUT_OBJECT\n    LIST\n    NON_NULL\n}\n\ntype __Directive {\n    name: String!\n    description: String\n    locations: [__DirectiveLocation!]!\n    args: [__InputValue!]!\n}\n\nenum __DirectiveLocation {\n    QUERY\n    MUTATION\n    SUBSCRIPTION\n    FIELD\n    FRAGMENT_DEFINITION\n    FRAGMENT_SPREAD\n    INLINE_FRAGMENT\n    SCHEMA\n    SCALAR\n    OBJECT\n    FIELD_DEFINITION\n    ARGUMENT_DEFINITION\n    INTERFACE\n    UNION\n    ENUM\n    ENUM_VALUE\n    INPUT_OBJECT\n    INPUT_FIELD_DEFINITION\n}\n",
	BuiltIn: true,
}

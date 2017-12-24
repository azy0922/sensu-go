// Code generated by graphql/generator package. DO NOT EDIT.

package schema

import (
	graphql "github.com/graphql-go/graphql"
	util "github.com/sensu/sensu-go/backend/apid/graphql/generator/util"
)

//
// MutationResolver represents a collection of methods whose products represent the
// response values of the 'Mutation' type.
//
// == Example SDL
//
//   """
//   Dog's are not hooman.
//   """
//   type Dog implements Pet {
//     "name of this fine beast."
//     name:  String!
//
//     "breed of this silly animal; probably shibe."
//     breed: [Breed]
//   }
//
// == Example generated interface
//
//   // DogResolver ...
//   type DogResolver interface {
//     // Name implements response to request for name field.
//     Name(context.Context, interface{}, graphql.Params) interface{}
//     // Breed implements response to request for breed field.
//     Breed(context.Context, interface{}, graphql.Params) interface{}
//     // IsTypeOf is used to determine if a given value is associated with the Dog type
//     IsTypeOf(interface{}, graphql.IsTypeOfParams) bool
//   }
//
// == Example implementation ...
//
//   // MyDogResolver implements DogResolver interface
//   type MyDogResolver struct {
//     logger logrus.LogEntry
//     store interface{
//       store.BreedStore
//       store.DogStore
//     }
//   }
//
//   // Name implements response to request for name field.
//   func (r *MyDogResolver) Name(p graphql.Params) (interface{}, error) {
//     // ... implementation details ...
//     dog := p.Source.(DogGetter)
//     return dog.GetName()
//   }
//
//   // Breed implements response to request for breed field.
//   func (r *MyDogResolver) Name(p graphql.Params) (interface{}, error) {
//     // ... implementation details ...
//     dog := p.Source.(DogGetter)
//     breed := r.store.GetBreed(dog.GetBreedName())
//     return breed
//   }
//
//   // IsTypeOf is used to determine if a given value is associated with the Dog type
//   func (r *MyDogResolver) IsTypeOf(p graphql.IsTypeOfParams) bool {
//     // ... implementation details ...
//     _, ok := p.Value.(DogGetter)
//     return ok
//   }
type MutationResolver interface {
	// CreateCheck implements response to request for 'createCheck' field.
	CreateCheck(graphql.ResolveParams) (interface{}, error)
	// IsTypeOf is used to determine if a given value is associated with the Mutation type
	IsTypeOf(graphql.IsTypeOfParams) bool
}

// MutationAliases implements all methods on MutationResolver interface by using reflection to
// match name of field to a field on the given value. Intent is reduce friction
// of writing new resolvers by removing all the instances where you would simply
// have the resolvers method return a field.
//
// == Example SDL
//
//    type Dog {
//      name:   String!
//      weight: Float!
//      dob:    DateTime
//      breed:  [Breed]
//    }
//
// == Example generated aliases
//
//   type DogAliases struct {}
//   func (_ DogAliases) Name(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Weight(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Dob(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Breed(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//
// == Example Implementation
//
//   type DogResolver struct { // Implements DogResolver
//     DogAliases
//     store store.BreedStore
//   }
//
//   // NOTE:
//   // All other fields are satisified by DogAliases but since this one
//   // requires hitting the store we implement it in our resolver.
//   func (r *DogResolver) Breed(p graphql.ResolveParams) interface{} {
//     dog := p.Source.(*Dog)
//     return r.BreedsById(dog.BreedIDs)
//   }
//
type MutationAliases struct{}

// CreateCheck implements response to request for 'createCheck' field.
func (_ MutationAliases) CreateCheck(p graphql.ResolveParams) {
	return util.DefaultResolver(p.Source, p.Info.FieldName)
}

// Mutation The root query for implementing GraphQL mutations.
func Mutation() graphql.ObjectConfig {
	return graphql.ObjectConfig{
		Description: "The root query for implementing GraphQL mutations.",
		Fields: graphql.Fields{"createCheck": &graphql.Field{
			Args: graphql.FieldConfigArgument{"input": &graphql.ArgumentConfig{
				Description: "self descriptive",
				Type:        graphql.NewNonNull(util.InputType("CreateCheckInput")),
			}},
			DeprecationReason: "",
			Description:       "Creates a new check.",
			Name:              "createCheck",
			Type:              util.OutputType("CreateCheckPayload"),
		}},
		Interfaces: []*graphql.Interface{},
		IsTypeOf: func(_ graphql.IsTypeOfParams) bool {
			// NOTE:
			// Panic by default. Intent is that when Service is invoked, values of
			// these fields are updated with instantiated resolvers. If these
			// defaults are called it is most certainly programmer err.
			// If you're see this comment then: 'Whoops! Sorry, my bad.'
			panic("Unimplemented; see MutationResolver.")
		},
		Name: "Mutation",
	}
}

// CreateCheckInput self descriptive
func CreateCheckInput() graphql.InputObjectConfig {
	return graphql.InputObjectConfig{
		Description: "self descriptive",
		Fields: graphql.InputObjectConfigFieldMap{
			"clientMutationId": &graphql.InputObjectFieldConfig{
				Description: "A unique identifier for the client performing the mutation.",
				Type:        graphql.String,
			},
			"command": &graphql.InputObjectFieldConfig{
				Description: "self descriptive",
				Type:        graphql.NewNonNull(graphql.String),
			},
			"name": &graphql.InputObjectFieldConfig{
				Description: "self descriptive",
				Type:        graphql.NewNonNull(graphql.String),
			},
			"ns": &graphql.InputObjectFieldConfig{
				DefaultValue: map[string]interface{}{
					"environment":  "default",
					"organization": "default",
				},
				Description: "namespace the resulting resource will belong to",
				Type:        util.InputType("Namespace"),
			},
		},
		Name: "CreateCheckInput",
	}
}

//
// CreateCheckPayloadResolver represents a collection of methods whose products represent the
// response values of the 'CreateCheckPayload' type.
//
//  == Example SDL
//
//    """
//    Dog's are not hooman.
//    """
//    type Dog implements Pet {
//      "name of this fine beast."
//      name:  String!
//
//      "breed of this silly animal; probably shibe."
//      breed: [Breed]
//    }
//
//  == Example generated interface
//
//   // DogResolver ...
//   type DogResolver interface {
//     // Name implements response to request for name field.
//     Name(context.Context, interface{}, graphql.Params) interface{}
//     // Breed implements response to request for breed field.
//     Breed(context.Context, interface{}, graphql.Params) interface{}
//     // IsTypeOf is used to determine if a given value is associated with the Dog type
//     IsTypeOf(interface{}, graphql.IsTypeOfParams) bool
//   }
//
//  == Example implementation ...
//
//   // MyDogResolver implements DogResolver interface
//   type MyDogResolver struct {
//     logger logrus.LogEntry
//     store interface{
//       store.BreedStore
//       store.DogStore
//     }
//   }
//
//   // Name implements response to request for name field.
//   func (r *MyDogResolver) Name(p graphql.Params) (interface{}, error) {
//     // ... implementation details ...
//     dog := p.Source.(DogGetter)
//     return dog.GetName()
//   }
//
//   // Breed implements response to request for breed field.
//   func (r *MyDogResolver) Name(p graphql.Params) (interface{}, error) {
//     // ... implementation details ...
//     dog := p.Source.(DogGetter)
//     breed := r.store.GetBreed(dog.GetBreedName())
//     return breed
//   }
//
//   // IsTypeOf is used to determine if a given value is associated with the Dog type
//   func (r *MyDogResolver) IsTypeOf(p graphql.IsTypeOfParams) bool {
//     // ... implementation details ...
//     _, ok := p.Value.(DogGetter)
//     return ok
//   }
type CreateCheckPayloadResolver interface {
	// ClientMutationId implements response to request for 'clientMutationId' field.
	ClientMutationId(graphql.ResolveParams) (interface{}, error)
	// Check implements response to request for 'check' field.
	Check(graphql.ResolveParams) (interface{}, error)
	// IsTypeOf is used to determine if a given value is associated with the CreateCheckPayload type
	IsTypeOf(graphql.IsTypeOfParams) bool
}

// CreateCheckPayloadAliases implements all methods on CreateCheckPayloadResolver interface by using reflection to
// match name of field to a field on the given value. Intent is reduce friction
// of writing new resolvers by removing all the instances where you would simply
// have the resolvers method return a field.
//
// == Example SDL
//
//    type Dog {
//      name:   String!
//      weight: Float!
//      dob:    DateTime
//      breed:  [Breed]
//    }
//
// == Example generated aliases
//
//   type DogAliases struct {}
//   func (_ DogAliases) Name(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Weight(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Dob(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Breed(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//
// == Example Implementation
//
//   type DogResolver struct { // Implements DogResolver
//     DogAliases
//     store store.BreedStore
//   }
//
//   // NOTE:
//   // All other fields are satisified by DogAliases but since this one
//   // requires hitting the store we implement it in our resolver.
//   func (r *DogResolver) Breed(p graphql.ResolveParams) interface{} {
//     dog := v.(*Dog)
//     return r.BreedsById(dog.BreedIDs)
//   }
type CreateCheckPayloadAliases struct{}

// ClientMutationId implements response to request for 'clientMutationId' field.
func (_ CreateCheckPayloadAliases) ClientMutationId(p graphql.ResolveParams) {
	return util.DefaultResolver(p.Source, p.Info.FieldName)
}

// Check implements response to request for 'check' field.
func (_ CreateCheckPayloadAliases) Check(p graphql.ResolveParams) {
	return util.DefaultResolver(p.Source, p.Info.FieldName)
}

// CreateCheckPayload self descriptive
func CreateCheckPayload() graphql.ObjectConfig {
	return graphql.ObjectConfig{
		Description: "self descriptive",
		Fields: graphql.Fields{
			"check": &graphql.Field{
				Args:              graphql.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "The new check.",
				Name:              "check",
				Type:              graphql.NewNonNull(util.OutputType("Check")),
			},
			"clientMutationId": &graphql.Field{
				Args:              graphql.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "A unique identifier for the client performing the mutation.",
				Name:              "clientMutationId",
				Type:              graphql.String,
			},
		},
		Interfaces: []*graphql.Interface{},
		IsTypeOf: func(_ graphql.IsTypeOfParams) bool {
			// NOTE:
			// Panic by default. Intent is that when Service is invoked, values of
			// these fields are updated with instantiated resolvers. If these
			// defaults are called it is most certainly programmer err.
			// If you're see this comment then: 'Whoops! Sorry, my bad.'
			panic("Unimplemented; see CreateCheckPayloadResolver.")
		},
		Name: "CreateCheckPayload",
	}
}

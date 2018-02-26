import {
    GraphQLSchema,
    GraphQLObjectType,
    GraphQLString,
    GraphQLInt,
    GraphQLFloat,
  } from 'graphql';
import randomResolver from './random';

const MenuType = new GraphQLObjectType({
    name: 'Menu',
    fields: {
        name: {
            type: GraphQLString,
        },
        type: {
            type: GraphQLString,
        },
        price: {
            type: GraphQLInt,
        },
        calorie: {
            type: GraphQLInt,
        },
        protein: {
            type: GraphQLFloat,
        },
        lipid: {
            type: GraphQLFloat,
        },
        carbohydrate: {
            type: GraphQLFloat,
        },
        sodium: {
            type: GraphQLInt,
        },
        saltEquivalent: {
            type: GraphQLFloat,
        },
        description: {
            type: GraphQLString,
        },
        imageURL: {
            type: GraphQLString,
        },
    }
});

const schema = new GraphQLSchema({
    query: new GraphQLObjectType({
      name: 'RootQueryType',
      fields: {
        random: {
          type: MenuType,
          args: {
            type: {
                type: GraphQLString,
            },
            name: {
                type: GraphQLString,
            },
          },
          resolve: randomResolver,
        },
      },
    }),
  });

export default schema;

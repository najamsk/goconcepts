# Graphql with Go
Using 99design Gqlgen library to have a schema and code generation. Followed offical [getting started](https://gqlgen.com/getting-started/) guide.

- [x] Playground working with query against Todos.
- [x] Create a new Todo using mutation.
- [x] Get realtime subscription that sends timestamps using websockets per second.


### Subscription
Just follow simple [subscription torial](https://gqlgen.com/recipes/subscriptions/)

## Important graphql snippetes for playground


### Query all todos

    query findTodos {
    todos {
        id
        text
        done
        user {
        id
        name
        }
        
    }
    }


### Create a new TODO

    mutation createTodo {
    createTodo(input: { text: "haleem", userId: "1234293923423423" }) {
        user {
        id
        }
        text
        done
    }
    }


### Subscription to time events

    subscription realTime{
    currentTime {
        unixTime
        timeStamp
    }
    }
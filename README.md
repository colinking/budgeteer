# budgeteer
✨ Make Budget Management Magical 🎩

## Setup

You'll need to install the `protoc` protobuf compiler.

macOS: run `brew install protobuf`.

For Windows or Linux, see [the protoc installation page](http://google.github.io/proto-lens/installing-protoc.html).

### Configuring localhost for HTTPS

You need to configure your browser to trust a self-generated HTTPs certificate, because this gRPC library relies on HTTP/2.

macOS: run `make install-certs-mac`.

For Windows or Linux, run `make certs` to generate the required certificates, then install them in your default browser.

## Stack

Wrap backend and app in Dockerfiles. Deploy to Heroku?

### app

> TypeScript -> types

> React -> JS state

> gRPC -> comms

> Auth0 -> auth

> ?? -> state management (maybe https://mobx.js.org/)?

> React Router -> client-side routing

> Jest -> frontend testing

### backend

> gRPC -> endpoints

> Swagger? -> API docs

> goDoc -> function docs

> ? -> testing

> Heroku (free plan) -> deployment

> ? -> database (S3? DynamoDB?)

> DataDog -> metrics

> dep -> dependency management (go modules are WIP)

> chamber/AWS? Heroku env? -> secrets

## Todos

### Short-Term

[X] Print out transactions via gRPC

[ ] Evergreen in front-end

[ ] Load transactions into table

### Long-Term

[ ] Configure app and backend to run in Docker containers

[ ] Add react-router routing

[ ] Launch on Heroku

[ ] ACM under moss.colinking.co

[ ] Connect Auth0 login

[ ] Store user data in DynamoDB

[ ] Setup goDoc and Swagger

[ ] [Add Jest-based tests](https://github.com/facebook/create-react-app/blob/master/packages/react-scripts/template/README.md#writing-tests)

[ ] Add performance monitoring + [speed-up rendering time](https://github.com/stereobooster/react-snap)

[ ] Finlize an Evergreen Typing file, upload to GitHub [info](http://definitelytyped.org/guides/best-practices.html)


## References

- [gRPC web/golang example client (from Improbable)](https://github.com/improbable-eng/grpc-web/tree/master/example)
- [React gRPC example](https://github.com/easyCZ/grpc-web-hacker-news)
- [Use Postman for gRPC](https://github.com/jnewmano/grpc-json-proxy)
- [gRPC CLI](https://github.com/njpatel/grpcc)
- [openssl command](https://letsencrypt.org/docs/certificates-for-localhost/#making-and-trusting-your-own-certificates)
- [react-router-v4](https://codeburst.io/react-router-v4-unofficial-migration-guide-5a370b8905a)

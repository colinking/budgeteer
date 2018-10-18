# budgeteer
✨ Make Budget Management Magical

### Configuring localhost for HTTPS

You need to configure your browser to trust a self-generated HTTPs certificate, because this project uses gRPC over HTTP/2.

macOS: run `make install-certs-mac`.

For Windows or Linux, run `make certs` to generate the required certificates, then install them in your default browser.

## References

- [gRPC web/golang example client (from Improbable)](https://github.com/improbable-eng/grpc-web/tree/master/example)
- [React gRPC example](https://github.com/easyCZ/grpc-web-hacker-news)
- [Use Postman for gRPC](https://github.com/jnewmano/grpc-json-proxy)
- [gRPC CLI](https://github.com/njpatel/grpcc)
- [openssl command](https://letsencrypt.org/docs/certificates-for-localhost/#making-and-trusting-your-own-certificates)
- [react-router-v4](https://codeburst.io/react-router-v4-unofficial-migration-guide-5a370b8905a)

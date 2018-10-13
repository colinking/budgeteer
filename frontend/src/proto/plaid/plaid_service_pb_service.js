// package: plaid
// file: proto/plaid/plaid_service.proto

var proto_plaid_plaid_service_pb = require("../../proto/plaid/plaid_service_pb");
var grpc = require("grpc-web-client").grpc;

var Plaid = (function () {
  function Plaid() {}
  Plaid.serviceName = "plaid.Plaid";
  return Plaid;
}());

Plaid.ExchangeToken = {
  methodName: "ExchangeToken",
  service: Plaid,
  requestStream: false,
  responseStream: false,
  requestType: proto_plaid_plaid_service_pb.ExchangeTokenRequest,
  responseType: proto_plaid_plaid_service_pb.ExchangeTokenResponse
};

Plaid.GetTransactions = {
  methodName: "GetTransactions",
  service: Plaid,
  requestStream: false,
  responseStream: false,
  requestType: proto_plaid_plaid_service_pb.GetTransactionsRequest,
  responseType: proto_plaid_plaid_service_pb.GetTransactionsResponse
};

exports.Plaid = Plaid;

function PlaidClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

PlaidClient.prototype.exchangeToken = function exchangeToken(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  grpc.unary(Plaid.ExchangeToken, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          callback(Object.assign(new Error(response.statusMessage), { code: response.status, metadata: response.trailers }), null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
};

PlaidClient.prototype.getTransactions = function getTransactions(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  grpc.unary(Plaid.GetTransactions, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          callback(Object.assign(new Error(response.statusMessage), { code: response.status, metadata: response.trailers }), null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
};

exports.PlaidClient = PlaidClient;


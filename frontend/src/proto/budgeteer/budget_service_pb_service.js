// package: budgeteer
// file: proto/budgeteer/budget_service.proto

var proto_budgeteer_budget_service_pb = require("../../proto/budgeteer/budget_service_pb");
var grpc = require("grpc-web-client").grpc;

var BudgetService = (function () {
  function BudgetService() {}
  BudgetService.serviceName = "budgeteer.BudgetService";
  return BudgetService;
}());

BudgetService.GetPurchase = {
  methodName: "GetPurchase",
  service: BudgetService,
  requestStream: false,
  responseStream: false,
  requestType: proto_budgeteer_budget_service_pb.GetPurchasesRequest,
  responseType: proto_budgeteer_budget_service_pb.Purchase
};

exports.BudgetService = BudgetService;

function BudgetServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

BudgetServiceClient.prototype.getPurchase = function getPurchase(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  grpc.unary(BudgetService.GetPurchase, {
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

exports.BudgetServiceClient = BudgetServiceClient;


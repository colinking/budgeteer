// package: user
// file: proto/user/user_service.proto

var proto_user_user_service_pb = require("../../proto/user/user_service_pb");
var grpc = require("grpc-web-client").grpc;

var UserService = (function () {
  function UserService() {}
  UserService.serviceName = "user.UserService";
  return UserService;
}());

UserService.UserLogin = {
  methodName: "UserLogin",
  service: UserService,
  requestStream: false,
  responseStream: false,
  requestType: proto_user_user_service_pb.UserLoginRequest,
  responseType: proto_user_user_service_pb.UserLoginResponse
};

exports.UserService = UserService;

function UserServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

UserServiceClient.prototype.userLogin = function userLogin(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  grpc.unary(UserService.UserLogin, {
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

exports.UserServiceClient = UserServiceClient;


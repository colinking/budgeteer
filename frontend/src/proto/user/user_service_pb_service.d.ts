// package: user
// file: proto/user/user_service.proto

import * as proto_user_user_service_pb from "../../proto/user/user_service_pb";
import {grpc} from "grpc-web-client";

type UserServiceUserLogin = {
  readonly methodName: string;
  readonly service: typeof UserService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_user_user_service_pb.UserLoginRequest;
  readonly responseType: typeof proto_user_user_service_pb.UserLoginResponse;
};

export class UserService {
  static readonly serviceName: string;
  static readonly UserLogin: UserServiceUserLogin;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }
export type ServiceClientOptions = { transport: grpc.TransportConstructor; debug?: boolean }

interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: () => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}

export class UserServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: ServiceClientOptions);
  userLogin(
    requestMessage: proto_user_user_service_pb.UserLoginRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError, responseMessage: proto_user_user_service_pb.UserLoginResponse|null) => void
  ): void;
  userLogin(
    requestMessage: proto_user_user_service_pb.UserLoginRequest,
    callback: (error: ServiceError, responseMessage: proto_user_user_service_pb.UserLoginResponse|null) => void
  ): void;
}


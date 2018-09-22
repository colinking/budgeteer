// package: plaid
// file: proto/plaid/plaid_service.proto

import * as proto_plaid_plaid_service_pb from "../../proto/plaid/plaid_service_pb";
import {grpc} from "grpc-web-client";

type PlaidExchangeToken = {
  readonly methodName: string;
  readonly service: typeof Plaid;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_plaid_plaid_service_pb.ExchangeTokenRequest;
  readonly responseType: typeof proto_plaid_plaid_service_pb.ExchangeTokenResponse;
};

type PlaidGetTransactions = {
  readonly methodName: string;
  readonly service: typeof Plaid;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof proto_plaid_plaid_service_pb.GetTransactionsRequest;
  readonly responseType: typeof proto_plaid_plaid_service_pb.GetTransactionsResponse;
};

export class Plaid {
  static readonly serviceName: string;
  static readonly ExchangeToken: PlaidExchangeToken;
  static readonly GetTransactions: PlaidGetTransactions;
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

export class PlaidClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: ServiceClientOptions);
  exchangeToken(
    requestMessage: proto_plaid_plaid_service_pb.ExchangeTokenRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError, responseMessage: proto_plaid_plaid_service_pb.ExchangeTokenResponse|null) => void
  ): void;
  exchangeToken(
    requestMessage: proto_plaid_plaid_service_pb.ExchangeTokenRequest,
    callback: (error: ServiceError, responseMessage: proto_plaid_plaid_service_pb.ExchangeTokenResponse|null) => void
  ): void;
  getTransactions(
    requestMessage: proto_plaid_plaid_service_pb.GetTransactionsRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError, responseMessage: proto_plaid_plaid_service_pb.GetTransactionsResponse|null) => void
  ): void;
  getTransactions(
    requestMessage: proto_plaid_plaid_service_pb.GetTransactionsRequest,
    callback: (error: ServiceError, responseMessage: proto_plaid_plaid_service_pb.GetTransactionsResponse|null) => void
  ): void;
}


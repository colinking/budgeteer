// package: budgeteer
// file: budgeteer/budget_service.proto

import * as budgeteer_budget_service_pb from "../budgeteer/budget_service_pb";
import {grpc} from "grpc-web-client";

type BudgetServiceGetPurchase = {
  readonly methodName: string;
  readonly service: typeof BudgetService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof budgeteer_budget_service_pb.GetPurchasesRequest;
  readonly responseType: typeof budgeteer_budget_service_pb.Purchase;
};

export class BudgetService {
  static readonly serviceName: string;
  static readonly GetPurchase: BudgetServiceGetPurchase;
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

export class BudgetServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: ServiceClientOptions);
  getPurchase(
    requestMessage: budgeteer_budget_service_pb.GetPurchasesRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError, responseMessage: budgeteer_budget_service_pb.Purchase|null) => void
  ): void;
  getPurchase(
    requestMessage: budgeteer_budget_service_pb.GetPurchasesRequest,
    callback: (error: ServiceError, responseMessage: budgeteer_budget_service_pb.Purchase|null) => void
  ): void;
}


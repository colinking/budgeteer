// package: plaid
// file: plaid/plaid_service.proto

import * as jspb from "google-protobuf";

export class ExchangeTokenRequest extends jspb.Message {
  getToken(): string;
  setToken(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ExchangeTokenRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ExchangeTokenRequest): ExchangeTokenRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ExchangeTokenRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ExchangeTokenRequest;
  static deserializeBinaryFromReader(message: ExchangeTokenRequest, reader: jspb.BinaryReader): ExchangeTokenRequest;
}

export namespace ExchangeTokenRequest {
  export type AsObject = {
    token: string,
  }
}

export class ExchangeTokenResponse extends jspb.Message {
  getAccessToken(): string;
  setAccessToken(value: string): void;

  getItemId(): string;
  setItemId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ExchangeTokenResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ExchangeTokenResponse): ExchangeTokenResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ExchangeTokenResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ExchangeTokenResponse;
  static deserializeBinaryFromReader(message: ExchangeTokenResponse, reader: jspb.BinaryReader): ExchangeTokenResponse;
}

export namespace ExchangeTokenResponse {
  export type AsObject = {
    accessToken: string,
    itemId: string,
  }
}

export class Transaction extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getAccountId(): string;
  setAccountId(value: string): void;

  clearCategoryList(): void;
  getCategoryList(): Array<string>;
  setCategoryList(value: Array<string>): void;
  addCategory(value: string, index?: number): string;

  getCategoryId(): string;
  setCategoryId(value: string): void;

  getType(): Transaction.Type;
  setType(value: Transaction.Type): void;

  getMerchantName(): string;
  setMerchantName(value: string): void;

  getAmount(): number;
  setAmount(value: number): void;

  getCurrencyType(): Transaction.Currency;
  setCurrencyType(value: Transaction.Currency): void;

  getDate(): string;
  setDate(value: string): void;

  hasLocation(): boolean;
  clearLocation(): void;
  getLocation(): Transaction.Location | undefined;
  setLocation(value?: Transaction.Location): void;

  hasPaymentMeta(): boolean;
  clearPaymentMeta(): void;
  getPaymentMeta(): Transaction.PaymentMeta | undefined;
  setPaymentMeta(value?: Transaction.PaymentMeta): void;

  getPending(): boolean;
  setPending(value: boolean): void;

  getPendingTransactionId(): string;
  setPendingTransactionId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Transaction.AsObject;
  static toObject(includeInstance: boolean, msg: Transaction): Transaction.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Transaction, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Transaction;
  static deserializeBinaryFromReader(message: Transaction, reader: jspb.BinaryReader): Transaction;
}

export namespace Transaction {
  export type AsObject = {
    id: string,
    accountId: string,
    categoryList: Array<string>,
    categoryId: string,
    type: Transaction.Type,
    merchantName: string,
    amount: number,
    currencyType: Transaction.Currency,
    date: string,
    location?: Transaction.Location.AsObject,
    paymentMeta?: Transaction.PaymentMeta.AsObject,
    pending: boolean,
    pendingTransactionId: string,
  }

  export class Location extends jspb.Message {
    getAddress(): string;
    setAddress(value: string): void;

    getCity(): string;
    setCity(value: string): void;

    getLat(): number;
    setLat(value: number): void;

    getLon(): number;
    setLon(value: number): void;

    getState(): string;
    setState(value: string): void;

    getStoreNumber(): string;
    setStoreNumber(value: string): void;

    getZip(): string;
    setZip(value: string): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Location.AsObject;
    static toObject(includeInstance: boolean, msg: Location): Location.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Location, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Location;
    static deserializeBinaryFromReader(message: Location, reader: jspb.BinaryReader): Location;
  }

  export namespace Location {
    export type AsObject = {
      address: string,
      city: string,
      lat: number,
      lon: number,
      state: string,
      storeNumber: string,
      zip: string,
    }
  }

  export class PaymentMeta extends jspb.Message {
    getByOrderOf(): string;
    setByOrderOf(value: string): void;

    getPayee(): string;
    setPayee(value: string): void;

    getPayer(): string;
    setPayer(value: string): void;

    getPaymentMethod(): string;
    setPaymentMethod(value: string): void;

    getPaymentProcessor(): string;
    setPaymentProcessor(value: string): void;

    getPpdid(): string;
    setPpdid(value: string): void;

    getReason(): string;
    setReason(value: string): void;

    getReferenceNumber(): string;
    setReferenceNumber(value: string): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): PaymentMeta.AsObject;
    static toObject(includeInstance: boolean, msg: PaymentMeta): PaymentMeta.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: PaymentMeta, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): PaymentMeta;
    static deserializeBinaryFromReader(message: PaymentMeta, reader: jspb.BinaryReader): PaymentMeta;
  }

  export namespace PaymentMeta {
    export type AsObject = {
      byOrderOf: string,
      payee: string,
      payer: string,
      paymentMethod: string,
      paymentProcessor: string,
      ppdid: string,
      reason: string,
      referenceNumber: string,
    }
  }

  export enum Type {
    UNRESOLVED = 0,
    DIGITAL = 1,
    PLACE = 2,
    SPECIAL = 3,
  }

  export enum Currency {
    USD = 0,
    CAD = 1,
  }
}

export class GetTransactionsRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetTransactionsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetTransactionsRequest): GetTransactionsRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetTransactionsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetTransactionsRequest;
  static deserializeBinaryFromReader(message: GetTransactionsRequest, reader: jspb.BinaryReader): GetTransactionsRequest;
}

export namespace GetTransactionsRequest {
  export type AsObject = {
  }
}

export class GetTransactionsResponse extends jspb.Message {
  clearTransactionsList(): void;
  getTransactionsList(): Array<Transaction>;
  setTransactionsList(value: Array<Transaction>): void;
  addTransactions(value?: Transaction, index?: number): Transaction;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetTransactionsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetTransactionsResponse): GetTransactionsResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetTransactionsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetTransactionsResponse;
  static deserializeBinaryFromReader(message: GetTransactionsResponse, reader: jspb.BinaryReader): GetTransactionsResponse;
}

export namespace GetTransactionsResponse {
  export type AsObject = {
    transactionsList: Array<Transaction.AsObject>,
  }
}


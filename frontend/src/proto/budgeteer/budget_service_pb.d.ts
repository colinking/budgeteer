// package: budgeteer
// file: budgeteer/budget_service.proto

import * as jspb from "google-protobuf";

export class Purchase extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getAmount(): number;
  setAmount(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Purchase.AsObject;
  static toObject(includeInstance: boolean, msg: Purchase): Purchase.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Purchase, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Purchase;
  static deserializeBinaryFromReader(message: Purchase, reader: jspb.BinaryReader): Purchase;
}

export namespace Purchase {
  export type AsObject = {
    id: string,
    amount: number,
  }
}

export class GetPurchasesRequest extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetPurchasesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetPurchasesRequest): GetPurchasesRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetPurchasesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetPurchasesRequest;
  static deserializeBinaryFromReader(message: GetPurchasesRequest, reader: jspb.BinaryReader): GetPurchasesRequest;
}

export namespace GetPurchasesRequest {
  export type AsObject = {
    id: string,
  }
}


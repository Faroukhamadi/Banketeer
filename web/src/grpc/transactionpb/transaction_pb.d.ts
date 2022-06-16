import * as jspb from 'google-protobuf'



export class Transaction extends jspb.Message {
  getId(): number;
  setId(value: number): Transaction;

  getSenderaccountid(): number;
  setSenderaccountid(value: number): Transaction;

  getReceiveraccountid(): number;
  setReceiveraccountid(value: number): Transaction;

  getTellerid(): number;
  setTellerid(value: number): Transaction;

  getAccountid(): number;
  setAccountid(value: number): Transaction;

  getCreatedat(): string;
  setCreatedat(value: string): Transaction;

  getUpdatedat(): string;
  setUpdatedat(value: string): Transaction;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Transaction.AsObject;
  static toObject(includeInstance: boolean, msg: Transaction): Transaction.AsObject;
  static serializeBinaryToWriter(message: Transaction, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Transaction;
  static deserializeBinaryFromReader(message: Transaction, reader: jspb.BinaryReader): Transaction;
}

export namespace Transaction {
  export type AsObject = {
    id: number,
    senderaccountid: number,
    receiveraccountid: number,
    tellerid: number,
    accountid: number,
    createdat: string,
    updatedat: string,
  }
}

export class GetTransactionsRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetTransactionsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetTransactionsRequest): GetTransactionsRequest.AsObject;
  static serializeBinaryToWriter(message: GetTransactionsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetTransactionsRequest;
  static deserializeBinaryFromReader(message: GetTransactionsRequest, reader: jspb.BinaryReader): GetTransactionsRequest;
}

export namespace GetTransactionsRequest {
  export type AsObject = {
  }
}

export class GetTransactionsResponse extends jspb.Message {
  getTransactionsList(): Array<Transaction>;
  setTransactionsList(value: Array<Transaction>): GetTransactionsResponse;
  clearTransactionsList(): GetTransactionsResponse;
  addTransactions(value?: Transaction, index?: number): Transaction;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetTransactionsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetTransactionsResponse): GetTransactionsResponse.AsObject;
  static serializeBinaryToWriter(message: GetTransactionsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetTransactionsResponse;
  static deserializeBinaryFromReader(message: GetTransactionsResponse, reader: jspb.BinaryReader): GetTransactionsResponse;
}

export namespace GetTransactionsResponse {
  export type AsObject = {
    transactionsList: Array<Transaction.AsObject>,
  }
}

export class GetTransactionRequest extends jspb.Message {
  getId(): number;
  setId(value: number): GetTransactionRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetTransactionRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetTransactionRequest): GetTransactionRequest.AsObject;
  static serializeBinaryToWriter(message: GetTransactionRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetTransactionRequest;
  static deserializeBinaryFromReader(message: GetTransactionRequest, reader: jspb.BinaryReader): GetTransactionRequest;
}

export namespace GetTransactionRequest {
  export type AsObject = {
    id: number,
  }
}

export class GetTransactionResponse extends jspb.Message {
  getTransaction(): Transaction | undefined;
  setTransaction(value?: Transaction): GetTransactionResponse;
  hasTransaction(): boolean;
  clearTransaction(): GetTransactionResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetTransactionResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetTransactionResponse): GetTransactionResponse.AsObject;
  static serializeBinaryToWriter(message: GetTransactionResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetTransactionResponse;
  static deserializeBinaryFromReader(message: GetTransactionResponse, reader: jspb.BinaryReader): GetTransactionResponse;
}

export namespace GetTransactionResponse {
  export type AsObject = {
    transaction?: Transaction.AsObject,
  }
}


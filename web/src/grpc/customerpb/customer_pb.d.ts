import * as jspb from 'google-protobuf'



export class Customer extends jspb.Message {
  getId(): number;
  setId(value: number): Customer;

  getFirstname(): string;
  setFirstname(value: string): Customer;

  getLastname(): string;
  setLastname(value: string): Customer;

  getCin(): string;
  setCin(value: string): Customer;

  getPhone(): number;
  setPhone(value: number): Customer;

  getAccountid(): number;
  setAccountid(value: number): Customer;

  getCreatedat(): string;
  setCreatedat(value: string): Customer;

  getUpdatedat(): string;
  setUpdatedat(value: string): Customer;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Customer.AsObject;
  static toObject(includeInstance: boolean, msg: Customer): Customer.AsObject;
  static serializeBinaryToWriter(message: Customer, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Customer;
  static deserializeBinaryFromReader(message: Customer, reader: jspb.BinaryReader): Customer;
}

export namespace Customer {
  export type AsObject = {
    id: number,
    firstname: string,
    lastname: string,
    cin: string,
    phone: number,
    accountid: number,
    createdat: string,
    updatedat: string,
  }
}

export class Account extends jspb.Message {
  getId(): number;
  setId(value: number): Account;

  getBalance(): number;
  setBalance(value: number): Account;

  getAccountnumber(): string;
  setAccountnumber(value: string): Account;

  getCreatedat(): string;
  setCreatedat(value: string): Account;

  getUpdatedat(): string;
  setUpdatedat(value: string): Account;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Account.AsObject;
  static toObject(includeInstance: boolean, msg: Account): Account.AsObject;
  static serializeBinaryToWriter(message: Account, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Account;
  static deserializeBinaryFromReader(message: Account, reader: jspb.BinaryReader): Account;
}

export namespace Account {
  export type AsObject = {
    id: number,
    balance: number,
    accountnumber: string,
    createdat: string,
    updatedat: string,
  }
}

export class GetCustomersRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetCustomersRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetCustomersRequest): GetCustomersRequest.AsObject;
  static serializeBinaryToWriter(message: GetCustomersRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetCustomersRequest;
  static deserializeBinaryFromReader(message: GetCustomersRequest, reader: jspb.BinaryReader): GetCustomersRequest;
}

export namespace GetCustomersRequest {
  export type AsObject = {
  }
}

export class GetCustomersResponse extends jspb.Message {
  getCustomersList(): Array<Customer>;
  setCustomersList(value: Array<Customer>): GetCustomersResponse;
  clearCustomersList(): GetCustomersResponse;
  addCustomers(value?: Customer, index?: number): Customer;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetCustomersResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetCustomersResponse): GetCustomersResponse.AsObject;
  static serializeBinaryToWriter(message: GetCustomersResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetCustomersResponse;
  static deserializeBinaryFromReader(message: GetCustomersResponse, reader: jspb.BinaryReader): GetCustomersResponse;
}

export namespace GetCustomersResponse {
  export type AsObject = {
    customersList: Array<Customer.AsObject>,
  }
}

export class GetCustomerRequest extends jspb.Message {
  getId(): number;
  setId(value: number): GetCustomerRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetCustomerRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetCustomerRequest): GetCustomerRequest.AsObject;
  static serializeBinaryToWriter(message: GetCustomerRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetCustomerRequest;
  static deserializeBinaryFromReader(message: GetCustomerRequest, reader: jspb.BinaryReader): GetCustomerRequest;
}

export namespace GetCustomerRequest {
  export type AsObject = {
    id: number,
  }
}

export class GetCustomerResponse extends jspb.Message {
  getCustomer(): Customer | undefined;
  setCustomer(value?: Customer): GetCustomerResponse;
  hasCustomer(): boolean;
  clearCustomer(): GetCustomerResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetCustomerResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetCustomerResponse): GetCustomerResponse.AsObject;
  static serializeBinaryToWriter(message: GetCustomerResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetCustomerResponse;
  static deserializeBinaryFromReader(message: GetCustomerResponse, reader: jspb.BinaryReader): GetCustomerResponse;
}

export namespace GetCustomerResponse {
  export type AsObject = {
    customer?: Customer.AsObject,
  }
}


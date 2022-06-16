import * as jspb from 'google-protobuf'



export class Teller extends jspb.Message {
  getId(): number;
  setId(value: number): Teller;

  getUsername(): string;
  setUsername(value: string): Teller;

  getPassword(): string;
  setPassword(value: string): Teller;

  getRole(): string;
  setRole(value: string): Teller;

  getCreatedat(): string;
  setCreatedat(value: string): Teller;

  getUpdatedat(): string;
  setUpdatedat(value: string): Teller;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Teller.AsObject;
  static toObject(includeInstance: boolean, msg: Teller): Teller.AsObject;
  static serializeBinaryToWriter(message: Teller, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Teller;
  static deserializeBinaryFromReader(message: Teller, reader: jspb.BinaryReader): Teller;
}

export namespace Teller {
  export type AsObject = {
    id: number,
    username: string,
    password: string,
    role: string,
    createdat: string,
    updatedat: string,
  }
}

export class GetTellersRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetTellersRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetTellersRequest): GetTellersRequest.AsObject;
  static serializeBinaryToWriter(message: GetTellersRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetTellersRequest;
  static deserializeBinaryFromReader(message: GetTellersRequest, reader: jspb.BinaryReader): GetTellersRequest;
}

export namespace GetTellersRequest {
  export type AsObject = {
  }
}

export class GetTellersResponse extends jspb.Message {
  getTellersList(): Array<Teller>;
  setTellersList(value: Array<Teller>): GetTellersResponse;
  clearTellersList(): GetTellersResponse;
  addTellers(value?: Teller, index?: number): Teller;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetTellersResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetTellersResponse): GetTellersResponse.AsObject;
  static serializeBinaryToWriter(message: GetTellersResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetTellersResponse;
  static deserializeBinaryFromReader(message: GetTellersResponse, reader: jspb.BinaryReader): GetTellersResponse;
}

export namespace GetTellersResponse {
  export type AsObject = {
    tellersList: Array<Teller.AsObject>,
  }
}

export class GetTellerRequest extends jspb.Message {
  getId(): number;
  setId(value: number): GetTellerRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetTellerRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetTellerRequest): GetTellerRequest.AsObject;
  static serializeBinaryToWriter(message: GetTellerRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetTellerRequest;
  static deserializeBinaryFromReader(message: GetTellerRequest, reader: jspb.BinaryReader): GetTellerRequest;
}

export namespace GetTellerRequest {
  export type AsObject = {
    id: number,
  }
}

export class GetTellerResponse extends jspb.Message {
  getTeller(): Teller | undefined;
  setTeller(value?: Teller): GetTellerResponse;
  hasTeller(): boolean;
  clearTeller(): GetTellerResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetTellerResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetTellerResponse): GetTellerResponse.AsObject;
  static serializeBinaryToWriter(message: GetTellerResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetTellerResponse;
  static deserializeBinaryFromReader(message: GetTellerResponse, reader: jspb.BinaryReader): GetTellerResponse;
}

export namespace GetTellerResponse {
  export type AsObject = {
    teller?: Teller.AsObject,
  }
}


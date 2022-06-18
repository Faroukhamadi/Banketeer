/**
 * @fileoverview gRPC-Web generated client stub for main
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as teller_pb from './teller_pb';


export class TellerServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'binary';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodDescriptorGetTeller = new grpcWeb.MethodDescriptor(
    '/main.TellerService/GetTeller',
    grpcWeb.MethodType.UNARY,
    teller_pb.GetTellerRequest,
    teller_pb.GetTellerResponse,
    (request: teller_pb.GetTellerRequest) => {
      return request.serializeBinary();
    },
    teller_pb.GetTellerResponse.deserializeBinary
  );

  getTeller(
    request: teller_pb.GetTellerRequest,
    metadata: grpcWeb.Metadata | null): Promise<teller_pb.GetTellerResponse>;

  getTeller(
    request: teller_pb.GetTellerRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: teller_pb.GetTellerResponse) => void): grpcWeb.ClientReadableStream<teller_pb.GetTellerResponse>;

  getTeller(
    request: teller_pb.GetTellerRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: teller_pb.GetTellerResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/main.TellerService/GetTeller',
        request,
        metadata || {},
        this.methodDescriptorGetTeller,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/main.TellerService/GetTeller',
    request,
    metadata || {},
    this.methodDescriptorGetTeller);
  }

  methodDescriptorGetTellers = new grpcWeb.MethodDescriptor(
    '/main.TellerService/GetTellers',
    grpcWeb.MethodType.UNARY,
    teller_pb.GetTellersRequest,
    teller_pb.GetTellersResponse,
    (request: teller_pb.GetTellersRequest) => {
      return request.serializeBinary();
    },
    teller_pb.GetTellersResponse.deserializeBinary
  );

  getTellers(
    request: teller_pb.GetTellersRequest,
    metadata: grpcWeb.Metadata | null): Promise<teller_pb.GetTellersResponse>;

  getTellers(
    request: teller_pb.GetTellersRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: teller_pb.GetTellersResponse) => void): grpcWeb.ClientReadableStream<teller_pb.GetTellersResponse>;

  getTellers(
    request: teller_pb.GetTellersRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: teller_pb.GetTellersResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/main.TellerService/GetTellers',
        request,
        metadata || {},
        this.methodDescriptorGetTellers,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/main.TellerService/GetTellers',
    request,
    metadata || {},
    this.methodDescriptorGetTellers);
  }

}


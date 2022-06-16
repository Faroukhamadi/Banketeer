/**
 * @fileoverview gRPC-Web generated client stub for main
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as customer_pb from './customer_pb';


export class CustomerServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodDescriptorGetCustomer = new grpcWeb.MethodDescriptor(
    '/main.CustomerService/GetCustomer',
    grpcWeb.MethodType.UNARY,
    customer_pb.GetCustomerRequest,
    customer_pb.GetCustomerResponse,
    (request: customer_pb.GetCustomerRequest) => {
      return request.serializeBinary();
    },
    customer_pb.GetCustomerResponse.deserializeBinary
  );

  getCustomer(
    request: customer_pb.GetCustomerRequest,
    metadata: grpcWeb.Metadata | null): Promise<customer_pb.GetCustomerResponse>;

  getCustomer(
    request: customer_pb.GetCustomerRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: customer_pb.GetCustomerResponse) => void): grpcWeb.ClientReadableStream<customer_pb.GetCustomerResponse>;

  getCustomer(
    request: customer_pb.GetCustomerRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: customer_pb.GetCustomerResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/main.CustomerService/GetCustomer',
        request,
        metadata || {},
        this.methodDescriptorGetCustomer,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/main.CustomerService/GetCustomer',
    request,
    metadata || {},
    this.methodDescriptorGetCustomer);
  }

  methodDescriptorGetCustomers = new grpcWeb.MethodDescriptor(
    '/main.CustomerService/GetCustomers',
    grpcWeb.MethodType.UNARY,
    customer_pb.GetCustomersRequest,
    customer_pb.GetCustomersResponse,
    (request: customer_pb.GetCustomersRequest) => {
      return request.serializeBinary();
    },
    customer_pb.GetCustomersResponse.deserializeBinary
  );

  getCustomers(
    request: customer_pb.GetCustomersRequest,
    metadata: grpcWeb.Metadata | null): Promise<customer_pb.GetCustomersResponse>;

  getCustomers(
    request: customer_pb.GetCustomersRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: customer_pb.GetCustomersResponse) => void): grpcWeb.ClientReadableStream<customer_pb.GetCustomersResponse>;

  getCustomers(
    request: customer_pb.GetCustomersRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: customer_pb.GetCustomersResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/main.CustomerService/GetCustomers',
        request,
        metadata || {},
        this.methodDescriptorGetCustomers,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/main.CustomerService/GetCustomers',
    request,
    metadata || {},
    this.methodDescriptorGetCustomers);
  }

}


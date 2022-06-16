/**
 * @fileoverview gRPC-Web generated client stub for main
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as transaction_pb from './transaction_pb';


export class TransactionServiceClient {
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

  methodDescriptorGetTransaction = new grpcWeb.MethodDescriptor(
    '/main.TransactionService/GetTransaction',
    grpcWeb.MethodType.UNARY,
    transaction_pb.GetTransactionRequest,
    transaction_pb.GetTransactionResponse,
    (request: transaction_pb.GetTransactionRequest) => {
      return request.serializeBinary();
    },
    transaction_pb.GetTransactionResponse.deserializeBinary
  );

  getTransaction(
    request: transaction_pb.GetTransactionRequest,
    metadata: grpcWeb.Metadata | null): Promise<transaction_pb.GetTransactionResponse>;

  getTransaction(
    request: transaction_pb.GetTransactionRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: transaction_pb.GetTransactionResponse) => void): grpcWeb.ClientReadableStream<transaction_pb.GetTransactionResponse>;

  getTransaction(
    request: transaction_pb.GetTransactionRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: transaction_pb.GetTransactionResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/main.TransactionService/GetTransaction',
        request,
        metadata || {},
        this.methodDescriptorGetTransaction,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/main.TransactionService/GetTransaction',
    request,
    metadata || {},
    this.methodDescriptorGetTransaction);
  }

  methodDescriptorGetTransactions = new grpcWeb.MethodDescriptor(
    '/main.TransactionService/GetTransactions',
    grpcWeb.MethodType.UNARY,
    transaction_pb.GetTransactionsRequest,
    transaction_pb.GetTransactionsResponse,
    (request: transaction_pb.GetTransactionsRequest) => {
      return request.serializeBinary();
    },
    transaction_pb.GetTransactionsResponse.deserializeBinary
  );

  getTransactions(
    request: transaction_pb.GetTransactionsRequest,
    metadata: grpcWeb.Metadata | null): Promise<transaction_pb.GetTransactionsResponse>;

  getTransactions(
    request: transaction_pb.GetTransactionsRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: transaction_pb.GetTransactionsResponse) => void): grpcWeb.ClientReadableStream<transaction_pb.GetTransactionsResponse>;

  getTransactions(
    request: transaction_pb.GetTransactionsRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: transaction_pb.GetTransactionsResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/main.TransactionService/GetTransactions',
        request,
        metadata || {},
        this.methodDescriptorGetTransactions,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/main.TransactionService/GetTransactions',
    request,
    metadata || {},
    this.methodDescriptorGetTransactions);
  }

}


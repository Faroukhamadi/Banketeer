import type {RequestHandlerOutput, RequestEvent} from '@sveltejs/kit'
import 'carbon-components-svelte/css/white.css';
import type * as grpcWeb from 'grpc-web';
import { GetCustomersRequest, GetCustomersResponse } from '../grpc/customerpb/customer_pb';
import {CustomerServiceClient} from '../grpc/customerpb/CustomerServiceClientPb'
import '../grpc/customerpb/getcustomersrequest'

export async function get({}: RequestEvent): Promise<RequestHandlerOutput> {
	console.log('hey from server');

	const customerService = new CustomerServiceClient('http://localhost:8000', null, null);

	const request = new GetCustomersRequest();

	console.log('This is request: ', request);

	customerService.getCustomers(
		request,
		{ 'custom-header-1': 'value1' },
		(err: grpcWeb.RpcError, response: GetCustomersResponse) => {
      if (err) {
        console.log('error: ', err);
      } else {
        console.log('response: ', response.toObject().customersList);
      }
		}
	);

	console.log('hey from server');
	
  return {
    status: 200,
    body: {'data': 'hey'}
  }
}
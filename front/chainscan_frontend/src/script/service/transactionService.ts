import request from './request';
import {
  ResponseType,
  TransactionsResponse,
  TransactionLogResponse,
  InternalTransactionsResponse,
  GethDebugTraceResponse,
} from '../model/index';
import { TransactionDetail, TransactionOverview, DailyTransactionCount } from '../model/transaction';

export const GetTransactions = function (
  pageNumber: number,
  pageSize: number,
  type: string,
  blockNumber: number
): Promise<ResponseType<TransactionsResponse>> {
  const offset = pageNumber * pageSize;
  const limit = pageSize;
  let url = '';
  if (blockNumber !== -1) {
    url =
      '/v1/txs?blockBegin=' +
      blockNumber +
      '&blockEnd=' +
      blockNumber +
      '&offset=' +
      offset +
      '&limit=' +
      limit +
      '&allView=true';
  } else if (type === 'all') {
    url = '/v1/txs?offset=' + offset + '&limit=' + limit + '&allView=true';
  } else {
    url = '/v1/tokens/txns/' + type + '?offset=' + offset + '&limit=' + limit;
  }
  return request<TransactionsResponse>({
    url: url,
    method: 'get',
  });
};

export const GetHomeTransactions = function (): Promise<ResponseType<TransactionsResponse>> {
  return request<TransactionsResponse>({
    url: '/v1/txs?offset=0&limit=10',
    method: 'get',
  });
};

export const GetTransactionsByAddress = function (
  pageNumber: number,
  pageSize: number,
  type: string,
  address: string
): Promise<ResponseType<TransactionsResponse>> {
  const offset = pageNumber * pageSize;
  const limit = pageSize;
  let url = '';
  if (type === 'internal') {
    type = 'erc20';
  }
  if (type === 'txs') {
    url = '/v1/accounts/' + address + '/txns?offset=' + offset + '&limit=' + limit;
  } else {
    url = '/v1/accounts/' + address + '/txns-' + type + '?offset=' + offset + '&limit=' + limit;
  }
  return request<TransactionsResponse>({
    url: url,
    method: 'get',
  });
};

export const GetInternalTransactionsByAddress = function (
  pageNumber: number,
  pageSize: number,
  address: string
): Promise<ResponseType<InternalTransactionsResponse>> {
  const offset = pageNumber * pageSize;
  const limit = pageSize;
  const url = '/v1/accounts/' + address + '/txns-internal' + '?offset=' + offset + '&limit=' + limit;
  return request<InternalTransactionsResponse>({
    url: url,
    method: 'get',
  });
};

export const GetTransactionsByToken = function (
  pageNumber: number,
  pageSize: number,
  type: string,
  address: string
): Promise<ResponseType<TransactionsResponse>> {
  const offset = pageNumber * pageSize;
  const limit = pageSize;
  const url = '/v1/tokens/txns/' + type + '?contract=' + address + '&offset=' + offset + '&limit=' + limit;
  return request<TransactionsResponse>({
    url: url,
    method: 'get',
  });
};

export const GetTxByHash = function (txHash: string): Promise<ResponseType<TransactionDetail>> {
  return request<TransactionDetail>({
    url: '/v1/txs/' + txHash,
    method: 'get',
  });
};

export const GetBaseTxByHash = function (txHash: string): Promise<ResponseType<TransactionDetail>> {
  return request<TransactionDetail>({
    url: '/v1/txs/' + txHash + '/base',
    method: 'get',
  });
};

export const GetTxLog = function (txHash: string): Promise<ResponseType<TransactionLogResponse>> {
  return request<TransactionLogResponse>({
    url: '/v1/txs/' + txHash + '/event-logs',
    method: 'get',
  });
};

export const GetGethDebugTrace = function (
  txHash: string,
  type: string
): Promise<ResponseType<GethDebugTraceResponse>> {
  return request<GethDebugTraceResponse>({
    url: '/v1/txs/' + txHash + '/' + type,
    method: 'get',
  });
};

export const GetTxTotal = function (
  beginTime: string,
  endTime: string
): Promise<ResponseType<DailyTransactionCount[]>> {
  return request<DailyTransactionCount[]>({
    url: '/v1/daily/trend',
    method: 'get',
    params: {
      beginTime: beginTime,
      endTime: endTime,
    },
  });
};

export const GetTxOverview = function (): Promise<ResponseType<TransactionOverview>> {
  return request<TransactionOverview>({
    url: '/v1/overview',
    method: 'get',
  });
};

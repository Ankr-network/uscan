import request from './request';
import { TokenTransfers, Token } from '../model/token';
import { ResponseType, TransactionsResponse } from '../model/index';
import { TokenHoldersResponse, TokenResponse } from '../model/index';

export const GetTokenTransfersByAddress = function (address: string): Promise<ResponseType<TokenTransfers>> {
  return request<TokenTransfers>({
    url: '/v1/tokens/' + address + '/type',
    method: 'get',
  });
};

export const GetTokenHoldersByAddress = function (
  address: string,
  type: string,
  pageNumber: number,
  pageSize: number
): Promise<ResponseType<TokenHoldersResponse>> {
  const offset = pageNumber * pageSize;
  const limit = pageSize;
  return request<TokenHoldersResponse>({
    url: '/v1/tokens/' + address + '/holders?type=' + type + '&offset=' + offset + '&limit=' + limit,
    method: 'get',
  });
};

export const GetTransactionsByToken = function (
  address: string,
  type: string,
  pageNumber: number,
  pageSize: number
): Promise<ResponseType<TransactionsResponse>> {
  const offset = pageNumber * pageSize;
  const limit = pageSize;
  const url = '/v1/tokens/' + address + '/transfers' + '?type=' + type + '&offset=' + offset + '&limit=' + limit;
  return request<TransactionsResponse>({
    url: url,
    method: 'get',
  });
};

export const GetTokenInventory = function (
  address: string,
  type: string,
  pageNumber: number,
  pageSize: number
): Promise<ResponseType<TokenResponse>> {
  const offset = pageNumber * pageSize;
  const limit = pageSize;
  const url = '/v1/tokens/' + address + '/inventory' + '?type=' + type + '&offset=' + offset + '&limit=' + limit;
  // console.log(url);
  return request<TokenResponse>({
    url: url,
    method: 'get',
  });
};

export const GetNFTDetailByID = function (
  address: string,
  tokenID: string,
  type: string
): Promise<ResponseType<Token>> {
  return request<Token>({
    url: '/v1/nfts/' + address + '/' + tokenID + '?type=' + type,
    method: 'get',
  });
};

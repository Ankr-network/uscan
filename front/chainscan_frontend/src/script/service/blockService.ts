import request from './request';
import { ResponseType, BlocksResponse } from '../model/index';
import { BlockDetail } from '../model/block';

export const GetBlocks = function (
  allView: boolean,
  pageNumber: number,
  pageSize: number
): Promise<ResponseType<BlocksResponse>> {
  const offset = pageNumber * pageSize;
  const limit = pageSize;
  let url = '/v1/blocks?offset=' + offset + '&limit=' + limit;
  if (allView) {
    url = url + '&allView=true';
  }
  return request<BlocksResponse>({
    url: url,
    method: 'get',
  });
};

export const GetBlockByNumber = function (blockNumber: number): Promise<ResponseType<BlockDetail>> {
  return request<BlockDetail>({
    url: '/v1/blocks/' + blockNumber,
    method: 'get',
  });
};

import request from './request';
import { ResponseType } from '../model/index';
import { AddressDetail } from '../model/address';

export const GetAddressInfo = function (address: string): Promise<ResponseType<AddressDetail>> {
  return request<AddressDetail>({
    url: '/v1/accounts/' + address,
    method: 'get',
  });
};

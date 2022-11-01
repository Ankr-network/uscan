/**
 * AddressDetail
 * @class
 */
export class AddressDetail {
  id: number;
  owner: string;
  balance: string;
  blockNumber: string;
  creator: string;
  txHash: string;
  code: string;
  name: string;
  symbol: string;
  tokenTotalSupply: number;
  nftTotalSupply: number;
  decimals: string;
  createdTime: number;
  /**
   * Create a AddressDetail.
   * @param {number} id
   * @param {string} owner
   * @param {string} balance
   * @param {string} blockNumber
   * @param {string} creator
   * @param {string} txHash
   * @param {string} code
   * @param {string} name
   * @param {string} symbol
   * @param {number} tokenTotalSupply
   * @param {number} nftTotalSupply
   * @param {string} decimals
   * @param {string} createdTime
   */
  constructor(
    id: number,
    owner: string,
    balance: string,
    blockNumber: string,
    creator: string,
    txHash: string,
    code: string,
    name: string,
    symbol: string,
    tokenTotalSupply: number,
    nftTotalSupply: number,
    decimals: string,
    createdTime: number
  ) {
    this.id = id;
    this.owner = owner;
    this.balance = balance;
    this.blockNumber = blockNumber;
    this.creator = creator;
    this.txHash = txHash;
    this.code = code;
    this.name = name;
    this.symbol = symbol;
    this.tokenTotalSupply = tokenTotalSupply;
    this.nftTotalSupply = nftTotalSupply;
    this.decimals = decimals;
    this.createdTime = createdTime;
  }
}

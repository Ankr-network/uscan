import { Overview } from './index';
import { formatNumber } from '../utils';
import { TableHeader } from './index';

/**
 * TokensTransferred
 * @class
 */
export class TokensTransferred {
  from: string;
  fromHex: string;
  to: string;
  toHex: string;
  address: string;
  addressName: string;
  addressSymbol: string;
  /**
   * Create a TokensTransferred.
   * @param {string} from
   * @param {string} fromHex
   * @param {string} to
   * @param {string} toHex
   * @param {string} address
   * @param {string} addressName
   * @param {string} addressSymbol
   */
  constructor(
    from: string,
    fromHex: string,
    to: string,
    toHex: string,
    address: string,
    addressName: string,
    addressSymbol: string
  ) {
    this.from = from;
    this.fromHex = fromHex;
    this.to = to;
    this.toHex = toHex;
    this.address = address;
    this.addressName = addressName;
    this.addressSymbol = addressSymbol;
  }
}

/**
 * Transaction Logs
 * @class
 */
export class TransactionLog {
  id: string;
  address: string;
  topics: string[];
  data: string;
  blockNumber: string;
  transactionHash: string;
  transactionIndex: number;
  blockHash: string;
  logIndex: number;
  removed: boolean;
  createdTime: number;
  /**
   * Create a TransactionLog.
   * @param {string} id
   * @param {string} address
   * @param {string[]} topics
   * @param {string} data
   * @param {string} blockNumber
   * @param {string} transactionHash
   * @param {number} transactionIndex
   * @param {string} blockHash
   * @param {number} logIndex
   * @param {boolean} removed
   * @param {number} createdTime
   */
  constructor(
    id: string,
    address: string,
    topics: string[],
    data: string,
    blockNumber: string,
    transactionHash: string,
    transactionIndex: number,
    blockHash: string,
    logIndex: number,
    removed: boolean,
    createdTime: number
  ) {
    this.id = id;
    this.address = address;
    this.topics = topics;
    this.data = data;
    this.blockNumber = blockNumber;
    this.transactionHash = transactionHash;
    this.transactionIndex = transactionIndex;
    this.blockHash = blockHash;
    this.logIndex = logIndex;
    this.removed = removed;
    this.createdTime = createdTime;
  }
}

/**
 * Transaction
 * @class
 */
export class TransactionDetail {
  hash: string;
  method: string;
  blockHash: string;
  blockNumber: string;
  from: string;
  fromName: string;
  fromSymbol: string;
  fromCode: string;
  to: string;
  toName: string;
  toSymbol: string;
  toCode: string;
  gas: string;
  gasPrice: number;
  value: string;
  createTime: number;
  createdTime: number;
  maxFeePerGas: string;
  maxPriorityFeePerGas: string;
  input: string;
  nonce: number;
  transactionIndex: string;
  type: string;
  chainID: string;
  v: string;
  r: string;
  s: string;
  totalLogs: number;
  tokensTransferred: TokensTransferred[];
  baseFeePerGas: string;
  gasLimit: number;
  contractAddress: string;
  contractAddressName: string;
  contractAddressSymbol: string;
  cumulativeGasUsed: number;
  effectiveGasPrice: string;
  gasUsed: number;
  root: string;
  errorReturn: string;
  status: number;
  tokenID: number;
  transactionHash: string;
  baseInfo: boolean = false;
  methodName: string;
  /**
   * Create a Transaction.
   * @param {string} hash
   * @param {string} method
   * @param {string} blockHash
   * @param {string} blockNumber
   * @param {string} from
   * @param {string} fromName
   * @param {string} fromSymbol
   * @param {string} fromCode
   * @param {string} to
   * @param {string} toName
   * @param {string} toSymbol
   * @param {string} toCode
   * @param {string} gas
   * @param {number} gasPrice
   * @param {string} value
   * @param {number} createTime
   * @param {number} createdTime
   * @param {string} maxFeePerGas
   * @param {string} maxPriorityFeePerGas
   * @param {string} input
   * @param {number} nonce
   * @param {string} transactionIndex
   * @param {string} type
   * @param {string} chainID
   * @param {string} v
   * @param {string} r
   * @param {string} s
   * @param {number} totalLogs
   * @param {TokensTransferred[]} tokensTransferred
   * @param {string} baseFeePerGas
   * @param {number} gasLimit
   * @param {string} contractAddress
   * @param {string} contractAddressName
   * @param {string} contractAddressSymbol
   * @param {number} cumulativeGasUsed
   * @param {string} effectiveGasPrice
   * @param {number} gasUsed
   * @param {string} root
   * @param {string} errorReturn
   * @param {number} status
   * @param {number} tokenID
   * @param {string} transactionHash
   * @param {boolean} baseInfo
   * @param {string} methodName
   */
  constructor(
    hash: string,
    method: string,
    blockHash: string,
    blockNumber: string,
    from: string,
    fromName: string,
    fromSymbol: string,
    fromCode: string,
    to: string,
    toName: string,
    toSymbol: string,
    toCode: string,
    gas: string,
    gasPrice: number,
    value: string,
    createTime: number,
    createdTime: number,
    maxFeePerGas: string,
    maxPriorityFeePerGas: string,
    input: string,
    nonce: number,
    transactionIndex: string,
    type: string,
    chainID: string,
    v: string,
    r: string,
    s: string,
    totalLogs: number,
    tokensTransferred: TokensTransferred[],
    baseFeePerGas: string,
    gasLimit: number,
    contractAddress: string,
    contractAddressName: string,
    contractAddressSymbol: string,
    cumulativeGasUsed: number,
    effectiveGasPrice: string,
    gasUsed: number,
    root: string,
    errorReturn: string,
    status: number,
    tokenID: number,
    transactionHash: string,
    baseInfo: boolean,
    methodName: string
  ) {
    this.hash = hash;
    this.method = method;
    this.blockHash = blockHash;
    this.blockNumber = blockNumber;
    this.from = from;
    this.fromName = fromName;
    this.fromSymbol = fromSymbol;
    this.fromCode = fromCode;
    this.to = to;
    this.toName = toName;
    this.toSymbol = toSymbol;
    this.toCode = toCode;
    this.gas = gas;
    this.gasPrice = gasPrice;
    this.value = value;
    this.createTime = createTime;
    this.createdTime = createTime;
    this.maxFeePerGas = maxFeePerGas;
    this.maxPriorityFeePerGas = maxPriorityFeePerGas;
    this.input = input;
    this.nonce = nonce;
    this.transactionIndex = transactionIndex;
    this.type = type;
    this.chainID = chainID;
    this.v = v;
    this.r = r;
    this.s = s;
    this.totalLogs = totalLogs;
    this.tokensTransferred = tokensTransferred;
    this.baseFeePerGas = baseFeePerGas;
    this.gasLimit = gasLimit;
    this.contractAddress = contractAddress;
    this.contractAddressName = contractAddressName;
    this.contractAddressSymbol = contractAddressSymbol;
    this.cumulativeGasUsed = cumulativeGasUsed;
    this.effectiveGasPrice = effectiveGasPrice;
    this.gasUsed = gasUsed;
    this.root = root;
    this.errorReturn = errorReturn;
    this.status = status;
    this.tokenID = tokenID;
    this.transactionHash = transactionHash;
    this.baseInfo = baseInfo;
    this.methodName = methodName;
  }
}

export interface InternalTransactionDetail {
  transactionHash: string;
  blockNumber: string;
  status: number;
  callType: string;
  depth: string;
  from: string;
  to: string;
  amount: string;
  gasLimit: string;
  createdTime: number;
}

export interface GethDebugTrace {
  pc: string;
  op: string;
  gas: string;
  gasCost: number;
  depth: number;
}

export interface TransactionCount {
  TxCount: number;
  Erc20Tx: number;
  Erc721Tx: number;
  AvgDifficulty: number;
  Difficult: number;
  BlockCount: number;
  Date: string;
}

export interface DailyTransactionCount {
  date: string;
  txCount: number;
}

export interface TransactionOverview {
  address: number;
  avgBlockTime: number;
  block: number;
  blockHeight: number;
  dailyTx: number;
  diff: number;
  tps: number;
  tx: number;
  erc20: number;
  erc721: number;
}

export const getTxOverviews = function (tx: TransactionDetail): Overview[] {
  const txParameterMap = new Map();
  txParameterMap.set('hash', [
    'Transaction Hash',
    // eslint-disable-next-line max-len
    'A TxHash or transaction hash is a unique 66-character identifier that is generated whenever a transaction is executed.',
  ]);
  txParameterMap.set('status', ['Status', 'The status of the transaction.']);
  txParameterMap.set('blockNumber', [
    'Block',
    // eslint-disable-next-line max-len
    'Number of the block in which the transaction is recorded. Block confirmations indicate how many blocks have been added since the transaction was mined.',
  ]);
  txParameterMap.set('createTime', ['Timestamp', 'The date and time at which a transaction is mined.']);
  txParameterMap.set('createdTime', ['Timestamp', 'The date and time at which a transaction is mined.']);
  txParameterMap.set('from', ['From', 'The sending party of the transaction.']);
  txParameterMap.set('to', ['To', 'The receiving party of the transaction (could be a contract address).']);
  txParameterMap.set('value', [
    'Value',
    // eslint-disable-next-line max-len
    'The value being transacted in Ether and fiat value. Note: You can click the fiat value (if available) to see historical value at the time of transaction.',
  ]);
  txParameterMap.set('gas', ['Transaction Fee', 'Amount paid to the miner for processing the transaction.']);
  txParameterMap.set('gasPrice', [
    'Gas Price',
    // eslint-disable-next-line max-len
    'Cost per unit of gas specified for the transaction, in Ether and Gwei. The higher the gas price the higher chance of getting included in a block.',
  ]);
  txParameterMap.set('gas', [
    'Gas Limit & Usage by Txn',
    // eslint-disable-next-line max-len
    'Maximum amount of gas allocated for the transaction & the amount eventually used. Normal ETH transfers involve " + res.gasLimit + " gas units while contracts involve higher values.',
  ]);
  txParameterMap.set('maxPriorityFeePerGas', ['Gas Fees', 'The amount eventually used.']);
  txParameterMap.set('tokensTransferred', ['Tokens Transferred', 'List of tokens transferred in the transaction.']);
  txParameterMap.set('input', [
    'Input Data',
    // eslint-disable-next-line max-len
    'Additional data included for this transaction. Commonly used as part of contract interaction or as a message sent to the recipient.',
  ]);
  const resList: Overview[] = [];
  for (const [key, value] of txParameterMap) {
    let valueDisplay: any = tx[key as keyof TransactionDetail] as string;
    if (valueDisplay === undefined || valueDisplay === null || valueDisplay.length === 0) {
      if (key !== 'to') {
        continue;
      }
    }
    // console.log('key', key, 'value', valueDisplay);
    if (key == 'from') {
      valueDisplay = {
        from: tx.from,
        fromCode: tx.fromCode,
        fromName: tx.fromName,
        fromSymbol: tx.fromSymbol,
      };
    } else if (key == 'to') {
      valueDisplay = {
        to: tx.to,
        toCode: tx.toCode,
        toName: tx.toName,
        toSymbol: tx.toSymbol,
        contractAddress: tx.contractAddress,
        contractAddressName: tx.contractAddressName,
        contractAddressSymbol: tx.contractAddressSymbol,
      };
    } else if (key == 'gas') {
      valueDisplay = {
        gasLimit: formatNumber(BigInt(tx.gasLimit)),
        gasUsed: formatNumber(BigInt(tx.gasUsed)),
        percent: Math.round((tx.gasUsed / tx.gasLimit) * 10000) / 100 + '%',
      };
    } else if (key == 'maxPriorityFeePerGas') {
      valueDisplay = {
        baseFeePerGas: tx.baseFeePerGas,
        maxFeePerGas: tx.maxFeePerGas,
        maxPriorityFeePerGas: tx.maxPriorityFeePerGas,
      };
    } else if (key == 'status') {
      valueDisplay = {
        status: tx.status,
        errorMsg: tx.errorReturn,
      };
    } else if (key == 'input') {
      valueDisplay = {
        inputContent: tx.input,
        methodName: tx.methodName,
      };
    }
    resList.push(new Overview(key, value[0] + ':', valueDisplay, value[1]));
  }
  // console.log('tx', tx);
  // console.log('resList', resList);
  return resList;
};

export const TransactionsHeaderList: TableHeader[] = [
  new TableHeader('Txn Hash', 'hash'),
  new TableHeader('Method', 'method'),
  new TableHeader('Block', 'blockNumber'),
  new TableHeader('Age', 'createTime'),
  new TableHeader('From', 'from'),
  new TableHeader('To', 'to'),
  new TableHeader('Value', 'value'),
  new TableHeader('Txn Fee', 'gas'),
];

export const Erc20TransactionsHeaderList: TableHeader[] = [
  new TableHeader('Txn Hash', 'transactionHash'),
  new TableHeader('Method', 'method'),
  new TableHeader('Block', 'blockNumber'),
  new TableHeader('Age', 'createdTime'),
  new TableHeader('From', 'from'),
  new TableHeader('To', 'to'),
  new TableHeader('Value(token)', 'value'),
  new TableHeader('Token', 'contract'),
];

export const Erc721TransactionsHeaderList: TableHeader[] = [
  new TableHeader('Txn Hash', 'transactionHash'),
  new TableHeader('Method', 'method'),
  new TableHeader('Block', 'blockNumber'),
  new TableHeader('Age', 'createdTime'),
  new TableHeader('From', 'from'),
  new TableHeader('To', 'to'),
  new TableHeader('TokenID', 'tokenID'),
];

export const InternalTransactionsHeaderList: TableHeader[] = [
  new TableHeader('Parent Txn Hash', 'transactionHash'),
  new TableHeader('Block', 'blockNumber'),
  new TableHeader('Age', 'createdTime'),
  new TableHeader('From', 'from'),
  new TableHeader('To', 'to'),
  new TableHeader('Value', 'amount'),
];

export const TokenErc20TransactionsHeaderList: TableHeader[] = [
  new TableHeader('Txn Hash', 'transactionHash'),
  new TableHeader('Method', 'method'),
  new TableHeader('Block', 'blockNumber'),
  new TableHeader('Age', 'createdTime'),
  new TableHeader('From', 'from'),
  new TableHeader('To', 'to'),
  new TableHeader('Quantity', 'value'),
];

export const TokenErcTransactionsHeaderList: TableHeader[] = [
  new TableHeader('Txn Hash', 'transactionHash'),
  new TableHeader('Method', 'method'),
  new TableHeader('Block', 'blockNumber'),
  new TableHeader('Age', 'createdTime'),
  new TableHeader('From', 'from'),
  new TableHeader('To', 'to'),
  new TableHeader('TokenID', 'tokenID'),
];

export const TokeHolderHeaderList: TableHeader[] = [
  new TableHeader('Address', 'owner'),
  new TableHeader('Quantity', 'quantity'),
];

export const TokeErcHolderHeaderList: TableHeader[] = [
  new TableHeader('Address', 'owner'),
  new TableHeader('token', 'tokenID'),
];

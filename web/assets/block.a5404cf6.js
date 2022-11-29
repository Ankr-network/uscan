import{k as l,b0 as n}from"./index.618b475f.js";import{T as a,O as r}from"./index.7d4bc839.js";const m=function(i,e,o){const s=e*o,c=o;let t="/blocks?offset="+s+"&limit="+c;return i&&(t=t+"&allView=true"),l({url:t,method:"get"})},u=function(i){return l({url:"/blocks/"+i,method:"get"})},b=function(i){const e=new Map;e.set("number",["Block Height","Also known as Block Number. The block height, which indicates the length of the blockchain, increases after the addition of the new block."]),e.set("timestamp",["Timestamp","The date and time at which a block is mined."]),e.set("transactionsTotal",["Transactions","The number of transactions in the block. Internal transaction is transactions as a result of contract execution that involves Ether value."]),e.set("miner",["Mined by","Miner who successfully include the block onto the blockchain."]),e.set("difficulty",["Difficulty","The amount of effort required to mine a new block. The difficulty algorithm may adjust according to time."]),e.set("totalDifficulty",["Total Difficulty","Total difficulty of the chain until this block."]),e.set("size",["Size","The block size is actually determined by the block gas limit."]),e.set("gasUsed",["Gas Used","The total gas used in the block and its percentage of gas filled in the block."]),e.set("gasLimit",["Gas Limit","Total gas limit provided by all transactions in the block."]),e.set("extraData",["Extra Data","Any data that can be included by the miner in the block."]),e.set("hash",["Hash","The hash of the block header of the current block."]),e.set("parentHash",["Parent Hash","The hash of the block from which this block was generated, also known as its parent block."]),e.set("sha3Uncles",["Sha3Uncles","The mechanism which Ethereum Javascript RLP encodes an empty string."]),e.set("stateRoot",["StateRoot","The root of the state trie."]),e.set("nonce",["Nonce","Block nonce is a value used during mining to demonstrate proof of work for a block."]);const o=[];for(const[s,c]of e){let t=i[s];s=="number"?t=parseInt(t).toString():s=="difficulty"?t=n(BigInt(parseInt(t))):s=="totalDifficulty"?t=n(BigInt(parseInt(t))):s=="size"?t=n(parseInt(t))+" bytes":s=="gasUsed"?t=n(parseInt(t)):s=="gasLimit"&&(t=n(parseInt(t))),o.push(new r(s,c[0]+":",t,c[1]))}return o},d=[new a("Block","number"),new a("Age","timestamp"),new a("Txn","transactionsTotal"),new a("Miner","miner"),new a("Gas Used","gasUsed"),new a("Gas Limit","gasLimit"),new a("Base Fee","baseFeePerGas")];export{d as B,u as G,m as a,b as g};

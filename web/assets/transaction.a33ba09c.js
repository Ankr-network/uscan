import{O as c,T as e}from"./index.7d4bc839.js";import{f as i}from"./utils.409af791.js";const h=function(a){const t=new Map;t.set("hash",["Transaction Hash","A TxHash or transaction hash is a unique 66-character identifier that is generated whenever a transaction is executed."]),t.set("status",["Status","The status of the transaction."]),t.set("blockNumber",["Block","Number of the block in which the transaction is recorded. Block confirmations indicate how many blocks have been added since the transaction was mined."]),t.set("createTime",["Timestamp","The date and time at which a transaction is mined."]),t.set("createdTime",["Timestamp","The date and time at which a transaction is mined."]),t.set("from",["From","The sending party of the transaction."]),t.set("to",["To","The receiving party of the transaction (could be a contract address)."]),t.set("value",["Value","The value being transacted in Ether and fiat value. Note: You can click the fiat value (if available) to see historical value at the time of transaction."]),t.set("gas",["Transaction Fee","Amount paid to the miner for processing the transaction."]),t.set("gasPrice",["Gas Price","Cost per unit of gas specified for the transaction, in Ether and Gwei. The higher the gas price the higher chance of getting included in a block."]),t.set("gas",["Gas Limit & Usage by Txn",'Maximum amount of gas allocated for the transaction & the amount eventually used. Normal ETH transfers involve " + res.gasLimit + " gas units while contracts involve higher values.']),parseInt(a.baseFeePerGas)!==0&&t.set("maxPriorityFeePerGas",["Gas Fees","The amount eventually used."]),t.set("tokensTransferred",["Tokens Transferred","List of tokens transferred in the transaction."]),t.set("input",["Input Data","Additional data included for this transaction. Commonly used as part of contract interaction or as a message sent to the recipient."]);const s=[];for(const[o,r]of t){let n=a[o];(n==null||n.length===0)&&o!=="to"||(o=="from"?n={from:a.from,fromCode:a.fromCode,fromName:a.fromName,fromSymbol:a.fromSymbol,fromContract:a.fromContract}:o=="to"?n={to:a.to,toCode:a.toCode,toName:a.toName,toSymbol:a.toSymbol,contractAddress:a.contractAddress,contractAddressName:a.contractAddressName,contractAddressSymbol:a.contractAddressSymbol,toContract:a.toContract}:o=="gas"?n={gas:i(BigInt(a.gas)),gasUsed:i(BigInt(a.gasUsed))}:o=="maxPriorityFeePerGas"?n={baseFeePerGas:a.baseFeePerGas,maxFeePerGas:a.maxFeePerGas,maxPriorityFeePerGas:a.maxPriorityFeePerGas}:o=="status"?n={status:a.status,errorMsg:a.errorReturn}:o=="input"&&(n={inputContent:a.input,methodName:a.methodName}),s.push(new c(o,r[0]+":",n,r[1])))}return s},l=[new e("Txn Hash","hash"),new e("Method","method"),new e("Block","blockNumber"),new e("Age","createTime"),new e("From","from"),new e("To","to"),new e("Value","value"),new e("Txn Fee","gas")],w=[new e("Txn Hash","transactionHash"),new e("Method","method"),new e("Block","blockNumber"),new e("Age","createdTime"),new e("From","from"),new e("To","to"),new e("Value","value"),new e("Token","contract")],u=[new e("Txn Hash","transactionHash"),new e("Method","method"),new e("Block","blockNumber"),new e("Age","createdTime"),new e("From","from"),new e("To","to"),new e("TokenID","tokenID"),new e("Token","contract")],T=[new e("Txn Hash","transactionHash"),new e("Method","method"),new e("Block","blockNumber"),new e("Age","createdTime"),new e("From","from"),new e("To","to"),new e("TokenID","tokenID"),new e("Token","contract"),new e("Quantity","value")],f=[new e("Parent Txn Hash","transactionHash"),new e("Block","blockNumber"),new e("Age","createdTime"),new e("From","from"),new e("To","to"),new e("Value","amount")],g=[new e("Txn Hash","transactionHash"),new e("Method","method"),new e("Block","blockNumber"),new e("Age","createdTime"),new e("From","from"),new e("To","to"),new e("Quantity","value")],k=[new e("Txn Hash","transactionHash"),new e("Method","method"),new e("Block","blockNumber"),new e("Age","createdTime"),new e("From","from"),new e("To","to"),new e("TokenID","tokenID")],b=[new e("Address","Address"),new e("Quantity","Quantity")];new e("Address","owner"),new e("token","tokenID");export{w as E,f as I,l as T,u as a,T as b,g as c,b as d,k as e,h as g};
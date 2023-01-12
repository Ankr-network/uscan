import{_ as S,d as D,bC as L,m as N,F as d,aM as _,j as M,J as B,i as R,o as x,c as F,e as f,t as g,n as h,w as H,W as P,a as V,b as q}from"./index.b0227044.js";import{E as I}from"./el-pagination.9514bc10.js";/* empty css               */import"./el-select.62bb2281.js";import{_ as G}from"./GenerateTransfers.45719357.js";import{_ as J}from"./GenerateTransactions.e879c032.js";import{a as v}from"./transactionService.c64f44ea.js";import{T as W,E as A,a as K,b as O}from"./transaction.5f4a60f1.js";import"./isEqual.de233551.js";import"./index.ae2625c8.js";import"./index.97fd9ae4.js";import"./el-table-column.1b3f6160.js";import"./objects.6215b246.js";import"./BaseTransactionInfo.163eba02.js";import"./index.eb369523.js";import"./_commonjsHelpers.b8add541.js";import"./utils.ec062d7c.js";import"./index.7d4bc839.js";const Q={class:"center-row"},U={class:"h4-title"},X={class:"tx-sub-title"},Y={style:{"margin-top":"1%",display:"flex","justify-content":"center"}},Z=D({__name:"TransactionsList",props:{txsType:String},setup(T){const t=T,y=L();document.title="Transactions | The "+N()+" Explorer";const n=d(1),i=d(25),e=_([]),r=_([]),l=d(0),c=y.query.block===void 0?-1:y.query.block,s=d("title"),p=d(""),m=d(!0);M(()=>{w()});const b=async()=>{if(t.txsType==="all"||t.txsType==="erc20"||t.txsType==="erc721"||t.txsType==="erc1155"){t.txsType==="all"?r.push(...W):t.txsType==="erc20"?r.push(...A):t.txsType==="erc721"?r.push(...K):t.txsType==="erc1155"&&r.push(...O);const a=await v(n.value-1,i.value,t.txsType,c);a.data.items.forEach(o=>{e.push(o)}),l.value=a.data.total,a.data.total==0&&(m.value=!1)}};b(),B(t,async()=>{w(),e.length=0,r.length=0,n.value=1,i.value=25,b()});const w=()=>{t.txsType==="all"?c!==-1?s.value="Transactions For Block "+c:s.value="Transactions":t.txsType==="erc20"?(s.value="Token Transfers",p.value="ERC-20"):t.txsType==="erc721"?(s.value="Non-Fungible Token Transfers",p.value="ERC-721"):t.txsType==="erc1155"&&(s.value="Multi-Token Token Tracker",p.value="ERC-1155")},k=async a=>{e.length=0,n.value=1,i.value=a;const o=await v(n.value-1,i.value,t.txsType,c);o.data.items.forEach(u=>{e.push(u)}),l.value=o.data.total},z=async a=>{e.length=0,n.value=a;const o=await v(n.value-1,i.value,t.txsType,c);o.data.items.forEach(u=>{e.push(u)}),l.value=o.data.total};return(a,o)=>{const u=R,E=J,j=G,C=I;return x(),F("div",null,[f("div",Q,[f("h4",U,g(s.value),1),p.value!==""?(x(),h(u,{key:0,style:{margin:"10px","font-weight":"bold"},color:"#DEE1E4",size:"small"},{default:H(()=>[q(g(p.value),1)]),_:1})):P("",!0)]),f("div",null,[f("h4",X,"(Showing the last "+g(l.value)+" records only)",1)]),t.txsType==="all"?(x(),h(E,{key:0,txsData:e,headerData:r,loadStatus:m.value},null,8,["txsData","headerData","loadStatus"])):(x(),h(j,{key:1,txsData:e,headerData:r,loadStatus:m.value},null,8,["txsData","headerData","loadStatus"])),f("div",Y,[V(C,{small:"",background:"",currentPage:n.value,"page-size":i.value,"page-sizes":[10,25,50,100],layout:"total, sizes, prev, pager, next, jumper",total:l.value,onSizeChange:k,onCurrentChange:z},null,8,["currentPage","page-size","total"])])])}}});var ht=S(Z,[["__scopeId","data-v-7a0174d6"]]);export{ht as default};
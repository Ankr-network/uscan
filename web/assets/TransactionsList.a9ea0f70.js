import{_ as S,d as N,bD as D,m as L,F as s,aM as y,j as M,J as B,i as R,o as w,c as F,e as f,t as x,n as P,w as V,W as H,a as _,b as q}from"./index.20bab758.js";import{E as I}from"./el-pagination.b6498350.js";/* empty css               */import"./el-select.d3b44f71.js";import{_ as G}from"./GenerateTransactions.2f1643cd.js";import{a as b}from"./transactionService.e41be4a8.js";import{T as J,E as W,a as T}from"./transaction.3b89b484.js";import"./isEqual.7330d602.js";import"./index.53af32c3.js";import"./index.f59113f4.js";import"./el-table-column.b9143553.js";import"./objects.5021561f.js";import"./index.eb369523.js";import"./_commonjsHelpers.b8add541.js";import"./index.7d4bc839.js";const A={class:"center-row"},K={class:"h4-title"},O={class:"tx-sub-title"},Q={style:{"margin-top":"1%",display:"flex","justify-content":"center"}},U=N({__name:"TransactionsList",props:{txsType:String},setup(k){const t=k,g=D();document.title="Transactions | The "+L+" Explorer";const o=s(1),d=s(25),n=y([]),r=y([]),l=s(0),c=g.query.block===void 0?-1:g.query.block,i=s("title"),p=s(""),m=s(!0);M(()=>{v()});const h=async()=>{if(t.txsType==="all"||t.txsType==="erc20"||t.txsType==="erc721"||t.txsType==="erc1155"){t.txsType==="all"?r.push(...J):t.txsType==="erc20"?r.push(...W):t.txsType==="erc721"?r.push(...T):t.txsType==="erc1155"&&r.push(...T);const e=await b(o.value-1,d.value,t.txsType,c);e.data.items.forEach(a=>{n.push(a)}),l.value=e.data.total,e.data.total==0&&(m.value=!1)}};h(),B(t,async()=>{v(),n.length=0,r.length=0,o.value=1,d.value=25,h()});const v=()=>{t.txsType==="all"?c!==-1?i.value="Transactions For Block "+c:i.value="Transactions":t.txsType==="erc20"?(i.value="Token Transfers",p.value="ERC-20"):t.txsType==="erc721"?(i.value="Non-Fungible Token Transfers",p.value="ERC-721"):t.txsType==="erc1155"&&(i.value="Multi-Token Token Tracker",p.value="ERC-1155")},z=async e=>{n.length=0,o.value=1,d.value=e;const a=await b(o.value-1,d.value,t.txsType,c);a.data.items.forEach(u=>{n.push(u)}),l.value=a.data.total},E=async e=>{n.length=0,o.value=e;const a=await b(o.value-1,d.value,t.txsType,c);a.data.items.forEach(u=>{n.push(u)}),l.value=a.data.total};return(e,a)=>{const u=R,j=G,C=I;return w(),F("div",null,[f("div",A,[f("h4",K,x(i.value),1),p.value!==""?(w(),P(u,{key:0,style:{margin:"10px","font-weight":"bold"},color:"#DEE1E4",size:"small"},{default:V(()=>[q(x(p.value),1)]),_:1})):H("",!0)]),f("div",null,[f("h4",O,"(Showing the last "+x(l.value)+" records only)",1)]),_(j,{txsData:n,headerData:r,loadStatus:m.value},null,8,["txsData","headerData","loadStatus"]),f("div",Q,[_(C,{small:"",background:"",currentPage:o.value,"page-size":d.value,"page-sizes":[10,25,50,100],layout:"total, sizes, prev, pager, next, jumper",total:l.value,onSizeChange:z,onCurrentChange:E},null,8,["currentPage","page-size","total"])])])}}});var pt=S(U,[["__scopeId","data-v-19dd7b30"]]);export{pt as default};

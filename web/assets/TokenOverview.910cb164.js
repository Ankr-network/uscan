import{_ as P,d as L,F as l,aM as z,l as T,J as O,o as F,c as G,e as n,a as t,w as a,p as H,h as R,b as h,t as E,f as x}from"./index.2c404f35.js";import{E as A,a as J}from"./el-tab-pane.95805e0e.js";import{E as U}from"./el-pagination.a591102c.js";/* empty css               */import"./el-select.a48810b9.js";import{_ as q}from"./GenerateTransactions.f832ab9e.js";import{E as K}from"./el-card.0711297d.js";import{E as Q,a as W}from"./el-col.7aeffd1d.js";import{d as X}from"./tokenService.879c035d.js";import{e as Y}from"./transactionService.7d903ab1.js";import{a as v,T as Z,E as $}from"./transaction.2bb97b38.js";import"./isEqual.231a092a.js";import"./index.a371a541.js";import"./index.6c8aa9e2.js";import"./el-table-column.3cca0966.js";import"./objects.a9ca95bb.js";import"./index.eb369523.js";import"./_commonjsHelpers.b8add541.js";import"./utils.409af791.js";import"./index.7d4bc839.js";const j=i=>(H("data-v-f3171e3c"),i=i(),R(),i),ee=j(()=>n("div",{class:"center-row"},[n("h2",null,"Token")],-1)),te=j(()=>n("div",{class:"card-header"},[n("span",null,"Overview")],-1)),ae={class:"card-content"},oe=h("Contract:"),ne=h("Token ID:"),re={style:{"margin-top":"1%",display:"flex","justify-content":"center"}},ie=L({__name:"TokenOverview",props:{address:String,tokenID:String,type:String},async setup(i){let r,f;const e=i,_=l("txs"),s=z([]),c=z([]),d=l(1),p=l(25),y=l(0),b=l(!0),w=([r,f]=T(()=>X(e.address,e.tokenID,e.type)),r=await r,f(),r),u=async()=>{if(e.type==="all"||e.type==="erc20"||e.type==="erc721"||e.type==="erc1155"){e.type==="all"?c.push(...Z):e.type==="erc20"?c.push(...$):e.type==="erc721"?c.push(...v):e.type==="erc1155"&&c.push(...v);const o=await Y(d.value-1,p.value,e.type,e.address);o.data.items.forEach(m=>{s.push(m)}),y.value=o.data.total,o.data.total==0&&(b.value=!1)}};[r,f]=T(()=>u()),await r,f(),O(e,async()=>{e.address!==void 0&&e.tokenID!==void 0&&e.type!==void 0&&(s.length=0,c.length=0,d.value=1,p.value=25,u())});const C=async o=>{s.length=0,d.value=1,p.value=o,u()},D=async o=>{s.length=0,d.value=o,u()};return(o,m)=>{const g=Q,k=W,S=K,I=q,N=U,V=A,B=J;return F(),G("div",null,[ee,n("div",null,[t(S,{class:"box-card"},{header:a(()=>[te]),default:a(()=>[n("div",ae,[t(k,null,{default:a(()=>[t(g,{span:10},{default:a(()=>[oe]),_:1}),t(g,{span:14},{default:a(()=>[h(E(x(w).data.contract),1)]),_:1})]),_:1}),t(k,null,{default:a(()=>[t(g,{span:10},{default:a(()=>[ne]),_:1}),t(g,{span:14},{default:a(()=>[h(E(x(w).data.tokenID),1)]),_:1})]),_:1})])]),_:1})]),n("div",null,[t(B,{modelValue:_.value,"onUpdate:modelValue":m[0]||(m[0]=M=>_.value=M)},{default:a(()=>[t(V,{label:"Transactions",name:"txs"},{default:a(()=>[t(I,{txsData:s,headerData:x(v),loadStatus:b.value},null,8,["txsData","headerData","loadStatus"]),n("div",re,[t(N,{small:"",background:"",currentPage:d.value,"page-size":p.value,"page-sizes":[10,25,50,100],layout:"total, sizes, prev, pager, next, jumper",total:y.value,onSizeChange:C,onCurrentChange:D},null,8,["currentPage","page-size","total"])])]),_:1})]),_:1},8,["modelValue"])])])}}});var je=P(ie,[["__scopeId","data-v-f3171e3c"]]);export{je as default};
import{_ as X,d as U,F as k,aM as A,l as P,y as me,r as W,o as f,c as x,e as o,V as q,ah as F,a as e,w as t,n as J,b as l,t as b,f as v,p as Y,h as Z,W as O,m as he}from"./index.14ec23e5.js";import{E as ge,a as ve}from"./el-tab-pane.591271f2.js";import{E as ee}from"./el-pagination.ac00537f.js";/* empty css               */import"./el-select.b1c77ca2.js";import{E as te,a as ae}from"./el-col.c7c36801.js";import{E as oe}from"./el-card.83574051.js";import{G as L,a as be,b as M,c as R}from"./tokenService.273c8720.js";import{E as xe,a as ye}from"./el-table-column.21dfafb8.js";import{_ as we}from"./GenerateTransactions.1040e00a.js";import{_ as ke}from"./CopyIcon.843d60dd.js";import{b as Te,c as ze,d as Ce,e as Ee}from"./transaction.ec615c86.js";import{G as Se}from"./addressService.e4d2f1c6.js";import"./isEqual.6ab61c57.js";import"./index.219237ce.js";import"./index.326a1e4c.js";import"./objects.361feb61.js";import"./transactionService.31a7d440.js";import"./index.eb369523.js";import"./_commonjsHelpers.b8add541.js";import"./index.7d4bc839.js";const Ie=d=>(Y("data-v-076662fe"),d=d(),Z(),d),De=Ie(()=>o("div",{style:{height:"100px",width:"100px","background-color":"#598df6","border-radius":"0.35rem"}},null,-1)),je={class:"text-secondary"},He=l(" TokenID: "),Ae={class:"text-secondary"},Pe=l(" Owner: "),$e={style:{"margin-top":"1%",display:"flex","justify-content":"center"}},Ge=U({__name:"GenerateInventory",props:{address:{type:String,require:!0},ercType:{type:String,require:!0}},async setup(d){let a,m;const r=d,g=k(1),_=k(25),p=A([]),u=k(1),s=([a,m]=P(()=>L(r.address,r.ercType,g.value-1,_.value)),a=await a,m(),a);u.value=s.data.total,s.data.items.forEach(c=>{p.push(c)});const D=me({get(){const c=[];return p.forEach((E,n)=>{const C=Math.floor(n/6);c[C]||(c[C]=[]),c[C].push(E)}),c},set(){}}),T=async c=>{p.length=0,g.value=1,_.value=c,(await L(r.address,r.ercType,g.value-1,_.value)).data.items.forEach(n=>{p.push(n)})},z=async c=>{p.length=0,g.value=c,(await L(r.address,r.ercType,g.value-1,_.value)).data.items.forEach(n=>{p.push(n)})};return(c,E)=>{const n=W("router-link"),C=oe,$=te,G=ae,B=ee;return f(),x("div",null,[o("div",null,[(f(!0),x(q,null,F(v(D),(j,N)=>(f(),x("div",{key:N,style:{"margin-top":"20px"}},[e(G,{gutter:20},{default:t(()=>[(f(!0),x(q,null,F(j,(y,H)=>(f(),J($,{span:4,key:H},{default:t(()=>[e(C,{"body-style":{padding:"10px"}},{default:t(()=>[De,o("div",je,[He,o("span",null,[e(n,{to:"/token/nfts/"+y.contract+"/"+y.tokenID+"/"+r.ercType},{default:t(()=>[l(b(y.tokenID),1)]),_:2},1032,["to"])])]),o("div",Ae,[Pe,o("span",null,[e(n,{to:"/address/"+y.owner},{default:t(()=>[l(b(y.owner.slice(0,18)+"..."),1)]),_:2},1032,["to"])])])]),_:2},1024)]),_:2},1024))),128))]),_:2},1024)]))),128)),o("div",$e,[e(B,{small:"",background:"",currentPage:g.value,"page-size":_.value,"page-sizes":[10,25,50,100],layout:"total, sizes, prev, pager, next, jumper",total:u.value,onSizeChange:T,onCurrentChange:z},null,8,["currentPage","page-size","total"])])])])}}});var Be=X(Ge,[["__scopeId","data-v-076662fe"]]);const Ne={key:0},Ve={key:1,style:{width:"180px"}},Le=U({__name:"GenerateHolders",props:{holdersData:{type:Array,require:!0},headerData:{type:Array,require:!0}},setup(d){const a=d;return(m,r)=>{const g=W("router-link"),_=xe,p=ye;return f(),x("div",null,[e(p,{class:"table-border",data:a.holdersData,"empty-text":"loading...","row-style":{height:"50px"}},{default:t(()=>[(f(!0),x(q,null,F(a.headerData,u=>(f(),J(_,{key:u.key,property:u.key,label:u.label},{default:t(s=>[s.column.property=="owner"?(f(),x("div",Ne,[e(g,{to:"/address/"+s.row[s.column.property]},{default:t(()=>[l(b(s.row[s.column.property]),1)]),_:2},1032,["to"])])):s.column.property=="quantity"?(f(),x("div",Ve,b(BigInt(parseInt(s.row[s.column.property]))),1)):O("",!0)]),_:2},1032,["property","label"]))),128))]),_:1},8,["data"])])}}});const ne=d=>(Y("data-v-7077b40c"),d=d(),Z(),d),Me={class:"center-row"},Re={class:"h4-title"},qe=l(" Token "),Fe={class:"small text-secondary"},Oe=l(" \xA0 "),Ue=ne(()=>o("div",{class:"card-header"},[o("span",null,"Overview")],-1)),We={class:"card-content"},Je=l("Max Total Supply:"),Ke=l("Holders:"),Qe=l("Transfers:"),Xe=ne(()=>o("div",{class:"card-header"},[o("span",null,"Profile Summary")],-1)),Ye={class:"card-content"},Ze=l("Contract:"),et={key:0},tt=l("Decimals:"),at={style:{"margin-top":"1%",display:"flex","justify-content":"center"}},ot={style:{"margin-top":"1%",display:"flex","justify-content":"center"}},nt=U({__name:"TokenAddress",props:{address:String},async setup(d){let a,m;const r=d;document.title="Token | The "+he+" Explorer";const g=k("transactions"),_=k(1),p=k(25),u=k(1),s=k(25),D=A([]),T=A([]),z=A([]),c=A([]),E=k(!0);let n="";const C=([a,m]=P(()=>be(r.address)),a=await a,m(),a);for(const i in C.data)C.data[i]!=0&&(n=i);n=="erc20"?(D.push(...Te),c.push(...ze)):(D.push(...Ce),c.push(...Ee));const $=([a,m]=P(()=>Se(r.address)),a=await a,m(),a),G=$.data.tokenTotalSupply,B=$.data.decimals,j=([a,m]=P(()=>M(r.address,n,_.value-1,p.value)),a=await a,m(),a),N=j.data.total;j.data.total!==0?j.data.items.forEach(i=>z.push(i)):E.value=!1;const y=([a,m]=P(()=>R(r.address,n,u.value-1,s.value)),a=await a,m(),a),H=y.data.total;y.data.total!==0&&y.data.items.forEach(i=>{T.push(i)});const re=async i=>{T.length=0,u.value=1,s.value=i,(await R(r.address,n,u.value-1,s.value)).data.items.forEach(w=>{T.push(w)})},se=async i=>{T.length=0,u.value=i,(await R(r.address,n,u.value-1,s.value)).data.items.forEach(w=>{T.push(w)})},le=async i=>{z.length=0,_.value=1,p.value=i,(await M(r.address,n,_.value-1,p.value)).data.items.forEach(w=>{z.push(w)})},de=async i=>{z.length=0,u.value=i,(await M(r.address,n,_.value-1,p.value)).data.items.forEach(w=>{z.push(w)})};return(i,S)=>{const w=ke,h=te,I=ae,K=oe,ce=W("router-link"),ie=we,Q=ee,V=ge,pe=Le,ue=Be,_e=ve;return f(),x("div",null,[o("div",Me,[o("h4",Re,[qe,o("span",Fe,"\xA0\xA0"+b(r.address),1)]),Oe,e(w,{text:r.address},null,8,["text"])]),o("div",null,[e(I,{gutter:20},{default:t(()=>[e(h,{span:12},{default:t(()=>[o("div",null,[e(K,{class:"box-card-address"},{header:t(()=>[Ue]),default:t(()=>[o("div",We,[e(I,null,{default:t(()=>[e(h,{span:9},{default:t(()=>[Je]),_:1}),e(h,{span:15},{default:t(()=>[l(b(v(G)),1)]),_:1})]),_:1}),e(I,null,{default:t(()=>[e(h,{span:9},{default:t(()=>[Ke]),_:1}),e(h,{span:15},{default:t(()=>[l(b(v(N)),1)]),_:1})]),_:1}),e(I,null,{default:t(()=>[e(h,{span:9},{default:t(()=>[Qe]),_:1}),e(h,{span:15},{default:t(()=>[l(b(v(H)),1)]),_:1})]),_:1})])]),_:1})])]),_:1}),e(h,{span:12},{default:t(()=>[o("div",null,[e(K,{class:"box-card-address"},{header:t(()=>[Xe]),default:t(()=>[o("div",Ye,[e(I,null,{default:t(()=>[e(h,{span:9},{default:t(()=>[Ze]),_:1}),e(h,{span:15},{default:t(()=>[e(ce,{to:"/address/"+d.address},{default:t(()=>[l(b(d.address),1)]),_:1},8,["to"])]),_:1})]),_:1}),v(n)=="erc20"?(f(),x("div",et,[e(I,null,{default:t(()=>[e(h,{span:9},{default:t(()=>[tt]),_:1}),e(h,{span:15},{default:t(()=>[l(b(v(B)),1)]),_:1})]),_:1})])):O("",!0)])]),_:1})])]),_:1})]),_:1})]),o("div",null,[e(_e,{modelValue:g.value,"onUpdate:modelValue":S[0]||(S[0]=fe=>g.value=fe)},{default:t(()=>[e(V,{label:"Transactions",name:"transactions"},{default:t(()=>[e(ie,{txsData:T,headerData:D,loadStatus:E.value},null,8,["txsData","headerData","loadStatus"]),o("div",at,[e(Q,{small:"",background:"",currentPage:u.value,"page-size":s.value,"page-sizes":[10,25,50,100],layout:"total, sizes, prev, pager, next, jumper",total:v(H),onSizeChange:re,onCurrentChange:se},null,8,["currentPage","page-size","total"])])]),_:1}),e(V,{label:"Holders",name:"holders"},{default:t(()=>[e(pe,{holdersData:z,headerData:c},null,8,["holdersData","headerData"]),o("div",ot,[e(Q,{small:"",background:"",currentPage:_.value,"page-size":p.value,"page-sizes":[10,25,50,100],layout:"total, sizes, prev, pager, next, jumper",total:v(H),onSizeChange:le,onCurrentChange:de},null,8,["currentPage","page-size","total"])])]),_:1}),v(n)==="erc721"||v(n)==="erc1155"?(f(),J(V,{key:0,label:"Inventory",name:"inventory"},{default:t(()=>[e(ue,{address:r.address,ercType:v(n)},null,8,["address","ercType"])]),_:1})):O("",!0)]),_:1},8,["modelValue"])])])}}});var Ct=X(nt,[["__scopeId","data-v-7077b40c"]]);export{Ct as default};
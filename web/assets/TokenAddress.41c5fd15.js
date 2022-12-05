import{_ as Y,d as U,F as w,aM as P,l as $,y as be,r as W,o as m,c as x,e as o,V as F,ah as O,a as e,w as t,n as J,b as s,t as v,f as h,p as Z,h as ee,W as Q,m as he}from"./index.aa130625.js";import{E as ge,a as ve}from"./el-tab-pane.2c73465d.js";import{E as te}from"./el-pagination.81330167.js";/* empty css               */import"./el-select.9226009c.js";import{E as ae,a as oe}from"./el-col.b1325d28.js";import{E as ne}from"./el-card.7b27de66.js";import{G as M,a as xe,b as R,c as q}from"./tokenService.9fe06fbf.js";import{E as ye,a as we}from"./el-table-column.d3ee0114.js";import{_ as ke}from"./GenerateTransactions.cc3cd391.js";import{_ as Te}from"./CopyIcon.a5ddd055.js";import{b as ze,c as Ce,d as Ee,e as Se}from"./transaction.2bb97b38.js";import{a as Ie}from"./addressService.24bfaa59.js";import"./isEqual.131f71a0.js";import"./index.404a38d5.js";import"./index.b0e21fc0.js";import"./objects.42eb58f6.js";import"./transactionService.fbf4207e.js";import"./index.eb369523.js";import"./_commonjsHelpers.b8add541.js";import"./utils.409af791.js";import"./index.7d4bc839.js";const De=d=>(Z("data-v-076662fe"),d=d(),ee(),d),je=De(()=>o("div",{style:{height:"100px",width:"100px","background-color":"#598df6","border-radius":"0.35rem"}},null,-1)),He={class:"text-secondary"},Ae=s(" TokenID: "),Pe={class:"text-secondary"},$e=s(" Owner: "),Be={style:{"margin-top":"1%",display:"flex","justify-content":"center"}},Ge=U({__name:"GenerateInventory",props:{address:{type:String,require:!0},ercType:{type:String,require:!0}},async setup(d){let a,b;const r=d,g=w(1),f=w(25),p=P([]),u=w(1),l=([a,b]=$(()=>M(r.address,r.ercType,g.value-1,f.value)),a=await a,b(),a);u.value=l.data.total,l.data.items.forEach(c=>{p.push(c)});const D=be({get(){const c=[];return p.forEach((C,n)=>{const z=Math.floor(n/6);c[z]||(c[z]=[]),c[z].push(C)}),c},set(){}}),k=async c=>{p.length=0,g.value=1,f.value=c,(await M(r.address,r.ercType,g.value-1,f.value)).data.items.forEach(n=>{p.push(n)})},T=async c=>{p.length=0,g.value=c,(await M(r.address,r.ercType,g.value-1,f.value)).data.items.forEach(n=>{p.push(n)})};return(c,C)=>{const n=W("router-link"),z=ne,j=ae,B=oe,G=te;return m(),x("div",null,[o("div",null,[(m(!0),x(F,null,O(h(D),(N,H)=>(m(),x("div",{key:H,style:{"margin-top":"20px"}},[e(B,{gutter:20},{default:t(()=>[(m(!0),x(F,null,O(N,(E,A)=>(m(),J(j,{span:4,key:A},{default:t(()=>[e(z,{"body-style":{padding:"10px"}},{default:t(()=>[je,o("div",He,[Ae,o("span",null,[e(n,{to:"/token/nfts/"+E.contract+"/"+E.tokenID+"/"+r.ercType},{default:t(()=>[s(v(E.tokenID),1)]),_:2},1032,["to"])])]),o("div",Pe,[$e,o("span",null,[e(n,{to:"/address/"+E.owner},{default:t(()=>[s(v(E.owner.slice(0,18)+"..."),1)]),_:2},1032,["to"])])])]),_:2},1024)]),_:2},1024))),128))]),_:2},1024)]))),128)),o("div",Be,[e(G,{small:"",background:"",currentPage:g.value,"page-size":f.value,"page-sizes":[10,25,50,100],layout:"total, sizes, prev, pager, next, jumper",total:u.value,onSizeChange:k,onCurrentChange:T},null,8,["currentPage","page-size","total"])])])])}}});var Ne=Y(Ge,[["__scopeId","data-v-076662fe"]]);const Ve={key:0},Le={key:1,style:{width:"380px"}},Me=U({__name:"GenerateHolders",props:{holdersData:{type:Array,require:!0},headerData:{type:Array,require:!0}},setup(d){const a=d;return(b,r)=>{const g=W("router-link"),f=ye,p=we;return m(),x("div",null,[e(p,{class:"table-border",data:a.holdersData,"empty-text":"loading...","row-style":{height:"50px"}},{default:t(()=>[(m(!0),x(F,null,O(a.headerData,u=>(m(),J(f,{key:u.key,property:u.key,label:u.label},{default:t(l=>[l.column.property=="Address"?(m(),x("div",Ve,[e(g,{to:"/address/"+l.row[l.column.property]},{default:t(()=>[s(v(l.row[l.column.property]),1)]),_:2},1032,["to"])])):l.column.property=="Quantity"?(m(),x("div",Le,v(BigInt(parseInt(l.row[l.column.property]))),1)):Q("",!0)]),_:2},1032,["property","label"]))),128))]),_:1},8,["data"])])}}});const re=d=>(Z("data-v-6cb83b52"),d=d(),ee(),d),Re={class:"center-row"},qe={class:"h4-title"},Fe=s(" Token "),Oe={class:"small text-secondary"},Qe=s(" \xA0 "),Ue=re(()=>o("div",{class:"card-header"},[o("span",null,"Overview")],-1)),We={class:"card-content"},Je=s("Max Total Supply:"),Ke=s("Holders:"),Xe=s("Transfers:"),Ye=re(()=>o("div",{class:"card-header"},[o("span",null,"Profile Summary")],-1)),Ze={class:"card-content"},et=s("Contract:"),tt={key:0},at=s("Decimals:"),ot=s("Symbol:"),nt={style:{"margin-top":"1%",display:"flex","justify-content":"center"}},rt={style:{"margin-top":"1%",display:"flex","justify-content":"center"}},st=U({__name:"TokenAddress",props:{address:String},async setup(d){let a,b;const r=d;document.title="Token | The "+he()+" Explorer";const g=w("transactions"),f=w(1),p=w(25),u=w(1),l=w(25),D=P([]),k=P([]),T=P([]),c=P([]),C=w(!0);let n="";const z=([a,b]=$(()=>xe(r.address)),a=await a,b(),a);for(const i in z.data)z.data[i]!=0&&(n=i);n=="erc20"?(D.push(...ze),c.push(...Ce)):(D.push(...Ee),c.push(...Se));const j=([a,b]=$(()=>Ie(r.address)),a=await a,b(),a),B=j.data.tokenTotalSupply,G=j.data.decimals,N=j.data.symbol,H=([a,b]=$(()=>R(r.address,n,f.value-1,p.value)),a=await a,b(),a),E=H.data.total;H.data.total!==0?H.data.items.forEach(i=>T.push(i)):C.value=!1;const A=([a,b]=$(()=>q(r.address,n,u.value-1,l.value)),a=await a,b(),a),V=A.data.total;A.data.total!==0&&A.data.items.forEach(i=>{k.push(i)});const se=async i=>{k.length=0,u.value=1,l.value=i,(await q(r.address,n,u.value-1,l.value)).data.items.forEach(y=>{k.push(y)})},le=async i=>{k.length=0,u.value=i,(await q(r.address,n,u.value-1,l.value)).data.items.forEach(y=>{k.push(y)})},de=async i=>{T.length=0,f.value=1,p.value=i,(await R(r.address,n,f.value-1,p.value)).data.items.forEach(y=>{T.push(y)})},ce=async i=>{T.length=0,u.value=i,(await R(r.address,n,f.value-1,p.value)).data.items.forEach(y=>{T.push(y)})};return(i,I)=>{const y=Te,_=ae,S=oe,K=ne,ie=W("router-link"),pe=ke,X=te,L=ge,ue=Me,_e=Ne,fe=ve;return m(),x("div",null,[o("div",Re,[o("h4",qe,[Fe,o("span",Oe,"\xA0\xA0"+v(r.address),1)]),Qe,e(y,{text:r.address},null,8,["text"])]),o("div",null,[e(S,{gutter:20},{default:t(()=>[e(_,{span:12},{default:t(()=>[o("div",null,[e(K,{class:"box-card-address"},{header:t(()=>[Ue]),default:t(()=>[o("div",We,[e(S,null,{default:t(()=>[e(_,{span:9},{default:t(()=>[Je]),_:1}),e(_,{span:15},{default:t(()=>[s(v(parseInt(h(B))),1)]),_:1})]),_:1}),e(S,null,{default:t(()=>[e(_,{span:9},{default:t(()=>[Ke]),_:1}),e(_,{span:15},{default:t(()=>[s(v(h(E)),1)]),_:1})]),_:1}),e(S,null,{default:t(()=>[e(_,{span:9},{default:t(()=>[Xe]),_:1}),e(_,{span:15},{default:t(()=>[s(v(h(V)),1)]),_:1})]),_:1})])]),_:1})])]),_:1}),e(_,{span:12},{default:t(()=>[o("div",null,[e(K,{class:"box-card-address"},{header:t(()=>[Ye]),default:t(()=>[o("div",Ze,[e(S,null,{default:t(()=>[e(_,{span:9},{default:t(()=>[et]),_:1}),e(_,{span:15},{default:t(()=>[e(ie,{to:"/address/"+d.address},{default:t(()=>[s(v(d.address),1)]),_:1},8,["to"])]),_:1})]),_:1}),h(n)=="erc20"?(m(),x("div",tt,[e(S,null,{default:t(()=>[e(_,{span:9},{default:t(()=>[at]),_:1}),e(_,{span:15},{default:t(()=>[s(v(h(G)),1)]),_:1})]),_:1}),e(S,null,{default:t(()=>[e(_,{span:9},{default:t(()=>[ot]),_:1}),e(_,{span:15},{default:t(()=>[s(v(h(N)),1)]),_:1})]),_:1})])):Q("",!0)])]),_:1})])]),_:1})]),_:1})]),o("div",null,[e(fe,{modelValue:g.value,"onUpdate:modelValue":I[0]||(I[0]=me=>g.value=me)},{default:t(()=>[e(L,{label:"Transactions",name:"transactions"},{default:t(()=>[e(pe,{txsData:k,headerData:D,loadStatus:C.value},null,8,["txsData","headerData","loadStatus"]),o("div",nt,[e(X,{small:"",background:"",currentPage:u.value,"page-size":l.value,"page-sizes":[10,25,50,100],layout:"total, sizes, prev, pager, next, jumper",total:h(V),onSizeChange:se,onCurrentChange:le},null,8,["currentPage","page-size","total"])])]),_:1}),e(L,{label:"Holders",name:"holders"},{default:t(()=>[e(ue,{holdersData:T,headerData:c},null,8,["holdersData","headerData"]),o("div",rt,[e(X,{small:"",background:"",currentPage:f.value,"page-size":p.value,"page-sizes":[10,25,50,100],layout:"total, sizes, prev, pager, next, jumper",total:h(V),onSizeChange:de,onCurrentChange:ce},null,8,["currentPage","page-size","total"])])]),_:1}),h(n)==="erc721"||h(n)==="erc1155"?(m(),J(L,{key:0,label:"Inventory",name:"inventory"},{default:t(()=>[e(_e,{address:r.address,ercType:h(n)},null,8,["address","ercType"])]),_:1})):Q("",!0)]),_:1},8,["modelValue"])])])}}});var It=Y(st,[["__scopeId","data-v-6cb83b52"]]);export{It as default};

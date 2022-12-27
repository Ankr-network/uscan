import{_ as y,d as _,aM as B,u as E,aq as V,r as C,o as r,c as a,a as e,w as t,e as p,t as i,f as u,aN as T,b as l,aO as I,aP as M,aQ as A,ac as $,E as S,i as O,aR as G,F as R,n as F,S as q}from"./index.692e6b45.js";import{E as D,a as L}from"./el-tab-pane.1c3f2e77.js";import{E as P,a as Q}from"./el-table-column.5a54df2b.js";/* empty css               */import{G as U,g as H}from"./block.95ca738b.js";import{g as J}from"./utils.fc402d12.js";import"./index.913cbcc8.js";import"./objects.2b4edb06.js";import"./index.7d4bc839.js";const K={class:"center-row"},W={style:{"max-width":"250px"}},X={key:0,style:{"font-weight":"900"}},Y={key:1,class:"center-row"},Z={key:2},ee={key:0},te={key:1},oe=l(" \xA0in this block "),re={key:3},ae={key:4},de=_({__name:"BlockOverview",props:{blockNumber:Number},setup(n){const m=n,c=B([]),g=E(),b=function(s){g.push("/block/"+s)},h=function(){g.push("/txs/all?block="+m.blockNumber)};return V(async()=>{c.length=0;const s=await U(m.blockNumber);H(s.data).forEach(d=>c.push(d))}),(s,d)=>{const f=$,v=S,w=P,x=O,k=G,z=C("router-link"),j=Q;return r(),a("div",null,[e(j,{class:"table-border",data:c,"show-header":!1,"empty-text":"loading...","row-style":{height:"50px"}},{default:t(()=>[e(w,{width:"240"},{default:t(o=>[p("div",K,[e(v,{effect:"dark",placement:"top"},{content:t(()=>[p("div",W,i(o.row.parameterExplain),1)]),default:t(()=>[e(f,null,{default:t(()=>[e(u(T))]),_:1})]),_:2},1024),l(" \xA0"+i(o.row.parameterDisplay),1)])]),_:1}),e(w,{prop:"parameterValue"},{default:t(o=>[o.row.parameterName=="number"?(r(),a("div",X,[l(i(o.row.parameterValue)+" \xA0 ",1),e(k,null,{default:t(()=>[e(x,{type:"primary",size:"small",style:{border:"0"},plain:"",onClick:d[0]||(d[0]=N=>b(parseInt(n.blockNumber)-1))},{default:t(()=>[e(f,null,{default:t(()=>[e(u(I))]),_:1})]),_:1}),e(x,{type:"primary",size:"small",style:{border:"0"},plain:"",onClick:d[1]||(d[1]=N=>b(parseInt(n.blockNumber)+1))},{default:t(()=>[e(f,null,{default:t(()=>[e(u(M))]),_:1})]),_:1})]),_:1})])):o.row.parameterName=="timestamp"?(r(),a("div",Y,[e(f,null,{default:t(()=>[e(u(A))]),_:1}),l("\xA0"+i(u(J)(o.row.parameterValue)),1)])):o.row.parameterName=="transactionsTotal"?(r(),a("div",Z,[o.row.parameterValue==0?(r(),a("div",ee,"0 transaction in this block")):(r(),a("div",te,[e(v,{class:"box-item",effect:"dark",content:"Click to view Transactions",placement:"right"},{default:t(()=>[e(x,{type:"primary",plain:"",size:"small",onClick:h,style:{border:"0"}},{default:t(()=>[l(i(o.row.parameterValue)+" transactions ",1)]),_:2},1024),oe]),_:2},1024)]))])):o.row.parameterName=="miner"?(r(),a("div",re,[e(z,{to:"/address/"+o.row.parameterValue},{default:t(()=>[l(i(o.row.parameterValue),1)]),_:2},1032,["to"])])):(r(),a("div",ae,i(o.row.parameterValue),1))]),_:1})]),_:1},8,["data"])])}}});var ie=y(de,[["__scopeId","data-v-de7634de"]]);const ne={style:{"margin-bottom":"-20px"}},le={class:"h4-title"},ce=l(" Block"),se={class:"small text-secondary"},fe=_({__name:"BlockIndex",props:{blockNumber:String},setup(n){const m=n;document.title="Blocks #"+m.blockNumber;const c=R("first");return(g,b)=>{const h=ie,s=D,d=L;return r(),a("div",null,[p("div",ne,[p("h4",le,[ce,p("span",se,"\xA0\xA0#"+i(n.blockNumber),1)])]),e(d,{modelValue:c.value,"onUpdate:modelValue":b[0]||(b[0]=f=>c.value=f),style:{}},{default:t(()=>[e(s,{label:"Overview",name:"first"},{default:t(()=>[(r(),F(q,null,{default:t(()=>[e(h,{blockNumber:parseInt(n.blockNumber)},null,8,["blockNumber"])]),_:1}))]),_:1})]),_:1},8,["modelValue"])])}}});var ye=y(fe,[["__scopeId","data-v-cf26eb68"]]);export{ye as default};

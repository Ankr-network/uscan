import{_ as w,d as _,aM as B,u as E,aq as V,r as T,o as r,c as a,a as t,w as e,e as p,t as n,f,aN as C,b as d,aO as I,aP as M,aQ as A,g as $,ac as S,E as O,i as G,aR as R,m as F,F as q,n as D,S as L}from"./index.20bab758.js";import{E as P,a as Q}from"./el-tab-pane.037a5237.js";import{E as U,a as H}from"./el-table-column.b9143553.js";/* empty css               */import{G as J,g as K}from"./block.efff100e.js";import"./index.53af32c3.js";import"./objects.5021561f.js";import"./index.7d4bc839.js";const W={class:"center-row"},X={style:{"max-width":"250px"}},Y={key:0,style:{"font-weight":"900"}},Z={key:1,class:"center-row"},tt={key:2},et={key:0},ot={key:1},rt=d(" \xA0in this block "),at={key:3},it={key:4},nt=_({__name:"BlockOverview",props:{blockNumber:Number},setup(l){const g=l,s=B([]),m=E(),u=function(c){m.push("/block/"+c)},h=function(){m.push("/txs/all?block="+g.blockNumber)};return V(async()=>{s.length=0;const c=await J(g.blockNumber);K(c.data).forEach(i=>s.push(i))}),(c,i)=>{const b=S,v=O,y=U,x=G,k=R,z=T("router-link"),j=H;return r(),a("div",null,[t(j,{class:"table-border",data:s,"empty-text":"loading...","row-style":{height:"50px"}},{default:e(()=>[t(y,{width:"240"},{default:e(o=>[p("div",W,[t(v,{effect:"dark",placement:"top"},{content:e(()=>[p("div",X,n(o.row.parameterExplain),1)]),default:e(()=>[t(b,null,{default:e(()=>[t(f(C))]),_:1})]),_:2},1024),d(" \xA0"+n(o.row.parameterDisplay),1)])]),_:1}),t(y,{prop:"parameterValue"},{default:e(o=>[o.row.parameterName=="number"?(r(),a("div",Y,[d(n(o.row.parameterValue)+" \xA0 ",1),t(k,null,{default:e(()=>[t(x,{type:"primary",size:"small",style:{border:"0"},plain:"",onClick:i[0]||(i[0]=N=>u(parseInt(l.blockNumber)-1))},{default:e(()=>[t(b,null,{default:e(()=>[t(f(I))]),_:1})]),_:1}),t(x,{type:"primary",size:"small",style:{border:"0"},plain:"",onClick:i[1]||(i[1]=N=>u(parseInt(l.blockNumber)+1))},{default:e(()=>[t(b,null,{default:e(()=>[t(f(M))]),_:1})]),_:1})]),_:1})])):o.row.parameterName=="timestamp"?(r(),a("div",Z,[t(b,null,{default:e(()=>[t(f(A))]),_:1}),d("\xA0"+n(f($)(o.row.parameterValue)),1)])):o.row.parameterName=="transactionsTotal"?(r(),a("div",tt,[o.row.parameterValue==0?(r(),a("div",et,"0 transaction in this block")):(r(),a("div",ot,[t(v,{class:"box-item",effect:"dark",content:"Click to view Transactions",placement:"right"},{default:e(()=>[t(x,{type:"primary",plain:"",size:"small",onClick:h,style:{border:"0"}},{default:e(()=>[d(n(o.row.parameterValue)+" transactions ",1)]),_:2},1024),rt]),_:2},1024)]))])):o.row.parameterName=="miner"?(r(),a("div",at,[t(z,{to:"/address/"+o.row.parameterValue},{default:e(()=>[d(n(o.row.parameterValue),1)]),_:2},1032,["to"])])):(r(),a("div",it,n(o.row.parameterValue),1))]),_:1})]),_:1},8,["data"])])}}});var lt=w(nt,[["__scopeId","data-v-05b7189a"]]);const dt={style:{"margin-bottom":"-20px"}},st={class:"h4-title"},ct=d(" Block"),bt={class:"small text-secondary"},ut=_({__name:"BlockIndex",props:{blockNumber:String},setup(l){const g=l;document.title="Blocks #"+g.blockNumber+" | The "+F+" Explorer";const s=q("first");return(m,u)=>{const h=lt,c=P,i=Q;return r(),a("div",null,[p("div",dt,[p("h4",st,[ct,p("span",bt,"\xA0\xA0#"+n(l.blockNumber),1)])]),t(i,{modelValue:s.value,"onUpdate:modelValue":u[0]||(u[0]=b=>s.value=b),style:{}},{default:e(()=>[t(c,{label:"Overview",name:"first"},{default:e(()=>[(r(),D(L,null,{default:e(()=>[t(h,{blockNumber:parseInt(l.blockNumber)},null,8,["blockNumber"])]),_:1}))]),_:1})]),_:1},8,["modelValue"])])}}});var wt=w(ut,[["__scopeId","data-v-6025f312"]]);export{wt as default};

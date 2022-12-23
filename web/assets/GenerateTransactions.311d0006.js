import{_ as z,d as C,F as j,aq as E,r as D,E as I,o as a,n as u,w as e,b as i,t as n,a as o,f as s,bD as T,c as r,W as c,ah as N,e as B,bE as H,bF as S,g as V,V as A,ac as F,i as M}from"./index.5c39bb93.js";import{E as q,a as U}from"./el-table-column.9ea7ff5a.js";/* empty css               */import{E as G,_ as L}from"./BaseTransactionInfo.98ad6198.js";import{g as x}from"./utils.fc402d12.js";import{e as h}from"./index.eb369523.js";const P={key:0},R={key:0,style:{width:"170px"}},W={key:1,style:{width:"170px"}},$={class:"span"},J={key:2,style:{width:"170px"}},K={key:3},O={key:4},Q={key:5},X={key:6,style:{width:"170px"}},Y={key:7,style:{width:"170px"}},Z={key:0,style:{display:"flex","align-items":"center"}},tt=i("Contract Creation"),et={key:1,style:{display:"flex","align-items":"center"}},at={key:8},ot={key:9},rt={key:10,style:{"font-size":"11px"}},nt={key:11,style:{width:"170px","font-size":"11px"}},dt=C({__name:"GenerateTransactions",props:{loadStatus:Boolean,txsData:{type:Array,require:!0},headerData:{type:Array,require:!0}},setup(_){const p=_,m=j("loading...");return E(()=>{p.loadStatus||(m.value="empty data")}),(it,lt)=>{const f=F,g=M,w=L,v=G,b=q,l=D("router-link"),y=I,k=U;return a(),u(k,{class:"table-border",data:p.txsData,"empty-text":"loading...","row-style":{height:"50px"}},{empty:e(()=>[i(n(m.value),1)]),default:e(()=>[o(b,{width:"37px"},{default:e(d=>[o(v,{placement:"right",title:"Additional Info",width:320,trigger:"click"},{reference:e(()=>[o(g,{style:{width:"5px"},type:"info",size:"small",onClick:t=>d.row.base=!0},{default:e(()=>[o(f,null,{default:e(()=>[o(s(T))]),_:1})]),_:2},1032,["onClick"])]),default:e(()=>[d.row.base==!0?(a(),r("div",P,[o(w,{txHash:d.row.transactionHash?d.row.transactionHash:d.row.hash},null,8,["txHash"])])):c("",!0)]),_:2},1024)]),_:1}),(a(!0),r(A,null,N(p.headerData,d=>(a(),u(b,{key:d.key,property:d.key,label:d.label},{default:e(t=>[t.column.property=="transactionHash"?(a(),r("div",R,[o(l,{to:"/tx/"+t.row[t.column.property]},{default:e(()=>[i(n(t.row[t.column.property].slice(0,15)+"..."),1)]),_:2},1032,["to"])])):c("",!0),t.column.property=="method"?(a(),r("div",W,[B("span",$,n(t.row[t.column.property]),1)])):c("",!0),t.column.property=="hash"?(a(),r("div",J,[o(l,{to:"/tx/"+t.row[t.column.property]},{default:e(()=>[i(n(t.row[t.column.property].slice(0,15)+"..."),1)]),_:2},1032,["to"])])):t.column.property=="blockNumber"?(a(),r("div",K,[o(l,{to:"/block/"+parseInt(t.row[t.column.property])},{default:e(()=>[i(n(parseInt(t.row[t.column.property])),1)]),_:2},1032,["to"])])):t.column.property=="createTime"?(a(),r("div",O,n(s(x)(t.row[t.column.property])),1)):t.column.property=="createdTime"?(a(),r("div",Q,n(s(x)(t.row[t.column.property])),1)):t.column.property=="from"?(a(),r("div",X,[o(l,{to:"/address/"+t.row[t.column.property]},{default:e(()=>[i(n(t.row[t.column.property].slice(0,15)+"..."),1)]),_:2},1032,["to"])])):t.column.property=="to"?(a(),r("div",Y,[t.row.method=="0x60806040"?(a(),r("div",Z,[o(f,null,{default:e(()=>[o(s(H))]),_:1}),o(y,{effect:"dark",placement:"top",content:"New Contract"},{default:e(()=>[o(l,{style:{"margin-left":"1.5px"},to:"/address/"+t.row[t.column.property]},{default:e(()=>[tt]),_:2},1032,["to"])]),_:2},1024)])):t.row.toContract?(a(),r("div",et,[o(f,null,{default:e(()=>[o(s(S))]),_:1}),o(y,{effect:"dark",placement:"top",content:"Contract"},{default:e(()=>[o(l,{style:{"margin-left":"1.5px"},to:"/address/"+t.row[t.column.property]},{default:e(()=>[i(n(t.row[t.column.property].slice(0,15)+"..."),1)]),_:2},1032,["to"])]),_:2},1024)])):(a(),u(l,{key:2,to:"/address/"+t.row[t.column.property]},{default:e(()=>[i(n(t.row[t.column.property].slice(0,15)+"..."),1)]),_:2},1032,["to"]))])):t.column.property=="value"?(a(),r("div",at,n(s(h).utils.formatUnits(t.row[t.column.property],t.row.contractDecimals))+" "+n(s(V)()),1)):t.column.property=="tokenID"?(a(),r("div",ot,n(parseInt(t.row[t.column.property])),1)):t.column.property=="gas"?(a(),r("div",rt,n(s(h).utils.formatUnits((parseInt(t.row[t.column.property],10)*parseInt(t.row.gasPrice,10)).toString(),18)),1)):t.column.property=="contract"?(a(),r("div",nt,[o(l,{to:"/token/"+t.row[t.column.property]},{default:e(()=>[i(n(t.row[t.column.property].slice(0,15)+"..."),1)]),_:2},1032,["to"])])):c("",!0)]),_:2},1032,["property","label"]))),128))]),_:1},8,["data"])}}});var bt=z(dt,[["__scopeId","data-v-44318bda"]]);export{bt as _};

import{_ as k,d as z,F as j,aq as C,r as D,o as a,n as m,w as e,b as l,t as r,a as n,f as s,bD as E,c as o,W as f,ah as I,e as T,g as B,V as N,ac as H,i as S}from"./index.d348e7fd.js";import{E as V,a as A}from"./el-table-column.7119721d.js";/* empty css               */import{E as M,_ as q}from"./BaseTransactionInfo.b7cf23aa.js";import{g as h}from"./utils.fc402d12.js";import{e as x}from"./index.eb369523.js";const F={key:0},U={key:0,style:{width:"170px"}},G={key:1,style:{width:"170px"}},L={class:"span"},P={key:2,style:{width:"170px"}},R={key:3},W={key:4},$={key:5},J={key:6,style:{width:"170px"}},K={key:7,style:{width:"170px"}},O={key:8},Q={key:9},X={key:10,style:{"font-size":"11px"}},Y={key:11,style:{width:"170px","font-size":"11px"}},Z=z({__name:"GenerateTransfers",props:{loadStatus:Boolean,txsData:{type:Array,require:!0},headerData:{type:Array,require:!0}},setup(y){const c=y,p=j("loading...");return C(()=>{c.loadStatus||(p.value="empty data")}),(tt,at)=>{const g=H,_=S,w=q,v=M,u=V,d=D("router-link"),b=A;return a(),m(b,{class:"table-border",data:c.txsData,"empty-text":"loading...","row-style":{height:"50px"}},{empty:e(()=>[l(r(p.value),1)]),default:e(()=>[n(u,{width:"37px"},{default:e(i=>[n(v,{placement:"right",title:"Additional Info",width:320,trigger:"click"},{reference:e(()=>[n(_,{style:{width:"5px"},type:"info",size:"small",onClick:t=>i.row.base=!0},{default:e(()=>[n(g,null,{default:e(()=>[n(s(E))]),_:1})]),_:2},1032,["onClick"])]),default:e(()=>[i.row.base==!0?(a(),o("div",F,[n(w,{txHash:i.row.transactionHash?i.row.transactionHash:i.row.hash},null,8,["txHash"])])):f("",!0)]),_:2},1024)]),_:1}),(a(!0),o(N,null,I(c.headerData,i=>(a(),m(u,{key:i.key,property:i.key,label:i.label},{default:e(t=>[t.column.property=="transactionHash"?(a(),o("div",U,[n(d,{to:"/tx/"+t.row[t.column.property]},{default:e(()=>[l(r(t.row[t.column.property].slice(0,15)+"..."),1)]),_:2},1032,["to"])])):f("",!0),t.column.property=="method"?(a(),o("div",G,[T("span",L,r(t.row[t.column.property]),1)])):f("",!0),t.column.property=="hash"?(a(),o("div",P,[n(d,{to:"/tx/"+t.row[t.column.property]},{default:e(()=>[l(r(t.row[t.column.property].slice(0,15)+"..."),1)]),_:2},1032,["to"])])):t.column.property=="blockNumber"?(a(),o("div",R,[n(d,{to:"/block/"+parseInt(t.row[t.column.property])},{default:e(()=>[l(r(parseInt(t.row[t.column.property])),1)]),_:2},1032,["to"])])):t.column.property=="createTime"?(a(),o("div",W,r(s(h)(t.row[t.column.property])),1)):t.column.property=="createdTime"?(a(),o("div",$,r(s(h)(t.row[t.column.property])),1)):t.column.property=="from"?(a(),o("div",J,[n(d,{to:"/address/"+t.row[t.column.property]},{default:e(()=>[l(r(t.row[t.column.property].slice(0,15)+"..."),1)]),_:2},1032,["to"])])):t.column.property=="to"?(a(),o("div",K,[n(d,{to:"/address/"+t.row[t.column.property]},{default:e(()=>[l(r(t.row[t.column.property].slice(0,15)+"..."),1)]),_:2},1032,["to"])])):t.column.property=="value"?(a(),o("div",O,r(s(x).utils.formatUnits(t.row[t.column.property],t.row.contractDecimals))+" "+r(s(B)()),1)):t.column.property=="tokenID"?(a(),o("div",Q,r(parseInt(t.row[t.column.property])),1)):t.column.property=="gas"?(a(),o("div",X,r(s(x).utils.formatUnits((parseInt(t.row[t.column.property],10)*parseInt(t.row.gasPrice,10)).toString(),18)),1)):t.column.property=="contract"?(a(),o("div",Y,[n(d,{to:"/token/"+t.row[t.column.property]},{default:e(()=>[l(r(t.row[t.column.property].slice(0,15)+"..."),1)]),_:2},1032,["to"])])):f("",!0)]),_:2},1032,["property","label"]))),128))]),_:1},8,["data"])}}});var dt=k(Z,[["__scopeId","data-v-faa13710"]]);export{dt as _};

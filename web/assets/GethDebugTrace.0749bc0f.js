import{_ as T,d as z,bB as j,aM as G,F as x,aq as C,r as E,o as f,c as b,e as o,a as t,w as i,f as c,t as u,p as S,h as q,b as m,W as g,O as D,be as N}from"./index.774cbafa.js";import{E as M,a as V}from"./el-tab-pane.92cfc2b2.js";import{E as O,a as B}from"./el-table-column.f61d754b.js";/* empty css               */import{f as I}from"./transactionService.ae79974f.js";import"./index.f445393f.js";import"./objects.9fcb7131.js";const v=n=>(S("data-v-0a802b0c"),n=n(),q(),n),L={style:{"margin-bottom":"-20px"}},R=v(()=>o("h4",{class:"h4-title"},"Geth VM Trace Transaction",-1)),J={style:{color:"#77838f !important","ont-size":"80%","font-weight":"400"}},F=m(" GETH Trace for Txn Hash "),H={class:"tx-sub-title"},P=v(()=>o("p",{class:"subtitle1"},"Raw traces",-1)),U={key:0},W={key:1},A=z({__name:"GethDebugTrace",setup(n){const e=j(),s=G([]),d=x(""),h=x(0);return C(async()=>{if(s.length=0,e.query.txhash){const r=await I(e.query.txhash,e.query.type);console.log("getGethDebugTrace",r.data.res),h.value=r.data.logNum,e.query.type=="tracetx"?JSON.parse(r.data.res).forEach(p=>{s.push(p)}):d.value=JSON.stringify(JSON.parse(r.data.res),null,4)}}),(r,l)=>{const p=E("router-link"),a=O,y=B,_=M,w=V;return f(),b("div",null,[o("div",L,[R,o("p",J,[F,t(p,{to:"/tx/"+c(e).query.txhash},{default:i(()=>[m(u(c(e).query.txhash),1)]),_:1},8,["to"])]),o("h4",H,"(Showing the last "+u(h.value)+" records only)",1)]),t(w,{style:{width:"100%","margin-top":"30px"}},{default:i(()=>[t(_,null,{label:i(()=>[P]),default:i(()=>[c(e).query.type=="tracetx"?(f(),b("div",U,[t(y,{data:s,style:{width:"100%"},"empty-text":"loading..."},{default:i(()=>[t(a,{type:"index",width:"100",label:"Step"}),t(a,{prop:"pc",label:"PC"}),t(a,{prop:"op",label:"Operation"}),t(a,{prop:"gas",label:"Gas"}),t(a,{prop:"gasCost",label:"GasCost"}),t(a,{prop:"depth",label:"Depth"})]),_:1},8,["data"])])):g("",!0),c(e).query.type=="tracetx2"?(f(),b("div",W,[D(o("textarea",{class:"byte-codes-text",style:{margin:"0px"},rows:"12","onUpdate:modelValue":l[0]||(l[0]=k=>d.value=k)}," ",512),[[N,d.value]])])):g("",!0)]),_:1})]),_:1})])}}});var et=T(A,[["__scopeId","data-v-0a802b0c"]]);export{et as default};
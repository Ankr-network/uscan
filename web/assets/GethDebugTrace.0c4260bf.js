import{_ as T,d as z,bB as j,aM as G,F as u,aq as C,r as E,o as p,c as h,e as o,a as t,w as i,f as c,t as g,p as S,h as q,b,W as m,O as D,be as N}from"./index.d5f49ade.js";import{E as M,a as V}from"./el-tab-pane.a2d47aac.js";import{E as O,a as B}from"./el-table-column.7af40814.js";/* empty css               */import{f as I}from"./transactionService.23982c24.js";import"./index.06f05625.js";import"./objects.b5077864.js";const y=n=>(S("data-v-ca21f786"),n=n(),q(),n),L={style:{"margin-bottom":"-20px"}},R=y(()=>o("h4",{class:"h4-title"},"Geth VM Trace Transaction",-1)),J={style:{color:"#77838f !important","ont-size":"80%","font-weight":"400"}},F=b(" GETH Trace for Txn Hash "),H={class:"tx-sub-title"},P=y(()=>o("p",{class:"subtitle1"},"Raw traces",-1)),U={key:0},W={key:1},A=z({__name:"GethDebugTrace",setup(n){const e=j(),s=G([]),d=u(""),x=u(0);return C(async()=>{if(s.length=0,e.query.txhash){const r=await I(e.query.txhash,e.query.type);console.log("getGethDebugTrace",r.data.res),x.value=r.data.logNum,e.query.type=="tracetx"?JSON.parse(r.data.res).forEach(f=>{s.push(f)}):d.value=JSON.stringify(JSON.parse(r.data.res),null,4)}}),(r,l)=>{const f=E("router-link"),a=O,v=B,_=M,w=V;return p(),h("div",null,[o("div",L,[R,o("p",J,[F,t(f,{to:"/tx/"+c(e).query.txhash},{default:i(()=>[b(g(c(e).query.txhash),1)]),_:1},8,["to"])]),o("h4",H,"(Showing the last "+g(x.value)+" records only)",1)]),t(w,{style:{width:"100%","margin-top":"30px"}},{default:i(()=>[t(_,null,{label:i(()=>[P]),default:i(()=>[c(e).query.type=="tracetx"?(p(),h("div",U,[t(v,{data:s,style:{width:"100%"},"empty-text":"loading..."},{default:i(()=>[t(a,{type:"index",width:"100",label:"Step"}),t(a,{prop:"pc",label:"PC"}),t(a,{prop:"op",label:"Operation"}),t(a,{prop:"gas",label:"Gas"}),t(a,{prop:"gasCost",label:"GasCost"}),t(a,{prop:"depth",label:"Depth"})]),_:1},8,["data"])])):m("",!0),c(e).query.type=="tracetx2"?(p(),h("div",W,[D(o("textarea",{class:"byte-codes-text",style:{margin:"0px"},rows:"12","onUpdate:modelValue":l[0]||(l[0]=k=>d.value=k),readonly:"readonly"},`
          `,512),[[N,d.value]])])):m("",!0)]),_:1})]),_:1})])}}});var et=T(A,[["__scopeId","data-v-ca21f786"]]);export{et as default};
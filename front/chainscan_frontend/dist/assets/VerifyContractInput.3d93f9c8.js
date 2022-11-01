import{_ as L,d as U,u as D,bA as F,m as G,k as n,w as J,bs as Y,b1 as H,j as K,o as i,c as u,a as s,f as e,V,b as _,p as Q,i as W,T as w,ag as C,e as j}from"./index.c5813a35.js";/* empty css               */import{E as X,a as Z}from"./el-select.95009e57.js";import{G as $}from"./contractService.a25e6223.js";import"./index.beb32d32.js";import"./isEqual.abf431c5.js";import"./index.19a8146b.js";const o=v=>(Q("data-v-3541035a"),v=v(),W(),v),ee={style:{display:"flex","flex-direction":"column"}},te=o(()=>e("div",{style:{"text-align":"center"}},[e("h3",{class:"content-sub"},"Verify & Publish Contract Source Code"),e("p",{style:{color:"#8c98a4","ont-weight":"700","font-size":"80%","font-weight":"bold"}}," COMPILER TYPE AND VERSION SELECTION ")],-1)),le=o(()=>e("div",{style:{display:"flex","justify-content":"center"}},[e("p",{style:{color:"#77838f",width:"80%"}},' \xA0 \xA0 Source code verification provides transparency for users interacting with smart contracts. By uploading the source code, Etherscan will match the compiled code with that on the blockchain. Just like contracts, a "smart contract" should provide end users with more information on what they are "digitally signing" for and give users an opportunity to audit the code to independently verify that it actually does what it is supposed to do. ')],-1)),oe=o(()=>e("br",null,null,-1)),se={style:{display:"flex","justify-content":"center"}},ae={class:"content-sub1"},ne={style:{width:"700px","font-size":"0.875rem","margin-bottom":"10px"}},ie=o(()=>e("p",null,"Please enter the Contract Address you would like to verify",-1)),ue={key:0,style:{color:"red"}},re=o(()=>e("p",null,"Required",-1)),ce=[re],de={style:{display:"flex","justify-content":"center"}},pe={class:"content-sub1"},_e=o(()=>e("p",null,"Please select Compiler Type",-1)),ve={key:0,style:{color:"red"}},ye=o(()=>e("p",null,"Required",-1)),fe=[ye],he={style:{display:"flex","justify-content":"center"}},me={class:"content-sub1"},be=o(()=>e("p",null,"Please select Compiler Version",-1)),Ve={key:0,style:{color:"red"}},ge=o(()=>e("p",null,"Required",-1)),xe=[ge],we={style:{display:"flex","justify-content":"center"}},Ce={class:"content-sub1"},Se=o(()=>e("p",null,"Please select Open Source License Type",-1)),Ee={key:0,style:{color:"red"}},ke=o(()=>e("p",null,"Required",-1)),Ie=[ke],Re={style:{display:"flex","justify-content":"center"}},Te={style:{display:"flex","justify-content":"center","margin-top":"10px"}},qe=j("Continue"),ze=j("Reset"),Oe=U({__name:"VerifyContractInput",async setup(v){let y,S;const N=D(),E=F();document.title="Verify & Publish Contract Source Code | The "+G+" Explorer";const r=n(""),f=n(!1),d=n(""),h=n(!1),k=[],c=n(""),m=n(!1),I=[],p=n(""),b=n(!1),R=[];E.query.a&&(r.value=E.query.a);const T=([y,S]=J(()=>$()),y=await y,S(),y);k.push({value:"solidity-single-file",label:"Solidity (Single file)"},{value:"solidity-standard-json-input",label:"Solidity (Standard-Json-Input)"});const q=new Map,z=new Map;T.data.compilerVersions.forEach(l=>{q.set(l.id,l.name),z.set(l.id,l.fileName),I.push({value:l.id,label:l.name})}),T.data.licenseTypes.forEach(l=>{R.push({value:l.id,label:l.name})});const P=()=>{r.value="",f.value=!1,d.value="",h.value=!1,c.value="",m.value=!1,p.value="",b.value=!1},A=()=>{if(r.value===""?f.value=!0:f.value=!1,d.value===""?h.value=!0:h.value=!1,c.value===""?m.value=!0:m.value=!1,p.value===""?b.value=!0:b.value=!1,r.value!==""&&d.value!==""&&c.value!==""&&p.value!==""){const l="/verifyContract/submit?a="+r.value+"&ct="+d.value+"&cv="+q.get(c.value)+"&cf="+z.get(c.value)+"&lictype="+p.value;N.push(l)}};return(l,a)=>{const M=Y,B=H,g=X,x=Z,O=K;return i(),u("div",ee,[te,s(M),le,oe,e("div",se,[e("div",ae,[e("div",ne,[ie,s(B,{modelValue:r.value,"onUpdate:modelValue":a[0]||(a[0]=t=>r.value=t),size:"large",placeholder:"0x...",clearable:""},null,8,["modelValue"]),f.value?(i(),u("div",ue,ce)):V("",!0)])])]),e("div",de,[e("div",pe,[_e,s(x,{modelValue:d.value,"onUpdate:modelValue":a[1]||(a[1]=t=>d.value=t),placeholder:"Select",size:"large",style:{width:"100%"}},{default:_(()=>[(i(),u(w,null,C(k,t=>s(g,{key:t.value,label:t.label,value:t.value},null,8,["label","value"])),64))]),_:1},8,["modelValue"]),h.value?(i(),u("div",ve,fe)):V("",!0)])]),e("div",he,[e("div",me,[be,s(x,{modelValue:c.value,"onUpdate:modelValue":a[2]||(a[2]=t=>c.value=t),placeholder:"Select",size:"large",style:{width:"100%"}},{default:_(()=>[(i(),u(w,null,C(I,t=>s(g,{key:t.value.name,label:t.label,value:t.value},null,8,["label","value"])),64))]),_:1},8,["modelValue"]),m.value?(i(),u("div",Ve,xe)):V("",!0)])]),e("div",we,[e("div",Ce,[Se,s(x,{modelValue:p.value,"onUpdate:modelValue":a[3]||(a[3]=t=>p.value=t),placeholder:"Select",size:"large",style:{width:"100%"}},{default:_(()=>[(i(),u(w,null,C(R,t=>s(g,{key:t.value,label:t.label,value:t.value},null,8,["label","value"])),64))]),_:1},8,["modelValue"]),b.value?(i(),u("div",Ee,Ie)):V("",!0)])]),e("div",Re,[e("div",Te,[s(O,{type:"primary",size:"large",onClick:A},{default:_(()=>[qe]),_:1}),s(O,{type:"info",size:"large",onClick:P},{default:_(()=>[ze]),_:1})])])])}}});var Ue=L(Oe,[["__scopeId","data-v-3541035a"]]);export{Ue as default};

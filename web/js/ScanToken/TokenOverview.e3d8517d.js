import{l as e,m as a,w as s,t,r,s as l}from"../element-plus/element-plus.e38f20bb.js";import{_ as o}from"../GenerateTransactions/GenerateTransactions.e422b84e.js";import{d as n}from"../tokenService/tokenService.66523378.js";import{e as i}from"../transactionService/transactionService.c5e8bb99.js";import{a as p,T as d,E as c}from"../transaction/transaction.08dab424.js";import{_ as u}from"../../assets/index.ce741f41.js";import{C as m,e as j,X as v,aq as f,d as h,o as b,c as y,a as g,S as _,Q as x,am as S,an as k,a4 as w,U as D,u as z}from"../@vue/@vue.f5a41070.js";import"../@vueuse/@vueuse.daa6c377.js";import"../@element-plus/@element-plus.f4b8dd72.js";import"../lodash-es/lodash-es.e287ec0c.js";import"../dayjs/dayjs.b9cd4081.js";import"../aes-js/aes-js.d869e6a6.js";import"../@ctrl/@ctrl.17547d95.js";import"../@popperjs/@popperjs.f1fb8f77.js";import"../escape-html/escape-html.d572c0fd.js";import"../normalize-wheel-es/normalize-wheel-es.db30398b.js";import"../ethers/ethers.273881c9.js";import"../@ethersproject/@ethersproject.b35f154f.js";import"../scrypt-js/scrypt-js.a8714fed.js";import"../bech32/bech32.8f2f2e26.js";import"../bn.js/bn.js.b1f68973.js";import"../hash.js/hash.js.0b6b30a5.js";import"../minimalistic-assert/minimalistic-assert.9b24acbb.js";import"../inherits/inherits.ea2611ce.js";import"../js-sha3/js-sha3.ea877be6.js";import"../index/index.9f3e1a29.js";import"../vue-router/vue-router.5c451dd1.js";import"../axios/axios.21f17a99.js";const T=e=>(S("data-v-f3171e3c"),e=e(),k(),e),I=T((()=>g("div",{class:"center-row"},[g("h2",null,"Token")],-1))),C=T((()=>g("div",{class:"card-header"},[g("span",null,"Overview")],-1))),V={class:"card-content"},E=w("Contract:"),G=w("Token ID:"),O={style:{"margin-top":"1%",display:"flex","justify-content":"center"}};var P=u(m({__name:"TokenOverview",props:{address:String,tokenID:String,type:String},async setup(u){let m,S;const k=u,T=j("txs"),P=v([]),U=v([]),q=j(1),Q=j(25),X=j(0),A=j(!0),B=([m,S]=f((()=>n(k.address,k.tokenID,k.type))),m=await m,S(),m),F=async()=>{if("all"===k.type||"erc20"===k.type||"erc721"===k.type||"erc1155"===k.type){"all"===k.type?U.push(...d):"erc20"===k.type?U.push(...c):("erc721"===k.type||"erc1155"===k.type)&&U.push(...p);const e=await i(q.value-1,Q.value,k.type,k.address);e.data.items.forEach((e=>{P.push(e)})),X.value=e.data.total,0==e.data.total&&(A.value=!1)}};[m,S]=f((()=>F())),await m,S(),h(k,(async()=>{void 0!==k.address&&void 0!==k.tokenID&&void 0!==k.type&&(P.length=0,U.length=0,q.value=1,Q.value=25,F())}));const H=async e=>{P.length=0,q.value=1,Q.value=e,F()},J=async e=>{P.length=0,q.value=e,F()};return(n,i)=>{const d=e,c=a,u=s,m=o,j=t,v=r,f=l;return b(),y("div",null,[I,g("div",null,[_(u,{class:"box-card"},{header:x((()=>[C])),default:x((()=>[g("div",V,[_(c,null,{default:x((()=>[_(d,{span:10},{default:x((()=>[E])),_:1}),_(d,{span:14},{default:x((()=>[w(D(z(B).data.contract),1)])),_:1})])),_:1}),_(c,null,{default:x((()=>[_(d,{span:10},{default:x((()=>[G])),_:1}),_(d,{span:14},{default:x((()=>[w(D(z(B).data.tokenID),1)])),_:1})])),_:1})])])),_:1})]),g("div",null,[_(f,{modelValue:T.value,"onUpdate:modelValue":i[0]||(i[0]=e=>T.value=e)},{default:x((()=>[_(v,{label:"Transactions",name:"txs"},{default:x((()=>[_(m,{txsData:P,headerData:z(p),loadStatus:A.value},null,8,["txsData","headerData","loadStatus"]),g("div",O,[_(j,{small:"",background:"",currentPage:q.value,"page-size":Q.value,"page-sizes":[10,25,50,100],layout:"total, sizes, prev, pager, next, jumper",total:X.value,onSizeChange:H,onCurrentChange:J},null,8,["currentPage","page-size","total"])])])),_:1})])),_:1},8,["modelValue"])])])}}}),[["__scopeId","data-v-f3171e3c"]]);export{P as default};

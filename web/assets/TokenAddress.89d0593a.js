import{_ as te,d as O,F as o,aM as N,l as be,y as we,r as Q,o as d,c as T,e as n,V as L,ah as q,a as t,w as e,n as I,b as s,t as m,f as P,p as ae,h as oe,aq as ke,W as U,m as Te,j as Ce,J as ze}from"./index.5c39bb93.js";import{E as Se,a as Ee}from"./el-tab-pane.0a7a86a6.js";import{E as ne}from"./el-pagination.4913e9bf.js";/* empty css               */import"./el-select.4a434427.js";import{E as se,a as re}from"./el-col.7dc061b9.js";import{E as le}from"./el-card.bcf7cb3e.js";import{G as B,a as Ae,b as R,c as F}from"./tokenService.26c2b5cb.js";import{E as Ie,a as je}from"./el-table-column.9ea7ff5a.js";import{e as Z}from"./index.eb369523.js";import{t as de}from"./utils.fc402d12.js";import{b as De,_ as He,a as $e}from"./ContractInfo.7f350045.js";import{_ as Ne}from"./GenerateTransfers.cfbfb829.js";import{_ as Pe}from"./CopyIcon.c912ebe0.js";import{c as Ge,d as ee,e as Me}from"./transaction.8e782377.js";import{G as Ve}from"./contractService.0787d90a.js";import"./isEqual.e6dd98ea.js";import"./index.66c13364.js";import"./index.386ebf73.js";import"./objects.6de92f3c.js";import"./_commonjsHelpers.b8add541.js";import"./el-collapse-item.867597a8.js";import"./BaseTransactionInfo.98ad6198.js";import"./transactionService.06cd2fa0.js";import"./index.7d4bc839.js";const Be=u=>(ae("data-v-712b3303"),u=u(),oe(),u),Re=Be(()=>n("div",{style:{height:"100px",width:"100px","background-color":"#598df6","border-radius":"0.35rem"}},null,-1)),Fe={class:"text-secondary"},Le=s(" TokenID: "),qe={class:"text-secondary"},Ue=s(" Owner: "),Oe={style:{"margin-top":"1%",display:"flex","justify-content":"center"}},Qe=O({__name:"GenerateInventory",props:{address:{type:String,require:!0},ercType:{type:String,require:!0}},async setup(u){let a,E;const r=u,f=o(1),_=o(25),p=N([]),z=o(1),x=([a,E]=be(()=>B(r.address,r.ercType,f.value-1,_.value)),a=await a,E(),a);z.value=x.data.total,x.data.items.forEach(c=>{p.push(c)});const w=we({get(){const c=[];return p.forEach((S,g)=>{const A=Math.floor(g/6);c[A]||(c[A]=[]),c[A].push(S)}),c},set(){}}),k=async c=>{p.length=0,f.value=1,_.value=c,(await B(r.address,r.ercType,f.value-1,_.value)).data.items.forEach(g=>{p.push(g)})},i=async c=>{p.length=0,f.value=c,(await B(r.address,r.ercType,f.value-1,_.value)).data.items.forEach(g=>{p.push(g)})};return(c,S)=>{const g=Q("router-link"),A=le,G=se,D=re,H=ne;return d(),T("div",null,[n("div",null,[(d(!0),T(L,null,q(P(w),(v,V)=>(d(),T("div",{key:V,style:{"margin-top":"20px"}},[t(D,{gutter:20},{default:e(()=>[(d(!0),T(L,null,q(v,(j,M)=>(d(),I(G,{span:4,key:M},{default:e(()=>[t(A,{"body-style":{padding:"10px"}},{default:e(()=>[Re,n("div",Fe,[Le,n("span",null,[t(g,{to:"/token/nfts/"+r.address+"/"+j.TokenID+"/"+r.ercType},{default:e(()=>[s(m(parseInt(j.TokenID)),1)]),_:2},1032,["to"])])]),n("div",qe,[Ue,n("span",null,[t(g,{to:"/address/"+j.Address},{default:e(()=>[s(m(j.Address.slice(0,18)+"..."),1)]),_:2},1032,["to"])])])]),_:2},1024)]),_:2},1024))),128))]),_:2},1024)]))),128)),n("div",Oe,[t(H,{small:"",background:"",currentPage:f.value,"page-size":_.value,"page-sizes":[10,25,50,100],layout:"total, sizes, prev, pager, next, jumper",total:z.value,onSizeChange:k,onCurrentChange:i},null,8,["currentPage","page-size","total"])])])])}}});var We=te(Qe,[["__scopeId","data-v-712b3303"]]);const Je={key:0},Ke={key:1},Xe={key:2},Ye={key:3},Ze=O({__name:"GenerateHolders",props:{loadStatus:Boolean,decimals:Number,holdersData:{type:Array,require:!0},headerData:{type:Array,require:!0},totalSupply:Number,pageSize:Number,pageIndex:Number},setup(u){const a=u,E=o("loading..."),r=o(0),f=_=>{let p="20%";switch(_){case"Address":p="40%"}return p};return ke(()=>{console.log("totalSupply",a.totalSupply*Math.pow(10,-a.decimals)),r.value=a.totalSupply*Math.pow(10,-a.decimals),a.loadStatus||(E.value="empty data")}),(_,p)=>{const z=Q("router-link"),x=Ie,w=je;return d(),T("div",null,[t(w,{class:"table-border",data:a.holdersData,"empty-text":"loading...","row-style":{height:"50px"}},{empty:e(()=>[s(m(E.value),1)]),default:e(()=>[(d(!0),T(L,null,q(a.headerData,k=>(d(),I(x,{key:k.key,property:k.key,label:k.label,"min-width":f(k.label)},{default:e(i=>[i.column.property=="rank"?(d(),T("div",Je,m((a.pageIndex-1)*a.pageSize+(i.$index+1)),1)):i.column.property=="Address"?(d(),T("div",Ke,[t(z,{to:"/address/"+i.row[i.column.property]},{default:e(()=>[s(m(i.row[i.column.property]),1)]),_:2},1032,["to"])])):i.column.property=="Quantity"?(d(),T("div",Xe,m(P(de)(P(Z).utils.formatUnits(i.row[i.column.property],a.decimals))),1)):i.column.property=="percentage"?(d(),T("div",Ye,m((parseFloat(P(Z).utils.formatUnits(i.row.Quantity,a.decimals))/parseFloat(r.value)*100).toFixed(4)+"%"),1)):U("",!0)]),_:2},1032,["property","label","min-width"]))),128))]),_:1},8,["data"])])}}});const et=u=>(ae("data-v-31e45160"),u=u(),oe(),u),tt={class:"center-row"},at={class:"h4-title"},ot=s(" Token "),nt={class:"small text-secondary"},st=s(" \xA0 "),rt={class:"card-header"},lt=s(" Overview\xA0\xA0\xA0\xA0"),dt={style:{"font-size":"0.96rem",color:"#9ba2aa","font-weight":"bold"}},it={class:"card-content"},ct=s("Max Total Supply:"),ut=s("Holders:"),pt=s("Transfers:"),mt=et(()=>n("div",{class:"card-header"},[n("span",null,"Profile Summary")],-1)),ft={class:"card-content"},_t=s("Contract:"),vt=s("Decimals:"),ht=s("Token Tracker:"),gt={style:{"margin-top":"1%",display:"flex","justify-content":"center"}},yt={style:{"margin-top":"1%",display:"flex","justify-content":"center"}},xt=O({__name:"TokenAddress",props:{address:String},setup(u){const a=u;document.title="Token | The "+Te()+" Explorer";const E=o("transactions"),r=o(1),f=o(25),_=o(1),p=o(25),z=N([]),x=N([]),w=N([]),k=N([]),i=o(!0),c=o(),S=o(0),g=o(0),A=o(""),G=o(""),D=o(0),H=o(0),v=o(""),V=o(!0),j=o(!1),M=o({}),W=o(""),J=o();Ce(async()=>{K()});const K=async()=>{const b=await Ae(a.address);for(const h in b.data)b.data[h]!=0&&(v.value=h);console.log("tokenType",b.data),v.value=="erc20"?(z.push(...Ge),k.push(...ee)):(z.push(...Me),k.push(...ee)),c.value=await De(a.address),S.value=c.value.data.tokenTotalSupply,g.value=c.value.data.decimals,A.value=c.value.data.symbol,G.value=c.value.data.name,J.value=c.value.data.code;const C=await R(a.address,v.value,r.value-1,f.value);H.value=C.data.total,C.data.total!==0?C.data.items.forEach(h=>w.push(h)):V.value=!1;const y=await F(a.address,v.value,_.value-1,p.value);D.value=y.data.total,y.data.total!==0?y.data.items.forEach(h=>{x.push(h)}):i.value=!1;const l=await Ve(a.address);M.value=l.data.contract,W.value=l.data.proxyContractAddress,l.data.contract&&(j.value=!0)};ze(a,async()=>{console.log("watch props"),x.length=0,_.value=1,p.value=25,D.value=0,H.value=0,z.length=0,w.length=0,r.value=1,f.value=25,k.length=0,K()});const ie=async b=>{x.length=0,_.value=1,p.value=b,(await F(a.address,v.value,_.value-1,p.value)).data.items.forEach(y=>{x.push(y)})},ce=async b=>{x.length=0,_.value=b,(await F(a.address,v.value,_.value-1,p.value)).data.items.forEach(y=>{x.push(y)})},ue=async b=>{w.length=0,r.value=1,f.value=b,(await R(a.address,v.value,r.value-1,f.value)).data.items.forEach(y=>{w.push(y)})},pe=async b=>{w.length=0,r.value=b,(await R(a.address,v.value,r.value-1,f.value)).data.items.forEach(y=>{w.push(y)})};return(b,C)=>{const y=Pe,l=se,h=re,X=le,me=Q("router-link"),fe=Ne,Y=ne,$=Se,_e=He,ve=$e,he=Ze,ge=We,ye=Ee;return d(),T("div",null,[n("div",tt,[n("h4",at,[ot,n("span",nt,"\xA0\xA0"+m(a.address),1)]),st,t(y,{text:a.address},null,8,["text"])]),n("div",null,[t(h,{gutter:20},{default:e(()=>[t(l,{span:12},{default:e(()=>[n("div",null,[t(X,{class:"box-card-address"},{header:e(()=>[n("div",rt,[n("span",null,[lt,n("span",dt,m(v.value),1)])])]),default:e(()=>[n("div",it,[t(h,null,{default:e(()=>[t(l,{span:9},{default:e(()=>[ct]),_:1}),v.value=="erc20"?(d(),I(l,{key:0,span:15},{default:e(()=>[s(m(P(de)(parseInt(S.value)*Math.pow(10,-g.value))),1)]),_:1})):(d(),I(l,{key:1,span:15},{default:e(()=>[s(m(parseInt(S.value)),1)]),_:1}))]),_:1}),t(h,null,{default:e(()=>[t(l,{span:9},{default:e(()=>[ut]),_:1}),t(l,{span:15},{default:e(()=>[s(m(H.value),1)]),_:1})]),_:1}),t(h,null,{default:e(()=>[t(l,{span:9},{default:e(()=>[pt]),_:1}),t(l,{span:15},{default:e(()=>[s(m(D.value),1)]),_:1})]),_:1})])]),_:1})])]),_:1}),t(l,{span:12},{default:e(()=>[n("div",null,[t(X,{class:"box-card-address"},{header:e(()=>[mt]),default:e(()=>[n("div",ft,[t(h,null,{default:e(()=>[t(l,{span:9},{default:e(()=>[_t]),_:1}),t(l,{span:15},{default:e(()=>[t(me,{to:"/address/"+u.address},{default:e(()=>[s(m(u.address),1)]),_:1},8,["to"])]),_:1})]),_:1}),v.value=="erc20"?(d(),I(h,{key:0},{default:e(()=>[t(l,{span:9},{default:e(()=>[vt]),_:1}),t(l,{span:15},{default:e(()=>[s(m(g.value),1)]),_:1})]),_:1})):U("",!0),t(h,null,{default:e(()=>[t(l,{span:9},{default:e(()=>[ht]),_:1}),t(l,{span:15},{default:e(()=>[s(m(G.value)+" ("+m(A.value)+")",1)]),_:1})]),_:1})])]),_:1})])]),_:1})]),_:1})]),n("div",null,[t(ye,{modelValue:E.value,"onUpdate:modelValue":C[0]||(C[0]=xe=>E.value=xe)},{default:e(()=>[t($,{label:"Transfers",name:"transactions"},{default:e(()=>[t(fe,{txsData:x,headerData:z,loadStatus:i.value},null,8,["txsData","headerData","loadStatus"]),n("div",gt,[t(Y,{small:"",background:"",currentPage:_.value,"page-size":p.value,"page-sizes":[10,25,50,100],layout:"total, sizes, prev, pager, next, jumper",total:D.value,onSizeChange:ie,onCurrentChange:ce},null,8,["currentPage","page-size","total"])])]),_:1}),j.value?(d(),I($,{key:1,label:"Contract(verified)",name:"contract-verified"},{default:e(()=>[t(ve,{contractAddress:u.address,contractInfo:M.value,proxyContractAddress:W.value},null,8,["contractAddress","contractInfo","proxyContractAddress"])]),_:1})):(d(),I($,{key:0,label:"Contract",name:"contract"},{default:e(()=>[t(_e,{contractAddress:u.address,codeContent:J.value},null,8,["contractAddress","codeContent"])]),_:1})),t($,{label:"Holders",name:"holders"},{default:e(()=>[t(he,{holdersData:w,headerData:k,loadStatus:i.value,decimals:g.value,totalSupply:parseInt(S.value),pageSize:f.value,pageIndex:r.value},null,8,["holdersData","headerData","loadStatus","decimals","totalSupply","pageSize","pageIndex"]),n("div",yt,[t(Y,{small:"",background:"",currentPage:r.value,"page-size":f.value,"page-sizes":[10,25,50,100],layout:"total, sizes, prev, pager, next, jumper",total:H.value,onSizeChange:ue,onCurrentChange:pe},null,8,["currentPage","page-size","total"])])]),_:1}),v.value==="erc721"||v.value==="erc1155"?(d(),I($,{key:2,label:"Inventory",name:"inventory"},{default:e(()=>[t(ge,{address:a.address,ercType:v.value},null,8,["address","ercType"])]),_:1})):U("",!0)]),_:1},8,["modelValue"])])])}}});var Ot=te(xt,[["__scopeId","data-v-31e45160"]]);export{Ot as default};

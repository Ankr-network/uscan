import{_ as C,d as D,r as R,o,c as r,e as l,V as I,ah as T,a as t,w as a,t as i,b as n,W as z,p as F,h as L,bs as te,F as N,aM as M,j as ae,bt as oe,n as S,f as h,bu as re,bv as ne,bw as ie,b1 as de,ac as G,i as O,bx as B,by as H,bz as U,m as le,aN as se,aQ as ce,g as pe,R as fe,bA as ue,bB as me,bC as xe,E as he,bD as ge,u as _e,J as ve,b3 as we}from"./index.498a182e.js";import{E as be,a as ye}from"./el-tab-pane.c149233f.js";import{E as q,a as J}from"./el-col.5f0c2281.js";import{E as Q,a as W}from"./el-table-column.5d99e046.js";/* empty css               */import{A as ke,e as A}from"./index.eb369523.js";import{_ as ze}from"./CopyIcon.0bb36f0a.js";import{G as Ve}from"./transactionService.ee6953b0.js";import{g as je}from"./transaction.60b0ea3a.js";import"./index.b9ac6d09.js";import"./objects.7354dc3a.js";import"./_commonjsHelpers.b8add541.js";import"./index.7d4bc839.js";const K=p=>(F("data-v-16369e1c"),p=p(),L(),p),Ae={class:"log"},Ne={class:"content"},Ce=K(()=>l("p",null,"Transaction Receipt Event Logs",-1)),De={class:"log-content"},$e={class:"icon-circle"},Ee=K(()=>l("div",{style:{"font-size":"15px","font-weight":"bold",color:"#4a4f55"}},"Address",-1)),Ie={key:0},Te={class:"center-row"},Me={class:"topic-index"},Se={style:{"margin-left":"10px"}},Fe=n("Data"),Le={class:"center-row",style:{"background-color":"#f8f9fa",width:"90%"}},Ge={style:{margin:"20px"}},Pe=D({__name:"TransactionLogs",props:{transactionLogs:{type:Array,require:!0}},setup(p){const u=p;return(w,y)=>{const c=q,b=R("router-link"),g=J,m=te;return o(),r("div",Ae,[l("div",Ne,[Ce,l("div",De,[(o(!0),r(I,null,T(u.transactionLogs,(x,f)=>(o(),r("div",{key:f},[l("div",null,[t(g,null,{default:a(()=>[t(c,{span:2},{default:a(()=>[l("div",$e,i(x.logIndex),1)]),_:2},1024),t(c,{span:22},{default:a(()=>[l("div",null,[t(g,{class:"log-row"},{default:a(()=>[t(c,{span:2},{default:a(()=>[Ee]),_:1}),t(c,{span:22},{default:a(()=>[l("div",null,[t(b,{to:"/address/"+x.address},{default:a(()=>[n(i(x.address),1)]),_:2},1032,["to"])])]),_:2},1024)]),_:2},1024),l("div",null,[(o(!0),r(I,null,T(x.topics,(d,e)=>(o(),r("div",{key:e},[t(g,{class:"log-row"},{default:a(()=>[t(c,{span:2},{default:a(()=>[e==0?(o(),r("div",Ie,"Topics")):z("",!0)]),_:2},1024),t(c,{span:22},{default:a(()=>[l("div",Te,[l("div",Me,i(e),1),l("div",Se,i(d),1)])]),_:2},1024)]),_:2},1024)]))),128))]),t(g,{class:"log-data-row"},{default:a(()=>[t(c,{span:2},{default:a(()=>[Fe]),_:1}),t(c,{span:22,style:{"word-break":"break-all"}},{default:a(()=>[l("div",Le,[l("div",Ge,i(x.data),1)])]),_:2},1024)]),_:2},1024)])]),_:2},1024)]),_:2},1024),l("div",null,[t(m)])])]))),128))])])])}}});var Re=C(Pe,[["__scopeId","data-v-16369e1c"]]);const Oe={key:0},Be={style:{"margin-top":"10px"}},He=n(" View Input As \xA0 "),Ue=n(" Default View "),qe=n("Original"),Je=n(" Decode Input Data \xA0"),Qe={key:1,style:{"margin-top":"10px","margin-bottom":"10px"}},We=n("\xA0 Switch Back "),Ke=D({__name:"InputDataInfo",props:{inputData:String,methodName:String},setup(p){const u=p,w=N(""),y=N(!0),c=M([]),b=(f,d)=>{c.length=0,new ke().decode(f,d).forEach((_,v)=>{c.push({index:v,type:f[v],data:_})})},g=f=>{y.value=f},m=()=>{const f=u.inputData.slice(10),d="Function: "+u.methodName,e="MethodID: "+u.inputData.slice(0,10);let s="",_=0;for(let v=0,V=f.length;v<V;v+=64)s+="["+_+"] :    "+f.slice(v,v+64)+`
`,_+=1;w.value=d+`

`+e+`
`+s},x=()=>{w.value=u.inputData};return ae(()=>{if(u.methodName!==""){const f=oe(u.methodName);f.length!==0&&b(f,"0x"+u.inputData.slice(10)),m()}else x()}),(f,d)=>{const e=de,s=G,_=O,v=B,V=H,$=U,j=Q,E=W;return o(),r("div",null,[y.value?(o(),r("div",Oe,[l("div",null,[t(e,{style:{"font-family":"Monaco"},modelValue:w.value,"onUpdate:modelValue":d[0]||(d[0]=k=>w.value=k),rows:"8",placeholder:"Please input","show-word-limit":"",type:"textarea",readonly:!0},null,8,["modelValue"])]),l("div",Be,[t($,null,{dropdown:a(()=>[t(V,null,{default:a(()=>[u.methodName!==""?(o(),S(v,{key:0,onClick:d[1]||(d[1]=k=>m())},{default:a(()=>[Ue]),_:1})):z("",!0),t(v,{onClick:d[2]||(d[2]=k=>x())},{default:a(()=>[qe]),_:1})]),_:1})]),default:a(()=>[t(_,{type:"info"},{default:a(()=>[He,t(s,null,{default:a(()=>[t(h(re))]),_:1})]),_:1})]),_:1}),u.methodName!==""?(o(),S(_,{key:0,type:"info",style:{"margin-left":"10px"},onClick:d[3]||(d[3]=k=>g(!1))},{default:a(()=>[Je,t(s,null,{default:a(()=>[t(h(ne))]),_:1})]),_:1})):z("",!0)])])):(o(),r("div",Qe,[t(E,{data:c,border:"",style:{width:"85%","font-size":"0.4rem","border-radius":"0.4rem"}},{default:a(()=>[t(j,{prop:"index",label:"#"}),t(j,{prop:"type",label:"Type"}),t(j,{prop:"data",label:"Data",width:"500"})]),_:1},8,["data"]),t(_,{style:{"margin-top":"10px"},type:"info",onClick:d[4]||(d[4]=k=>g(!0))},{default:a(()=>[t(s,null,{default:a(()=>[t(h(ie))]),_:1}),We]),_:1})]))])}}});var Xe=C(Ke,[["__scopeId","data-v-37117e86"]]);const X=p=>(F("data-v-65919186"),p=p(),L(),p),Ye={class:"center-row"},Ze={style:{"max-width":"250px"}},et={key:0,style:{"font-weight":"900"}},tt={key:1,class:"center-row"},at={key:2,class:"center-row"},ot={key:3,class:"center-row"},rt={key:0,class:"center-row"},nt=n(" Contract \xA0 "),it=n(" \xA0 "),dt={key:1,class:"center-row"},lt=n(" \xA0 "),st={key:4},ct={key:0,class:"center-row"},pt=n(" Contract \xA0 "),ft=n(" Created \xA0 "),ut={key:1,class:"center-row"},mt=n(" Contract \xA0 "),xt=n(" \xA0 "),ht={key:2,class:"center-row"},gt=n(" \xA0 "),_t={key:5},vt={key:6},wt={class:"center-row"},bt=X(()=>l("div",{style:{"font-weight":"bold"}},"From",-1)),yt=n(" \xA0\xA0\xA0 "),kt=n(" \xA0\xA0\xA0 "),zt=X(()=>l("div",{style:{"font-weight":"bold"}},"To",-1)),Vt=n(" \xA0\xA0\xA0 "),jt=n(" \xA0\xA0\xA0 "),At={key:0},Nt={key:1},Ct={key:7},Dt={key:8},$t={key:9},Et={key:10},It={key:11},Tt={key:0},Mt={class:"success-status"},St=n(" \xA0 Success "),Ft={key:1},Lt={class:"fail-status"},Gt=n(" \xA0 Fail "),Pt={key:2},Rt={class:"pending-status"},Ot=n(" \xA0 Pending "),Bt={key:12},Ht=D({__name:"TransactionOverview",props:{txOverviews:{type:Array,require:!0}},setup(p){const u=p;return document.title="Transaction overview | The "+le+" Explorer",(w,y)=>{const c=G,b=he,g=Q,m=R("router-link"),x=ze,f=Xe,d=W;return o(),r("div",null,[t(d,{class:"table-border",data:u.txOverviews,"empty-text":"loading...","row-style":{height:"50px"}},{default:a(()=>[t(g,{width:"240"},{default:a(e=>[l("div",Ye,[t(b,{effect:"dark",placement:"top"},{content:a(()=>[l("div",Ze,i(e.row.parameterExplain),1)]),default:a(()=>[t(c,null,{default:a(()=>[t(h(se))]),_:1})]),_:2},1024),n(" \xA0"+i(e.row.parameterDisplay),1)])]),_:1}),t(g,{prop:"parameterValue"},{default:a(e=>[e.row.parameterName=="blockNumber"?(o(),r("div",et,[t(m,{to:"/block/"+parseInt(e.row.parameterValue)},{default:a(()=>[n(i(parseInt(e.row.parameterValue)),1)]),_:2},1032,["to"])])):e.row.parameterName=="createTime"?(o(),r("div",tt,[t(c,null,{default:a(()=>[t(h(ce))]),_:1}),n("\xA0"+i(h(pe)(e.row.parameterValue))+"\xA0 ("+i(new Date(e.row.parameterValue*1e3).toUTCString())+") ",1)])):e.row.parameterName=="hash"?(o(),r("div",at,[n(i(e.row.parameterValue)+" \xA0 ",1),t(x,{text:e.row.parameterValue},null,8,["text"])])):e.row.parameterName=="from"?(o(),r("div",ot,[e.row.parameterValue.fromContract?(o(),r("div",rt,[nt,t(m,{to:"/address/"+e.row.parameterValue.from},{default:a(()=>[n(i(e.row.parameterValue.from)+" \xA0 "+i(e.row.parameterValue.fromName),1)]),_:2},1032,["to"]),it,t(x,{text:e.row.parameterValue.from},null,8,["text"])])):(o(),r("div",dt,[t(m,{to:"/address/"+e.row.parameterValue.from},{default:a(()=>[n(i(e.row.parameterValue.from),1)]),_:2},1032,["to"]),lt,t(x,{text:e.row.parameterValue.from},null,8,["text"])]))])):e.row.parameterName=="to"?(o(),r("div",st,[e.row.parameterValue.to==""?(o(),r("div",ct,[pt,t(m,{to:"/address/"+e.row.parameterValue.contractAddress},{default:a(()=>[n(i(e.row.parameterValue.contractAddress),1)]),_:2},1032,["to"]),ft,t(x,{text:e.row.parameterValue.contractAddress},null,8,["text"])])):e.row.parameterValue.toContract?(o(),r("div",ut,[mt,t(m,{to:"/address/"+e.row.parameterValue.to},{default:a(()=>[n(i(e.row.parameterValue.to)+" \xA0 "+i(e.row.parameterValue.toName),1)]),_:2},1032,["to"]),xt,t(x,{text:e.row.parameterValue.to},null,8,["text"])])):(o(),r("div",ht,[t(m,{to:"/address/"+e.row.parameterValue.to},{default:a(()=>[n(i(e.row.parameterValue.to),1)]),_:2},1032,["to"]),gt,t(x,{text:e.row.parameterValue.to},null,8,["text"])]))])):e.row.parameterName=="value"?(o(),r("div",_t,i(h(A).utils.formatEther(e.row.parameterValue))+" Eth ",1)):e.row.parameterName=="tokensTransferred"?(o(),r("div",vt,[l("div",{class:fe(e.row.parameterValue.length>=3?"rolling":"")},[(o(!0),r(I,null,T(e.row.parameterValue,(s,_)=>(o(),r("div",{key:_},[l("div",wt,[bt,yt,t(m,{to:"/address/"+s.from},{default:a(()=>[n(i(s.from.slice(0,18)+"..."),1)]),_:2},1032,["to"]),kt,zt,Vt,t(m,{to:"/address/"+s.to},{default:a(()=>[n(i(s.to.slice(0,18)+"..."),1)]),_:2},1032,["to"]),jt,t(m,{to:"/address/"+s.address},{default:a(()=>[s.addressName?(o(),r("div",At,i(s.addressName),1)):(o(),r("div",Nt,i(s.address.slice(0,18)+"..."),1))]),_:2},1032,["to"])])]))),128))],2)])):e.row.parameterName=="gas"?(o(),r("div",Ct,i(e.row.parameterValue.gasUsed)+" Gwei | "+i(e.row.parameterValue.gasLimit)+" Gwei("+i(e.row.parameterValue.percent)+") ",1)):e.row.parameterName=="gasPrice"?(o(),r("div",Dt,i(h(A).utils.formatEther(e.row.parameterValue))+" Eth ",1)):e.row.parameterName=="maxPriorityFeePerGas"?(o(),r("div",$t," Base: "+i(h(A).utils.formatEther(e.row.parameterValue.baseFeePerGas,"gwei"))+" Gwei | Max: "+i(h(A).utils.formatEther(e.row.parameterValue.maxFeePerGas,"gwei"))+" Gwei | MaxPriority: "+i(h(A).utils.formatEther(e.row.parameterValue.maxPriorityFeePerGas,"gwei"))+" Gwei ",1)):e.row.parameterName=="input"?(o(),r("div",Et,[t(f,{inputData:e.row.parameterValue.inputContent,methodName:e.row.parameterValue.methodName},null,8,["inputData","methodName"])])):e.row.parameterName=="status"?(o(),r("div",It,[e.row.parameterValue.status==1?(o(),r("div",Tt,[l("div",Mt,[t(c,{color:"green"},{default:a(()=>[t(h(ue))]),_:1}),St])])):z("",!0),e.row.parameterValue.status==0?(o(),r("div",Ft,[l("div",Lt,[t(c,{color:"red"},{default:a(()=>[t(h(me))]),_:1}),Gt]),n(" Fail reason : "+i(e.row.parameterValue.errorMsg),1)])):z("",!0),e.row.parameterValue.status==3?(o(),r("div",Pt,[l("div",Rt,[t(c,null,{default:a(()=>[t(h(xe))]),_:1}),Ot])])):z("",!0)])):(o(),r("div",Bt,i(e.row.parameterValue),1))]),_:1})]),_:1},8,["data"])])}}});var Ut=C(Ht,[["__scopeId","data-v-65919186"]]);const qt=p=>(F("data-v-358a7cae"),p=p(),L(),p),Jt=qt(()=>l("div",null,[l("h4",{class:"h4-title"},"Transactions Details")],-1)),Qt={class:"more-button"},Wt=n("Geth Debug Trace"),Kt=D({__name:"TransactionIndex",props:{txHash:String},setup(p){const u=p,w=N("txs"),y=N(0),c=M([]),b=M([]),g=ge(),m=_e(),x=d=>{console.log("command",d),(d=="tracetx2"||d=="tracetx")&&m.push("/vmtrace?txhash="+u.txHash+"&type="+d)},f=async d=>{c.length=0,b.length=0;const e=await Ve(d);y.value=e.data.logs.length,je(e.data).forEach(s=>{c.push(s)}),e.data.logs.forEach(s=>{b.push(s)})};return f(u.txHash),ve(()=>g.params,async d=>{d.txHash&&f(d.txHash)}),(d,e)=>{const s=q,_=G,v=O,V=B,$=H,j=U,E=J,k=Ut,P=be,Y=Re,Z=ye;return o(),r("div",null,[t(E,null,{default:a(()=>[t(s,{span:12},{default:a(()=>[Jt]),_:1}),t(s,{span:12,class:"more-info"},{default:a(()=>[l("div",Qt,[t(j,{onCommand:x},{dropdown:a(()=>[t($,null,{default:a(()=>[t(V,{command:"tracetx"},{default:a(()=>[Wt]),_:1})]),_:1})]),default:a(()=>[t(v,{style:{width:"5px"},type:"info",size:"small"},{default:a(()=>[t(_,null,{default:a(()=>[t(h(we))]),_:1})]),_:1})]),_:1})])]),_:1})]),_:1}),t(Z,{modelValue:w.value,"onUpdate:modelValue":e[0]||(e[0]=ee=>w.value=ee)},{default:a(()=>[t(P,{label:"Overview",name:"txs"},{default:a(()=>[t(k,{txOverviews:c},null,8,["txOverviews"])]),_:1}),y.value!=0?(o(),S(P,{key:0,name:"logs"},{label:a(()=>[l("span",null,"Logs("+i(y.value)+")",1)]),default:a(()=>[t(Y,{transactionLogs:b},null,8,["transactionLogs"])]),_:1})):z("",!0)]),_:1},8,["modelValue"])])}}});var ca=C(Kt,[["__scopeId","data-v-358a7cae"]]);export{ca as default};
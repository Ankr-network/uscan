import{_ as C,d as D,r as B,o,c as r,e as l,V as I,ah as T,a as t,w as a,t as i,b as n,W as z,p as F,h as L,br as ae,F as N,aM as M,j as oe,n as S,f as h,bs as re,bt as ne,bu as ie,b0 as de,ac as P,i as G,bv as H,bw as U,bx as q,m as le,aN as se,aQ as ce,g as O,R as pe,by as fe,bz as ue,bA as me,E as xe,bB as he,u as ge,J as be,b2 as _e}from"./index.c920246b.js";import{E as ve,a as we}from"./el-tab-pane.cadda52e.js";import{E as J,a as Q}from"./el-col.96d245e9.js";import{E as W,a as K}from"./el-table-column.ac49457f.js";/* empty css               */import{a as ye,g as ke}from"./utils.409af791.js";import{A as ze,e as A}from"./index.eb369523.js";import{_ as Ve}from"./CopyIcon.0b3aa05b.js";import{G as je}from"./transactionService.c25014f2.js";import{g as Ae}from"./transaction.2e77694f.js";import"./index.dfd0bcc2.js";import"./objects.463bde18.js";import"./_commonjsHelpers.b8add541.js";import"./index.7d4bc839.js";const X=p=>(F("data-v-16369e1c"),p=p(),L(),p),Ne={class:"log"},Ce={class:"content"},De=X(()=>l("p",null,"Transaction Receipt Event Logs",-1)),$e={class:"log-content"},Ee={class:"icon-circle"},Ie=X(()=>l("div",{style:{"font-size":"15px","font-weight":"bold",color:"#4a4f55"}},"Address",-1)),Te={key:0},Me={class:"center-row"},Se={class:"topic-index"},Fe={style:{"margin-left":"10px"}},Le=n("Data"),Pe={class:"center-row",style:{"background-color":"#f8f9fa",width:"90%"}},Re={style:{margin:"20px"}},Oe=D({__name:"TransactionLogs",props:{transactionLogs:{type:Array,require:!0}},setup(p){const u=p;return(v,y)=>{const c=J,w=B("router-link"),g=Q,m=ae;return o(),r("div",Ne,[l("div",Ce,[De,l("div",$e,[(o(!0),r(I,null,T(u.transactionLogs,(x,f)=>(o(),r("div",{key:f},[l("div",null,[t(g,null,{default:a(()=>[t(c,{span:2},{default:a(()=>[l("div",Ee,i(x.logIndex),1)]),_:2},1024),t(c,{span:22},{default:a(()=>[l("div",null,[t(g,{class:"log-row"},{default:a(()=>[t(c,{span:2},{default:a(()=>[Ie]),_:1}),t(c,{span:22},{default:a(()=>[l("div",null,[t(w,{to:"/address/"+x.address},{default:a(()=>[n(i(x.address),1)]),_:2},1032,["to"])])]),_:2},1024)]),_:2},1024),l("div",null,[(o(!0),r(I,null,T(x.topics,(d,e)=>(o(),r("div",{key:e},[t(g,{class:"log-row"},{default:a(()=>[t(c,{span:2},{default:a(()=>[e==0?(o(),r("div",Te,"Topics")):z("",!0)]),_:2},1024),t(c,{span:22},{default:a(()=>[l("div",Me,[l("div",Se,i(e),1),l("div",Fe,i(d),1)])]),_:2},1024)]),_:2},1024)]))),128))]),t(g,{class:"log-data-row"},{default:a(()=>[t(c,{span:2},{default:a(()=>[Le]),_:1}),t(c,{span:22,style:{"word-break":"break-all"}},{default:a(()=>[l("div",Pe,[l("div",Re,i(x.data),1)])]),_:2},1024)]),_:2},1024)])]),_:2},1024)]),_:2},1024),l("div",null,[t(m)])])]))),128))])])])}}});var Be=C(Oe,[["__scopeId","data-v-16369e1c"]]);const Ge={key:0},He={style:{"margin-top":"10px"}},Ue=n(" View Input As \xA0 "),qe=n(" Default View "),Je=n("Original"),Qe=n(" Decode Input Data \xA0"),We={key:1,style:{"margin-top":"10px","margin-bottom":"10px"}},Ke=n("\xA0 Switch Back "),Xe=D({__name:"InputDataInfo",props:{inputData:String,methodName:String},setup(p){const u=p,v=N(""),y=N(!0),c=M([]),w=(f,d)=>{c.length=0,new ze().decode(f,d).forEach((b,_)=>{c.push({index:_,type:f[_],data:b})})},g=f=>{y.value=f},m=()=>{const f=u.inputData.slice(10),d="Function: "+u.methodName,e="MethodID: "+u.inputData.slice(0,10);let s="",b=0;for(let _=0,V=f.length;_<V;_+=64)s+="["+b+"] :    "+f.slice(_,_+64)+`
`,b+=1;v.value=d+`

`+e+`
`+s},x=()=>{v.value=u.inputData};return oe(()=>{if(u.methodName!==""){const f=ye(u.methodName);f.length!==0&&w(f,"0x"+u.inputData.slice(10)),m()}else x()}),(f,d)=>{const e=de,s=P,b=G,_=H,V=U,$=q,j=W,E=K;return o(),r("div",null,[y.value?(o(),r("div",Ge,[l("div",null,[t(e,{style:{"font-family":"Monaco"},modelValue:v.value,"onUpdate:modelValue":d[0]||(d[0]=k=>v.value=k),rows:"8",placeholder:"Please input","show-word-limit":"",type:"textarea",readonly:!0},null,8,["modelValue"])]),l("div",He,[t($,null,{dropdown:a(()=>[t(V,null,{default:a(()=>[u.methodName!==""?(o(),S(_,{key:0,onClick:d[1]||(d[1]=k=>m())},{default:a(()=>[qe]),_:1})):z("",!0),t(_,{onClick:d[2]||(d[2]=k=>x())},{default:a(()=>[Je]),_:1})]),_:1})]),default:a(()=>[t(b,{type:"info"},{default:a(()=>[Ue,t(s,null,{default:a(()=>[t(h(re))]),_:1})]),_:1})]),_:1}),u.methodName!==""?(o(),S(b,{key:0,type:"info",style:{"margin-left":"10px"},onClick:d[3]||(d[3]=k=>g(!1))},{default:a(()=>[Qe,t(s,null,{default:a(()=>[t(h(ne))]),_:1})]),_:1})):z("",!0)])])):(o(),r("div",We,[t(E,{data:c,border:"",style:{width:"85%","font-size":"0.4rem","border-radius":"0.4rem"}},{default:a(()=>[t(j,{prop:"index",label:"#"}),t(j,{prop:"type",label:"Type"}),t(j,{prop:"data",label:"Data",width:"500"})]),_:1},8,["data"]),t(b,{style:{"margin-top":"10px"},type:"info",onClick:d[4]||(d[4]=k=>g(!0))},{default:a(()=>[t(s,null,{default:a(()=>[t(h(ie))]),_:1}),Ke]),_:1})]))])}}});var Ye=C(Xe,[["__scopeId","data-v-37117e86"]]);const Y=p=>(F("data-v-97aacbe6"),p=p(),L(),p),Ze={class:"center-row"},et={style:{"max-width":"250px"}},tt={key:0,style:{"font-weight":"900"}},at={key:1,class:"center-row"},ot={key:2,class:"center-row"},rt={key:3,class:"center-row"},nt={key:0,class:"center-row"},it=n(" Contract \xA0 "),dt=n(" \xA0 "),lt={key:1,class:"center-row"},st=n(" \xA0 "),ct={key:4},pt={key:0,class:"center-row"},ft=n(" Contract \xA0 "),ut=n(" Created \xA0 "),mt={key:1,class:"center-row"},xt=n(" Contract \xA0 "),ht=n(" \xA0 "),gt={key:2,class:"center-row"},bt=n(" \xA0 "),_t={key:5},vt={key:6},wt={class:"center-row"},yt=Y(()=>l("div",{style:{"font-weight":"bold"}},"From",-1)),kt=n(" \xA0\xA0\xA0 "),zt=n(" \xA0\xA0\xA0 "),Vt=Y(()=>l("div",{style:{"font-weight":"bold"}},"To",-1)),jt=n(" \xA0\xA0\xA0 "),At=n(" \xA0\xA0\xA0 "),Nt={key:0},Ct={key:1},Dt={key:7},$t={key:8},Et={key:9},It={key:10},Tt={key:11},Mt={key:0},St={class:"success-status"},Ft=n(" \xA0 Success "),Lt={key:1},Pt={class:"fail-status"},Rt=n(" \xA0 Fail "),Ot={key:2},Bt={class:"pending-status"},Gt=n(" \xA0 Pending "),Ht={key:12},Ut=D({__name:"TransactionOverview",props:{txOverviews:{type:Array,require:!0}},setup(p){const u=p;return document.title="Transaction overview | The "+le()+" Explorer",(v,y)=>{const c=P,w=xe,g=W,m=B("router-link"),x=Ve,f=Ye,d=K;return o(),r("div",null,[t(d,{class:"table-border",data:u.txOverviews,"empty-text":"loading...","row-style":{height:"50px"}},{default:a(()=>[t(g,{width:"240"},{default:a(e=>[l("div",Ze,[t(w,{effect:"dark",placement:"top"},{content:a(()=>[l("div",et,i(e.row.parameterExplain),1)]),default:a(()=>[t(c,null,{default:a(()=>[t(h(se))]),_:1})]),_:2},1024),n(" \xA0"+i(e.row.parameterDisplay),1)])]),_:1}),t(g,{prop:"parameterValue"},{default:a(e=>[e.row.parameterName=="blockNumber"?(o(),r("div",tt,[t(m,{to:"/block/"+parseInt(e.row.parameterValue)},{default:a(()=>[n(i(parseInt(e.row.parameterValue)),1)]),_:2},1032,["to"])])):e.row.parameterName=="createTime"?(o(),r("div",at,[t(c,null,{default:a(()=>[t(h(ce))]),_:1}),n("\xA0"+i(h(ke)(e.row.parameterValue))+"\xA0 ("+i(new Date(e.row.parameterValue*1e3).toUTCString())+") ",1)])):e.row.parameterName=="hash"?(o(),r("div",ot,[n(i(e.row.parameterValue)+" \xA0 ",1),t(x,{text:e.row.parameterValue},null,8,["text"])])):e.row.parameterName=="from"?(o(),r("div",rt,[e.row.parameterValue.fromContract?(o(),r("div",nt,[it,t(m,{to:"/address/"+e.row.parameterValue.from},{default:a(()=>[n(i(e.row.parameterValue.from)+" \xA0 "+i(e.row.parameterValue.fromName),1)]),_:2},1032,["to"]),dt,t(x,{text:e.row.parameterValue.from},null,8,["text"])])):(o(),r("div",lt,[t(m,{to:"/address/"+e.row.parameterValue.from},{default:a(()=>[n(i(e.row.parameterValue.from),1)]),_:2},1032,["to"]),st,t(x,{text:e.row.parameterValue.from},null,8,["text"])]))])):e.row.parameterName=="to"?(o(),r("div",ct,[e.row.parameterValue.to==""?(o(),r("div",pt,[ft,t(m,{to:"/address/"+e.row.parameterValue.contractAddress},{default:a(()=>[n(i(e.row.parameterValue.contractAddress),1)]),_:2},1032,["to"]),ut,t(x,{text:e.row.parameterValue.contractAddress},null,8,["text"])])):e.row.parameterValue.toContract?(o(),r("div",mt,[xt,t(m,{to:"/address/"+e.row.parameterValue.to},{default:a(()=>[n(i(e.row.parameterValue.to)+" \xA0 "+i(e.row.parameterValue.toName),1)]),_:2},1032,["to"]),ht,t(x,{text:e.row.parameterValue.to},null,8,["text"])])):(o(),r("div",gt,[t(m,{to:"/address/"+e.row.parameterValue.to},{default:a(()=>[n(i(e.row.parameterValue.to),1)]),_:2},1032,["to"]),bt,t(x,{text:e.row.parameterValue.to},null,8,["text"])]))])):e.row.parameterName=="value"?(o(),r("div",_t,i(h(A).utils.formatEther(e.row.parameterValue))+" "+i(h(O)()),1)):e.row.parameterName=="tokensTransferred"?(o(),r("div",vt,[l("div",{class:pe(e.row.parameterValue.length>=3?"rolling":"")},[(o(!0),r(I,null,T(e.row.parameterValue,(s,b)=>(o(),r("div",{key:b},[l("div",wt,[yt,kt,t(m,{to:"/address/"+s.from},{default:a(()=>[n(i(s.from.slice(0,18)+"..."),1)]),_:2},1032,["to"]),zt,Vt,jt,t(m,{to:"/address/"+s.to},{default:a(()=>[n(i(s.to.slice(0,18)+"..."),1)]),_:2},1032,["to"]),At,t(m,{to:"/address/"+s.address},{default:a(()=>[s.addressName?(o(),r("div",Nt,i(s.addressName),1)):(o(),r("div",Ct,i(s.address.slice(0,18)+"..."),1))]),_:2},1032,["to"])])]))),128))],2)])):e.row.parameterName=="gas"?(o(),r("div",Dt,i(e.row.parameterValue.gas)+" | "+i(e.row.parameterValue.gasUsed),1)):e.row.parameterName=="gasPrice"?(o(),r("div",$t,i(h(A).utils.formatEther(e.row.parameterValue))+" "+i(h(O)()),1)):e.row.parameterName=="maxPriorityFeePerGas"?(o(),r("div",Et," Base: "+i(h(A).utils.formatEther(e.row.parameterValue.baseFeePerGas,"gwei"))+" | Max: "+i(h(A).utils.formatEther(e.row.parameterValue.maxFeePerGas,"gwei"))+" | MaxPriority: "+i(h(A).utils.formatEther(e.row.parameterValue.maxPriorityFeePerGas,"gwei")),1)):e.row.parameterName=="input"?(o(),r("div",It,[t(f,{inputData:e.row.parameterValue.inputContent,methodName:e.row.parameterValue.methodName},null,8,["inputData","methodName"])])):e.row.parameterName=="status"?(o(),r("div",Tt,[e.row.parameterValue.status==1?(o(),r("div",Mt,[l("div",St,[t(c,{color:"green"},{default:a(()=>[t(h(fe))]),_:1}),Ft])])):z("",!0),e.row.parameterValue.status==0?(o(),r("div",Lt,[l("div",Pt,[t(c,{color:"red"},{default:a(()=>[t(h(ue))]),_:1}),Rt]),n(" Fail reason : "+i(e.row.parameterValue.errorMsg),1)])):z("",!0),e.row.parameterValue.status==3?(o(),r("div",Ot,[l("div",Bt,[t(c,null,{default:a(()=>[t(h(me))]),_:1}),Gt])])):z("",!0)])):(o(),r("div",Ht,i(e.row.parameterValue),1))]),_:1})]),_:1},8,["data"])])}}});var qt=C(Ut,[["__scopeId","data-v-97aacbe6"]]);const Jt=p=>(F("data-v-09d9b034"),p=p(),L(),p),Qt=Jt(()=>l("div",null,[l("h4",{class:"h4-title"},"Transactions Details")],-1)),Wt={class:"more-button"},Kt=n("Geth Debug Trace_2"),Xt=D({__name:"TransactionIndex",props:{txHash:String},setup(p){const u=p,v=N("txs"),y=N(0),c=M([]),w=M([]),g=he(),m=ge(),x=d=>{console.log("command",d),(d=="tracetx2"||d=="tracetx")&&m.push("/vmtrace?txhash="+u.txHash+"&type="+d)},f=async d=>{c.length=0,w.length=0;const e=await je(d);y.value=e.data.logs.length,Ae(e.data).forEach(s=>{c.push(s)}),e.data.logs.forEach(s=>{w.push(s)})};return f(u.txHash),be(()=>g.params,async d=>{d.txHash&&f(d.txHash)}),(d,e)=>{const s=J,b=P,_=G,V=H,$=U,j=q,E=Q,k=qt,R=ve,Z=Be,ee=we;return o(),r("div",null,[t(E,null,{default:a(()=>[t(s,{span:12},{default:a(()=>[Qt]),_:1}),t(s,{span:12,class:"more-info"},{default:a(()=>[l("div",Wt,[t(j,{onCommand:x},{dropdown:a(()=>[t($,null,{default:a(()=>[t(V,{command:"tracetx2"},{default:a(()=>[Kt]),_:1})]),_:1})]),default:a(()=>[t(_,{style:{width:"5px"},type:"info",size:"small"},{default:a(()=>[t(b,null,{default:a(()=>[t(h(_e))]),_:1})]),_:1})]),_:1})])]),_:1})]),_:1}),t(ee,{modelValue:v.value,"onUpdate:modelValue":e[0]||(e[0]=te=>v.value=te)},{default:a(()=>[t(R,{label:"Overview",name:"txs"},{default:a(()=>[t(k,{txOverviews:c},null,8,["txOverviews"])]),_:1}),y.value!=0?(o(),S(R,{key:0,name:"logs"},{label:a(()=>[l("span",null,"Logs("+i(y.value)+")",1)]),default:a(()=>[t(Z,{transactionLogs:w},null,8,["transactionLogs"])]),_:1})):z("",!0)]),_:1},8,["modelValue"])])}}});var fa=C(Xt,[["__scopeId","data-v-09d9b034"]]);export{fa as default};

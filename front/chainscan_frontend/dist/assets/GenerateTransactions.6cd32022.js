import{ax as R,bF as A,bG as S,bd as v,C as F,d as T,M as G,k as P,y as k,g as a,bH as U,o as n,n as E,b as r,c as s,t as d,Q as V,V as x,R as B,e as p,bI as O,E as q,X as J,bJ as K,a0 as Q,_ as D,w as X,r as H,a as o,bB as W,bC as Y,bD as Z,bK as ee,p as te,i as oe,f as z,ab as L,bs as ae,ap as re,bL as ne,ag as ie,h as I,T as le,j as se}from"./index.c5813a35.js";import{E as de,a as pe}from"./el-table-column.788b5075.js";/* empty css               */import{f as ce}from"./transactionService.171d171d.js";import{e as $}from"./index.eb369523.js";const fe=R({trigger:A.trigger,placement:S.placement,disabled:A.disabled,visible:v.visible,transition:v.transition,popperOptions:S.popperOptions,tabindex:S.tabindex,content:v.content,popperStyle:v.popperStyle,popperClass:v.popperClass,enterable:{...v.enterable,default:!0},effect:{...v.effect,default:"light"},teleported:v.teleported,title:String,width:{type:[String,Number],default:150},offset:{type:Number,default:void 0},showAfter:{type:Number,default:0},hideAfter:{type:Number,default:200},autoClose:{type:Number,default:0},showArrow:{type:Boolean,default:!0},persistent:{type:Boolean,default:!0}}),ue={"update:visible":i=>F(i),"before-enter":()=>!0,"before-leave":()=>!0,"after-enter":()=>!0,"after-leave":()=>!0},he={name:"ElPopover"},me=T({...he,props:fe,emits:ue,setup(i,{expose:l,emit:c}){const u=i,y=G("popover"),f=P(),_=k(()=>{var t;return(t=a(f))==null?void 0:t.popperRef}),w=k(()=>[{width:U(u.width)},u.popperStyle]),g=k(()=>[y.b(),u.popperClass,{[y.m("plain")]:!!u.content}]),m=k(()=>u.transition==="el-fade-in-linear"),h=()=>{var t;(t=f.value)==null||t.hide()},C=()=>{c("before-enter")},b=()=>{c("before-leave")},e=()=>{c("after-enter")},M=()=>{c("update:visible",!1),c("after-leave")};return l({popperRef:_,hide:h}),(t,qe)=>(n(),E(a(q),O({ref_key:"tooltipRef",ref:f},t.$attrs,{trigger:t.trigger,placement:t.placement,disabled:t.disabled,visible:t.visible,transition:t.transition,"popper-options":t.popperOptions,tabindex:t.tabindex,content:t.content,offset:t.offset,"show-after":t.showAfter,"hide-after":t.hideAfter,"auto-close":t.autoClose,"show-arrow":t.showArrow,"aria-label":t.title,effect:t.effect,enterable:t.enterable,"popper-class":a(g),"popper-style":a(w),teleported:t.teleported,persistent:t.persistent,"gpu-acceleration":a(m),onBeforeShow:C,onBeforeHide:b,onShow:e,onHide:M}),{content:r(()=>[t.title?(n(),s("div",{key:0,class:V(a(y).e("title")),role:"title"},d(t.title),3)):x("v-if",!0),B(t.$slots,"default",{},()=>[p(d(t.content),1)])]),default:r(()=>[t.$slots.reference?B(t.$slots,"reference",{key:0}):x("v-if",!0)]),_:3},16,["trigger","placement","disabled","visible","transition","popper-options","tabindex","content","offset","show-after","hide-after","auto-close","show-arrow","aria-label","effect","enterable","popper-class","popper-style","teleported","persistent","gpu-acceleration"]))}});var be=J(me,[["__file","/home/runner/work/element-plus/element-plus/packages/components/popover/src/popover.vue"]]);const N=(i,l)=>{const c=l.arg||l.value,u=c==null?void 0:c.popperRef;u&&(u.triggerRef=i)};var ge={mounted(i,l){N(i,l)},updated(i,l){N(i,l)}};const ve="popover",xe=K(ge,ve),ye=Q(be,{directive:xe});const j=i=>(te("data-v-7e8591b0"),i=i(),oe(),i),_e=j(()=>z("h4",null,"Status:",-1)),we={key:0,class:"center-row"},ke=p(" \xA0 Success "),ze={key:1,class:"center-row"},je=p(" \xA0 Fail "),Ce={key:2,class:"center-row"},Se=p(" \xA0 Pending "),Ee=j(()=>z("h4",null,"Transaction Fee:",-1)),$e=j(()=>z("h4",null,"Gas Info:",-1)),Te=j(()=>z("h4",null,"Nonce:",-1)),Ae=p(" See more details \xA0 "),Be=T({__name:"BaseTransactionInfo",props:{txHash:String},async setup(i){let l,c;const u=i,f=([l,c]=X(()=>ce(u.txHash)),l=await l,c(),l).data;return(_,w)=>{const g=L,m=ae,h=H("router-link");return n(),s("div",null,[_e,a(f).status==1?(n(),s("div",we,[o(g,{color:"green"},{default:r(()=>[o(a(W))]),_:1}),ke])):x("",!0),a(f).status==0?(n(),s("div",ze,[o(g,{color:"red"},{default:r(()=>[o(a(Y))]),_:1}),je])):x("",!0),a(f).status==3?(n(),s("div",Ce,[o(g,null,{default:r(()=>[o(a(Z))]),_:1}),Se])):x("",!0),o(m),Ee,p(" "+d(a($).utils.formatEther(a(f).gasLimit*a(f).gasUsed))+" Eth ",1),o(m),$e,p(" "+d(a(f).gasUsed)+" Used From "+d(a(f).gasLimit)+" GasLimit ",1),o(m),Te,p(" "+d(a(f).nonce)+" ",1),o(m),o(h,{class:"center-row",to:"/tx/"+i.txHash},{default:r(()=>[Ae,o(g,{size:"large"},{default:r(()=>[o(a(ee))]),_:1})]),_:1},8,["to"])])}}});var Ie=D(Be,[["__scopeId","data-v-7e8591b0"]]);const Ne={key:0},Pe={key:0,style:{width:"170px"}},De={key:1,style:{width:"170px"}},He={key:2},Le={key:3},Me={key:4},Re={key:5,style:{width:"170px"}},Fe={key:6,style:{width:"170px"}},Ge={key:7},Ue={key:8,style:{"font-size":"11px"}},Ve={key:9,style:{width:"170px","font-size":"11px"}},Oe=T({__name:"GenerateTransactions",props:{loadStatus:Boolean,txsData:{type:Array,require:!0},headerData:{type:Array,require:!0}},setup(i){const l=i,c=P("loading...");return re(()=>{l.loadStatus||(c.value="empty data")}),(u,y)=>{const f=L,_=se,w=Ie,g=ye,m=de,h=H("router-link"),C=pe;return n(),E(C,{class:"table-border",data:l.txsData,"empty-text":"loading...","row-style":{height:"50px"}},{empty:r(()=>[p(d(c.value),1)]),default:r(()=>[o(m,{width:"37px"},{default:r(b=>[o(g,{placement:"right",title:"Additional Info",width:320,trigger:"click"},{reference:r(()=>[o(_,{style:{width:"5px"},type:"info",size:"small",onClick:e=>b.row.base=!0},{default:r(()=>[o(f,null,{default:r(()=>[o(a(ne))]),_:1})]),_:2},1032,["onClick"])]),default:r(()=>[b.row.base==!0?(n(),s("div",Ne,[o(w,{txHash:b.row.hash},null,8,["txHash"])])):x("",!0)]),_:2},1024)]),_:1}),(n(!0),s(le,null,ie(l.headerData,b=>(n(),E(m,{key:b.key,property:b.key,label:b.label},{default:r(e=>[e.column.property=="transactionHash"?(n(),s("div",Pe,[o(h,{to:"/tx/"+e.row[e.column.property]},{default:r(()=>[p(d(e.row[e.column.property].slice(0,15)+"..."),1)]),_:2},1032,["to"])])):x("",!0),e.column.property=="hash"?(n(),s("div",De,[o(h,{to:"/tx/"+e.row[e.column.property]},{default:r(()=>[p(d(e.row[e.column.property].slice(0,15)+"..."),1)]),_:2},1032,["to"])])):e.column.property=="blockNumber"?(n(),s("div",He,[o(h,{to:"/block/"+parseInt(e.row[e.column.property])},{default:r(()=>[p(d(parseInt(e.row[e.column.property])),1)]),_:2},1032,["to"])])):e.column.property=="createTime"?(n(),s("div",Le,d(a(I)(e.row[e.column.property])),1)):e.column.property=="createdTime"?(n(),s("div",Me,d(a(I)(e.row[e.column.property])),1)):e.column.property=="from"?(n(),s("div",Re,[o(h,{to:"/address/"+e.row[e.column.property]},{default:r(()=>[p(d(e.row[e.column.property].slice(0,15)+"..."),1)]),_:2},1032,["to"])])):e.column.property=="to"?(n(),s("div",Fe,[o(h,{to:"/address/"+e.row[e.column.property]},{default:r(()=>[p(d(e.row[e.column.property].slice(0,15)+"..."),1)]),_:2},1032,["to"])])):e.column.property=="value"?(n(),s("div",Ge,d(a($).utils.formatUnits(e.row[e.column.property],18)),1)):e.column.property=="gas"?(n(),s("div",Ue,d(a($).utils.formatUnits(e.row[e.column.property],18)),1)):e.column.property=="contract"?(n(),s("div",Ve,[o(h,{to:"/token/"+e.row[e.column.property]},{default:r(()=>[p(d(e.row[e.column.property].slice(0,15)+"..."),1)]),_:2},1032,["to"])])):x("",!0)]),_:2},1032,["property","label"]))),128))]),_:1},8,["data"])}}});var Ye=D(Oe,[["__scopeId","data-v-531e75ae"]]);export{Ye as _};

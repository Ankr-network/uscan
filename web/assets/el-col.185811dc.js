import{ay as d,az as f,aA as p,d as g,s as E,y as r,N as u,B as h,aB as O,o as v,n as w,w as y,T as b,R as k,f as a,Z as j,X as $,Y as _,a1 as N,$ as S}from"./index.20bab758.js";const C=Symbol("rowContextKey"),B=d({tag:{type:String,default:"div"},span:{type:Number,default:24},offset:{type:Number,default:0},pull:{type:Number,default:0},push:{type:Number,default:0},xs:{type:f([Number,Object]),default:()=>p({})},sm:{type:f([Number,Object]),default:()=>p({})},md:{type:f([Number,Object]),default:()=>p({})},lg:{type:f([Number,Object]),default:()=>p({})},xl:{type:f([Number,Object]),default:()=>p({})}}),P={name:"ElCol"},A=g({...P,props:B,setup(n){const e=n,{gutter:s}=E(C,{gutter:r(()=>0)}),t=u("col"),m=r(()=>{const l={};return s.value&&(l.paddingLeft=l.paddingRight=`${s.value/2}px`),l}),i=r(()=>{const l=[];return["span","offset","pull","push"].forEach(o=>{const c=e[o];h(c)&&(o==="span"?l.push(t.b(`${e[o]}`)):c>0&&l.push(t.b(`${o}-${e[o]}`)))}),["xs","sm","md","lg","xl"].forEach(o=>{h(e[o])?l.push(t.b(`${o}-${e[o]}`)):O(e[o])&&Object.entries(e[o]).forEach(([c,x])=>{l.push(c!=="span"?t.b(`${o}-${c}-${x}`):t.b(`${o}-${x}`))})}),s.value&&l.push(t.is("guttered")),l});return(l,R)=>(v(),w($(l.tag),{class:k([a(t).b(),a(i)]),style:j(a(m))},{default:y(()=>[b(l.$slots,"default")]),_:3},8,["class","style"]))}});var K=_(A,[["__file","/home/runner/work/element-plus/element-plus/packages/components/col/src/col.vue"]]);const Z=N(K),L=["start","center","end","space-around","space-between","space-evenly"],T=["top","middle","bottom"],z=d({tag:{type:String,default:"div"},gutter:{type:Number,default:0},justify:{type:String,values:L,default:"start"},align:{type:String,values:T,default:"top"}}),D={name:"ElRow"},I=g({...D,props:z,setup(n){const e=n,s=u("row"),t=r(()=>e.gutter);S(C,{gutter:t});const m=r(()=>{const i={};return e.gutter&&(i.marginRight=i.marginLeft=`-${e.gutter/2}px`),i});return(i,l)=>(v(),w($(i.tag),{class:k([a(s).b(),a(s).is(`justify-${e.justify}`,i.justify!=="start"),a(s).is(`align-${e.align}`,i.align!=="top")]),style:j(a(m))},{default:y(()=>[b(i.$slots,"default")]),_:3},8,["class","style"]))}});var J=_(I,[["__file","/home/runner/work/element-plus/element-plus/packages/components/row/src/row.vue"]]);const q=N(J);export{Z as E,q as a};
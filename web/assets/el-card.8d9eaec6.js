import{ay as c,az as t,d as n,N as i,o as d,c as l,R as o,f as r,T as s,W as p,e as b,Z as h,Y as v,b as u,t as y,a1 as g}from"./index.20bab758.js";const f=c({header:{type:String,default:""},bodyStyle:{type:t([String,Object,Array]),default:""},shadow:{type:String,values:["always","hover","never"],default:"always"}}),m={name:"ElCard"},w=n({...m,props:f,setup(x){const e=i("card");return(a,S)=>(d(),l("div",{class:o([r(e).b(),r(e).is(`${a.shadow}-shadow`)])},[a.$slots.header||a.header?(d(),l("div",{key:0,class:o(r(e).e("header"))},[s(a.$slots,"header",{},()=>[u(y(a.header),1)])],2)):p("v-if",!0),b("div",{class:o(r(e).e("body")),style:h(a.bodyStyle)},[s(a.$slots,"default")],6)],2))}});var _=v(w,[["__file","/home/runner/work/element-plus/element-plus/packages/components/card/src/card.vue"]]);const C=g(_);export{C as E};

import{o as e,b as s}from"../element-plus/element-plus.e38f20bb.js";import{u as a}from"../vue-clipboard3/vue-clipboard3.52d51301.js";import{F as o}from"../@element-plus/@element-plus.f4b8dd72.js";import{_ as t}from"../../assets/index.ce741f41.js";import{C as l,e as p,o as r,c,S as n,Q as u,a4 as i,U as m,u as v}from"../@vue/@vue.f5a41070.js";const d={class:"copy-content"};var b=t(l({__name:"CopyIcon",props:{text:String},setup(t){const l=t,b=p("Copy to clipboard"),f=p(!1),{toClipboard:y}=a(),_=async()=>{b.value="Copy Success";try{await y(l.text)}catch(e){}};return(a,t)=>{const l=s,p=e;return r(),c("div",d,[n(p,{placement:"right",visible:f.value},{content:u((()=>[i(m(b.value),1)])),default:u((()=>[n(l,{class:"copy-icon",onClick:_,onMouseenter:t[0]||(t[0]=e=>f.value=!0),onMouseleave:t[1]||(t[1]=e=>{f.value=!1,b.value="Copy to clipboard"})},{default:u((()=>[n(v(o))])),_:1})])),_:1},8,["visible"])])}}}),[["__scopeId","data-v-546336e4"]]);export{b as _};

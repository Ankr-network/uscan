import{k as f}from"./index.e16d46e2.js";const l=function(t,s,n,o){console.log("type",n,"blockNumber",o);const e=t*s,c=s;let i="";return o!==-1?i="/blocks/"+o+"/txs?offset="+e+"&limit="+c:n==="all"?i="/txs?offset="+e+"&limit="+c:i="/tokens/txns/"+n+"?offset="+e+"&limit="+c,f({url:i,method:"get"})},a=function(t,s,n,o){const e=t*s,c=s;let i="";return n==="internal"&&(n="erc20"),n==="txs"?i="/accounts/"+o+"/txns?offset="+e+"&limit="+c:i="/accounts/"+o+"/txns-"+n+"?offset="+e+"&limit="+c,f({url:i,method:"get"})},u=function(t,s,n){const o=t*s,e=s,c="/accounts/"+n+"/txns-internal?offset="+o+"&limit="+e;return f({url:c,method:"get"})},m=function(t,s,n,o){const e=t*s,c=s,i="/tokens/txns/"+n+"?contract="+o+"&offset="+e+"&limit="+c;return f({url:i,method:"get"})},x=function(t){return f({url:"/txs/"+t,method:"get"})},d=function(t){return f({url:"/txs/"+t+"/base",method:"get"})},h=function(t,s){return f({url:"/txs/"+t+"/"+s,method:"get"})};export{x as G,l as a,d as b,a as c,u as d,m as e,h as f};

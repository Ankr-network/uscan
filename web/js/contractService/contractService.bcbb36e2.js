import{r as t}from"../../assets/index.bb287700.js";const e=function(){return t({url:"/contracts/metadata",method:"get"})},r=function(e){return t({url:"/contracts-verify/"+e+"/status",method:"get"})},a=function(e,r){return t({url:"/contracts/"+e+"/verify",headers:{"Content-Type":"multipart/form-data"},method:"post",data:r})},n=function(e){return t({url:"/contracts/"+e+"/content",method:"get"})},o=function(e,r,a){let n=1;switch(e){case"erc20":n=1;break;case"erc721":n=2}return t({url:"/accounts?type="+n+"&offset="+r*a+"&limit="+a,method:"get"})};export{e as G,a as S,n as a,r as b,o as c};

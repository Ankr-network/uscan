import{aG as p,aH as v,aI as d,aJ as P,aK as _,a8 as O}from"./index.03ebd84b.js";import{i as c}from"./index.9a8c641f.js";var m=function(){try{var r=p(Object,"defineProperty");return r({},"",{}),r}catch{}}(),l=m;function x(r,n,t){n=="__proto__"&&l?l(r,n,{configurable:!0,enumerable:!0,value:t,writable:!0}):r[n]=t}var I=Object.prototype,h=I.hasOwnProperty;function K(r,n,t){var a=r[n];(!(h.call(r,n)&&v(a,t))||t===void 0&&!(n in r))&&x(r,n,t)}function q(r,n,t,a){if(!d(r))return r;n=P(n,r);for(var s=-1,u=n.length,g=u-1,i=r;i!=null&&++s<u;){var e=_(n[s]),f=t;if(e==="__proto__"||e==="constructor"||e==="prototype")return r;if(s!=g){var o=i[e];f=a?a(o,e,i):void 0,f===void 0&&(f=d(o)?o:c(n[s+1])?[]:{})}K(i,e,f),i=i[e]}return r}function w(r,n,t){return r==null?r:q(r,n,t)}const H=r=>Object.entries(r),J=(r,n,t)=>({get value(){return O(r,n,t)},set value(a){w(r,n,a)}});export{H as e,J as g};

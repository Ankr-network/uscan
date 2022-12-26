import{_ as j,d as C,r as M,o as e,n as _,w as g,c as i,ah as S,a as x,b as k,t as a,W as z,f as n,g as b,V as I,m as L,F as v,aM as N,l as F,e as B,p as P,h as T}from"./index.774cbafa.js";import{E as V}from"./el-pagination.1900ee1d.js";/* empty css               */import"./el-select.f372dc0a.js";import{E as A,a as G}from"./el-table-column.f61d754b.js";import{g as $}from"./utils.fc402d12.js";import{e as y}from"./index.eb369523.js";import{B as q,a as w}from"./block.0ed52f69.js";import"./isEqual.c1107be0.js";import"./index.f445393f.js";import"./index.261af684.js";import"./objects.9fcb7131.js";import"./_commonjsHelpers.b8add541.js";import"./index.7d4bc839.js";const R={key:0},U={key:1},H={key:2},W={key:3},J={key:4},K={key:5,style:{width:"180px"}},O=C({__name:"GenerateBlocks",props:{blocksData:{type:Array,require:!0},headerData:{type:Array,require:!0}},setup(s){const d=s;return(h,c)=>{const o=M("router-link"),r=A,u=G;return e(),_(u,{class:"table-border",data:d.blocksData,"empty-text":"loading...","row-style":{height:"50px"}},{default:g(()=>[(e(!0),i(I,null,S(d.headerData,l=>(e(),_(r,{key:l.key,property:l.key,label:l.label},{default:g(t=>[t.column.property=="number"?(e(),i("div",R,[x(o,{to:"/block/"+parseInt(t.row[t.column.property])},{default:g(()=>[k(a(parseInt(t.row[t.column.property])),1)]),_:2},1032,["to"])])):z("",!0),t.column.property=="timestamp"?(e(),i("div",U,a(n($)(t.row[t.column.property])),1)):t.column.property=="gasUsed"?(e(),i("div",H,a(n(y).utils.formatEther(t.row[t.column.property]))+" "+a(n(b)()),1)):t.column.property=="gasLimit"?(e(),i("div",W,a(n(y).utils.formatEther(t.row[t.column.property]))+" "+a(n(b)()),1)):t.column.property=="baseFeePerGas"?(e(),i("div",J,a(n(y).utils.formatEther(t.row[t.column.property]))+" "+a(n(b)()),1)):t.column.property=="miner"?(e(),i("div",K,[x(o,{to:"/address/"+t.row.miner},{default:g(()=>[k(a(t.row.miner.slice(0,18)+"..."),1)]),_:2},1032,["to"])])):z("",!0)]),_:2},1032,["property","label"]))),128))]),_:1},8,["data"])}}});var Q=j(O,[["__scopeId","data-v-cb9a198c"]]);const X=s=>(P("data-v-209fd945"),s=s(),T(),s),Y=X(()=>B("h4",{class:"h4-title"},"Blocks",-1)),Z={style:{"margin-top":"1%",display:"flex","justify-content":"center"}},tt=C({__name:"BlocksList",async setup(s){let d,h;document.title="Blocks | The "+L()+" Explorer";const c=v(1),o=v(25),r=N([]),u=([d,h]=F(()=>w(!0,c.value-1,o.value)),d=await d,h(),d);u.data.items.forEach(f=>{r.push(f)});const l=v(u.data.total),t=async f=>{r.length=0,c.value=1,o.value=f;const p=await w(!0,c.value-1,o.value);p.data.items.forEach(m=>{r.push(m)}),l.value=p.data.total},E=async f=>{r.length=0,c.value=f;const p=await w(!0,c.value-1,o.value);p.data.items.forEach(m=>{r.push(m)}),l.value=p.data.total};return(f,p)=>{const m=Q,D=V;return e(),i("div",null,[Y,x(m,{blocksData:r,headerData:n(q)},null,8,["blocksData","headerData"]),B("div",Z,[x(D,{small:"",background:"",currentPage:c.value,"page-size":o.value,"page-sizes":[10,25,50,100],layout:"total, sizes, prev, pager, next, jumper",total:l.value,onSizeChange:t,onCurrentChange:E},null,8,["currentPage","page-size","total"])])])}}});var gt=j(tt,[["__scopeId","data-v-209fd945"]]);export{gt as default};
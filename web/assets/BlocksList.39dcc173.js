import{_,d as k,r as E,o as i,n as b,w as u,c as f,ah as D,a as g,b as y,t as x,W as w,f as z,g as M,V as S,m as I,F as h,aM as L,l as N,e as j,p as T,h as V}from"./index.e25a3bc2.js";import{E as A}from"./el-pagination.d494e440.js";/* empty css               */import"./el-select.37c13417.js";import{E as F,a as P}from"./el-table-column.5363a24e.js";import{B as G,a as v}from"./block.5c440c3e.js";import"./isEqual.4b7cfd62.js";import"./index.2ccb97f8.js";import"./index.c2e25f46.js";import"./objects.dbdf6c99.js";import"./index.7d4bc839.js";const $={key:0},q={key:1},R={key:2,style:{width:"180px"}},H=k({__name:"GenerateBlocks",props:{blocksData:{type:Array,require:!0},headerData:{type:Array,require:!0}},setup(d){const o=d;return(m,r)=>{const t=E("router-link"),a=F,p=P;return i(),b(p,{class:"table-border",data:o.blocksData,"empty-text":"loading...","row-style":{height:"50px"}},{default:u(()=>[(i(!0),f(S,null,D(o.headerData,n=>(i(),b(a,{key:n.key,property:n.key,label:n.label},{default:u(e=>[e.column.property=="number"?(i(),f("div",$,[g(t,{to:"/block/"+parseInt(e.row[e.column.property])},{default:u(()=>[y(x(parseInt(e.row[e.column.property])),1)]),_:2},1032,["to"])])):w("",!0),e.column.property=="timestamp"?(i(),f("div",q,x(z(M)(e.row[e.column.property])),1)):e.column.property=="miner"?(i(),f("div",R,[g(t,{to:"/address/"+e.row.miner},{default:u(()=>[y(x(e.row.miner.slice(0,18)+"..."),1)]),_:2},1032,["to"])])):w("",!0)]),_:2},1032,["property","label"]))),128))]),_:1},8,["data"])}}});var W=_(H,[["__scopeId","data-v-581e313d"]]);const J=d=>(T("data-v-1f4739a6"),d=d(),V(),d),K=J(()=>j("h4",{class:"h4-title"},"Blocks",-1)),O={style:{"margin-top":"1%",display:"flex","justify-content":"center"}},Q=k({__name:"BlocksList",async setup(d){let o,m;document.title="Blocks | The "+I+" Explorer";const r=h(1),t=h(25),a=L([]),p=([o,m]=N(()=>v(!0,r.value-1,t.value)),o=await o,m(),o);p.data.items.forEach(l=>{a.push(l)});const n=h(p.data.total),e=async l=>{a.length=0,r.value=1,t.value=l;const s=await v(!0,r.value-1,t.value);s.data.items.forEach(c=>{a.push(c)}),n.value=s.data.total},C=async l=>{a.length=0,r.value=l;const s=await v(!0,r.value-1,t.value);s.data.items.forEach(c=>{a.push(c)}),n.value=s.data.total};return(l,s)=>{const c=W,B=A;return i(),f("div",null,[K,g(c,{blocksData:a,headerData:z(G)},null,8,["blocksData","headerData"]),j("div",O,[g(B,{small:"",background:"",currentPage:r.value,"page-size":t.value,"page-sizes":[10,25,50,100],layout:"total, sizes, prev, pager, next, jumper",total:n.value,onSizeChange:e,onCurrentChange:C},null,8,["currentPage","page-size","total"])])])}}});var de=_(Q,[["__scopeId","data-v-1f4739a6"]]);export{de as default};
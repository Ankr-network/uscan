import{_,d as k,r as E,o as n,n as b,w as m,c as f,ah as D,a as u,b as y,t as x,W as w,f as z,V as M,m as S,F as h,aM as I,l as L,e as j,p as N,h as T}from"./index.b80e27f4.js";import{E as V}from"./el-pagination.57f78bf1.js";/* empty css               */import"./el-select.4683c41f.js";import{E as A,a as F}from"./el-table-column.a1d1fbd4.js";import{g as P}from"./utils.409af791.js";import{B as G,a as v}from"./block.70da773a.js";import"./isEqual.3ed7cd22.js";import"./index.143c4a8f.js";import"./index.499bc329.js";import"./objects.422c4ffa.js";import"./index.7d4bc839.js";const $={key:0},q={key:1},R={key:2,style:{width:"180px"}},H=k({__name:"GenerateBlocks",props:{blocksData:{type:Array,require:!0},headerData:{type:Array,require:!0}},setup(i){const o=i;return(g,r)=>{const t=E("router-link"),a=A,p=F;return n(),b(p,{class:"table-border",data:o.blocksData,"empty-text":"loading...","row-style":{height:"50px"}},{default:m(()=>[(n(!0),f(M,null,D(o.headerData,d=>(n(),b(a,{key:d.key,property:d.key,label:d.label},{default:m(e=>[e.column.property=="number"?(n(),f("div",$,[u(t,{to:"/block/"+parseInt(e.row[e.column.property])},{default:m(()=>[y(x(parseInt(e.row[e.column.property])),1)]),_:2},1032,["to"])])):w("",!0),e.column.property=="timestamp"?(n(),f("div",q,x(z(P)(e.row[e.column.property])),1)):e.column.property=="miner"?(n(),f("div",R,[u(t,{to:"/address/"+e.row.miner},{default:m(()=>[y(x(e.row.miner.slice(0,18)+"..."),1)]),_:2},1032,["to"])])):w("",!0)]),_:2},1032,["property","label"]))),128))]),_:1},8,["data"])}}});var W=_(H,[["__scopeId","data-v-581e313d"]]);const J=i=>(N("data-v-209fd945"),i=i(),T(),i),K=J(()=>j("h4",{class:"h4-title"},"Blocks",-1)),O={style:{"margin-top":"1%",display:"flex","justify-content":"center"}},Q=k({__name:"BlocksList",async setup(i){let o,g;document.title="Blocks | The "+S()+" Explorer";const r=h(1),t=h(25),a=I([]),p=([o,g]=L(()=>v(!0,r.value-1,t.value)),o=await o,g(),o);p.data.items.forEach(l=>{a.push(l)});const d=h(p.data.total),e=async l=>{a.length=0,r.value=1,t.value=l;const s=await v(!0,r.value-1,t.value);s.data.items.forEach(c=>{a.push(c)}),d.value=s.data.total},C=async l=>{a.length=0,r.value=l;const s=await v(!0,r.value-1,t.value);s.data.items.forEach(c=>{a.push(c)}),d.value=s.data.total};return(l,s)=>{const c=W,B=V;return n(),f("div",null,[K,u(c,{blocksData:a,headerData:z(G)},null,8,["blocksData","headerData"]),j("div",O,[u(B,{small:"",background:"",currentPage:r.value,"page-size":t.value,"page-sizes":[10,25,50,100],layout:"total, sizes, prev, pager, next, jumper",total:d.value,onSizeChange:e,onCurrentChange:C},null,8,["currentPage","page-size","total"])])])}}});var le=_(Q,[["__scopeId","data-v-209fd945"]]);export{le as default};

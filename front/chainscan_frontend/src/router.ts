import { createWebHistory, createRouter } from 'vue-router';
import { RouteRecordRaw } from 'vue-router';

const routes: Array<RouteRecordRaw> = [
  { path: '/', component: () => import('./components/ScanHome/ScanHome.vue') },
  { path: '/block/:blockNumber', component: () => import('./components/ScanBlock/BlockIndex.vue'), props: true },
  { path: '/blocks', component: () => import('./components/ScanBlock/BlocksList.vue') },
  { path: '/tx/:txHash', component: () => import('./components/ScanTransaction/TransactionIndex.vue'), props: true },
  { path: '/txs/:txsType', component: () => import('./components/ScanTransaction/TransactionsList.vue'), props: true },
  { path: '/address/:address', component: () => import('./components/ScanAddress/AddressIndex.vue'), props: true },
  { path: '/token/:address', component: () => import('./components/ScanToken/TokenAddress.vue'), props: true },
  { path: '/verifyContract/input', component: () => import('./components/ScanContract/VerifyContractInput.vue') },
  { path: '/verifyContract/submit', component: () => import('./components/ScanContract/VerifyContractSubmit.vue') },
  {
    path: '/token/nfts/:address/:tokenID/:type',
    component: () => import('./components/ScanToken/TokenOverview.vue'),
    props: true,
  },
  { path: '/search/:searchText', component: () => import('./components/ScanHeader/SearchRes.vue'), props: true },
  { path: '/search/not-found', component: () => import('./components/ScanHeader/SearchNotFound.vue') },
  { path: '/vmtrace', component: () => import('./components/ScanTransaction/GethDebugTrace.vue') },
  { path: '/exportData', component: () => import('./components/ScanTransaction/ExportData.vue') },
  {
    path: '/contracts/:contractType',
    component: () => import('./components/ScanContract/ContractList.vue'),
    props: true,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;

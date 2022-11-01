<template lang="">
  <div>
    <div class="center-row">
      <h2>Token</h2>
    </div>
    <div>
      <el-card class="box-card">
        <template #header>
          <div class="card-header">
            <span>Overview</span>
          </div>
        </template>
        <div class="card-content">
          <el-row>
            <el-col :span="10">Contract:</el-col>
            <el-col :span="14">{{ nftDetail.data.contract }}</el-col>
          </el-row>
          <el-row>
            <el-col :span="10">Token ID:</el-col>
            <el-col :span="14">{{ nftDetail.data.tokenID }}</el-col>
          </el-row>
        </div>
      </el-card>
    </div>
    <div>
      <el-tabs v-model="activeName">
        <el-tab-pane label="Transactions" name="txs">
          <generate-transactions
            :txsData="txsData"
            :headerData="Erc721TransactionsHeaderList"
            :loadStatus="isEmpty"
          ></generate-transactions>
          <div style="margin-top: 1%; display: flex; justify-content: center">
            <el-pagination
              small
              background
              :currentPage="currentPageIndex"
              :page-size="pageSizeNumber"
              :page-sizes="[10, 25, 50, 100]"
              layout="total, sizes, prev, pager, next, jumper"
              :total="total"
              @size-change="handleSizeChange"
              @current-change="handleCurrentChange"
            />
          </div>
        </el-tab-pane>
      </el-tabs>
    </div>
  </div>
</template>
<script lang="ts" setup>
import { GetNFTDetailByID } from '../../script/service/tokenService';
import { GetTransactionsByToken } from '../../script/service/transactionService';
import {
  TransactionDetail,
  TransactionsHeaderList,
  Erc721TransactionsHeaderList,
  Erc20TransactionsHeaderList,
} from '../../script/model/transaction';
import { TableHeader } from '../../script/model/index';
import { ref, reactive, watch } from 'vue';

const props = defineProps({
  address: String,
  tokenID: String,
  type: String,
});

const activeName = ref('txs');
const txsData: TransactionDetail[] = reactive([]);
const headerData: TableHeader[] = reactive([]);
const currentPageIndex = ref(1);
const pageSizeNumber = ref(25);
const total = ref(0);
const isEmpty = ref(true);

const nftDetail = await GetNFTDetailByID(props.address as string, props.tokenID as string, props.type as string);
// console.log('nftDetail', nftDetail.data);

const GetTransactions = async () => {
  if (props.type === 'all' || props.type === 'erc20' || props.type === 'erc721' || props.type === 'erc1155') {
    if (props.type === 'all') {
      headerData.push(...TransactionsHeaderList);
    } else if (props.type === 'erc20') {
      headerData.push(...Erc20TransactionsHeaderList);
    } else if (props.type === 'erc721') {
      headerData.push(...Erc721TransactionsHeaderList);
    } else if (props.type === 'erc1155') {
      headerData.push(...Erc721TransactionsHeaderList);
    }
    const nftTransactions = await GetTransactionsByToken(
      currentPageIndex.value - 1,
      pageSizeNumber.value,
      props.type as string,
      props.address as string
    );
    // console.log('nftTransactions', nftTransactions);
    nftTransactions.data.items.forEach((element) => {
      // console.log('nftTransactions', element);
      txsData.push(element);
    });
    total.value = nftTransactions.data.total;
    if (nftTransactions.data.total == 0) {
      isEmpty.value = false;
    }
  }
};

await GetTransactions();

// console.log('txsData', txsData);

watch(props, async () => {
  if (props.address !== undefined && props.tokenID !== undefined && props.type !== undefined) {
    txsData.length = 0;
    headerData.length = 0;
    currentPageIndex.value = 1;
    pageSizeNumber.value = 25;
    GetTransactions();
  }
});

const handleSizeChange = async (pageSizeArg: number) => {
  txsData.length = 0;
  currentPageIndex.value = 1;
  pageSizeNumber.value = pageSizeArg;
  GetTransactions();
};

const handleCurrentChange = async (currentPageArg: number) => {
  txsData.length = 0;
  currentPageIndex.value = currentPageArg;
  GetTransactions();
};
</script>
<style lang="less" scoped>
@import '../../css/style.css';
</style>

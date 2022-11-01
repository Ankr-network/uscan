<template lang="">
  <div>
    <div class="center-row">
      <h4 class="h4-title">
        {{ tableTitle }}
      </h4>

      <el-button v-if="txType !== ''" style="margin: 10px; font-weight: bold" color="#DEE1E4" size="small">
        {{ txType }}
      </el-button>
    </div>

    <div>
      <h4 class="tx-sub-title">(Showing the last {{ total }} records only)</h4>
    </div>

    <generate-transactions :txsData="txsData" :headerData="headerData" :loadStatus="isEmpty"></generate-transactions>
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
  </div>
</template>
<script lang="ts" setup>
import { GetTransactions } from '../../script/service/transactionService';
import {
  TransactionDetail,
  TransactionsHeaderList,
  Erc721TransactionsHeaderList,
  Erc20TransactionsHeaderList,
} from '../../script/model/transaction';
import { TableHeader } from '../../script/model/index';
import { reactive, ref, watch, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import { getTitle } from '../../script/utils';

const route = useRoute();

document.title = 'Transactions | The ' + getTitle + ' Explorer';

const props = defineProps({
  txsType: String,
});

const currentPageIndex = ref(1);
const pageSizeNumber = ref(25);
const txsData: TransactionDetail[] = reactive([]);
const headerData: TableHeader[] = reactive([]);
const total = ref(0);
const blockNumber: number = route.query.block === undefined ? -1 : (route.query.block as unknown as number);
const tableTitle = ref('title');
const txType = ref('');
const isEmpty = ref(true);

onMounted(() => {
  setTitle();
});

const getTransactions = async () => {
  if (
    props.txsType === 'all' ||
    props.txsType === 'erc20' ||
    props.txsType === 'erc721' ||
    props.txsType === 'erc1155'
  ) {
    if (props.txsType === 'all') {
      headerData.push(...TransactionsHeaderList);
    } else if (props.txsType === 'erc20') {
      headerData.push(...Erc20TransactionsHeaderList);
    } else if (props.txsType === 'erc721') {
      headerData.push(...Erc721TransactionsHeaderList);
    } else if (props.txsType === 'erc1155') {
      headerData.push(...Erc721TransactionsHeaderList);
    }
    const res = await GetTransactions(currentPageIndex.value - 1, pageSizeNumber.value, props.txsType, blockNumber);
    // console.log(res);
    res.data.items.forEach((element) => {
      txsData.push(element);
    });
    total.value = res.data.total;
    if (res.data.total == 0) {
      isEmpty.value = false;
    }
  }
};

getTransactions();

watch(props, async () => {
  // console.log(props.txsType);
  setTitle();
  txsData.length = 0;
  headerData.length = 0;
  currentPageIndex.value = 1;
  pageSizeNumber.value = 25;
  getTransactions();
});

const setTitle = () => {
  if (props.txsType === 'all') {
    if (blockNumber !== -1) {
      tableTitle.value = 'Transactions For Block' + blockNumber;
    } else {
      tableTitle.value = 'Transactions';
    }
  } else if (props.txsType === 'erc20') {
    tableTitle.value = 'Token Transfers';
    txType.value = 'ERC-20';
  } else if (props.txsType === 'erc721') {
    tableTitle.value = 'Non-Fungible Token Transfers';
    txType.value = 'ERC-721';
  } else if (props.txsType === 'erc1155') {
    tableTitle.value = 'Multi-Token Token Tracker';
    txType.value = 'ERC-1155';
  }
};

const handleSizeChange = async (pageSizeArg: number) => {
  txsData.length = 0;
  currentPageIndex.value = 1;
  pageSizeNumber.value = pageSizeArg;
  const res = await GetTransactions(
    currentPageIndex.value - 1,
    pageSizeNumber.value,
    props.txsType as string,
    blockNumber
  );
  res.data.items.forEach((element) => {
    txsData.push(element);
  });
  total.value = res.data.total;
};

const handleCurrentChange = async (currentPageArg: number) => {
  txsData.length = 0;
  currentPageIndex.value = currentPageArg;
  const res = await GetTransactions(
    currentPageIndex.value - 1,
    pageSizeNumber.value,
    props.txsType as string,
    blockNumber
  );
  res.data.items.forEach((element) => {
    txsData.push(element);
  });
  total.value = res.data.total;
};
</script>
<style lang="less" scoped>
@import '../../css/style.css';
</style>

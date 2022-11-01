<template lang="">
  <div>
    <div class="center-row">
      <h4 class="h4-title">
        Address <span class="small text-secondary">&nbsp;&nbsp;{{ props.address }}</span>
      </h4>
      &nbsp;
      <copy-icon :text="props.address"></copy-icon>
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
            <el-col :span="10">Balance:</el-col>
            <el-col :span="14">{{ ethers.utils.formatUnits(props.addressInfo.balance, 18) }} Eth</el-col>
          </el-row>
        </div>
      </el-card>
    </div>
    <el-tabs v-model="activeName">
      <el-tab-pane label="Transactions" name="txs">
        <generate-transactions
          :txsData="txsData"
          :headerData="headerData"
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
        <div style="float: right; margin-top: -25px">
          <div class="download">[ Download</div>
          <router-link class="download" :to="'/exportData?type=txns' + '&a=' + props.address">
            excel Export
          </router-link>
          <div class="download">
            <el-icon><Download /></el-icon>]
          </div>
        </div>
      </el-tab-pane>
      <el-tab-pane v-if="internalCount != 0" name="internal">
        <template #label>
          <span>Internal Txns({{ internalCount }})</span>
        </template>
        <internal-transactions
          :txsData="internalTxsData"
          :headerData="InternalTransactionsHeaderList"
        ></internal-transactions>
        <div style="margin-top: 1%; display: flex; justify-content: center">
          <el-pagination
            small
            background
            :currentPage="currentPageIndex"
            :page-size="pageSizeNumber"
            :page-sizes="[10, 25, 50, 100]"
            layout="total, sizes, prev, pager, next, jumper"
            :total="internalCount"
            @size-change="handleInternalSizeChange"
            @current-change="handleInternalCurrentChange"
          />
        </div>
        <div style="float: right; margin-top: -25px">
          <div class="download">[ Download</div>
          <router-link class="download" :to="'/exportData?type=txns-internal' + '&a=' + props.address">
            excel Export
          </router-link>
          <div class="download">
            <el-icon><Download /></el-icon>]
          </div>
        </div>
      </el-tab-pane>
      <el-tab-pane v-if="erc20count != 0" name="erc20">
        <template #label>
          <span>Erc20 Token Txns({{ erc20count }})</span>
        </template>
        <generate-transactions
          :txsData="txsData"
          :headerData="headerData"
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

        <div style="float: right; margin-top: -25px">
          <div class="download">[ Download</div>
          <router-link class="download" :to="'/exportData?type=txns-erc20' + '&a=' + props.address">
            excel Export
          </router-link>
          <div class="download">
            <el-icon><Download /></el-icon>]
          </div>
        </div>
      </el-tab-pane>
      <el-tab-pane v-if="erc721count != 0" name="erc721">
        <template #label>
          <span>Erc721 Token Txns({{ erc721count }})</span>
        </template>
        <generate-transactions
          :txsData="txsData"
          :headerData="headerData"
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
        <div style="float: right; margin-top: -25px">
          <div class="download">[ Download</div>
          <router-link class="download" :to="'/exportData?type=txns-erc721' + '&a=' + props.address">
            excel Export
          </router-link>
          <div class="download">
            <el-icon><Download /></el-icon>]
          </div>
        </div>
      </el-tab-pane>
      <el-tab-pane v-if="erc1155count != 0" name="erc1155">
        <template #label>
          <span>Erc1155 Token Txns({{ erc1155count }})</span>
        </template>
        <generate-transactions
          :txsData="txsData"
          :headerData="headerData"
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
        <div style="float: right; margin-top: -25px">
          <div class="download">[ Download</div>
          <router-link class="download" :to="'/exportData?type=txns-erc1155' + '&a=' + props.address">
            excel Export
          </router-link>
          <div class="download">
            <el-icon><Download /></el-icon>]
          </div>
        </div>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>
<script lang="ts" setup>
import { ref, reactive, watch } from 'vue';
import { AddressDetail } from '../../script/model/address';
import { ethers } from 'ethers';
import {
  TransactionDetail,
  TransactionsHeaderList,
  Erc721TransactionsHeaderList,
  Erc20TransactionsHeaderList,
  InternalTransactionsHeaderList,
  InternalTransactionDetail,
} from '../../script/model/transaction';
import { TableHeader } from '../../script/model/index';
import { GetTransactionsByAddress, GetInternalTransactionsByAddress } from '../../script/service/transactionService';
import { Download } from '@element-plus/icons-vue';

const props = defineProps({
  address: String,
  addressInfo: {
    type: Object as () => AddressDetail,
  },
});

const activeName = ref('txs');
const txsData: TransactionDetail[] = reactive([]);
const internalTxsData: InternalTransactionDetail[] = reactive([]);
const headerData: TableHeader[] = reactive([]);
const currentPageIndex = ref(1);
const pageSizeNumber = ref(25);
const total = ref(0);
const erc20count = ref(0);
const erc721count = ref(0);
const erc1155count = ref(0);
const internalCount = ref(0);
const isEmpty = ref(true);

watch(props, async () => {
  currentPageIndex.value = 1;
  pageSizeNumber.value = 25;
  if (props.addressInfo?.id !== undefined) {
    txsData.length = 0;
    if (activeName.value === 'txs') {
      headerData.push(...TransactionsHeaderList);
    }
    const res = await GetTransactionsByAddress(
      currentPageIndex.value - 1,
      pageSizeNumber.value,
      'txs',
      props.address as string
    );
    res.data.items.forEach((element) => {
      txsData.push(element);
    });
    total.value = res.data.total;

    const resErc20 = await GetTransactionsByAddress(
      currentPageIndex.value - 1,
      pageSizeNumber.value,
      'erc20',
      props.address as string
    );
    erc20count.value = resErc20.data.total;

    const resErc721 = await GetTransactionsByAddress(
      currentPageIndex.value - 1,
      pageSizeNumber.value,
      'erc721',
      props.address as string
    );
    erc721count.value = resErc721.data.total;

    const resErc1155 = await GetTransactionsByAddress(
      currentPageIndex.value - 1,
      pageSizeNumber.value,
      'erc1155',
      props.address as string
    );
    erc1155count.value = resErc1155.data.total;

    const resInternal = await GetInternalTransactionsByAddress(
      currentPageIndex.value - 1,
      pageSizeNumber.value,
      props.address as string
    );
    internalCount.value = resInternal.data.total;
    // console.log('resInternal', resInternal);
  }
});

watch(activeName, async (currentValue) => {
  currentPageIndex.value = 1;
  pageSizeNumber.value = 25;

  // console.log('switch', currentValue);
  txsData.length = 0;
  headerData.length = 0;
  if (activeName.value === 'txs') {
    headerData.push(...TransactionsHeaderList);
  } else if (activeName.value === 'erc20') {
    headerData.push(...Erc20TransactionsHeaderList);
  } else if (activeName.value === 'erc721') {
    headerData.push(...Erc721TransactionsHeaderList);
  } else if (activeName.value === 'erc1155') {
    headerData.push(...Erc721TransactionsHeaderList);
  } else if (activeName.value === 'internal') {
    headerData.push(...InternalTransactionsHeaderList);
  }

  if (activeName.value == 'internal') {
    const resInternal = await GetInternalTransactionsByAddress(
      currentPageIndex.value - 1,
      pageSizeNumber.value,
      props.address as string
    );
    resInternal.data.items.forEach((element) => {
      internalTxsData.push(element);
    });
    internalCount.value = resInternal.data.total;
  } else {
    const res = await GetTransactionsByAddress(
      currentPageIndex.value - 1,
      pageSizeNumber.value,
      activeName.value,
      props.address as string
    );
    res.data.items.forEach((element) => {
      txsData.push(element);
    });
    total.value = res.data.total;
    if (res.data.total == 0) {
      isEmpty.value = false;
    }
  }
});

const handleSizeChange = async (pageSizeArg: number) => {
  txsData.length = 0;
  currentPageIndex.value = 1;
  pageSizeNumber.value = pageSizeArg;
  const res = await GetTransactionsByAddress(
    currentPageIndex.value - 1,
    pageSizeNumber.value,
    activeName.value,
    props.address as string
  );
  res.data.items.forEach((element) => {
    txsData.push(element);
  });
  total.value = res.data.total;
};

const handleCurrentChange = async (currentPageArg: number) => {
  txsData.length = 0;
  currentPageIndex.value = currentPageArg;
  const res = await GetTransactionsByAddress(
    currentPageIndex.value - 1,
    pageSizeNumber.value,
    activeName.value,
    props.address as string
  );
  res.data.items.forEach((element) => {
    txsData.push(element);
  });
  total.value = res.data.total;
};

const handleInternalSizeChange = async (pageSizeArg: number) => {
  internalTxsData.length = 0;
  currentPageIndex.value = 1;
  pageSizeNumber.value = pageSizeArg;
  const res = await GetInternalTransactionsByAddress(
    currentPageIndex.value - 1,
    pageSizeNumber.value,
    props.address as string
  );
  res.data.items.forEach((element) => {
    internalTxsData.push(element);
  });
  internalCount.value = res.data.total;
};

const handleInternalCurrentChange = async (currentPageArg: number) => {
  internalTxsData.length = 0;
  currentPageIndex.value = currentPageArg;
  const res = await GetInternalTransactionsByAddress(
    currentPageIndex.value - 1,
    pageSizeNumber.value,
    props.address as string
  );
  res.data.items.forEach((element) => {
    internalTxsData.push(element);
  });
  internalCount.value = res.data.total;
};
</script>
<style lang="less" scoped>
@import '../../css/style.css';
</style>

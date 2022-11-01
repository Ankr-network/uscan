<template lang="">
  <div>
    <div class="center-row">
      <h4 class="h4-title">
        Token <span class="small text-secondary">&nbsp;&nbsp;{{ props.address }}</span>
      </h4>
      &nbsp;
      <copy-icon :text="props.address"></copy-icon>
    </div>
    <div>
      <el-row :gutter="20">
        <el-col :span="12">
          <div>
            <el-card class="box-card-address">
              <template #header>
                <div class="card-header">
                  <span>Overview</span>
                </div>
              </template>
              <div class="card-content">
                <el-row>
                  <el-col :span="9">Max Total Supply:</el-col>
                  <el-col :span="15">{{ totalSupply }}</el-col>
                </el-row>
                <el-row>
                  <el-col :span="9">Holders:</el-col>
                  <el-col :span="15">{{ holdersTotal }}</el-col>
                </el-row>
                <el-row>
                  <el-col :span="9">Transfers:</el-col>
                  <el-col :span="15">{{ transfersTotal }}</el-col>
                </el-row>
              </div>
            </el-card>
          </div>
        </el-col>
        <el-col :span="12">
          <div>
            <el-card class="box-card-address">
              <template #header>
                <div class="card-header">
                  <span>Profile Summary</span>
                </div>
              </template>
              <div class="card-content">
                <el-row>
                  <el-col :span="9">Contract:</el-col>
                  <el-col :span="15">
                    <router-link :to="'/address/' + address">{{ address }}</router-link>
                  </el-col>
                </el-row>
                <div v-if="tokenType == 'erc20'">
                  <el-row>
                    <el-col :span="9">Decimals:</el-col>
                    <el-col :span="15">{{ decimals }}</el-col>
                  </el-row>
                </div>
              </div>
            </el-card>
          </div>
        </el-col>
      </el-row>
    </div>
    <div>
      <el-tabs v-model="activeName">
        <el-tab-pane label="Transactions" name="transactions">
          <generate-transactions
            :txsData="txsData"
            :headerData="headerDataTx"
            :loadStatus="isEmpty"
          ></generate-transactions>
          <div style="margin-top: 1%; display: flex; justify-content: center">
            <el-pagination
              small
              background
              :currentPage="currentPageIndexTx"
              :page-size="pageSizeNumberTx"
              :page-sizes="[10, 25, 50, 100]"
              layout="total, sizes, prev, pager, next, jumper"
              :total="transfersTotal"
              @size-change="handleSizeChangeTx"
              @current-change="handleCurrentChangeTx"
            />
          </div>
        </el-tab-pane>
        <el-tab-pane label="Holders" name="holders">
          <generate-holders :holdersData="holdersData" :headerData="headerDataHolder"></generate-holders>
          <div style="margin-top: 1%; display: flex; justify-content: center">
            <el-pagination
              small
              background
              :currentPage="currentPageIndexHolder"
              :page-size="pageSizeNumberHolder"
              :page-sizes="[10, 25, 50, 100]"
              layout="total, sizes, prev, pager, next, jumper"
              :total="transfersTotal"
              @size-change="handleSizeChangeHolder"
              @current-change="handleCurrentChangeHolder"
            />
          </div>
        </el-tab-pane>
        <el-tab-pane v-if="tokenType === 'erc721' || tokenType === 'erc1155'" label="Inventory" name="inventory">
          <generate-Inventory :address="props.address" :ercType="tokenType"></generate-Inventory>
        </el-tab-pane>
      </el-tabs>
    </div>
  </div>
</template>
<script lang="ts" setup>
import {
  GetTokenTransfersByAddress,
  GetTokenHoldersByAddress,
  GetTransactionsByToken,
} from '../../script/service/tokenService';
import {
  TokenErc20TransactionsHeaderList,
  TokenErcTransactionsHeaderList,
  TokeHolderHeaderList,
  TokeErcHolderHeaderList,
  TransactionDetail,
} from '../../script/model/transaction';
import { TokenTransfers, TokenHolder } from '../../script/model/token';
import { TableHeader } from '../../script/model/index';
import { GetAddressInfo } from '../../script/service/addressService';
import { ref, reactive } from 'vue';
import { getTitle } from '../../script/utils';

document.title = 'Token | The ' + getTitle + ' Explorer';

const props = defineProps({
  address: String,
});

const activeName = ref('transactions');
const currentPageIndexHolder = ref(1);
const pageSizeNumberHolder = ref(25);
const currentPageIndexTx = ref(1);
const pageSizeNumberTx = ref(25);
const headerDataTx: TableHeader[] = reactive([]);
const txsData: TransactionDetail[] = reactive([]);
const holdersData: TokenHolder[] = reactive([]);
const headerDataHolder: TableHeader[] = reactive([]);
const isEmpty = ref(true);

let tokenType = '';

const tokenTransfersRes = await GetTokenTransfersByAddress(props.address as string);
// console.log(tokenTransfersRes);

// eslint-disable-next-line guard-for-in
for (const i in tokenTransfersRes.data) {
  if (tokenTransfersRes.data[i as keyof TokenTransfers] != 0) {
    tokenType = i;
  }
}
// console.log('tokenType', tokenType);

if (tokenType == 'erc20') {
  headerDataTx.push(...TokenErc20TransactionsHeaderList);
  headerDataHolder.push(...TokeHolderHeaderList);
} else {
  headerDataTx.push(...TokenErcTransactionsHeaderList);
  headerDataHolder.push(...TokeErcHolderHeaderList);
}

const addressInfoRes = await GetAddressInfo(props.address as string);
const totalSupply = addressInfoRes.data.tokenTotalSupply;
const decimals = addressInfoRes.data.decimals;
// console.log(addressInfoRes);

const tokenHoldersByAddressRes = await GetTokenHoldersByAddress(
  props.address as string,
  tokenType,
  currentPageIndexHolder.value - 1,
  pageSizeNumberHolder.value
);
const holdersTotal = tokenHoldersByAddressRes.data.total;
// console.log('tokenHoldersByAddressRes', tokenHoldersByAddressRes);
tokenHoldersByAddressRes.data.items.forEach((element) => holdersData.push(element));
if (tokenHoldersByAddressRes.data.total == 0) {
  isEmpty.value = false;
}

const tokenTransactionRes = await GetTransactionsByToken(
  props.address as string,
  tokenType,
  currentPageIndexTx.value - 1,
  pageSizeNumberTx.value
);
const transfersTotal = tokenTransactionRes.data.total;
tokenTransactionRes.data.items.forEach((element) => {
  // console.log('element', element);
  txsData.push(element);
});
// console.log('tokenTransactionRes', tokenTransactionRes);

const handleSizeChangeTx = async (pageSizeArg: number) => {
  txsData.length = 0;
  currentPageIndexTx.value = 1;
  pageSizeNumberTx.value = pageSizeArg;
  const res = await GetTransactionsByToken(
    props.address as string,
    tokenType,
    currentPageIndexTx.value - 1,
    pageSizeNumberTx.value
  );
  res.data.items.forEach((element) => {
    txsData.push(element);
  });
};

const handleCurrentChangeTx = async (currentPageArg: number) => {
  txsData.length = 0;
  currentPageIndexTx.value = currentPageArg;
  const res = await GetTransactionsByToken(
    props.address as string,
    tokenType,
    currentPageIndexTx.value - 1,
    pageSizeNumberTx.value
  );
  res.data.items.forEach((element) => {
    txsData.push(element);
  });
};

const handleSizeChangeHolder = async (pageSizeArg: number) => {
  holdersData.length = 0;
  currentPageIndexHolder.value = 1;
  pageSizeNumberHolder.value = pageSizeArg;
  const res = await GetTokenHoldersByAddress(
    props.address as string,
    tokenType,
    currentPageIndexHolder.value - 1,
    pageSizeNumberHolder.value
  );
  res.data.items.forEach((element) => {
    holdersData.push(element);
  });
};

const handleCurrentChangeHolder = async (currentPageArg: number) => {
  holdersData.length = 0;
  currentPageIndexTx.value = currentPageArg;
  const res = await GetTokenHoldersByAddress(
    props.address as string,
    tokenType,
    currentPageIndexHolder.value - 1,
    pageSizeNumberHolder.value
  );
  res.data.items.forEach((element) => {
    holdersData.push(element);
  });
};
</script>
<style lang="less" scoped>
@import '../../css/style.css';
</style>

<template lang="">
  <div>
    <el-table class="table-border" :data="tableData" empty-text="loading..." :row-style="{ height: '75px' }">
      <el-table-column label="Latest Transactions" width="250">
        <template v-slot:default="scope">
          <el-row>
            <el-col :span="6">
              <div class="list-icon-circle">
                <p style="font-size: 15px; font-weight: bold">Tx</p>
              </div>
            </el-col>
            <el-col :span="18">
              <router-link :to="'/tx/' + scope.row.hash"> {{ scope.row.hash.slice(0, 15) + '...' }}</router-link>
              <div>{{ getAge(scope.row.createTime) }}</div>
            </el-col>
          </el-row>
        </template>
      </el-table-column>
      <el-table-column width="250">
        <template v-slot:default="scope">
          <div>
            <div>
              From
              <router-link :to="'/address/' + scope.row.from"> {{ scope.row.from.slice(0, 19) + '...' }}</router-link>
            </div>
            <div>
              To <router-link :to="'/address/' + scope.row.to"> {{ scope.row.to.slice(0, 19) + '...' }}</router-link>
            </div>
          </div>
        </template>
      </el-table-column>
      <el-table-column>
        <template v-slot:default="scope">
          <div>
            <el-tooltip effect="dark" content="gasPrice" placement="right">
              <div style="text-align: right">
                <!-- <el-tag type="info">{{ this.$wei2gwei(scope.row.gas) }} Gwei</el-tag> -->
                <el-tag type="info">
                  {{
                    ethers.utils.formatUnits((parseInt(scope.row.gas) * parseInt(scope.row.gasPrice)).toString(), 18)
                  }}
                  Eth
                  <!-- {{ scope.row.gas }} Eth -->
                </el-tag>
              </div>
            </el-tooltip>
          </div>
        </template>
      </el-table-column>
    </el-table>
    <div class="center">
      <el-button class="home-bottom-button" type="primary" plain @click="moveToTxs">View all Transactions</el-button>
    </div>
  </div>
</template>
<script lang="ts" setup>
import { GetHomeTransactions } from '../../script/service/transactionService';
import { getAge } from '../../script/utils';
import { ethers } from 'ethers';
import { useRouter } from 'vue-router';

const router = useRouter();
const res = await GetHomeTransactions();
const tableData = res.data.items;

const moveToTxs = () => {
  router.push('/txs/all');
};
</script>
<style lang="less" scoped>
@import '../../css/style.css';
</style>

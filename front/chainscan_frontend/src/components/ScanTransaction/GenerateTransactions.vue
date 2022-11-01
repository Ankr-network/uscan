<template lang="">
  <el-table class="table-border" :data="props.txsData" empty-text="loading..." :row-style="{ height: '50px' }">
    <template #empty>{{ emptyText }}</template>
    <el-table-column width="37px">
      <template v-slot:default="scope">
        <el-popover placement="right" title="Additional Info" :width="320" trigger="click">
          <template #reference>
            <el-button style="width: 5px" type="info" size="small" @click="scope.row.base = true">
              <el-icon><View /></el-icon>
            </el-button>
          </template>
          <div v-if="scope.row.base == true">
            <base-transaction-info :txHash="scope.row.hash"></base-transaction-info>
          </div>
        </el-popover>
      </template>
    </el-table-column>
    <el-table-column v-for="info in props.headerData" :key="info.key" :property="info.key" :label="info.label">
      <template v-slot:default="scope">
        <div v-if="scope.column.property == 'transactionHash'" style="width: 170px">
          <router-link :to="'/tx/' + scope.row[scope.column.property]">{{
            scope.row[scope.column.property].slice(0, 15) + '...'
          }}</router-link>
        </div>
        <div v-if="scope.column.property == 'hash'" style="width: 170px">
          <router-link :to="'/tx/' + scope.row[scope.column.property]">{{
            scope.row[scope.column.property].slice(0, 15) + '...'
          }}</router-link>
        </div>
        <div v-else-if="scope.column.property == 'blockNumber'">
          <router-link :to="'/block/' + parseInt(scope.row[scope.column.property])">{{
            parseInt(scope.row[scope.column.property])
          }}</router-link>
        </div>
        <div v-else-if="scope.column.property == 'createTime'">
          {{ getAge(scope.row[scope.column.property]) }}
        </div>
        <div v-else-if="scope.column.property == 'createdTime'">
          {{ getAge(scope.row[scope.column.property]) }}
        </div>
        <div v-else-if="scope.column.property == 'from'" style="width: 170px">
          <router-link :to="'/address/' + scope.row[scope.column.property]">{{
            scope.row[scope.column.property].slice(0, 15) + '...'
          }}</router-link>
        </div>
        <div v-else-if="scope.column.property == 'to'" style="width: 170px">
          <router-link :to="'/address/' + scope.row[scope.column.property]">{{
            scope.row[scope.column.property].slice(0, 15) + '...'
          }}</router-link>
        </div>
        <div v-else-if="scope.column.property == 'value'">
          {{ ethers.utils.formatUnits(scope.row[scope.column.property], 18) }}
        </div>
        <div v-else-if="scope.column.property == 'gas'" style="font-size: 11px">
          {{ ethers.utils.formatUnits(scope.row[scope.column.property], 18) }}
        </div>
        <div v-else-if="scope.column.property == 'contract'" style="width: 170px; font-size: 11px">
          <router-link :to="'/token/' + scope.row[scope.column.property]">{{
            scope.row[scope.column.property].slice(0, 15) + '...'
          }}</router-link>
        </div>
      </template>
    </el-table-column>
  </el-table>
</template>
<script lang="ts" setup>
import { TransactionDetail } from '../../script/model/transaction';
import { TableHeader } from '../../script/model/index';
import { getAge } from '../../script/utils';
import { ethers } from 'ethers';
import { ref, watchEffect } from 'vue';
import { View } from '@element-plus/icons-vue';

const emptyText = ref('loading...');
const props = defineProps({
  loadStatus: Boolean,
  txsData: {
    type: Array as () => Array<TransactionDetail>,
    require: true,
  },
  headerData: {
    type: Array as () => Array<TableHeader>,
    require: true,
  },
});
watchEffect(() => {
  if (!props.loadStatus) {
    emptyText.value = 'empty data';
  }
});

// console.log('txsData', props.txsData);
</script>
<style lang="less" scoped>
@import '../../css/style.css';
</style>

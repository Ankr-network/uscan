<template lang="">
  <el-table class="table-border" :data="props.txsData" empty-text="loading..." :row-style="{ height: '50px' }">
    <template #empty>{{ emptyText }}</template>
    <el-table-column v-for="info in props.headerData" :key="info.key" :property="info.key" :label="info.label">
      <template v-slot:default="scope">
        <div v-if="scope.column.property == 'transactionHash'" style="width: 170px">
          <router-link :to="'/tx/' + scope.row[scope.column.property]">{{
            scope.row[scope.column.property].slice(0, 15) + '...'
          }}</router-link>
        </div>
        <div v-else-if="scope.column.property == 'blockNumber'">
          <router-link :to="'/block/' + parseInt(scope.row[scope.column.property])">{{
            parseInt(scope.row[scope.column.property])
          }}</router-link>
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
        <div v-else-if="scope.column.property == 'amount'">{{ parseInt(scope.row[scope.column.property]) }}</div>
      </template>
    </el-table-column>
  </el-table>
</template>
<script lang="ts" setup>
import { TransactionDetail } from '../../script/model/transaction';
import { TableHeader } from '../../script/model/index';
import { ref } from 'vue';
import { getAge } from '../../script/utils';
// import { ethers } from 'ethers';
// import { View } from '@element-plus/icons-vue';

const emptyText = ref('loading...');
const props = defineProps({
  txsData: {
    type: Array as () => Array<TransactionDetail>,
    require: true,
  },
  headerData: {
    type: Array as () => Array<TableHeader>,
    require: true,
  },
});

// console.log('txsData', props.txsData);
</script>
<style lang="less" scoped>
@import '../../css/style.css';
</style>

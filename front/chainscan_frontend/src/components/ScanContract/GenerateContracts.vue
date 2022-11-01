<template lang="">
  <el-table class="table-border" :data="props.contractsData" empty-text="loading..." :row-style="{ height: '50px' }">
    <el-table-column v-for="info in props.headerData" :key="info.key" :property="info.key" :label="info.label">
      <template v-slot:default="scope">
        <div v-if="scope.column.property == 'owner'">
          <router-link :to="'/address/' + scope.row[scope.column.property]">
            {{ scope.row[scope.column.property].slice(0, 25) + '...' }}
          </router-link>
        </div>
        <div v-else-if="scope.column.property == 'creator'">
          <router-link :to="'/address/' + scope.row[scope.column.property]">
            {{ scope.row[scope.column.property] ? scope.row[scope.column.property].slice(0, 25) + '...' : '' }}
          </router-link>
        </div>
      </template>
    </el-table-column>
  </el-table>
</template>
<script lang="ts" setup>
import { ContractContent } from '../../script/model/contract';
import { TableHeader } from '../../script/model/index';

const props = defineProps({
  contractsData: {
    type: Array as () => Array<ContractContent>,
    require: true,
  },
  headerData: {
    type: Array as () => Array<TableHeader>,
    require: true,
  },
});
</script>
<style lang="less" scoped>
@import '../../css/style.css';
</style>

<template lang="">
  <div>
    <el-table class="table-border" :data="props.holdersData" empty-text="loading..." :row-style="{ height: '50px' }">
      <el-table-column v-for="info in props.headerData" :key="info.key" :property="info.key" :label="info.label">
        <template v-slot:default="scope">
          <div v-if="scope.column.property == 'owner'">
            <router-link :to="'/address/' + scope.row[scope.column.property]">{{
              scope.row[scope.column.property]
            }}</router-link>
          </div>
          <div v-else-if="scope.column.property == 'quantity'" style="width: 180px">
            {{ BigInt(parseInt(scope.row[scope.column.property])) }}
          </div>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>
<script lang="ts" setup>
import { TokenHolder } from '../../script/model/token';
import { TableHeader } from '../../script/model/index';

const props = defineProps({
  holdersData: {
    type: Array as () => Array<TokenHolder>,
    require: true,
  },
  headerData: {
    type: Array as () => Array<TableHeader>,
    require: true,
  },
});
</script>
<style lang=""></style>

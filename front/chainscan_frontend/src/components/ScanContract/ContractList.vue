<template lang="">
  <div>
    <div class="center-row">
      <h4 class="h4-title">Contracts</h4>

      <el-button style="margin: 10px; font-weight: bold" color="#DEE1E4" size="small">
        {{ props.contractType }}
      </el-button>
    </div>

    <div>
      <generate-contracts :contractsData="contractsData" :headerData="ContractsHeaderList"></generate-contracts>
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
  </div>
</template>
<script lang="ts" setup>
import { getTitle } from '../../script/utils';
import { ref, reactive, watchEffect } from 'vue';
import { ContractsHeaderList, ContractContent } from '../../script/model/contract';
import { GetContracts } from '../../script/service/contractService';

document.title = 'Contracts | The ' + getTitle + ' Explorer';

const props = defineProps({
  contractType: String,
});

const currentPageIndex = ref(1);
const pageSizeNumber = ref(25);
const contractsData: ContractContent[] = reactive([]);
const total = ref(0);

watchEffect(async () => {
  const res = await GetContracts(props.contractType as string, currentPageIndex.value - 1, pageSizeNumber.value);
  res.data.items.forEach((element) => {
    contractsData.push(element);
  });
  total.value = res.data.total;
});

const handleSizeChange = async (pageSizeArg: number) => {
  contractsData.length = 0;
  currentPageIndex.value = 1;
  pageSizeNumber.value = pageSizeArg;
  const res = await GetContracts(props.contractType as string, currentPageIndex.value - 1, pageSizeNumber.value);
  res.data.items.forEach((element) => {
    contractsData.push(element);
  });
  total.value = res.data.total;
};

const handleCurrentChange = async (currentPageArg: number) => {
  contractsData.length = 0;
  currentPageIndex.value = currentPageArg;
  const res = await GetContracts(props.contractType as string, currentPageIndex.value - 1, pageSizeNumber.value);
  res.data.items.forEach((element) => {
    contractsData.push(element);
  });
  total.value = res.data.total;
};
</script>
<style lang="less" scoped>
@import '../../css/style.css';
</style>

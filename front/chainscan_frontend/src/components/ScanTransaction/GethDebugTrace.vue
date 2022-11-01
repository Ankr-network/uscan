<template lang="">
  <div>
    <div style="margin-bottom: -20px">
      <h4 class="h4-title">Geth VM Trace Transaction</h4>
      <p style="color: #77838f !important; ont-size: 80%; font-weight: 400">
        GETH Trace for Txn Hash <router-link :to="'/tx/' + route.query.txhash">{{ route.query.txhash }}</router-link>
      </p>
      <h4 class="tx-sub-title">(Showing the last {{ total }} records only)</h4>
    </div>

    <el-tabs style="width: 100%; margin-top: 30px">
      <el-tab-pane>
        <template #label>
          <p class="subtitle1">Raw traces</p>
        </template>
        <div v-if="route.query.type == 'tracetx'">
          <el-table :data="resTraceList" style="width: 100%" empty-text="loading...">
            <el-table-column type="index" width="100" label="Step" />
            <el-table-column prop="pc" label="PC" />
            <el-table-column prop="op" label="Operation" />
            <el-table-column prop="gas" label="Gas" />
            <el-table-column prop="gasCost" label="GasCost" />
            <el-table-column prop="depth" label="Depth" />
          </el-table>
        </div>
        <div v-if="route.query.type == 'tracetx2'">
          <textarea class="byte-codes-text" style="margin: 0px" rows="12" v-model="resTrace2"> </textarea>
        </div>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>
<script lang="ts" setup>
import { useRoute } from 'vue-router';
import { ref, reactive, watchEffect } from 'vue';
import { GetGethDebugTrace } from '../../script/service/transactionService';
import { GethDebugTrace } from '../../script/model/transaction';

const route = useRoute();
const resTraceList: GethDebugTrace[] = reactive([]);
const resTrace2 = ref('');
const total = ref(0);

watchEffect(async () => {
  resTraceList.length = 0;

  if (route.query.txhash as string) {
    const getGethDebugTraceRes = await GetGethDebugTrace(route.query.txhash as string, route.query.type as string);
    // console.log('getGethDebugTrace', getGethDebugTraceRes.data);

    total.value = getGethDebugTraceRes.data.logNum;

    if (route.query.type == 'tracetx') {
      const resList: GethDebugTrace[] = JSON.parse(getGethDebugTraceRes.data.res);
      resList.forEach((element) => {
        resTraceList.push(element);
      });
    } else {
      resTrace2.value = JSON.stringify(getGethDebugTraceRes.data.res, null, 2);
    }
  }
  // resTrace2.value = 'hahahahah';
});

// console.log('params', route.query);
</script>
<style lang="less" scoped>
@import '../../css/style.css';
</style>

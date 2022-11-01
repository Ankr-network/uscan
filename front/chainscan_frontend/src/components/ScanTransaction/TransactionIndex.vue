<template lang="">
  <div>
    <el-row>
      <el-col :span="12">
        <div>
          <h4 class="h4-title">Transactions Details</h4>
        </div>
      </el-col>
      <el-col :span="12" class="more-info">
        <div class="more-button">
          <el-dropdown @command="handleCommand">
            <el-button style="width: 5px" type="info" size="small">
              <el-icon><MoreFilled /></el-icon>
            </el-button>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item command="tracetx">Geth Debug Trace</el-dropdown-item>
                <el-dropdown-item command="tracetx2">Geth Debug Trace_2</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>
      </el-col>
    </el-row>

    <el-tabs v-model="activeName">
      <el-tab-pane label="Overview" name="txs">
        <transaction-overview :txHash="props.txHash"></transaction-overview>
      </el-tab-pane>
      <el-tab-pane v-if="logCount != 0" name="logs">
        <template #label>
          <span>Logs({{ logCount }})</span>
        </template>
        <transaction-logs :transactionLogs="transactionLogsData"></transaction-logs>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>
<script lang="ts" setup>
import { ref } from 'vue';
import { MoreFilled } from '@element-plus/icons-vue';
import { GetTxLog } from '../../script/service/transactionService';
import { useRouter } from 'vue-router';

const props = defineProps({
  txHash: String,
});
const activeName = ref('txs');
const logCount = ref(0);

const res = await GetTxLog(props.txHash as string);
logCount.value = res.data.total;
const transactionLogsData = res.data.items;

const router = useRouter();

const handleCommand = (command: string | number | object) => {
  console.log('command', command);
  if (command == 'tracetx2') {
    router.push((('/vmtrace?txhash=' + props.txHash) as string) + '&type=' + command);
  } else if (command == 'tracetx') {
    router.push((('/vmtrace?txhash=' + props.txHash) as string) + '&type=' + command);
  }
};
</script>
<style lang="less" scoped>
@import '../../css/style.css';
</style>

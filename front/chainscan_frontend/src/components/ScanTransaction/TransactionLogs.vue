<template lang="">
  <div class="log">
    <div class="content">
      <p>Transaction Receipt Event Logs</p>
      <div class="log-content">
        <div v-for="(log, index) in props.transactionLogs" :key="index">
          <div>
            <el-row>
              <el-col :span="2">
                <div class="icon-circle">{{ log.logIndex }}</div>
              </el-col>
              <el-col :span="22">
                <div>
                  <el-row class="log-row">
                    <el-col :span="2">
                      <div style="font-size: 15px; font-weight: bold; color: #4a4f55">Address</div>
                    </el-col>
                    <el-col :span="22">
                      <div>
                        <router-link :to="'/address/' + log.address"> {{ log.address }} </router-link>
                      </div>
                    </el-col>
                  </el-row>
                  <div>
                    <div v-for="(topic, index) in log.topics" :key="index">
                      <el-row class="log-row">
                        <el-col :span="2"><div v-if="index == 0">Topics</div></el-col>
                        <el-col :span="22">
                          <div class="center-row">
                            <div class="topic-index">{{ index }}</div>
                            <div style="margin-left: 10px">{{ topic }}</div>
                          </div>
                        </el-col>
                      </el-row>
                    </div>
                  </div>
                  <el-row class="log-data-row">
                    <el-col :span="2">Data</el-col>
                    <el-col :span="22" style="word-break: break-all">
                      <div class="center-row" style="background-color: #f8f9fa; width: 90%">
                        <div style="margin: 20px">{{ log.data }}</div>
                      </div>
                    </el-col>
                  </el-row>
                </div>
              </el-col>
            </el-row>
            <div>
              <el-divider />
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts" setup>
import { TransactionLog } from '../../script/model/transaction';

const props = defineProps({
  transactionLogs: {
    type: Array as () => Array<TransactionLog>,
    require: true,
  },
});
</script>
<style lang="less" scoped>
@import '../../css/style.css';
.log {
  width: 100%;
  margin-bottom: 60px;
  background-color: white;
  border-radius: 0.35rem;
}
.log-content {
  margin-top: 30px;
  margin-left: 15px;
}
.content {
  margin-top: 10px;
  margin-bottom: 10px;
  margin-left: 17px;
}
.icon-circle {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 45px;
  width: 45px;
  background-color: rgb(238, 249, 246);
  color: #00c9a7;
  border-radius: 50%;
}
.log-row {
  display: flex;
  align-items: center;
  height: 42px;
}

.log-data-row {
  display: flex;
  align-items: center;
}
.topic-index {
  color: #77838f;
  background-color: rgba(119, 131, 143, 0.1);
  padding: 0.2rem 0.5rem;
  border-radius: 30%;
}
</style>

<template lang="">
  <div>
    <div style="text-align: center">
      <h3 class="h3-title">Download Data({{ title }})</h3>
      <div style="display: flex; justify-content: center">
        <p style="color: #77838f; width: 50%">
          The information you requested can be downloaded from this page. But before continuing please verify that you
          are not a robot by completing the captcha below.
        </p>
      </div>
      <div style="display: flex; justify-content: center">
        <el-card class="box-card">
          <p>Export the earliest 5000 records starting from</p>
          <el-date-picker
            v-model="value"
            type="datetimerange"
            range-separator="To"
            start-placeholder="Start date"
            end-placeholder="End date"
          />
          <div style="margin-top: 20px">
            <el-button type="primary" @click="exportData">Download</el-button>
          </div>
        </el-card>
      </div>
      <div style="display: flex; justify-content: center">
        <p style="color: #77838f; width: 50%">Tip: To avoid the 30 secs timeout, please select a small Date range</p>
      </div>
    </div>
  </div>
</template>
<script lang="ts" setup>
import { ref, watchEffect } from 'vue';
import axios from 'axios';
import { useRoute } from 'vue-router';
import { getTitle } from '../../script/utils';

const route = useRoute();

const value = ref([0, 0]);
const type = ref(route.query.type as string);
const address = ref(route.query.a as string);
const title = ref('Transactions');

document.title = 'Export Data | The ' + getTitle + ' Explorer';

watchEffect(() => {
  console.log(type.value);
  switch (type.value) {
    case 'txns-erc20':
      title.value = 'ERC-20 Token Transfers';
      break;
    case 'txns-erc721':
      title.value = 'ERC-721 Token Transfers';
      break;
    case 'txns-erc721':
      title.value = 'ERC-721 Token Transfers';
      break;
    case 'txns-internal':
      title.value = 'Internal Transactions';
      break;
  }
});

const exportData = async () => {
  if (value.value[0] == 0 || value.value[1] == 0) {
    value.value[0] = Date.now() - 7 * 24 * 60 * 60 * 1000;
    value.value[1] = Date.now();
  }

  const beginTime = parseInt((value.value[0] / 1000) as any as string);
  const endTime = parseInt((value.value[1] / 1000) as any as string);

  // console.log('beginTime', beginTime, 'endTime', endTime, 'address', address.value, 'type', type.value);

  axios({
    method: 'get',
    url:
      import.meta.env.VITE_BASE_URL +
      '/v1/accounts/' +
      address.value +
      '/' +
      type.value +
      '/download' +
      '?beginTime=' +
      beginTime +
      '&endTime=' +
      endTime,
    responseType: 'blob',
  })
    .then((res) => {
      console.log(res);
      const data = res.data;
      if (!data) {
        return;
      }
      const url = window.URL.createObjectURL(new Blob([data]));
      const a = document.createElement('a');
      a.style.display = 'none';
      a.href = url;
      a.setAttribute('download', 'export-' + address.value + '.xls');
      document.body.appendChild(a);
      a.click();
      window.URL.revokeObjectURL(a.href);
      document.body.removeChild(a);
    })
    .catch((error) => {
      console.log(error);
    });
};
</script>
<style lang="less" scoped>
@import '../../css/style.css';
.box-card {
  width: 50%;
  margin-top: 30px;
  margin-bottom: 30px;
}
</style>

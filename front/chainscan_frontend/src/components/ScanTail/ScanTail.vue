<template lang="">
  <div class="footer-container">
    <div class="item" style="width: 400px">
      <div style="width: 340px">
        <p style="font-size: 20px; color: white">Powered by Ankr</p>
        <p style="font-size: 13px; color: white">
          {{ getTitle }}scan is a Block Explorer and Analytics Platform for {{ getTitle }}, a decentralized smart
          contracts platform.
        </p>
      </div>
    </div>
    <div class="item">
      <div style="width: 200px">
        <p style="font-size: 14px; color: white">Products</p>
        <el-divider border-style="dotted" />
        <div class="item-content">
          <p>RPC Service</p>
          <p>Advanced APIs/SDKs</p>
          <p>Gaming SDKs</p>
          <p>App Chains</p>
        </div>
      </div>
    </div>
    <div class="item">
      <div style="width: 200px">
        <p style="font-size: 14px; color: white">Developers</p>
        <el-divider border-style="dotted" />
        <div class="item-content">
          <p>Docs</p>
          <p>Grants & Programs</p>
          <p>Hash node</p>
          <p>Github</p>
        </div>
      </div>
    </div>
    <div class="item">
      <div style="width: 200px">
        <p style="font-size: 14px; color: white">Company</p>
        <el-divider border-style="dotted" />
        <div class="item-content">
          <p>About Us</p>
          <p>Blog</p>
          <p>Careers</p>
          <p>Brand Assets</p>
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts" setup>
import { getTitle } from '../../script/utils';
</script>
<style lang="less" scoped>
@import '../../css/style.css';
.footer-container {
  margin-top: 13px;
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
}
.item {
  width: 50px;
  // background-color: red;
  margin: 5px;
  width: 250px;
  height: 200px;
}

.item-content {
  display: flex;
  flex-direction: column;
  margin-top: 10px;
}

.item-content > p {
  margin-top: 1px;
  font-size: 14px;
  color: #b3b6c4;
}
</style>

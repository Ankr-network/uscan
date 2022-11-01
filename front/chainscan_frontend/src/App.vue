<template>
  <div class="index">
    <el-container>
      <el-header>
        <div :class="isHome ? 'home-header' : 'info-header'">
          <component :is="isHome ? HomeHeaderVue : InfoHeaderVue"></component>
        </div>
      </el-header>
      <el-main>
        <Suspense><router-view class="content"></router-view></Suspense>
      </el-main>
      <div class="footer">
        <el-footer> <scan-tail></scan-tail> </el-footer>
      </div>
    </el-container>
  </div>
</template>

<script lang="ts" setup>
import { useRoute } from 'vue-router';
import { watch, ref } from 'vue';
import HomeHeaderVue from './components/ScanHeader/HomeHeader.vue';
import InfoHeaderVue from './components/ScanHeader/InfoHeader.vue';
import { getTitle } from './script/utils';

document.title = 'The ' + getTitle + ' Explorer';

const route = useRoute();
const isHome = ref(false);
// console.log('index', route.path);
if (route.path === '/') {
  isHome.value = true;
}

watch(
  () => route.path,
  (val) => {
    // console.log('val', val);
    if (val === '/') {
      isHome.value = true;
    } else {
      isHome.value = false;
    }
  }
);
</script>

<style lang="less">
.index {
  height: 100%;
  width: 100%;
  position: absolute;
  top: 0;
  left: 0;
}

.content {
  max-width: 1350px;
  margin: 0 auto;
  margin-top: 15px;
}
.el-header {
  padding: 0;
  height: 100%;
}

.el-main {
  min-height: 100%;
  width: 100%;
  padding: 0;
}

.home-header {
  background-color: #263258;
  height: 250px;
}

.info-header {
  background-color: white;
  height: 125px;
}

.el-footer {
  display: flex;
  height: 250px;
  max-width: 1350px;
  background-color: transparent;
  justify-content: center;
  margin: 0 auto;
}

.footer {
  background-color: #263258;
  margin-top: 2%;
  bottom: 0;
}

@media screen and (max-width: 500px) {
  .el-footer {
    height: 1000px;
  }
  .footer {
    margin-top: 10%;
  }
  .info-header {
    height: 170px;
  }
}

.el-container {
  width: 100%;
  min-height: 100vh;
  // display: flex;
  // flex-direction: column;
  // justify-content: space-between;
}
</style>

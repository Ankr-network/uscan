<template lang="">
  <div class="home-content">
    <div class="statistic">
      <scan-statistic></scan-statistic>
    </div>

    <div class="home-left">
      <Suspense>
        <scan-home-blocks></scan-home-blocks>
      </Suspense>
    </div>
    <div class="home-right">
      <Suspense>
        <scan-home-transactions></scan-home-transactions>
      </Suspense>
    </div>
  </div>
</template>
<script lang="ts" setup>
import { getTitle } from '../../script/utils';
document.title = 'Home | The ' + getTitle + ' Explorer';
</script>
<style>
.home-content {
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
}

.home-right {
  justify-content: flex-end;
  margin-left: auto;
  min-width: 665px;
}

.home-left {
  min-width: 665px;
}

@media screen and (max-width: 500px) {
  .home-right {
    min-width: 200px;
  }
  .home-left {
    min-width: 200px;
  }
}

.statistic {
  margin-bottom: 15px;
}
</style>

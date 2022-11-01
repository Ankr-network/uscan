<template lang="">
  <div>Searching... {{ searchText }}</div>
</template>
<script lang="ts" setup>
import { SearchByType } from '../../script/service/searchService';
import { useRouter } from 'vue-router';

const router = useRouter();

const props = defineProps({
  searchText: String,
});

interface LinkItem {
  value: string;
  link: string;
}

const querySearch = async (queryString: string) => {
  let searchRes: LinkItem = {} as LinkItem;
  const typeMap = {
    1: { display: 'Null', route: '' },
    2: { display: 'Address', route: '/address/' },
    3: { display: 'Block', route: '/block/' },
    4: { display: 'Transaction', route: '/tx/' },
  };
  const resList: LinkItem[] = [];
  if (queryString.trim() !== '') {
    if (queryString.trim().length == 42 || !isNaN(Number(queryString.trim()))) {
      const searchTypeRes = await SearchByType(1, queryString.trim());
      // console.log('queryResType.value', searchTypeRes.data.type);
      if (searchTypeRes.data.type == 1 || searchTypeRes.data.type == 0) {
        resList.push({ value: 'Not Found', link: '404' });
        searchRes = {} as LinkItem;
      } else {
        const res: LinkItem = {
          value: typeMap[searchTypeRes.data.type as keyof typeof typeMap].display + ' : ' + queryString.trim(),
          link: typeMap[searchTypeRes.data.type as keyof typeof typeMap].route + queryString.trim(),
        };
        resList.push(res);
        searchRes = res;
      }
    } else {
      searchRes = {} as LinkItem;
    }
  } else {
    searchRes = {} as LinkItem;
  }
  // console.log('searchRes', searchRes);
  if (searchRes.value) {
    router.push(searchRes.link);
  } else {
    router.push('/search/not-found');
  }
};

querySearch(props.searchText as string);
</script>
<style lang=""></style>

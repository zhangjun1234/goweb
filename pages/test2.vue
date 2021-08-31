<template>
  <div>
    <div v-for="(item, index) in arr" :key="index">{{ item }}</div>
    
    <div >1</div>
    
    <div >2</div>
    
    <div >3</div>
    
    <div >4</div>
    
    <div >5</div>
    
  </div>
</template>

<script>
export default {
  data() {
    return {
      arr: [1, 3, 4, 5, 6],
    };
  },
};
</script>

<template>
  <div>
    <div v-for="(item, index) in arr" :key="index">{{ item }}</div>
    {[range .arr ]}
    <div >{[.]}</div>
    {[end]}
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

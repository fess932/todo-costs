<template>
  <div
    class="
      p-3
      max-w-xs
      bg-white
      rounded-xl
      border-2
      flex flex-col
      text-center
      cursor-text
    "
  >
    
    <h3 class="text-lg font-medium text-black">{{ costName }}</h3>
    <p class="text-gray-500">{{ value }} hour</p>
  </div>
</template>

<script>
export default {
  props: {
    costName: '',
    value: 0,
  },
}
</script>

<style>
</style>
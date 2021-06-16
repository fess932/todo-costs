<template>
  <div
    class="
      transition
      duration-300
      ease-in-out
      p-6
      max-w-xs
      bg-white
      rounded-xl
      border-2
      flex flex-col
      text-center
      cursor-pointer
    "
    v-bind:class="{
      'bg-purple-500 cursor-text': isActive,
    }"
  >
    <h3 class="text-xl font-medium text-black">{{ taskName }}</h3>
    <p v-bind:class="[isActive ? 'text-white' : 'text-gray-500']">
      Total {{ total }}hours
    </p>
    <p v-bind:class="[isActive ? 'text-white' : 'text-gray-500']">
      Average{{ average }}
    </p>
  </div>
</template>

<script>
export default {
  props: {
    isActive: false,
    taskName: '',
    total: 0,
    average: 0,
  },
}
</script>

<style>
</style>
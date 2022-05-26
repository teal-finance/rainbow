<template>
  <tr>
    <th colspan="3" class="border-none"></th>
    <th colspan="6" class="text-center">CALL</th>
    <th class="border-none"></th>
    <th colspan="6" class="text-center">PUT</th>
  </tr>
</template>

<style lang="sass" scoped>
tr:nth-child(1)
  th
    @apply pb-3
tr:nth-child(2)
  th:nth-child(2), th:nth-child(5)
    @apply secondary p-3
  th:nth-child(3), th:nth-child(6)
    @apply bg-rainbow-800 text-white
</style>

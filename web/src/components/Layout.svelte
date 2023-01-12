<script lang="ts">
  import Search from "./Search.svelte";
</script>

<div class="container mx-auto mt-10">
  <div>
    <h1 class="text-4xl font-bold text-gray-700">
      <img
        src="/icons/notification.svg"
        alt=""
        width="30"
        height="30"
        align="left"
      /> Notify
    </h1>
  </div>
  <div class="flex justify-between">
    <Search />
    <div class="pt-5">
      <button
        class="p-2 rounded-md bg-blue-500 border-0 text-white hover:bg-blue-400"
        >Add Service +</button
      >
    </div>
  </div>
  <div class="text-gray-600 text-sm">
    Slack <input type="checkbox" />
    Email <input type="checkbox" />
    TWillio <input type="checkbox" />
  </div>
  <div class="shadow-md rounded-md mt-2">
    <slot />
  </div>
</div>

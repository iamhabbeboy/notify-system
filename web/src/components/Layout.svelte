<script lang="ts">
  import { modal } from "../store/modal";
  import Filter from "./Filter.svelte";

  const handleModalOpen = () => {
    modal.update(value => value = true)
  }
</script>

<div class="container mx-auto mt-10">
  <div>
    <h1 class="text-4xl font-bold text-blue-900">
      <img
        src="/icons/notification.svg"
        alt=""
        width="30"
        height="30"
        align="left"
      /> Alert Notify
    </h1>
  </div>
  <div class="flex justify-between">
    <Filter />
    <div class="pt-5">
      <button on:click={handleModalOpen}
        class="py-2 px-5 rounded-md bg-blue-900  border-0 text-white hover:bg-blue-800"
        >Service +</button
      >
    </div>
  </div>

  <div class="shadow-md rounded-md mt-2 my-3 overflow-y mb-20">
    <slot />
  </div>
</div>

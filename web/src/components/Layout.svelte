<script lang="ts">
  import Filter from "./Filter.svelte";
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
      /> Notify
    </h1>
  </div>
  <div class="flex justify-between">
    <Filter />
    <div class="pt-5">
      <button
        class="py-2 px-5 rounded-md bg-blue-900  border-0 text-white hover:bg-blue-800"
        >Service +</button
      >
    </div>
  </div>
  <div class="bg-white my-5 p-5 rounded-md">
    <div class="flex w-3/3">
      <div class="p-3">
        <p class="text-sm">
          Lorem ipsum dolor sit amet consectetur adipisicing elit. Iure,
          corporis! Aliquam repellendus omnis aspernatur totam ab corporis nobis
          necessitatibus, labore eaque quo voluptatibus, sed maxime maiores ex
          optio error corrupti.
        </p>
      </div>
      <div class="w-full">
        <h1 class="font-bold ">Setting</h1>
        <hr class="mb-3"/>
        <div>
          <div >
            <input
              type="text"
              placeholder="Slack"
              class="border-2 mb-3 w-full border-gray-300 bg-white h-10 px-5 pr-5 rounded-lg text-sm focus:outline-none"
            />
          </div>
          <div>
            <textarea
              placeholder="Config"
              class="border-2 h-100 w-full border-gray-300 bg-white h-10 px-5 pr-5 rounded-lg text-sm focus:outline-none"
            ></textarea>
          </div>
        </div>
      </div>
    </div>
  </div>
  <div class="shadow-md rounded-md mt-2 my-3">
    <slot />
  </div>
</div>

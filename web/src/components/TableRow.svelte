<script lang="ts">
  import type { LogType } from "../InterfaceTypes/LogType";
  export let payload: LogType;
  let elem;

  const handleModal = () => {
    elem = document.getElementById("myModal").showModal();
  };
</script>

<tr class="hover:bg-gray-50">
  <th class="flex gap-3 px-6 py-4 font-normal text-gray-900">
    <div class="text-sm">
      <div class="font-medium text-gray-700 flex">
        <span>
          <img
            src="/icons/{payload.serviceName}.svg"
            width="20"
            height="20"
            alt=""
          /></span
        >
        {payload.serviceName}
      </div>
    </div>
  </th>
  <td class="px-6 py-4">{payload.environment}</td>
  <!-- <td class="px-6 py-4">
    <span
      class="inline-flex items-center gap-1 rounded-full bg-green-50 px-2 py-1 text-xs font-semibold text-green-600"
    >
      <span class="h-1.5 w-1.5 rounded-full bg-green-600" />
      Delivered
    </span>
  </td> -->
  <td class="px-6 py-4">
    <p>{payload.message}</p>
  </td>
  <td class="px-6 py-4">
    <p>{payload.date}</p>
  </td>
  <td class="px-6 py-4">
    <div class="flex justify-end gap-4">
      <!-- <button id="btn" class="py-3 px-10 bg-gray-800 text-white rounded text shadow-xl">Open</button> -->
      <button on:click={handleModal}>
        <svg
          xmlns="http://www.w3.org/2000/svg"
          class="icon icon-tabler icon-tabler-eyeglass"
          width="24"
          height="24"
          viewBox="0 0 24 24"
          stroke-width="1.5"
          stroke="#597e8d"
          fill="none"
          stroke-linecap="round"
          stroke-linejoin="round"
        >
          <path stroke="none" d="M0 0h24v24H0z" fill="none" />
          <path d="M8 4h-2l-3 10" />
          <path d="M16 4h2l3 10" />
          <line x1="10" y1="16" x2="14" y2="16" />
          <path d="M21 16.5a3.5 3.5 0 0 1 -7 0v-2.5h7v2.5" />
          <path d="M10 16.5a3.5 3.5 0 0 1 -7 0v-2.5h7v2.5" />
        </svg>
      </button>
    </div>
  </td>
</tr>

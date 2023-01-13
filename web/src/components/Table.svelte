<script lang="ts">
      import TableRow from "./TableRow.svelte";
import logs from "../store/logs";
</script>

<table
    class="w-full border-collapse bg-white text-left text-sm text-gray-500"
  >
    <thead class="bg-gray-50">
      <tr>
        <th scope="col" class="px-6 py-4 font-medium text-gray-900">Service</th>
        <th scope="col" class="px-6 py-4 font-medium text-gray-900"
          >Environment</th
        >
        <th scope="col" class="px-6 py-4 font-medium text-gray-900">Status</th>
        <th scope="col" class="px-6 py-4 font-medium text-gray-900">Message</th>
        <th scope="col" class="px-6 py-4 font-medium text-gray-900">Date</th>
        <th scope="col" class="px-6 py-4 font-medium text-gray-900" />
      </tr>
    </thead>
    <tbody class="divide-y divide-gray-100 border-t border-gray-100">
      {#each $logs as log}
        <TableRow payload={log} />
      {/each}
    </tbody>
  </table>
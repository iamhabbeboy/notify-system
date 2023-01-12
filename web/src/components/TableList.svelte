<script>
  import { onMount } from "svelte";
  import TableRow from "./TableRow.svelte";
  // import { LogType } from "../types/LogType.ts";
  let logs;

  // onMount(() => {
     logs = [
     {
      id: 1,
       serviceName: "slack",
       environment: "production",
       message: "error occured at line 1092",
       date: "2023-03-02",
     },
     {
      id: 2,
       serviceName: "email",
       environment: "production",
       message: "error occured at line 1092",
       date: "2023-03-02",
     }
    ]
  // })
</script>
<div class="overflow-hidden rounded-lg border border-gray-200 shadow-md">
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
        <th scope="col" class="px-6 py-4 font-medium text-gray-900" />
      </tr>
    </thead>
    <tbody class="divide-y divide-gray-100 border-t border-gray-100">
      {#each logs as log (log.id)}
        <TableRow payload={log}/>
      {/each}
    </tbody>
  </table>
</div>

<script>
  import { onMount } from "svelte";
  import Table from "./Table.svelte";
  import logs from "../store/logs";
  import {BASE_URL} from "../config/config";
  
  let isLoading = false
  let logStatus = "false"
  onMount(async () => {
    let logList = await fetch(`${BASE_URL}/logs`).then((data) => data.json());
    logStatus = logList.status
    if(logStatus === "true") {
      logs.addLog(logList.data)
    }
    isLoading = true;
  });
</script>

<div class="overflow-hidden rounded-lg border bg-white border-gray-200 shadow-md">
  {#if isLoading === false || logStatus === "false"}
    <div class="p-10 text-center text-gray-500 flex place-items-center">
      <img src="/icons/gear.svg" alt="" width="100" height="100"/>
      <h1 class="text-2xl">{logStatus === "false" ? "Error Occured while retrieving your data." : "Please wait, while we retrieve your data."}</h1>
    </div>
  {:else}
    {#if $logs.length === 0}
    <div class="p-10 text-gray-500">
      <h1 class="text-2xl">No log record available</h1>
    </div>
    {:else}
      <Table />
    {/if}
  {/if}
</div>

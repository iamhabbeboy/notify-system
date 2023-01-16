<script lang="ts">
  import { onMount } from "svelte";
  import { BASE_URL } from "../config/config";
  let services = [
    {
      name: "slack",
      config: {
        webhook: "http://",
        channel: "CHANNEL_ID",
      },
    },
    {
      name: "email",
      config: {
        host: "smtp.gmail.com",
        sender_address: "iamhabbeboy@gmail.com",
        username: "iamhabbeboy",
        password: "password",
      },
    },
  ];
  //   onMount(async () => {
  //     const response = await fetch(`${BASE_URL}/setting`).then((data) =>
  //       data.json()
  //     );
  //     services = response;
  //     console.log(services);
  //   });
</script>

<div>
  {#each services as service}
    <div>
        <label for="" class="text-left">{service.name}</label>
      <textarea
        placeholder="Config"
        style="height: 200px"
        class="border-2 w-full border-gray-300 bg-white p-5 rounded-lg text-sm focus:outline-none"
      >{JSON.stringify(service.config)}</textarea>
      
    </div>
  {/each}
  <button class="py-2 px-5 rounded-md bg-blue-900  border-0 text-white hover:bg-blue-800"> Save </button>
</div>

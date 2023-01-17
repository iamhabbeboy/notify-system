<script lang="ts">
  import { modal } from "../store/modal";

  const handleCloseModal = () => {
    modal.update( value => value = false);
  }

</script>

<div class="backdrop { $modal ? '' : 'hidden'}">
  <div
    id="myModal"
    class="h-auto mx-auto w-11/12 md:w-2/6 p-5  bg-white rounded-md "
  >
    <div class="flex flex-col w-full h-auto ">
      <!-- Header -->
      <div class="flex w-full h-auto justify-center items-center">
        <div
          class="flex w-10/12 h-auto py-3 justify-center items-center text-2xl font-bold"
        >
          Service Setting
        </div>
        <div class="flex w-1/12 h-auto justify-center cursor-pointer">
          <button on:click={handleCloseModal}>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="24"
              height="24"
              viewBox="0 0 24 24"
              fill="none"
              stroke="#000000"
              stroke-width="2"
              stroke-linecap="round"
              stroke-linejoin="round"
              class="feather feather-x"
              ><line x1="18" y1="6" x2="6" y2="18" /><line
                x1="6"
                y1="6"
                x2="18"
                y2="18"
              /></svg
            ></button
          >
        </div>
        <!--Header End-->
      </div>
      <!-- Modal Content-->
      <div
        class="w-full h-auto py-10 px-2 bg-gray-200 rounded text-center text-gray-500"
      >
        <div>
          <div>
            <input
              type="text"
              placeholder="Slack"
              class="border-2 mb-3 w-full border-gray-300 bg-white h-10 px-5 pr-5 rounded-lg text-sm focus:outline-none"
            />
          </div>
          <div>
            <textarea
              placeholder="Config"
              style="height: 200px"
              class="border-2 w-full border-gray-300 bg-white p-5 rounded-lg text-sm focus:outline-none"
            />
          </div>
        </div>
      </div>
      <!-- End of Modal Content-->
    </div>
  </div>
</div>

<!-- component -->
<style>
  /* div#myModal {
    position: absolute;
    top: 0;
    z-index: 1000;
  } */

  div#myModal[open] {
    animation: appear 0.15s cubic-bezier(0, 1.8, 1, 1.8);
  }

  .backdrop {
    width: 100%;
    height: 100vh;
    z-index: 1000;
    position: fixed;
    top: 0;
    padding-top: 20px;
    /* display: none; */
    margin: auto;
    background: rgba(0, 0, 0, 0.5);
    /* backdrop-filter: blur(3px); */
  }

  @keyframes appear {
    from {
      opacity: 0;
      transform: translateX(-3rem);
    }

    to {
      opacity: 1;
      transform: translateX(0);
    }
  }
</style>

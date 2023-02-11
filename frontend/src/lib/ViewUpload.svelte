<script>
  import Box from "./Box.svelte";
  import CheckCircle from "svelte-material-icons/CheckCircle.svelte";
  import { createEventDispatcher } from "svelte";
  import { fade, fly } from "svelte/transition";

  export let urlPath = "";

  let toast = false;

  const dispatch = createEventDispatcher();

  function handleCopy() {
    navigator.clipboard.writeText(urlPath);
    toast = true;
    setTimeout(() => {
      toast = false;
    }, 3000);
  }
</script>

{#if toast}
  <div
    class="absolute top-10 inset-x-0"
    in:fly={{ y: -200, duration: 500 }}
    out:fade
  >
    <div class="py-4 mx-auto w-fit px-10 rounded-xl bg-green-500">
      <p class="text-white inline-flex items-center gap-2">
        URL copied to clipboard!
      </p>
    </div>
  </div>
{/if}

<Box
  ><div class="flex flex-col justify-center items-center gap-4">
    <CheckCircle size="44" color="#219653" />
    <h1 class="font-medium text-3xl text-gray-600 pb-2">
      Upload Successfully!
    </h1>
    <div class="max-w-[30rem] max-h-[20rem] overflow-hidden rounded-xl">
      <img src={urlPath} alt="uploaded" />
    </div>
    <div
      class="w-full bg-gray-100 rounded-xl border-2 max-w-[30rem] flex items-center justify-between gap-1"
    >
      <p class="truncate text-sm ml-4">
        {urlPath}
      </p>
      <button
        class="rounded-xl bg-blue-500 text-white py-2 px-4"
        on:click={handleCopy}>Copy</button
      >
    </div>
    <button
      class="rounded-xl bg-blue-500 text-white py-2 px-4"
      on:click={() => dispatch("home")}
    >
      Back to home
    </button>
  </div>
</Box>

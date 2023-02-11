<script lang="ts">
  import Alert from "svelte-material-icons/Alert.svelte";
  import { fade, fly } from "svelte/transition";
  import { createEventDispatcher } from "svelte";

  const dispatch = createEventDispatcher();
</script>

<div
  class="absolute top-10 inset-x-0"
  in:fly={{ y: -200, duration: 500 }}
  out:fade
>
  <div class="py-4 mx-auto w-fit px-10 rounded-xl bg-red-500">
    <p class="text-white inline-flex items-center gap-2">
      <Alert />
      An error has occured while uploading your file. Please try again.
      <button class="underline" on:click={() => dispatch("dismiss")}
        >Dismiss</button
      >
    </p>
  </div>
</div>

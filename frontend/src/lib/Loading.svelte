<script>
  import Box from "./Box.svelte";
</script>

<Box>
  <p class="text-gray-500 mb-4">Uploading...</p>
  <div class="relative h-2 w-[20rem] bg-gray-200 rounded-xl overflow-hidden">
    <div class="absolute h-2 bg-blue-500 rounded-xl w-1/3 loader" />
  </div>
</Box>

<style>
  .loader {
    left: -100%;
    animation: load 2s linear infinite;
  }

  @keyframes load {
    0% {
      left: -50%;
    }
    100% {
      left: 100%;
    }
  }
</style>

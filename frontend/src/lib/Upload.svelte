<script lang="ts">
  import { createEventDispatcher } from "svelte";
  import Box from "./Box.svelte";
  import Background from "../assets/image.svg";

  const dispatch = createEventDispatcher();

  let dragging = false;

  let error = false;

  const validTypes = ["image/jpeg", "image/png"];

  function dropHandler(e: DragEvent) {
    let file = e.dataTransfer.files[0];
    if (file && validTypes.includes(file.type)) {
      dispatch("upload", { file });
    } else {
      error = true;
    }
  }

  function changeHandler(e: Event) {
    let file = (e.target as HTMLInputElement).files[0];
    if (file && validTypes.includes(file.type)) {
      dispatch("upload", { file });
    } else {
      error = true;
    }
  }
</script>

<Box>
  <div class="flex flex-col justify-center items-center">
    <h1 class="font-medium text-3xl">Upload your image</h1>
    <p class="font-medium text-gray-600 my-4 ">File should be Jpeg, Png...</p>

    <div
      class="grid place-content-center place-items-center gap-10 w-[24rem] h-[18rem] border-dashed border border-blue-500 rounded-xl {dragging
        ? 'bg-blue-100'
        : ' bg-blue-50'}"
      id="drop-area"
      on:drop|preventDefault={dropHandler}
      on:dragover|preventDefault
      on:dragenter|preventDefault={() => (dragging = true)}
      on:dragleave|preventDefault={() => (dragging = false)}
    >
      <img
        src={Background}
        alt="upload"
        class="bg-gray-100"
        width="150px"
        height="100px"
      />
      <p class="text-gray-400">Drag & Drop your image here</p>
    </div>
    {#if error}
      <p class="text-red-500 text-sm">Please provide a valid image file</p>
    {/if}

    <p class="text-gray-400 my-4">or</p>

    <label
      for="file-upload"
      class="px-[1rem] py-[0.5rem] bg-blue-500 hover:bg-blue-600 text-white rounded-xl font-medium cursor-pointer"
      >Choose a file</label
    >

    <input
      id="file-upload"
      type="file"
      class="hidden"
      on:change={changeHandler}
    />
  </div>
</Box>

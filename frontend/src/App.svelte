<script lang="ts">
  import ErrorToast from "./lib/ErrorToast.svelte";
  import Loading from "./lib/Loading.svelte";
  import Upload from "./lib/Upload.svelte";
  import ViewUpload from "./lib/ViewUpload.svelte";

  let loading = false;
  let uploaded = false;
  let error = false;

  let urlPath = "";

  async function handleUpload(e: CustomEvent) {
    loading = true;
    const file = e.detail.file as File;
    const formData = new FormData();
    formData.append("file", file, file.name);
    try {
      // Github pages cannot make proxy request
      const path =
        import.meta.env.MODE === "development"
          ? "http://localhost:3000/api/image"
          : "https://go-server-p05m.onrender.com/api/image";
      const response = await fetch(path, {
        method: "POST",
        body: formData,
      });
      const result = await response.json();
      urlPath = result.output;
      uploaded = true;
    } catch (err) {
      error = true;
    } finally {
      setTimeout(() => {
        loading = false;
      }, 2000);
    }
  }
</script>

<main class="h-screen w-full bg-neutral-200 grid place-content-center">
  {#if loading}
    <Loading />
  {:else if uploaded}
    <ViewUpload {urlPath} on:home={() => (uploaded = false)} />
  {:else}
    <Upload on:upload={handleUpload} />
  {/if}
</main>

{#if error}
  <ErrorToast on:dismiss={() => (error = false)} />
{/if}

<footer class="flex justify-center py-4 text-gray-600 fixed bottom-0 inset-x-0">
  <p>
    Created by <a
      href="https://github.com/Manethpak"
      target="_blank"
      rel="noopener noreferrer"
      class="text-gray-700 underline">Manethpak</a
    > - devChallenge.io
  </p>
</footer>

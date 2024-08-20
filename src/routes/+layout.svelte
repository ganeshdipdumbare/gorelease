<script>
  import "../app.css";
  import ToggleMode from "./ToggleMode.svelte";
</script>

<div class="flex min-h-screen flex-col">
  <ToggleMode />
  <div class="flex min-h-screen items-center justify-center p-4">
    <slot />
  </div>
</div>

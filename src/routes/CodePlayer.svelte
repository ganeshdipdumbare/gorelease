<script lang="ts">
  import { Button } from "$lib/components/ui/button/index.js";
  import * as Card from "$lib/components/ui/card/index.js";
  import "highlight.js/styles/github-dark.css";
  import { Label } from "$lib/components/ui/label/index.js";
  import { CodeBlock } from "svhighlight";

  export let code = ``;
  export let title = ``;
  export let description = ``;
  export let highlightLines = ``;

  let data = { output: "", error: "" };
  let loader = false;
  async function playCode() {
    loader = true;
    const url = `https://gorelease-1-v5611824.deta.app/backend/api/v1/playcode`;
    const requestBody = {
      code: code,
    };

    const response = await fetch(url, {
      method: "POST",
      body: JSON.stringify(requestBody),
      headers: {
        Origin: "https://gorelease-1-v5611824.deta.app",
      },
    });

    const respData = await response.json();
    data.output = respData.output;
    data.error = respData.error;
    loader = false;
  }
</script>

<Card.Root class="w-200 h-200">
  <Card.Header>
    <Card.Title>{title}</Card.Title>
    <Card.Description class="overflow-auto max-h-20">
      <pre>{description}</pre>
    </Card.Description>
  </Card.Header>
  <Card.Content>
    <form>
      <CodeBlock
        {code}
        rounded="rounded-lg"
        language="go"
        showFocusButtons={false}
        {highlightLines}
      />
    </form>
  </Card.Content>
  <Card.Footer class="flex justify-end">
    <Button on:click={playCode} class="w-10px">
      {#if loader}
        <div
          class="animate-spin inline-block size-6 border-[3px] border-current border-t-transparent text-white rounded-full mx-auto"
          role="status"
          aria-label="loading"
        >
          <span class="sr-only">Loading...</span>
        </div>
      {:else}
        ▶ Run
      {/if}
    </Button>
  </Card.Footer>

  {#if data.output !== "" && data.output !== "undefined" && loader === false}
    <Card.Footer class="flex justify-between">
      <Label>Output:</Label>
      <pre class="overflow-auto max-h-20">{data.output}</pre>
    </Card.Footer>
  {:else if data.error !== "" && data.error !== "undefined" && loader === false}
    <Card.Footer class="flex justify-between">
      <Label>Error</Label>
      <pre>{data.error}</pre>
    </Card.Footer>
  {/if}
</Card.Root>

<script lang="ts">
  import Button from "$lib/components/ui/button/button.svelte";

  let count: number = 0;
  const increment = () => {
    count += 1;
  };
</script>

<Button on:click={increment}>
  count is {count}
</Button>

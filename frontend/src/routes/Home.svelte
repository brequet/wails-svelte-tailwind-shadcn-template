<script lang="ts">
  import Counter from "$lib/Counter.svelte";
  import Button from "$lib/components/ui/button/button.svelte";
  import { Greet } from "$wailsjs/hello/HelloService.js";
  import logo from "../assets/images/logo-universal.png";

  let resultText: string = "Please enter your name below 👇";
  let name: string;

  function greet(): void {
    Greet(name).then((result) => (resultText = result));
  }
</script>

<div class="flex flex-col items-center space-y-4">
  <img alt="Wails logo" id="logo" src={logo} class="w-64 h-64" />
  <div class="result text-center text-xl font-bold" id="result">
    {resultText}
  </div>
  <div class="input-box flex flex-row items-center space-x-2">
    <input
      autocomplete="off"
      bind:value={name}
      class="input w-full p-2 border border-gray-300 rounded-md focus:outline-none focus:border-blue-500"
      id="name"
      type="text"
      placeholder="Enter your name"
    />
    <Button
      class="btn bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
      on:click={greet}>Greet</Button
    >
  </div>

  <Counter />
</div>

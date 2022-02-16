<script lang="ts">
  import type { TWorld } from '$lib/components/World.svelte'
  import World from '$lib/components/World.svelte'
  import { onMount } from 'svelte'
  import axios from 'axios'

  let world: TWorld | null = null

  onMount(async () => {
    const { data } = await axios({
      url: 'http://localhost:8000/random',
    })
    world = data
  })

  async function next() {
    if (!world) return

    const { data } = await axios({
      url: 'http://localhost:8000/next',
      method: 'POST',
      data: world,
    })
    world = data
  }
</script>

<h1>Langton Ant</h1>
<h2>Camao</h2>

{#if world}
  <World {world} />
{/if}

<button on:click={next}>Next</button>

<style>
  button {
    appearance: none;
    padding: 0.5rem 2rem;
    outline: none;
    border: none;
    border-radius: 0.1rem;
    margin-top: 1rem;
    cursor: pointer;
  }
</style>

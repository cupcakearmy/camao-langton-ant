<script lang="ts" context="module">
  export enum TDirection {
    Up = 0,
    Right = 1,
    Down = 2,
    Left = 3,
  }

  export type TWorld = {
    field: boolean[][]
    ant: {
      x: number
      y: number
      direction: TDirection
    }
  }
</script>

<script lang="ts">
  export let world: TWorld
</script>

{#each world.field as row}
  <div class="row">
    {#each row as cell}
      {#if cell}
        <div class="cell white" />
      {:else}
        <div class="cell" />
      {/if}
    {/each}
  </div>
{/each}

<style>
  .row {
    --light: rgb(243, 239, 239);
    --dark: rgb(21, 20, 20);
    display: flex;
  }
  .cell {
    display: inline-block;
    margin: 0.05rem;
    width: 2rem;
    aspect-ratio: 1/1;

    border: 0.15rem solid rgb(7, 7, 7);
    border-radius: 0.1rem;

    background-color: var(--dark);
    color: var(--light);
  }

  .cell.white {
    background-color: var(--light);
    color: var(--dark);
  }
</style>

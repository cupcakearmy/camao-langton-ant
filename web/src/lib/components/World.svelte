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

  const mapping = {
    0: '↑',
    1: '→',
    2: '↓',
    3: '←',
  }

  function getContentForCell(world: World, x: number, y: number): string {
    if (world.ant.x === x && world.ant.y === y) {
      return mapping[world.ant.direction]
    }
    return ''
  }
</script>

{#each world.field as row, y}
  <div class="row">
    {#each row as cell, x}
      <div class="cell" class:white={cell}>
        {getContentForCell(world, x, y)}
      </div>
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
    display: inline-flex;
    justify-content: center;
    align-items: center;

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
